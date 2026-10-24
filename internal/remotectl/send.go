package remotectl

import (
	"encoding/json"
	"fmt"
	"net"
)

// CommandData encodes an action that can be sent to a Server.
type CommandData struct {
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

type Sender struct {
	conn     net.Conn
	requests chan req
}

func Client(ep Endpoint) (*Sender, error) {
	p, err := ep.path()
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("unix", p)
	if err != nil {
		return nil, err
	}

	s := &Sender{
		conn:     conn,
		requests: make(chan req),
	}
	go s.run()

	if err := s.Send(&CommandData{Action: "ping"}); err != nil {
		_ = s.Close()
		return nil, err
	}

	return s, nil
}

func (s *Sender) Send(cmd *CommandData) error {
	res := make(chan error, 1)
	s.requests <- req{cmd: cmd, res: res}
	return <-res
}

func (s *Sender) Close() error {
	close(s.requests)
	return s.conn.Close()
}

func (s *Sender) run() {
	for r := range s.requests {
		r.res <- sendCommand(s.conn, r.cmd)
	}
}

type req struct {
	cmd *CommandData
	res chan<- error
}

func sendCommand(conn net.Conn, cmd *CommandData) error {
	if err := json.NewEncoder(conn).Encode(cmd); err != nil {
		return err
	}
	_, err := fmt.Fprintln(conn)
	if err != nil {
		return err
	}

	var rsp CommandData
	err = json.NewDecoder(conn).Decode(&rsp)
	if err != nil {
		return fmt.Errorf("didn't get ack: %w", err)
	}
	if rsp.Action != "ack" {
		return fmt.Errorf("got %q instead of ack", rsp.Action)
	}
	return nil
}
