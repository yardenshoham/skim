package images

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKafkaStrimziIOV1Beta2Kafka(t *testing.T) {
	t.Parallel()
	kafka := map[string]any{
		"apiVersion": "kafka.strimzi.io/v1beta2",
		"kind":       "Kafka",
		"spec": map[string]any{
			"kafka": map[string]any{
				"image": nginxLatest,
			},
		},
	}
	output := make(map[string]struct{})
	err := KafkaStrimziIOV1Beta2Kafka(kafka, output)
	require.NoError(t, err)
	require.Equal(t, map[string]struct{}{
		nginxLatest: {},
	}, output)
}
