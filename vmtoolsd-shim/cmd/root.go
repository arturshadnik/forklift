package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/konveyor/forklift-controller/vmtoolsd-shim/pkg/utils"
)

var (
	rootCmd  *cobra.Command
	ovfMount = filepath.Join("mnt", "ovfenv")
	ovfXml   = filepath.Join(ovfMount, "ovfenv.xml")
)

func init() {
	initRootCmd()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initRootCmd() *cobra.Command {
	var cmd string

	rootCmd = &cobra.Command{
		Use:          "vmtoolsd",
		Short:        "VMware Tools Shim",
		Long:         "VMware Tools Shim",
		SilenceUsage: false,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			// create ovfenv ConfigMap mount if not found
			if _, err := os.Stat(ovfXml); os.IsNotExist(err) {
				if err := os.MkdirAll(ovfMount, os.ModePerm); err != nil {
					return err
				}
				params := []string{
					"/dev/disk/by-id/ata-QEMU_HARDDISK_ovfenv", ovfMount,
				}
				cmd := exec.Command("mount", params...) //#nosec G204
				if _, stderr, err := utils.Execute(true, cmd); err != nil {
					return errors.Wrap(err, stderr)
				}
			}

			// print the OVF properties
			data, err := os.ReadFile(ovfXml) //#nosec
			if err != nil {
				return err
			}
			fmt.Print(string(data))
			return nil
		},
	}

	flags := rootCmd.Flags()
	flags.StringVar(&cmd, "cmd", "", "Command")

	return rootCmd
}
