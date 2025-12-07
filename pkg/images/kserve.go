package images

import "fmt"

// ServingKserveIOV1alpha1ClusterServingRuntime extracts images from a serving.kserve.io/v1alpha1.ClusterServingRuntime manifest placing them in the output map as keys.
func ServingKserveIOV1alpha1ClusterServingRuntime(clusterServingRuntime map[string]any, output map[string]struct{}) error {
	return servingRuntimeSpec(clusterServingRuntime, output)
}

// ServingKserveIOV1alpha1ServingRuntime extracts images from a serving.kserve.io/v1alpha1.ServingRuntime manifest placing them in the output map as keys.
func ServingKserveIOV1alpha1ServingRuntime(servingRuntime map[string]any, output map[string]struct{}) error {
	return servingRuntimeSpec(servingRuntime, output)
}

// ServingKserveIOV1alpha1ClusterStorageContainer extracts images from a serving.kserve.io/v1alpha1.ClusterStorageContainer manifest placing them in the output map as keys.
func ServingKserveIOV1alpha1ClusterStorageContainer(clusterStorageContainer map[string]any, output map[string]struct{}) error {
	spec, ok := clusterStorageContainer["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in clusterStorageContainer: %+v", clusterStorageContainer)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, clusterStorageContainer: %+v", clusterStorageContainer)
	}
	container, ok := specMap["container"]
	if !ok {
		return nil
	}
	containerMap, ok := container.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert container to map, clusterStorageContainer spec: %+v", specMap)
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
	return nil
}

// ServingKserveIOV1beta1InferenceService extracts images from a serving.kserve.io/v1beta1.InferenceService manifest placing them in the output map as keys.
func ServingKserveIOV1beta1InferenceService(inferenceService map[string]any, output map[string]struct{}) error {
	spec, ok := inferenceService["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in inferenceService: %+v", inferenceService)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, inferenceService: %+v", inferenceService)
	}

	// Extract from predictor
	if predictor, ok := specMap["predictor"]; ok {
		if predictorMap, ok := predictor.(map[string]any); ok {
			if err := inferenceServiceComponentSpec(predictorMap, output); err != nil {
				return err
			}
		}
	}

	// Extract from explainer
	if explainer, ok := specMap["explainer"]; ok {
		if explainerMap, ok := explainer.(map[string]any); ok {
			if err := inferenceServiceComponentSpec(explainerMap, output); err != nil {
				return err
			}
		}
	}

	// Extract from transformer
	if transformer, ok := specMap["transformer"]; ok {
		if transformerMap, ok := transformer.(map[string]any); ok {
			if err := inferenceServiceComponentSpec(transformerMap, output); err != nil {
				return err
			}
		}
	}

	return nil
}

// servingRuntimeSpec extracts images from a ServingRuntimeSpec (used by ClusterServingRuntime and ServingRuntime).
func servingRuntimeSpec(manifest map[string]any, output map[string]struct{}) error {
	spec, ok := manifest["spec"]
	if !ok {
		return fmt.Errorf("failed to find spec field in manifest: %+v", manifest)
	}
	specMap, ok := spec.(map[string]any)
	if !ok {
		return fmt.Errorf("failed to convert spec to map, manifest: %+v", manifest)
	}

	// ServingRuntimeSpec embeds containers directly in spec
	return v1PodSpec(specMap, output)
}

// inferenceServiceComponentSpec extracts images from a PredictorSpec/ExplainerSpec/TransformerSpec which embed PodSpec.
func inferenceServiceComponentSpec(componentSpec map[string]any, output map[string]struct{}) error {
	// The component spec embeds PodSpec fields directly (containers, initContainers)
	return v1PodSpec(componentSpec, output)
}
