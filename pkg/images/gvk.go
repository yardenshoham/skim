package images

import "fmt"

var imagelessGVKs = map[string]struct{}{
	"v1.Namespace":                                                   {},
	"storage.k8s.io/v1.StorageClass":                                 {},
	"apiextensions.k8s.io/v1.CustomResourceDefinition":               {},
	"v1.ServiceAccount":                                              {},
	"v1.Service":                                                     {},
	"networking.k8s.io/v1.Ingress":                                   {},
	"networking.k8s.io/v1.IngressClass":                              {},
	"autoscaling/v1.HorizontalPodAutoscaler":                         {},
	"v1.ConfigMap":                                                   {},
	"monitoring.coreos.com/v1.PodMonitor":                            {},
	"networking.k8s.io/v1.NetworkPolicy":                             {},
	"v1.Secret":                                                      {},
	"rbac.authorization.k8s.io/v1.ClusterRole":                       {},
	"rbac.authorization.k8s.io/v1.RoleBinding":                       {},
	"monitoring.coreos.com/v1.PrometheusRule":                        {},
	"policy/v1.PodDisruptionBudget":                                  {},
	"monitoring.coreos.com/v1.ServiceMonitor":                        {},
	"autoscaling/v2.HorizontalPodAutoscaler":                         {},
	"v1.PersistentVolumeClaim":                                       {},
	"rbac.authorization.k8s.io/v1.ClusterRoleBinding":                {},
	"rbac.authorization.k8s.io/v1.Role":                              {},
	"kafka.strimzi.io/v1beta2.KafkaNodePool":                         {},
	"tekton.dev/v1beta1.TaskRun":                                     {},
	"triggers.tekton.dev/v1beta1.TriggerBinding":                     {},
	"triggers.tekton.dev/v1beta1.ClusterTriggerBinding":              {},
	"monitoring.coreos.com/v1alpha1.AlertmanagerConfig":              {},
	"triggers.tekton.dev/v1alpha1.ClusterInterceptor":                {},
	"admissionregistration.k8s.io/v1.MutatingWebhookConfiguration":   {},
	"admissionregistration.k8s.io/v1.ValidatingWebhookConfiguration": {},
	"kyverno.io/v1.ClusterPolicy":                                    {},
	"kyverno.io/v1.Policy":                                           {},
	"cert-manager.io/v1.Certificate":                                 {},
	"cert-manager.io/v1.ClusterIssuer":                               {},
	"cert-manager.io/v1.Issuer":                                      {},
}

// UnknownGVKError is returned when a manifest has an unexpected GVK.
type UnknownGVKError struct {
	GVK      string
	Manifest map[string]any
}

func (e *UnknownGVKError) Error() string {
	return fmt.Sprintf("failed to detect Group Version Kind: %s, manifest: %+v", e.GVK, e.Manifest)
}

func v1PodSpec(podSpec map[string]any, output map[string]struct{}) error {
	containers, ok := podSpec["containers"]
	if !ok {
		return fmt.Errorf("failed to find containers field in podSpec: %+v", podSpec)
	}
	containerList, ok := containers.([]any)
	if !ok {
		return nil
	}
	for _, container := range containerList {
		containerMap, ok := container.(map[string]any)
		if !ok {
			return fmt.Errorf("failed to convert container to map, container: %+v", container)
		}
		image, ok := containerMap["image"]
		if !ok {
			return nil
		}
		imageStr, ok := image.(string)
		if !ok {
			return fmt.Errorf("failed to convert image to string, container: %+v", containerMap)
		}
		output[imageStr] = struct{}{}
	}
	initContainers, ok := podSpec["initContainers"]
	if !ok {
		return nil
	}
	initContainerList, ok := initContainers.([]any)
	if !ok {
		return nil
	}
	for _, initContainer := range initContainerList {
		initContainerMap, ok := initContainer.(map[string]any)
		if !ok {
			return fmt.Errorf("failed to convert initContainer to map, initContainer: %+v", initContainer)
		}
		image, ok := initContainerMap["image"]
		if !ok {
			return nil
		}
		imageStr, ok := image.(string)
		if !ok {
			return fmt.Errorf("failed to convert image to string, initContainer: %+v", initContainerMap)
		}
		output[imageStr] = struct{}{}
	}
	return nil
}

func v1PodTemplateSpec(podTemplate map[string]any, output map[string]struct{}) error {
	spec, ok := podTemplate["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in podTemplate: %+v", podTemplate)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, podTemplate: %+v", podTemplate)
	}
	return v1PodSpec(specMap, output)
}

