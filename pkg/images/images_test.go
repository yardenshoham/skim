package images

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromManifests(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "deployment.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	extractor := NewExtractor()
	err = extractor.ExtractFromManifests(ctx, file, images)
	require.NoError(t, err)
	require.Len(t, images, 1)
	require.Contains(t, images, "example.com/processor:1.2.3")
}

func TestFromManifestsUnknownGVK(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	extractor := NewExtractor()
	err = extractor.ExtractFromManifests(ctx, file, images)
	require.Error(t, err)
	extractorSkip := NewExtractor()
	_, err = file.Seek(0, io.SeekStart)
	require.NoError(t, err)
	extractorSkip.UnknownGVKBehavior = UnknownGVKSkip
	err = extractorSkip.ExtractFromManifests(ctx, file, images)
	require.NoError(t, err)
	require.Empty(t, images)
}

func TestFromManifestsCustomGVK(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	images := make(map[string]struct{})
	extractor := NewExtractor()
	extractor.GVKMappings = map[string]func(map[string]any, map[string]struct{}) error{
		"v1.Podonkadonk": func(manifest map[string]any, output map[string]struct{}) error {
			return V1Pod(manifest, output)
		},
	}
	err = extractor.ExtractFromManifests(ctx, file, images)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{"nginx:1.21.0": {}, "busybox:1.35": {}}, images)
}

func TestExtractImagesFromFreeText(t *testing.T) {
	t.Parallel()
	file, err := os.ReadFile(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	images := make(map[string]struct{})
	err = extractImagesFromFreeText(string(file), images)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		"nginx:1.21.0": {},
		"busybox:1.35": {},
	}, images)
}

func TestFromManifestsUnknownGVKFreeText(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "unknown_gvk.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	extractor := NewExtractor()
	extractor.UnknownGVKBehavior = UnknownGVKFreeText
	err = extractor.ExtractFromManifests(ctx, file, images)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		"nginx:1.21.0": {},
		"busybox:1.35": {},
	}, images)
}

func TestFromManifestsMultipleObjects(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	file, err := os.Open(filepath.Join("..", "..", "testdata", "multi_object.yaml"))
	require.NoError(t, err)
	defer file.Close()
	images := make(map[string]struct{})
	extractor := NewExtractor()
	err = extractor.ExtractFromManifests(ctx, file, images)
	require.NoError(t, err)
	require.Len(t, images, 2)
	require.Contains(t, images, "nginx:1.21.0")
	require.Contains(t, images, "redis:7.0")
}
