package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPostgresqlCNPGIOV1Cluster(t *testing.T) {
	t.Parallel()
	cluster := map[string]any{
		"apiVersion": "postgresql.cnpg.io/v1",
		"kind":       "Cluster",
		"spec": map[string]any{
			"imageName": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := PostgresqlCNPGIOV1Cluster(cluster, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}
