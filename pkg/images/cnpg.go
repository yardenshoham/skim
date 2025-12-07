package images

import "fmt"

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
