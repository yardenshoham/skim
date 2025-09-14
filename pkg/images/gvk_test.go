package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const nginxLatest = "nginx:latest"
const busybox128 = "example.com/busybox:1.28"

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

func TestElasticsearchK8sElasticCoV1Elasticsearch(t *testing.T) {
	t.Parallel()
	elasticsearch := map[string]any{
		"apiVersion": "elasticsearch.k8s.elastic.co/v1",
		"kind":       "Elasticsearch",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := ElasticsearchK8sElasticCoV1Elasticsearch(elasticsearch, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestKafkaStrimziIOV1Beta2Kafka(t *testing.T) {
	t.Parallel()
	kafka := map[string]any{
		"apiVersion": "kafka.strimzi.io/v1beta2",
		"kind":       "Kafka",
		"spec": map[string]any{
			"kafka": map[string]any{
				"image": nginxLatest,
			},
		},
	}
	output := make(map[string]struct{})
	err := KafkaStrimziIOV1Beta2Kafka(kafka, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestKibanaK8sElasticCoV1Kibana(t *testing.T) {
	t.Parallel()
	kibana := map[string]any{
		"apiVersion": "kibana.k8s.elastic.co/v1",
		"kind":       "Kibana",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := KibanaK8sElasticCoV1Kibana(kibana, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

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
		},
	}
	output := make(map[string]struct{})
	err := MonitoringCoreosComV1Prometheus(prometheus, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}
