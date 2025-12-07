package images

import "fmt"

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
