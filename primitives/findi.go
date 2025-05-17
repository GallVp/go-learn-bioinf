package primitives

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
	funk "github.com/thoas/go-funk"
	"golang.org/x/exp/slices"
)

func findIndicesOfMin(numbers []int) []int {

	var indices []int

	minVal := slices.Min(numbers)
	for i, v := range numbers {
		if v == minVal {
			indices = append(indices, i)
		}
	}
	return indices
}

func findICmdMain(cmd *cobra.Command, args []string) {

	verbose, _ := cmd.Root().PersistentFlags().GetBool("verbose")

	// Set up logger
	log.SetLevel(log.ErrorLevel)

	if verbose {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug logging enabled")
	}

	var numbersString string
	var err error = nil

	if len(args) > 0 {
		log.Info("Input numbers are ", args)
		numbersString = args[0]
	}

	if len(args) == 0 {
		log.Info("Reading from stdin")
		var data []byte
		data, err = io.ReadAll(os.Stdin)
		numbersString = string(data)
	}

	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
	}

	re := regexp.MustCompile(`\s+`)
	numbersStrArray := re.Split(strings.TrimSpace(string(numbersString)), -1)

	numbers := funk.Map(numbersStrArray, func(numStr string) int {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		return num
	}).([]int)

	log.Info("Input numbers are ", numbers)

	indices := findIndicesOfMin(numbers)

	log.Info("Indices are ", indices)

	outputStr := strings.Join(funk.Map(indices, func(i int) string {
		return strconv.Itoa(i)
	}).([]string), " ")

	fmt.Fprintln(cmd.OutOrStdout(), outputStr)
}

func NewFindICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "findi [<integers as a string> | <stdin>]",
		Short: "Find the indices of the minimum value in a list of integers",
		Run: func(cmd *cobra.Command, args []string) {
			findICmdMain(cmd, args)
		},
		Args:    cobra.MaximumNArgs(1),
		Example: "findi \"4 -2 3\"",
	}
}
