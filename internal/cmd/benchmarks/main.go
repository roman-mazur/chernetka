// Command benchmarks runs the benchmarks of the module and records their results
// in .benchmarks/results.json in the module root.
//
// The results file holds up to the last 3 measurements: every run of the command
// adds one and drops the oldest. TestNoDegradation compares the last measurement
// with the previous ones recorded on the same machine.
//
// A noisy benchmark can override the tolerated increase of its time with a log line:
//
//	b.Log("tolerance=20%")
//
// Usage:
//
//	go run ./internal/cmd/benchmarks [-bench regexp] [-count n] [packages]
//
// or with go generate ./internal/cmd/benchmarks.
//
// The packages are relative to the module root and default to ./...
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	bench := flag.String("bench", ".", "run only the benchmarks matching the regular expression")
	count := flag.Int("count", 3, "run each benchmark n times in one measurement")
	flag.Parse()
	pkgs := flag.Args()
	if len(pkgs) == 0 {
		pkgs = []string{"./..."}
	}

	root, err := moduleRoot()
	if err != nil {
		log.Fatal(err)
	}
	goVersion, err := goEnv("GOVERSION")
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"test", "-run=^$", "-bench=" + *bench, "-benchmem", "-count=" + strconv.Itoa(*count)}
	cmd := exec.Command("go", append(args, pkgs...)...)
	cmd.Dir = root
	var out bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &out)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("running the benchmarks: %s", err)
	}

	m := Measurement{Time: time.Now().UTC().Truncate(time.Second), GoVersion: goVersion}
	if err := parseOutput(&out, &m); err != nil {
		log.Fatalf("parsing the benchmark output: %s", err)
	}
	if len(m.Results) == 0 {
		log.Fatal("no benchmark results")
	}

	path := filepath.Join(root, resultsFile)
	ms, err := load(path)
	if err != nil {
		log.Fatal(err)
	}
	ms = add(ms, m)
	if err := save(path, ms); err != nil {
		log.Fatal(err)
	}
	log.Printf("recorded %d results in %s (%d/%d measurements)", len(m.Results), resultsFile, len(ms), maxMeasurements)
}

func moduleRoot() (string, error) {
	gomod, err := goEnv("GOMOD")
	if err != nil {
		return "", err
	}
	if gomod == "" || gomod == os.DevNull {
		return "", fmt.Errorf("not in a module")
	}
	return filepath.Dir(gomod), nil
}

func goEnv(name string) (string, error) {
	out, err := exec.Command("go", "env", name).Output()
	if err != nil {
		return "", fmt.Errorf("go env %s: %w", name, err)
	}
	return strings.TrimSpace(string(out)), nil
}
