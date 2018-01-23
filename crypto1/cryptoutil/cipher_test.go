package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCipher(t *testing.T) {
	t.Run("String()", func(t *testing.T) {
		c := NewCipher("a23c")
		s := c.String()
		assert.Equal(t, "a23c", s)
	})
	t.Run("Binary()", func(t *testing.T) {
		c := NewCipher("a23c")
		s := c.Binary()
		assert.Equal(t, "1010001000111100", s)
	})
	t.Run("XOR() same length", func(t *testing.T) {
		c1 := NewCipher("f3")
		c2 := NewCipher("26")

		c3, err := c1.XOR(c2)
		assert.NoError(t, err)
		assert.Equal(t, "d5", c3.String())
	})
	t.Run("XOR() different length", func(t *testing.T) {
		c1 := NewCipher("f342")
		c2 := NewCipher("26")

		c3, err := c1.XOR(c2)
		assert.NoError(t, err)
		assert.Equal(t, "d542", c3.String())
	})
}