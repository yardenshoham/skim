package images

import "fmt"

// UnknownGVKError is returned when a manifest has an unexpected GVK.
type UnknownGVKError struct {
	GVK      string
	Manifest map[string]any
}

func (e *UnknownGVKError) Error() string {
	return fmt.Sprintf("failed to detect Group Version Kind: %s, manifest: %+v", e.GVK, e.Manifest)
}
