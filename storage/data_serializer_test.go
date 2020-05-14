package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarshalFloats(t *testing.T) {
	t.Run("general case", func(t*testing.T) {
		array := []float32{3.14, 42.42, -23, 12}
		data, err := MarshalFloats(array)
		require.NoError(t, err)
		arrayUnm, err := UnmarshallFloats(data)
		require.NoError(t, err)
		assert.Equal(t, arrayUnm, array)
	})

	t.Run("zero array", func(t *testing.T) {
		var array []float32
		data, err := MarshalFloats(array)
		require.NoError(t, err)
		arrayUnm, err := UnmarshallFloats(data)
		require.NoError(t, err)
		assert.Equal(t, len(arrayUnm), 0)
	})
}
