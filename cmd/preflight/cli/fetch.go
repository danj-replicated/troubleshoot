package cli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/replicatedhq/troubleshoot/pkg/oci"
	"github.com/spf13/cobra"
)

func FetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Download the preflight file from URL",
		RunE: func(cmd *cobra.Command, args []string) error {

				content, err := oci.PullPreflightFromOCI(args[0])
				if err != nil {
					if err == oci.ErrNoRelease {
						return errors.Errorf("no release found for %s.\nCheck the oci:// uri for errors or contact the application vendor for support.", args[0])
					}

					return err
				}

				fmt.Println(string(content))

			return nil
		},
	}
	return cmd
}