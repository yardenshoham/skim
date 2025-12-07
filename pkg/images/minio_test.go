package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinIOMinIOV2Tenant(t *testing.T) {
	t.Parallel()
	tenant := map[string]any{
		"apiVersion": "minio.min.io/v2",
		"kind":       "Tenant",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := MinIOMinIOV2Tenant(tenant, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}
