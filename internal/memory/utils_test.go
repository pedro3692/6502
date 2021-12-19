package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLittleEndian(t *testing.T) {
	expected := uint16(0xfffc)
	actual := LittleEndian([2]byte{0xfc, 0xff})
	assert.Equal(t, expected, actual)
}
