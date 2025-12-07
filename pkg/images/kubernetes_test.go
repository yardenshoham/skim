package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestV1Pod(t *testing.T) {
	t.Parallel()
	pod := map[string]any{
		"apiVersion": "v1",
		"kind":       "Pod",
		"spec": map[string]any{
			"initContainers": []any{
				map[string]any{
					"name":  "ubuntu",
					"image": "ubuntu:latest",
				},
			},
			"containers": []any{
				map[string]any{
					"name":  "nginx",
					"image": nginxLatest,
				},
				map[string]any{
					"name":  "busybox",
					"image": busybox128,
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := V1Pod(pod, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest:     {},
		busybox128:      {},
		"ubuntu:latest": {},
	}, output)
}

func TestAppsV1Deployment(t *testing.T) {
	t.Parallel()
	deployment := map[string]any{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"spec": map[string]any{
			"template": map[string]any{
				"spec": map[string]any{
					"containers": []any{
						map[string]any{
							"name":  "nginx",
							"image": nginxLatest,
						},
						map[string]any{
							"name":  "busybox",
							"image": busybox128,
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := AppsV1Deployment(deployment, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestAppsV1StatefulSet(t *testing.T) {
	t.Parallel()
	statefulSet := map[string]any{
		"apiVersion": "apps/v1",
		"kind":       "StatefulSet",
		"spec": map[string]any{
			"template": map[string]any{
				"spec": map[string]any{
					"containers": []any{
						map[string]any{
							"name":  "nginx",
							"image": nginxLatest,
						},
						map[string]any{
							"name":  "busybox",
							"image": busybox128,
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := AppsV1StatefulSet(statefulSet, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestAppsV1DaemonSet(t *testing.T) {
	t.Parallel()
	daemonSet := map[string]any{
		"apiVersion": "apps/v1",
		"kind":       "DaemonSet",
		"spec": map[string]any{
			"template": map[string]any{
				"spec": map[string]any{
					"containers": []any{
						map[string]any{
							"name":  "nginx",
							"image": nginxLatest,
						},
						map[string]any{
							"name":  "busybox",
							"image": busybox128,
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := AppsV1DaemonSet(daemonSet, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestBatchV1Job(t *testing.T) {
	t.Parallel()
	job := map[string]any{
		"apiVersion": "batch/v1",
		"kind":       "Job",
		"spec": map[string]any{
			"template": map[string]any{
				"spec": map[string]any{
					"containers": []any{
						map[string]any{
							"name":  "nginx",
							"image": nginxLatest,
						},
						map[string]any{
							"name":  "busybox",
							"image": busybox128,
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := BatchV1Job(job, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestBatchV1CronJob(t *testing.T) {
	t.Parallel()
	cronJob := map[string]any{
		"apiVersion": "batch/v1",
		"kind":       "CronJob",
		"spec": map[string]any{
			"jobTemplate": map[string]any{
				"spec": map[string]any{
					"template": map[string]any{
						"spec": map[string]any{
							"containers": []any{
								map[string]any{
									"name":  "nginx",
									"image": nginxLatest,
								},
								map[string]any{
									"name":  "busybox",
									"image": busybox128,
								},
							},
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := BatchV1CronJob(cronJob, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}
