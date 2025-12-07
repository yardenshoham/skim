package images

// MonitoringCoreosComV1Alertmanager extracts images from a monitoring.coreos.com/v1.Alertmanager manifest placing them in the output map as keys.
func MonitoringCoreosComV1Alertmanager(alertmanager map[string]any, output map[string]struct{}) error {
	return specImage(alertmanager, output)
}

// MonitoringCoreosComV1Prometheus extracts images from a monitoring.coreos.com/v1.Prometheus manifest placing them in the output map as keys.
func MonitoringCoreosComV1Prometheus(prometheus map[string]any, output map[string]struct{}) error {
	err := specImage(prometheus, output)
	if err != nil {
		return err
	}
	_ = v1PodTemplateSpec(prometheus, output) // optional so we ignore errors
	return nil
}
