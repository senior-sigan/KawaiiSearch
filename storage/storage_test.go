package storage

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestStorage_GetFeatures(t *testing.T) {
	storage := NewInMemStorage()
	defer storage.Close()

	features := make([]float32, 1024)
	for i := range features {
		features[i] = rand.Float32()
	}
	err := storage.UpsertFeatures("test", features)
	require.NoError(t, err)

	featuresRes, err := storage.GetFeatures("test")
	require.NoError(t, err)

	assert.Equal(t, featuresRes, features)
}

func TestStorage_GetFiles(t *testing.T) {
	storage := NewInMemStorage()
	defer storage.Close()

	err := storage.UpsertFeatures("file1", []float32{1,2,3})
	require.NoError(t, err)
	err = storage.UpsertFeatures("file2", []float32{1,2,3})
	require.NoError(t, err)

	keys, err := storage.GetFiles()
	require.NoError(t, err)
	assert.Equal(t, len(keys), 2)
	assert.Equal(t, keys[0], "file1")
	assert.Equal(t, keys[1], "file2")
}
