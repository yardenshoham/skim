package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListCmdStdin(t *testing.T) {
	t.Parallel()
	listCmd := newListCmd()
	file, err := os.Open("../testdata/deployment.yaml")
	require.NoError(t, err)
	defer file.Close()

	var stdout bytes.Buffer
	listCmd.SetIn(file)
	listCmd.SetOut(&stdout)
	listCmd.SetArgs([]string{"-"})
	err = listCmd.Execute()
	require.NoError(t, err)
	require.Contains(t, stdout.String(), "example.com/processor:1.2.3")
}

func TestListCmdMixedFilesAndStdin(t *testing.T) {
	t.Parallel()
	listCmd := newListCmd()
	file, err := os.Open("../testdata/pod.yaml")
	require.NoError(t, err)
	defer file.Close()

	var stdout bytes.Buffer
	listCmd.SetIn(file)
	listCmd.SetOut(&stdout)
	listCmd.SetArgs([]string{"../testdata/deployment.yaml", "-"})
	err = listCmd.Execute()
	require.NoError(t, err)
	// Should contain images from both the deployment.yaml file and pod.yaml from stdin
	require.Contains(t, stdout.String(), "example.com/processor:1.2.3") // from deployment.yaml
	require.Contains(t, stdout.String(), "nginx:1.21.0")                // from pod.yaml via stdin
	require.Contains(t, stdout.String(), "busybox:1.35")                // from pod.yaml via stdin
}
