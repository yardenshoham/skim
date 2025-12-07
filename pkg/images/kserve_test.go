package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServingKserveIOV1alpha1ClusterServingRuntime(t *testing.T) {
	t.Parallel()
	clusterServingRuntime := map[string]any{
		"apiVersion": "serving.kserve.io/v1alpha1",
		"kind":       "ClusterServingRuntime",
		"spec": map[string]any{
			"containers": []any{
				map[string]any{
					"name":  "kserve-container",
					"image": nginxLatest,
				},
				map[string]any{
					"name":  "sidecar",
					"image": busybox128,
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1alpha1ClusterServingRuntime(clusterServingRuntime, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestServingKserveIOV1alpha1ServingRuntime(t *testing.T) {
	t.Parallel()
	servingRuntime := map[string]any{
		"apiVersion": "serving.kserve.io/v1alpha1",
		"kind":       "ServingRuntime",
		"spec": map[string]any{
			"containers": []any{
				map[string]any{
					"name":  "kserve-container",
					"image": nginxLatest,
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1alpha1ServingRuntime(servingRuntime, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestServingKserveIOV1alpha1ClusterStorageContainer(t *testing.T) {
	t.Parallel()
	clusterStorageContainer := map[string]any{
		"apiVersion": "serving.kserve.io/v1alpha1",
		"kind":       "ClusterStorageContainer",
		"spec": map[string]any{
			"container": map[string]any{
				"name":  "storage-initializer",
				"image": nginxLatest,
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1alpha1ClusterStorageContainer(clusterStorageContainer, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestServingKserveIOV1beta1InferenceService(t *testing.T) {
	t.Parallel()
	inferenceService := map[string]any{
		"apiVersion": "serving.kserve.io/v1beta1",
		"kind":       "InferenceService",
		"spec": map[string]any{
			"predictor": map[string]any{
				"containers": []any{
					map[string]any{
						"name":  "predictor",
						"image": nginxLatest,
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1beta1InferenceService(inferenceService, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestServingKserveIOV1beta1InferenceServiceWithAllComponents(t *testing.T) {
	t.Parallel()
	inferenceService := map[string]any{
		"apiVersion": "serving.kserve.io/v1beta1",
		"kind":       "InferenceService",
		"spec": map[string]any{
			"predictor": map[string]any{
				"containers": []any{
					map[string]any{
						"name":  "predictor",
						"image": nginxLatest,
					},
				},
				"initContainers": []any{
					map[string]any{
						"name":  "init-predictor",
						"image": "init:predictor",
					},
				},
			},
			"explainer": map[string]any{
				"containers": []any{
					map[string]any{
						"name":  "explainer",
						"image": busybox128,
					},
				},
			},
			"transformer": map[string]any{
				"containers": []any{
					map[string]any{
						"name":  "transformer",
						"image": "transformer:v1",
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1beta1InferenceService(inferenceService, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest:      {},
		"init:predictor": {},
		busybox128:       {},
		"transformer:v1": {},
	}, output)
}

func TestServingKserveIOV1beta1InferenceServicePredictorOnly(t *testing.T) {
	t.Parallel()
	// Test with only predictor spec (minimal InferenceService)
	inferenceService := map[string]any{
		"apiVersion": "serving.kserve.io/v1beta1",
		"kind":       "InferenceService",
		"spec": map[string]any{
			"predictor": map[string]any{
				"containers": []any{
					map[string]any{
						"name":  "sklearn",
						"image": "kserve/sklearnserver:v0.8.0",
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := ServingKserveIOV1beta1InferenceService(inferenceService, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		"kserve/sklearnserver:v0.8.0": {},
	}, output)
}
