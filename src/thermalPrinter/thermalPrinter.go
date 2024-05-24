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
	// Message + feed n lines
	processed = append(processed, []byte{0x1B, 0x64, 0x03}...)

	// Message + partial cut command
	// processed = append(processed, []byte{0x1B, 0x69}...)

	// Message + full cut command
	processed = append(processed, []byte{0x1B, 0x6D}...)

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
