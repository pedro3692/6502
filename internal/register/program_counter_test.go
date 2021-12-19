package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPCLoad(t *testing.T) {
	var pc ProgramCounter

	expected := [2]byte{0x34, 0xf1}

	pc.Load(expected)

	actual := pc.data
	assert.Equal(t, expected, actual)
}

func TestPCRead(t *testing.T) {
	var pc ProgramCounter

	expected := [2]byte{0xfc, 0xff}

	pc.data = expected

	actual := pc.Read()
	assert.Equal(t, expected, actual)

	expected = [2]byte{0x00, 0x01}

	pc.data = expected

	actual = pc.Read()
	assert.Equal(t, expected, actual)
}

func TestPCReset(t *testing.T) {
	var pc ProgramCounter

	pc.Reset()

	expected := [2]byte{0xfc, 0xff}
	actual := pc.data
	assert.Equal(t, expected, actual)
}
