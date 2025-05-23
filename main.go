package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/gallvp/glb/primitives"

	cobra "github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{Use: "glb"}

	rootCmd.PersistentFlags().BoolP("verbose", "V", false, "Enable verbose logging")

	rootCmd.AddCommand(primitives.NewSkewCmd())
	rootCmd.AddCommand(primitives.NewFindICmd())
	rootCmd.AddCommand(primitives.NewMismatchCmd())

	log.SetLevel(log.ErrorLevel)

	if err := rootCmd.Execute(); err != nil {
		log.Errorf("Error executing command: %v", err)
	}
}
