package cmd

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yardenshoham/skim/pkg/images"
)

func newListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:     "list PATH [PATH...]",
		Short:   "List container images from Kubernetes resources",
		Example: "skim list path/to/k8s-manifest.yaml",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.New(slog.NewTextHandler(cmd.ErrOrStderr(), nil))
			outputStream := cmd.OutOrStdout()
			imagesOutput := make(map[string]struct{})

			// Process each argument - can be files or stdin (-)
			for _, arg := range args {
				if arg == "-" {
					// Process stdin
					logger.Info("Processing stdin")
					inputStream := cmd.InOrStdin()
					err := images.FromManifests(inputStream, imagesOutput)
					if err != nil {
						return fmt.Errorf("failed to extract images from stdin: %w", err)
					}
					continue
				}

				// Process file path(s) - could be files or directories
				err := filepath.WalkDir(arg, func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						return err
					}
					if d.Type().IsRegular() {
						logger.Info("Processing file", "path", path)
						file, err := os.Open(path)
						if err != nil {
							return fmt.Errorf("failed to open file %s: %w", path, err)
						}
						defer file.Close()
						err = images.FromManifests(file, imagesOutput)
						if err != nil {
							return fmt.Errorf("failed to extract images from file %s: %w", path, err)
						}
					}
					return nil
				})
				if err != nil {
					return fmt.Errorf("failed to walk path %s: %w", arg, err)
				}
			}
			outputSlice := make([]string, 0, len(imagesOutput))
			for image := range imagesOutput {
				outputSlice = append(outputSlice, image)
			}
			if len(outputSlice) == 0 {
				logger.Warn("No images found")
				return nil
			}
			slices.Sort(outputSlice)
			resultList := strings.Join(outputSlice, "\n")
			_, err := fmt.Fprintln(outputStream, resultList)
			if err != nil {
				return fmt.Errorf("failed to write output: %w", err)
			}
			return nil
		},
	}
	return listCmd
}
