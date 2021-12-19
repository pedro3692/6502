package register

import (
	"testing"

	"github.com/pedro3692/6502/internal/memory"
	"github.com/stretchr/testify/assert"
)

func TestRLoad(t *testing.T) {
	var r Register

	expected := byte(0x45)

	r.Load(expected)
	assert.Equal(t, expected, r.data)
}

func TestRStore(t *testing.T) {
	var m memory.Memory
	r := New(&m)

	expected := byte(0x13)

	r.data = expected

	r.Store([2]byte{0x05, 0x00})
	assert.Equal(t, expected, r.memory.Read([2]byte{0x05, 0x00}))
}

func TestRRead(t *testing.T) {
	var r Register

	expected := byte(0x13)

	r.data = expected

	actual := r.Read()
	assert.Equal(t, expected, actual)
}
