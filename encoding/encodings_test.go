package encodings

import (
	"bytes"
	"testing"
	// "github.com/stretchr/testify/mock"
)

func TestProcessWithASCII(t *testing.T) {
	data := "m"
	expect := []byte("m")

	var p Parser
	res := p.Process(data)

	if !bytes.Equal(res, expect) {
		t.Errorf("%x is different of %x\n", res, expect)
	}
}

func TestProcessWithUTF8(t *testing.T) {
	data := "€"
	expect := []byte{0x1B, 0x74, 0x10, 0x80}

	var p Parser
	p.Initialize()
	res := p.Process(data)

	if !bytes.Equal(res, expect) {
		t.Errorf("%x is different of %x\n", res, expect)
	}
}

func TestProcessUnavailableChar(t *testing.T) {
	data := "�"
	expect := []byte("X")

	var p Parser
	res := p.Process(data)

	if !bytes.Equal(res, expect) {
		t.Errorf("%x is different of %x\n", res, expect)
	}
}

func TestProcessNewLineChar(t *testing.T) {
	data := "\\n"
	expect := []byte{0x0A}

	var p Parser
	res := p.Process(data)

	if !bytes.Equal(res, expect) {
		t.Errorf("%x is different of %x\n", res, expect)
	}
}

func TestProcessNewTabChar(t *testing.T) {
	data := "\\t"
	expect := []byte{0x09}

	var p Parser
	res := p.Process(data)

	if !bytes.Equal(res, expect) {
		t.Errorf("%x is different of %x\n", res, expect)
	}
}
