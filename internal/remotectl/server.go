package remotectl

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"
	"slices"
	"strings"
	"sync"

	"rmazur.io/chernetka/internal/logger"
)

// Server implements a local socket server to receive editor commands.
type Server struct {
	l net.Listener

	mu                sync.Mutex
	activeConnections []net.Conn
}

type Executor interface {
	ExecuteCommand(cmd CommandData)
}

// NewServer starts listening on the endpoint socket.
// A socket file left by a process that did not shut down gracefully is replaced.
func NewServer(ep Endpoint) (*Server, error) {
	p, err := ep.path()
	if err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", p)
	if err != nil && isStaleSocket(p) {
		_ = os.Remove(p)
		l, err = net.Listen("unix", p)
	}
	if err != nil {
		return nil, err
	}
	return &Server{l: l}, nil
}

// isStaleSocket checks whether the socket file exists but nobody listens on it.
func isStaleSocket(p string) bool {
	if _, err := os.Stat(p); err != nil {
		return false
	}
	conn, err := net.Dial("unix", p)
	if err != nil {
		return true
	}
	_ = conn.Close()
	return false
}

func (c *Server) Close() error {
	c.mu.Lock()
	for _, conn := range c.activeConnections {
		_ = conn.Close()
	}
	c.activeConnections = nil
	c.mu.Unlock()
	return c.l.Close()
}

func (c *Server) handle(e Executor, commands <-chan string, logf logger.Func) {
	for line := range commands {
		var cmd CommandData
		if err := json.Unmarshal([]byte(line), &cmd); err != nil {
			logf("cmd decode error: %s, data: %s", err, line)
			continue
		}
		if cmd.Action == "ping" {
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
		conn, err := c.accept()
		if err != nil {
			if !errors.Is(err, io.EOF) && !errors.Is(err, net.ErrClosed) {
				logf("ctl accept error: %s", err)
			}
			break
		}

		scan := bufio.NewScanner(conn)
		scan.Split(bufio.ScanLines)
		wg.Go(func() {
			defer c.close(conn)

			for scan.Scan() {
				line := strings.TrimSpace(scan.Text())
				if line != "" {
					commands <- line
					_ = json.NewEncoder(conn).Encode(CommandData{Action: "ack"})
				}
			}
		})
	}

	wg.Wait()
}

func (c *Server) accept() (net.Conn, error) {
	conn, err := c.l.Accept()
	if err == nil {
		c.mu.Lock()
		c.activeConnections = append(c.activeConnections, conn)
		c.mu.Unlock()
	}
	return conn, err
}

func (c *Server) close(conn net.Conn) {
	_ = conn.Close()
	c.mu.Lock()
	c.activeConnections = slices.DeleteFunc(c.activeConnections, func(c net.Conn) bool { return c == conn })
	c.mu.Unlock()
}
