package device

import (
  "os"
)

type PrinterUSB struct {
	Path string
}

func (p *PrinterUSB) Write(data []byte) (int, error) {
	printer, err := os.OpenFile(p.Path, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer printer.Close()

  return printer.Write(data)
}