func v1PodSpecTemplateSpec(manifest map[string]any, output map[string]struct{}) error {
	spec, ok := manifest["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in manifest: %+v", manifest)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, manifest: %+v", manifest)
	}
	template, ok := specMap["template"]
	if !ok {
		return fmt.Errorf("failed to find template field in manifest spec: %+v", specMap)
	}
	templateMap, ok := template.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert template to map, manifest spec: %+v", specMap)
	}
	return v1PodTemplateSpec(templateMap, output)
}

// V1Pod extracts images from a v1.Pod manifest placing them in the output map as keys.
func V1Pod(pod map[string]any, output map[string]struct{}) error {
	return v1PodTemplateSpec(pod, output)
}

// AppsV1Deployment extracts images from an apps/v1.Deployment manifest placing them in the output map as keys.
func AppsV1Deployment(deployment map[string]any, output map[string]struct{}) error {
	return v1PodSpecTemplateSpec(deployment, output)
}

// AppsV1StatefulSet extracts images from an apps/v1.StatefulSet manifest placing them in the output map as keys.
func AppsV1StatefulSet(statefulSet map[string]any, output map[string]struct{}) error {
	return v1PodSpecTemplateSpec(statefulSet, output)
}

// AppsV1DaemonSet extracts images from an apps/v1.DaemonSet manifest placing them in the output map as keys.
func AppsV1DaemonSet(daemonSet map[string]any, output map[string]struct{}) error {
	return v1PodSpecTemplateSpec(daemonSet, output)
}

// BatchV1Job extracts images from a batch/v1.Job manifest placing them in the output map as keys.
func BatchV1Job(job map[string]any, output map[string]struct{}) error {
	return v1PodSpecTemplateSpec(job, output)
}

// BatchV1CronJob extracts images from a batch/v1.CronJob manifest placing them in the output map as keys.
func BatchV1CronJob(cronJob map[string]any, output map[string]struct{}) error {
	spec, ok := cronJob["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in cronJob: %+v", cronJob)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, cronJob: %+v", cronJob)
	}
	jobTemplate, ok := specMap["jobTemplate"]
	if !ok {
		return fmt.Errorf("failed to find jobTemplate field in cronJob spec: %+v", specMap)
	}
	jobTemplateMap, ok := jobTemplate.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert jobTemplate to map, cronJob spec: %+v", specMap)
	}
	return v1PodSpecTemplateSpec(jobTemplateMap, output)
}

// PostgresqlCNPGIOV1Cluster extracts images from a postgresql.cnpg.io/v1.Cluster manifest placing them in the output map as keys.
func PostgresqlCNPGIOV1Cluster(cluster map[string]any, output map[string]struct{}) error {
	spec, ok := cluster["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in cluster: %+v", cluster)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, cluster: %+v", cluster)
	}
	imageName, ok := specMap["imageName"]
	if !ok {
		return nil
	}
	imageNameStr, ok := imageName.(string)
	if !ok {
		return fmt.Errorf("failed to convert imageName to string, cluster spec: %+v", specMap)
	}
	output[imageNameStr] = struct{}{}
	return nil
}

func specImage(manifest map[string]any, output map[string]struct{}) error {
	spec, ok := manifest["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in manifest: %+v", manifest)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, manifest: %+v", manifest)
	}
	image, ok := specMap["image"]
	if !ok {
		return nil
	}
	imageStr, ok := image.(string)
	if !ok {
		return fmt.Errorf("failed to convert image to string, manifest spec: %+v", specMap)
	}
	output[imageStr] = struct{}{}
	return nil
}

// ElasticsearchK8sElasticCoV1Elasticsearch extracts images from an elasticsearch.k8s.elastic.co/v1.Elasticsearch manifest placing them in the output map as keys.
func ElasticsearchK8sElasticCoV1Elasticsearch(elasticsearch map[string]any, output map[string]struct{}) error {
	return specImage(elasticsearch, output)
}

// KafkaStrimziIOV1Beta2Kafka extracts images from a kafka.strimzi.io/v1beta2.Kafka manifest placing them in the output map as keys.
func KafkaStrimziIOV1Beta2Kafka(kafka map[string]any, output map[string]struct{}) error {
	spec, ok := kafka["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in kafka: %+v", kafka)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, kafka: %+v", kafka)
	}
	kafkaSpec, ok := specMap["kafka"]
	if !ok {
		return nil
	}
	kafkaSpecMap, ok := kafkaSpec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert kafka spec to map, kafka spec: %+v", specMap)
	}
	image, ok := kafkaSpecMap["image"]
	if !ok {
		return nil
	}
	imageStr, ok := image.(string)
	if !ok {
		return fmt.Errorf("failed to convert image to string, kafka spec: %+v", kafkaSpecMap)
	}
	output[imageStr] = struct{}{}
	return nil
}

