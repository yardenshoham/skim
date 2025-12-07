package images

import "fmt"

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
