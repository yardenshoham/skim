package images

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/goccy/go-yaml"
)

type UnknownGVKBehavior int

const (
	// UnknownGVKFail indicates that the extractor should fail when encountering an unknown GVK.
	UnknownGVKFail UnknownGVKBehavior = iota
	// UnknownGVKSkip indicates that the extractor should skip manifests with unknown GVKs.
	UnknownGVKSkip
	// UnknownGVKFreeText indicates that the extractor should attempt to extract image references from the entire input (all manifests) as free text.
	UnknownGVKFreeText
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
func (e *Extractor) ExtractFromManifests(ctx context.Context, r io.Reader, images map[string]struct{}) error {
	var bufferForYAML bytes.Buffer
	var entireInput string
	if e.UnknownGVKBehavior == UnknownGVKFreeText {
		// we should buffer the input in case we need to parse it as free text
		_, err := io.Copy(&bufferForYAML, r)
		if err != nil {
			return fmt.Errorf("failed to buffer input: %w", err)
		}
		entireInput = bufferForYAML.String()
		r = bytes.NewBufferString(entireInput)
	}
	decoder := yaml.NewDecoder(r, yaml.AllowDuplicateMapKey())
	for {
		var manifest map[string]any
		if err := decoder.DecodeContext(ctx, &manifest); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode manifest: %w", err)
		}
		err := fromManifest(manifest, images, e.GVKMappings)
		if err != nil {
			if unknownGVKError, ok := errors.AsType[*UnknownGVKError](err); ok {
				switch e.UnknownGVKBehavior {
				case UnknownGVKFail:
					return fmt.Errorf("failed to extract images from manifest: %w", err)
				case UnknownGVKSkip:
					e.Logger.WarnContext(ctx, "Skipping unknown GVK", "group-version-kind", unknownGVKError.GVK, "manifest", unknownGVKError.Manifest)
					continue
				case UnknownGVKFreeText:
					e.Logger.WarnContext(ctx, "Unknown GVK, extracting images as free text from the input", "group-version-kind", unknownGVKError.GVK, "manifest", unknownGVKError.Manifest)
					err := extractImagesFromFreeText(entireInput, images)
					if err != nil {
						return fmt.Errorf("failed to extract images from free text: %w", err)
					}
					continue
				default:
					panic("unhandled UnknownGVKBehavior")
				}
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
	case "serving.kserve.io/v1alpha1.ClusterServingRuntime":
		return ServingKserveIOV1alpha1ClusterServingRuntime(manifest, output)
	case "serving.kserve.io/v1alpha1.ServingRuntime":
		return ServingKserveIOV1alpha1ServingRuntime(manifest, output)
	case "serving.kserve.io/v1alpha1.ClusterStorageContainer":
		return ServingKserveIOV1alpha1ClusterStorageContainer(manifest, output)
	case "serving.kserve.io/v1beta1.InferenceService":
		return ServingKserveIOV1beta1InferenceService(manifest, output)
	}
	return &UnknownGVKError{
		GVK:      gvkString,
		Manifest: manifest,
	}
}

var imagePrefixes = []string{
	"image: ",
	"imageName: ",
}

// extractImagesFromFreeText finds lines of the form "image: <image>" or "imageName: <image>" in the input string and adds <image> to the output map.
func extractImagesFromFreeText(manifest string, output map[string]struct{}) error {
	lines := strings.SplitSeq(manifest, "\n")
	for line := range lines {
		for _, prefix := range imagePrefixes {
			if _, after, ok := strings.Cut(line, prefix); ok {
				// Extract everything after the prefix
				imageName := strings.TrimSpace(after)
				// Remove quotes if present
				imageName = strings.Trim(imageName, `"'`)
				if imageName != "" {
					output[imageName] = struct{}{}
				}
				break // Found a match, no need to check other prefixes
			}
		}
	}
	return nil
}
