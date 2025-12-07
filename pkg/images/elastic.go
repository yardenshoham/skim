package images

// ElasticsearchK8sElasticCoV1Elasticsearch extracts images from an elasticsearch.k8s.elastic.co/v1.Elasticsearch manifest placing them in the output map as keys.
func ElasticsearchK8sElasticCoV1Elasticsearch(elasticsearch map[string]any, output map[string]struct{}) error {
	return specImage(elasticsearch, output)
}

// KibanaK8sElasticCoV1Kibana extracts images from a kibana.k8s.elastic.co/v1.Kibana manifest placing them in the output map as keys.
func KibanaK8sElasticCoV1Kibana(kibana map[string]any, output map[string]struct{}) error {
	return specImage(kibana, output)
}
