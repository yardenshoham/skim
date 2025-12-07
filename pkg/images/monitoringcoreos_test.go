package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMonitoringCoreosComV1Alertmanager(t *testing.T) {
	t.Parallel()
	alertmanager := map[string]any{
		"apiVersion": "monitoring.coreos.com/v1",
		"kind":       "Alertmanager",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := MonitoringCoreosComV1Alertmanager(alertmanager, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestMonitoringCoreosComV1Prometheus(t *testing.T) {
	t.Parallel()
	prometheus := map[string]any{
		"apiVersion": "monitoring.coreos.com/v1",
		"kind":       "Prometheus",
		"spec": map[string]any{
			"image": nginxLatest,
			"containers": []any{
				map[string]any{
					"name":  "sidecar",
					"image": busybox128,
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := MonitoringCoreosComV1Prometheus(prometheus, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}
