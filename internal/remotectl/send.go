package remotectl

import (
	"encoding/json"
	"fmt"
	"net"
	"path/filepath"

	"rmazur.io/chernetka/internal"
)

// CommandData encodes an action that can be sent to a Server.
type CommandData struct {
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

func SendCommand(cmd *CommandData) error {
	p, err := internal.UserDir()
	if err != nil {
		return err
	}
	conn, err := net.Dial("unix", filepath.Join(p, socketName))
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
