package images

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImages(t *testing.T) {
	t.Parallel()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "deployment.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	err = FromManifests(file, images)
	require.NoError(t, err)
	require.Len(t, images, 1)
	require.Contains(t, images, "example.com/processor:1.2.3")
}
