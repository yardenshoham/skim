package images

import "fmt"

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