// KibanaK8sElasticCoV1Kibana extracts images from a kibana.k8s.elastic.co/v1.Kibana manifest placing them in the output map as keys.
func KibanaK8sElasticCoV1Kibana(kibana map[string]any, output map[string]struct{}) error {
	return specImage(kibana, output)
}

// TektonDevV1beta1Task extracts images from a tekton.dev/v1beta1.Task manifest placing them in the output map as keys.
func TektonDevV1beta1Task(task map[string]any, output map[string]struct{}) error {
	spec, ok := task["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in task: %+v", task)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, task: %+v", task)
	}
	steps, ok := specMap["steps"]
	if !ok {
		return nil
	}
	stepsList, ok := steps.([]any)
	if !ok {
		return fmt.Errorf("failed to convert steps to array, task spec: %+v", specMap)
	}
	for _, step := range stepsList {
		stepMap, ok := step.(map[string]any)
		if !ok {
			return fmt.Errorf("failed to convert step to map, step: %+v", step)
		}
		image, ok := stepMap["image"]
		if !ok {
			continue
		}
		imageStr, ok := image.(string)
		if !ok {
			return fmt.Errorf("failed to convert image to string, step: %+v", stepMap)
		}
		output[imageStr] = struct{}{}
	}
	return nil
}

// MinIOMinIOV2Tenant extracts images from a minio.min.io/v2.Tenant manifest placing them in the output map as keys.
func MinIOMinIOV2Tenant(tenant map[string]any, output map[string]struct{}) error {
	return specImage(tenant, output)
}

// TriggersTektonDevV1beta1EventListener extracts images from a triggers.tekton.dev/v1beta1.EventListener manifest placing them in the output map as keys.
func TriggersTektonDevV1beta1EventListener(eventListener map[string]any, output map[string]struct{}) error {
	spec, ok := eventListener["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in eventListener: %+v", eventListener)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, eventListener: %+v", eventListener)
	}
	resources, ok := specMap["resources"]
	if !ok {
		return nil
	}
	resourcesMap, ok := resources.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert resources to map, eventListener spec: %+v", specMap)
	}
	kubernetesResource, ok := resourcesMap["kubernetesResource"]
	if !ok {
		return nil
	}
	kubernetesResourceMap, ok := kubernetesResource.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert kubernetesResource to map, resources: %+v", resourcesMap)
	}
	return v1PodSpecTemplateSpec(kubernetesResourceMap, output)
}

// TriggersTektonDevV1beta1TriggerTemplate extracts images from a triggers.tekton.dev/v1beta1.TriggerTemplate manifest placing them in the output map as keys.
func TriggersTektonDevV1beta1TriggerTemplate(triggerTemplate map[string]any, output map[string]struct{}) error {
	spec, ok := triggerTemplate["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in triggerTemplate: %+v", triggerTemplate)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, triggerTemplate: %+v", triggerTemplate)
	}
	resourceTemplates, ok := specMap["resourcetemplates"]
	if !ok {
		return nil
	}
	resourceTemplatesList, ok := resourceTemplates.([]any)
	if !ok {
		return fmt.Errorf("failed to convert resourcetemplates to array, triggerTemplate spec: %+v", specMap)
	}
	for _, resourceTemplate := range resourceTemplatesList {
		resourceTemplateMap, ok := resourceTemplate.(map[string]any)
		if !ok {
			return fmt.Errorf("failed to convert resourceTemplate to map, resourceTemplate: %+v", resourceTemplate)
		}
		err := fromManifest(resourceTemplateMap, output, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// MonitoringCoreosComV1Alertmanager extracts images from a monitoring.coreos.com/v1.Alertmanager manifest placing them in the output map as keys.
func MonitoringCoreosComV1Alertmanager(alertmanager map[string]any, output map[string]struct{}) error {
	return specImage(alertmanager, output)
}

// MonitoringCoreosComV1Prometheus extracts images from a monitoring.coreos.com/v1.Prometheus manifest placing them in the output map as keys.
func MonitoringCoreosComV1Prometheus(prometheus map[string]any, output map[string]struct{}) error {
	err := specImage(prometheus, output)
	if err != nil {
		return err
	}
	_ = v1PodTemplateSpec(prometheus, output) // optional so we ignore errors
	return nil
}
