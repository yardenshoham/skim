package images

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromManifests(t *testing.T) {
	t.Parallel()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "deployment.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	extractor := NewExtractor()
	err = extractor.FromManifests(file, images)
	require.NoError(t, err)
	require.Len(t, images, 1)
	require.Contains(t, images, "example.com/processor:1.2.3")
}

func TestFromManifestsUnknownGVK(t *testing.T) {
	t.Parallel()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	images := make(map[string]struct{})
	extractor := NewExtractor()
	err = extractor.FromManifests(file, images)
	require.Error(t, err)
	extractorSkip := NewExtractor()
	file.Close()
	file, err = os.Open(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	defer file.Close()
	extractorSkip.UnknownGVKBehavior = UnknownGVKSkip
	err = extractorSkip.FromManifests(file, images)
	require.NoError(t, err)
	require.Empty(t, images)
}
