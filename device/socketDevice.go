package device

import (
	"errors"
	"fmt"
	"net"
)

type PrinterSocket struct {
	Host string
	Port string
}

func (p *PrinterSocket) Write(data []byte) (int, error) {
	conn, err := net.Dial("tcp", p.Host+":"+p.Port)
	if err != nil {
		errStr := fmt.Sprintf("Error: %s", err)
		return 0, errors.New(errStr)
	}
	defer conn.Close()

	written, err := conn.Write(data)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	return written, err
}
