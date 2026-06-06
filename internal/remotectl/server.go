package remotectl

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"path/filepath"
	"strings"
	"sync"

	"rmazur.io/chernetka/internal"
	"rmazur.io/chernetka/internal/logger"
)

// Server implements a local socket server to receive editor commands.
type Server struct {
	l net.Listener
}

type Executor interface {
	ExecuteCommand(cmd CommandData)
}

const socketName = "ctl.socket"

func NewServer() (*Server, error) {
	p, err := internal.UserDir()
	if err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", filepath.Join(p, socketName))
	if err != nil {
		return nil, err
	}
	return &Server{l: l}, nil
}

func (c *Server) Close() error {
	return c.l.Close()
}

func (c *Server) handle(e Executor, commands <-chan string, logf logger.Func) {
	for line := range commands {
		var cmd CommandData
		if err := json.Unmarshal([]byte(line), &cmd); err != nil {
			logf("cmd decode error: %s, data: %s", err, line)
			continue
		}
		e.ExecuteCommand(cmd)
	}
}

func (c *Server) Run(e Executor, logf logger.Func) {
	commands := make(chan string)
	defer close(commands)
	go c.handle(e, commands, logf)

	var wg sync.WaitGroup

	for {
		conn, err := c.l.Accept()
		if err != nil {
			if !errors.Is(err, io.EOF) && !errors.Is(err, net.ErrClosed) {
				logf("ctl accept error: %s", err)
			}
			break
		}

		scan := bufio.NewScanner(conn)
		scan.Split(bufio.ScanLines)
		wg.Go(func() {
			defer conn.Close()

			for scan.Scan() {
				line := strings.TrimSpace(scan.Text())
				if line != "" {
					commands <- line
				}
			}
		})
	}

	wg.Wait()
}
