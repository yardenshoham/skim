package images

import (
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
)

// FromManifests extracts image references from a YAML stream placing them in the images map as keys.
func FromManifests(r io.Reader, images map[string]struct{}) error {
	decoder := yaml.NewDecoder(r, yaml.AllowDuplicateMapKey())
	for {
		var manifest map[string]any
		if err := decoder.Decode(&manifest); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode manifest: %w", err)
		}
		err := FromManifest(manifest, images)
		if err != nil {
			return fmt.Errorf("failed to extract images from manifest: %w", err)
		}
	}
	return nil
}

// FromManifest extracts image references from a Kubernetes manifest placing them in the output map as keys.
func FromManifest(manifest map[string]any, output map[string]struct{}) error {
	apiVersion, ok := manifest["apiVersion"]
	if !ok {
		return fmt.Errorf("failed to find apiVersion field, manifest: %+v", manifest)
	}
	apiVersionStr, ok := apiVersion.(string)
	if !ok {
		return fmt.Errorf("failed to convert apiVersion to string, manifest: %+v", manifest)
	}
	kind, ok := manifest["kind"]
	if !ok {
		return fmt.Errorf("failed to find kind field, manifest: %+v", manifest)
	}
	kindStr, ok := kind.(string)
	if !ok {
		return fmt.Errorf("failed to convert kind to string, manifest: %+v", manifest)
	}
	gvkString := fmt.Sprintf("%s.%s", apiVersionStr, kindStr)
	if _, ok := imagelessGVKs[gvkString]; ok {
		return nil
	}
	switch gvkString {
	case "v1.Pod":
		return V1Pod(manifest, output)
	case "apps/v1.Deployment":
		return AppsV1Deployment(manifest, output)
	case "apps/v1.StatefulSet":
		return AppsV1StatefulSet(manifest, output)
	case "apps/v1.DaemonSet":
		return AppsV1DaemonSet(manifest, output)
	case "batch/v1.Job":
		return BatchV1Job(manifest, output)
	case "batch/v1.CronJob":
		return BatchV1CronJob(manifest, output)
	case "postgresql.cnpg.io/v1.Cluster":
		return PostgresqlCNPGIOV1Cluster(manifest, output)
	case "elasticsearch.k8s.elastic.co/v1.Elasticsearch":
		return ElasticsearchK8sElasticCoV1Elasticsearch(manifest, output)
	case "kafka.strimzi.io/v1beta2.Kafka":
		return KafkaStrimziIOV1Beta2Kafka(manifest, output)
	case "kibana.k8s.elastic.co/v1.Kibana":
		return KibanaK8sElasticCoV1Kibana(manifest, output)
	case "tekton.dev/v1beta1.Task":
		return TektonDevV1beta1Task(manifest, output)
	case "minio.min.io/v2.Tenant":
		return MinIOMinIOV2Tenant(manifest, output)
	case "triggers.tekton.dev/v1beta1.EventListener":
		return TriggersTektonDevV1beta1EventListener(manifest, output)
	case "triggers.tekton.dev/v1beta1.TriggerTemplate":
		return TriggersTektonDevV1beta1TriggerTemplate(manifest, output)
	case "monitoring.coreos.com/v1.Alertmanager":
		return MonitoringCoreosComV1Alertmanager(manifest, output)
	case "monitoring.coreos.com/v1.Prometheus":
		return MonitoringCoreosComV1Prometheus(manifest, output)
	}
	return &UnknownGVKError{
		GVK:      gvkString,
		Manifest: manifest,
	}
}
