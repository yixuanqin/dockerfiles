package cmd

import (
	"errors"

	"github.com/laincloud/dockerfiles/src/core"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push docker images",
	RunE:  push,
}

func init() {
	pushCmd.Flags().StringVar(&commit1, "commit1", "origin/master", "previous commit")
	pushCmd.Flags().StringVar(&commit2, "commit2", "HEAD", "current commit")
	pushCmd.Flags().StringVarP(&registryHost, "registry-host", "r", "", "the registry host who serves this image")
	pushCmd.Flags().StringVarP(&organization, "organization", "o", "laincloud", "the organization build this image")
	rootCmd.AddCommand(pushCmd)
}

func push(cmd *cobra.Command, args []string) error {
	if commit1 == "" {
		return errors.New("--commit1 is required")
	}

	if commit2 == "" {
		return errors.New("--commit2 is required")
	}

	if organization == "" {
		return errors.New("--organization is required")
	}

	return core.Make(core.Args{
		Command:      core.Push,
		Commit1:      commit1,
		Commit2:      commit2,
		Organization: organization,
		RegistryHost: registryHost,
	})
}
