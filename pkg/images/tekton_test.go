package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTektonDevV1beta1Task(t *testing.T) {
	t.Parallel()
	task := map[string]any{
		"apiVersion": "tekton.dev/v1beta1",
		"kind":       "Task",
		"spec": map[string]any{
			"steps": []any{
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
	err := TektonDevV1beta1Task(task, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestTriggersTektonDevV1beta1EventListener(t *testing.T) {
	t.Parallel()
	eventListener := map[string]any{
		"apiVersion": "triggers.tekton.dev/v1beta1",
		"kind":       "EventListener",
		"spec": map[string]any{
			"resources": map[string]any{
				"kubernetesResource": map[string]any{
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
		},
	}
	output := make(map[string]struct{})
	err := TriggersTektonDevV1beta1EventListener(eventListener, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
		busybox128:  {},
	}, output)
}

func TestTriggersTektonDevV1beta1TriggerTemplate(t *testing.T) {
	t.Parallel()
	triggerTemplate := map[string]any{
		"apiVersion": "triggers.tekton.dev/v1beta1",
		"kind":       "TriggerTemplate",
		"metadata": map[string]any{
			"name": "video-process-tt",
		},
		"spec": map[string]any{
			"params": []any{
				map[string]any{
					"name": "video",
				},
				map[string]any{
					"name": "output",
				},
			},
			"resourcetemplates": []any{
				map[string]any{
					"apiVersion": "tekton.dev/v1beta1",
					"kind":       "TaskRun",
					"metadata": map[string]any{
						"generateName": "video-process-run-",
					},
					"spec": map[string]any{
						"podTemplate": map[string]any{
							"runtimeClassName": "nvidia",
							"imagePullSecrets": nil,
						},
						"serviceAccountName": "video-process",
						"tolerations": []any{
							map[string]any{
								"effect":   "NoSchedule",
								"key":      "nvidia.com/gpu",
								"operator": "Exists",
							},
						},
						"workspaces": []any{
							map[string]any{
								"name":     "workdir",
								"emptyDir": map[string]any{},
							},
							map[string]any{
								"name":     "results",
								"emptyDir": map[string]any{},
							},
							map[string]any{
								"name": "conf",
								"configMap": map[string]any{
									"name": "video-process-config",
								},
							},
						},
						"taskRef": map[string]any{
							"name": "video-process-tekton-gcp-task",
						},
						"params": []any{
							map[string]any{
								"name":  "video",
								"value": "$(tt.params.video)",
							},
							map[string]any{
								"name":  "output",
								"value": "$(tt.params.output)",
							},
						},
					},
				},
			},
		},
	}
	output := make(map[string]struct{})
	err := TriggersTektonDevV1beta1TriggerTemplate(triggerTemplate, output)
	require.NoError(t, err)
}
