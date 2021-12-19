package register

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegative(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b10000000

	expected := true
	actual := ps.Negative()
	assert.Equal(t, expected, actual)

	ps.data = 0b01101100

	expected = false
	actual = ps.Negative()
	assert.Equal(t, expected, actual)
}

func TestOverflow(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b01000000

	expected := true
	actual := ps.Overflow()
	assert.Equal(t, expected, actual)

	ps.data = 0b10101100

	expected = false
	actual = ps.Overflow()
	assert.Equal(t, expected, actual)
}

func TestBreak(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b00010000

	expected := true
	actual := ps.Break()
	assert.Equal(t, expected, actual)

	ps.data = 0b10101100

	expected = false
	actual = ps.Break()
	assert.Equal(t, expected, actual)
}

func TestDecimalMode(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b00001000

	expected := true
	actual := ps.DecimalMode()

	assert.Equal(t, expected, actual)

	ps.data = 0b10100101

	expected = false
	actual = ps.DecimalMode()
	assert.Equal(t, expected, actual)
}

func TestIRQBDisable(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b00000100

	expected := true
	actual := ps.IRQBDisable()
	assert.Equal(t, expected, actual)

	ps.data = 0b10101010

	expected = false
	actual = ps.IRQBDisable()
	assert.Equal(t, expected, actual)
}

func TestZero(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b00000010

	expected := true
	actual := ps.Zero()
	assert.Equal(t, expected, actual)

	ps.data = 0b10101100

	expected = false
	actual = ps.Zero()
	assert.Equal(t, expected, actual)
}

func TestCarry(t *testing.T) {
	var ps ProcessorStatus

	ps.data = 0b00000001

	expected := true
	actual := ps.Carry()
	assert.Equal(t, expected, actual)

	ps.data = 0b10101100

	expected = false
	actual = ps.Carry()
	assert.Equal(t, expected, actual)
}

//TODO: Write set and reset tests
