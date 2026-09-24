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

// SendCommand delivers the command to the Server listening on the endpoint.
func SendCommand(ep Endpoint, cmd *CommandData) error {
	p, err := ep.path()
	if err != nil {
		return err
	}
	conn, err := net.Dial("unix", p)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(cmd); err != nil {
		return err
	}
	_, err = fmt.Fprintln(conn)
	return err
}
