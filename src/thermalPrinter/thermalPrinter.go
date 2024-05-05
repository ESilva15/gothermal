package thermalprinter

import (
	"fmt"
	"sync"

	"thermalPrinter/config"
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
	_, err := p.Device.Write(processed)
	if err != nil {
		fmt.Println("Err: ", err)
	}
}

func GetInstance(customDev device.WriteableDevice) *TP {
	once.Do(func() {
		var p encodings.Parser
		p.Initialize()

		cfg := config.GetInstance()
		device := cfg.GetWriteableDevice()

		instance = &TP{
			Device:        device,
			TextProcessor: p,
		}
	})

	return instance
}
