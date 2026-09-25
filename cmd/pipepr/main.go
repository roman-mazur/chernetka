package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
	"rmazur.io/chernetka/internal/vt/escape"
)

var (
	size = flag.Int("size", 0, "size of the whole content")
)

func main() {
	flag.Parse()
	p := make(chan int, 16)      // progress from the io.Copy
	updates := make(chan int, 1) // total number of bytes sent for the visualization
	visDone := make(chan struct{})

	go visualizeProgress(visDone, updates)
	go receiveUpdates(updates, p)

	obs := &observer{
		Writer:          os.Stdout,
		progressChannel: p,
	}
	n, err := io.Copy(obs, os.Stdin)
	close(p)

	<-visDone
	log.Printf("copied %d bytes", n)

	if err != nil {
		log.Fatal(err)
	}
}

func receiveUpdates(updates chan<- int, p <-chan int) {
	defer close(updates)
	total := 0
	for update := range p {
		if update > 0 {
			total += update
			updates <- total
		}
	}
}

func visualizeProgress(visDone chan<- struct{}, updates <-chan int) {
	defer close(visDone)
	if *size == 0 {
		return
	}
	if !term.IsTerminal(int(os.Stderr.Fd())) {
		return
	}
	w, _, err := term.GetSize(int(os.Stderr.Fd()))
	if err != nil {
		return
	}
	pw := w - 16
	var lastUpdate time.Time
	for val := range updates {
		if time.Since(lastUpdate) < time.Millisecond*10 {
			continue
		}
		lastUpdate = time.Now()
		updateProgressBar(float64(val)/float64(*size), pw)
	}

	updateProgressBar(1, pw)
	escape.UpdateProgress(os.Stderr, escape.ProgressStateHidden, 0)
	fmt.Fprintln(os.Stderr)
}

func updateProgressBar(p float64, pw int) {
	barsCount := int(p * float64(pw))
	fmt.Fprint(os.Stderr,
		"\r[",
		strings.Repeat("*", barsCount),
		strings.Repeat(".", pw-barsCount),
		"] ",
		fmt.Sprintf("%.2f%%", p*100),
	)
	escape.UpdateProgress(os.Stderr, escape.ProgressStateDefault, int(p*100))
}

type observer struct {
	io.Writer
	progressChannel chan int
}

func (o *observer) Write(b []byte) (int, error) {
	n, err := o.Writer.Write(b)
	o.progressChannel <- n
	return n, err
}
