package primitives

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	funk "github.com/thoas/go-funk"

	cobra "github.com/spf13/cobra"
)

func readSequenceFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 100_000)
	scanner.Buffer(buf, 100_000)
	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return builder.String(), nil
}

func computeGCskew(seq string) []int {

	skew := 0
	test := funk.Map([]rune(seq), func(base rune) int {

		if base == 'G' || base == 'g' {
			skew++
		}

		if base == 'C' || base == 'c' {
			skew--
		}
		return skew
	})

	return append([]int{0}, test.([]int)...)
}

func skewCmdMain(cmd *cobra.Command) {
	// Flags
	sequenceInput := cmd.Flag("sequence").Value.String()
	fileInput := cmd.Flag("file").Value.String()
	verbose, _ := cmd.Root().PersistentFlags().GetBool("verbose")

	// Set up logger
	log.SetLevel(log.ErrorLevel)

	if verbose {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug logging enabled")
	}

	// Validate flags
	if sequenceInput == "" && fileInput == "" {
		log.Fatal("You must provide either -s (sequence) or -f (file)")
	}

	if sequenceInput != "" && fileInput != "" {
		log.Fatal("You cannot provide both -s (sequence) and -f (file)")
	}

	// Read sequence
	var sequence string
	var err error
	if sequenceInput != "" {
		sequence = sequenceInput
		err = nil
	}

	if fileInput != "" {
		sequence, err = readSequenceFromFile(fileInput)
	}

	if err != nil {
		log.Fatalf("Failed to read sequence from file: %v", err)
	}

	log.Infof("Input sequence (length %d)", len(sequence))

	skews := computeGCskew(sequence)

	for _, val := range skews {
		fmt.Printf("%d\n", val)
	}
}

var SkewCmd = &cobra.Command{
	Use:   "skew [-s <sequence> | -f <file>]",
	Short: "Find the skew of a sequence",
	Run: func(cmd *cobra.Command, args []string) {
		skewCmdMain(cmd)
	},
}

func InitSkewCmd() {
	SkewCmd.Flags().StringP("sequence", "s", "", "DNA sequence string")
	SkewCmd.Flags().StringP("file", "f", "", "File path containing DNA sequence")
}
