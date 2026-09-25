//go:build windows

package vt

import "os"

// openTTY opens the console input and output buffers.
func openTTY() (in, out *os.File, closeFn func(), err error) {
	// Console mode calls need both read and write access to the handles.
	in, err = os.OpenFile("CONIN$", os.O_RDWR, 0)
	if err != nil {
		return nil, nil, nil, err
	}
	out, err = os.OpenFile("CONOUT$", os.O_RDWR, 0)
	if err != nil {
		_ = in.Close()
		return nil, nil, nil, err
	}
	return in, out, func() {
		_ = in.Close()
		_ = out.Close()
	}, nil
}
