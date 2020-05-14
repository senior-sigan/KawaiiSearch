package files

import (
	"context"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

func TestFilesStorage_Connect(t *testing.T) {
	var files []string

	root, err := ioutil.TempDir("", "files_storage_test")
	require.NoError(t, err)
	_, err = ioutil.TempFile(root, "*.jpg")
	require.NoError(t, err)

	fs := NewFilesStorage(root)

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	fs.Connect(ctx)
	require.NoError(t, err)
	defer fs.Close()

	go func() {
		for {
			select {
			case file, ok := <-fs.OnNewFile:
				if !ok {
					return
				}
				files = append(files, file)
			case <- ctx.Done():
				err := fs.Close()
				require.NoError(t, err)
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, len(files), 1)

	_, err = ioutil.TempFile(root, "*.jpg")
	require.NoError(t, err)

	time.Sleep(500 * time.Millisecond)
	assert.Equal(t, len(files), 2)

	<-ctx.Done()

	assert.Equal(t, len(files), 2)
}
