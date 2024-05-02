package thermalprinter

import (
	"sync"

	"thermalPrinter/device"
	encodings "thermalPrinter/encoding"
)

var (
	instance *TP
	once     sync.Once
)

type TP struct {
	TextProcessor encodings.Parser
	Device        device.WriteableDevice
}

func (p *TP) Print(message string) {
	processed := p.TextProcessor.Process(message)
	p.Device.Write(processed)
}

func GetInstance(customDev device.WriteableDevice) *TP {
	once.Do(func() {
		var p encodings.Parser
		p.Initialize()

		devSocket := device.PrinterSocket{
			Host: "",
			Port: "",
		}

		instance = &TP{
			Device:        &devSocket,
			TextProcessor: p,
		}
	})

	return instance
}
