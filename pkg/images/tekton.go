package images

import "fmt"

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
