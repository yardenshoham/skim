package images

import (
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/goccy/go-yaml"
)

type UnknownGVKBehavior int

const (
	// UnknownGVKFail indicates that the extractor should fail when encountering an unknown GVK.
	UnknownGVKFail UnknownGVKBehavior = iota
	// UnknownGVKSkip indicates that the extractor should skip manifests with unknown GVKs.
	UnknownGVKSkip
	// TODO: add free-text option grepping "image: " and "imageName: " fields from unknown GVKs.
)

// Extractor extracts image references from Kubernetes manifests.
type Extractor struct {
	// UnknownGVKBehavior defines the behavior when encountering unknown GVKs.
	UnknownGVKBehavior UnknownGVKBehavior
	// Logger is used for logging messages. Make sure you initialize it or use [NewExtractor].
	Logger *slog.Logger
	// GVKMappings maps custom GVK strings to their corresponding extraction functions. You can add custom GVKs here.
	// The key is the GVK string in the format "apiVersion.kind", e.g. "apps/v1.Deployment".
	// The value is a function that takes a manifest and an output map, and extracts image references from the manifest.
	GVKMappings map[string]func(map[string]any, map[string]struct{}) error
}

// NewExtractor creates a new Extractor with the provided options.
func NewExtractor() *Extractor {
	e := &Extractor{
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	return e
}

// ExtractFromManifests extracts image references from a YAML stream placing them in the images map as keys.
func (e *Extractor) ExtractFromManifests(r io.Reader, images map[string]struct{}) error {
	decoder := yaml.NewDecoder(r, yaml.AllowDuplicateMapKey())
	for {
		var manifest map[string]any
		if err := decoder.Decode(&manifest); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode manifest: %w", err)
		}
		err := fromManifest(manifest, images, e.GVKMappings)
		if err != nil {
			var unknownGVKError *UnknownGVKError
			if e.UnknownGVKBehavior == UnknownGVKSkip && errors.As(err, &unknownGVKError) {
				e.Logger.Warn("Skipping unknown GVK", "group-version-kind", unknownGVKError.GVK, "manifest", unknownGVKError.Manifest)
				continue
			}
			return fmt.Errorf("failed to extract images from manifest: %w", err)
		}
	}
	return nil
}

// fromManifest extracts image references from a Kubernetes manifest placing them in the output map as keys.
func fromManifest(manifest map[string]any, output map[string]struct{}, gvkMappings map[string]func(map[string]any, map[string]struct{}) error) error {
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
	if gvkMappings != nil {
		if extractorFunc, found := gvkMappings[gvkString]; found {
			return extractorFunc(manifest, output)
		}
	}
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
