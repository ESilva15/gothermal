package config

import (
	"log"
	"os"
	"slices"
	"sync"
	"thermalPrinter/device"

	"gopkg.in/yaml.v2"
)

var (
	instance       *Configuration
	once           sync.Once
	validUseValues = []string{
		"socket",
		"usb",
	}
)

type usbDev struct {
	Path string `yaml:"path"`
}

type socketDev struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Configuration struct {
	Use    string     `yaml:"use"`
	Usb    *usbDev    `yaml:"usb"`
	Socket *socketDev `yaml:"socket"`
}

func GetInstance() *Configuration {
	once.Do(func() {
		instance = &Configuration{}
		instance.loadConfiguration()
	})

	return instance
}

func (c *Configuration) loadConfiguration() {
	confPath := "/home/esilva/Desktop/projetos/gothermal/src/config.yaml"

	file, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Unable to open configuration file [%s]: %s", confPath,
			err.Error())
	}

	err = yaml.Unmarshal(file, &instance)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err.Error())
	}

	if !slices.Contains(validUseValues, instance.Use) {
		log.Fatalf("Unusual value for `Use`.")
	}
}

func (c *Configuration) GetWriteableDevice() device.WriteableDevice {
	if c.Use == "usb" {
		return &device.PrinterUSB{
			Path: c.Usb.Path,
		}
	} else if c.Use == "socket" {
		return &device.PrinterSocket{
			Host: c.Socket.Host,
			Port: c.Socket.Port,
		}
	}

	return nil
}
