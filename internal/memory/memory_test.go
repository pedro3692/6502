package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	var m Memory

	expected := byte(0x55)
	m.memory[0x145] = expected

	actual := m.Read([2]byte{0x45, 0x01})
	assert.Equal(t, expected, actual)
}

func TestStore(t *testing.T) {
	var m Memory

	expected := byte(0x55)

	m.Store([2]byte{0x45, 0x01}, expected)
	assert.Equal(t, expected, m.memory[0x145])
}

func TestSet(t *testing.T) {
	var (
		m  Memory
		ma [Size]byte
	)

	ma[0] = 0x01

	expected := [Size]byte{0x01}

	m.Set(ma)
	assert.Equal(t, expected, m.memory)
}
