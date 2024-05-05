package device

type WriteableDevice interface {
	Write(data []byte) (int, error)
}

type PrinterDevice struct {
	Device WriteableDevice
}
