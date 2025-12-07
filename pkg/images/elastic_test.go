package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElasticsearchK8sElasticCoV1Elasticsearch(t *testing.T) {
	t.Parallel()
	elasticsearch := map[string]any{
		"apiVersion": "elasticsearch.k8s.elastic.co/v1",
		"kind":       "Elasticsearch",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := ElasticsearchK8sElasticCoV1Elasticsearch(elasticsearch, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}

func TestKibanaK8sElasticCoV1Kibana(t *testing.T) {
	t.Parallel()
	kibana := map[string]any{
		"apiVersion": "kibana.k8s.elastic.co/v1",
		"kind":       "Kibana",
		"spec": map[string]any{
			"image": nginxLatest,
		},
	}
	output := make(map[string]struct{})
	err := KibanaK8sElasticCoV1Kibana(kibana, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}
