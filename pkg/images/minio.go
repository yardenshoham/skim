package images

// MinIOMinIOV2Tenant extracts images from a minio.min.io/v2.Tenant manifest placing them in the output map as keys.
func MinIOMinIOV2Tenant(tenant map[string]any, output map[string]struct{}) error {
	return specImage(tenant, output)
}
