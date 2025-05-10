package primitives

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
	funk "github.com/thoas/go-funk"

	utils "github.com/gallvp/glb/utils"
)

const (
	maxCharsPerLine = 120
)

func BaseWiseMatch(
	seq1, seq2 string,
) ([]bool, error) {
	if len(seq1) != len(seq2) {
		return nil, fmt.Errorf("sequences are not of equal length")
	}

	seq1Chars := strings.Split(seq1, "")
	seq2Chars := strings.Split(seq2, "")

	return funk.Map(funk.Zip(seq1Chars, seq2Chars), func(pair funk.Tuple) bool {
		return strings.EqualFold(pair.Element1.(string), pair.Element2.(string))
	}).([]bool), nil
}

func NumberOfMatches(seq1, seq2 string) (int, error) {
	matches, err := BaseWiseMatch(seq1, seq2)
	if err != nil {
		return 0, err
	}

	ints := funk.Map(matches, func(val bool) int {
		if val {
			return 1
		} else {
			return 0
		}
	}).([]int)

	return int(funk.Sum(ints)), nil
}

func NumberOfMismatches(seq1, seq2 string) (int, error) {

	len := len(seq1)

	numMatches, err := NumberOfMatches(seq1, seq2)

	if err != nil {
		return 0, err
	}

	return len - numMatches, nil

}

func PrintMismatch(seq1, seq2 string) error {
	matches, err := BaseWiseMatch(seq1, seq2)
	if err != nil {
		return err
	}

	seq1Chunks := utils.SplitByRunes(seq1, maxCharsPerLine)
	seq2Chunks := utils.SplitByRunes(seq2, maxCharsPerLine)

	// fmt.Println(seq1)
	matchesAsString := strings.Join(funk.Map(matches, func(match bool) string {
		if match {
			return "|"
		} else {
			return " "
		}
	}).([]string), "")

	matchesChunks := utils.SplitByRunes(matchesAsString, maxCharsPerLine)

	funk.ForEach(funk.Zip(funk.Zip(seq1Chunks, seq2Chunks), matchesChunks), func(pair funk.Tuple) {
		seq1Chunk := pair.Element1.(funk.Tuple).Element1.(string)
		seq2Chunk := pair.Element1.(funk.Tuple).Element2.(string)
		matchesChunk := pair.Element2.(string)
		fmt.Printf("%s\n%s\n%s\n\n", seq1Chunk, matchesChunk, seq2Chunk)
	})

	return nil

}

func printMismatchMainCmd(cmd *cobra.Command, args []string) {

	verbose, _ := cmd.Root().PersistentFlags().GetBool("verbose")
	noPrint, _ := cmd.Flags().GetBool("no-print")

	// Set up logger
	log.SetLevel(log.ErrorLevel)

	if verbose {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug logging enabled")
	}

	var mismatches int
	var err error
	mismatches, err = NumberOfMismatches(args[0], args[1])

	if err != nil {
		log.Fatalf("Error calculating mismatches: %v", err)
	}

	fmt.Printf("Number of mismatches: %d (%.2f%%)\n", mismatches, float64(mismatches)/float64(len(args[0]))*100.0)

	if !noPrint {
		err = PrintMismatch(args[0], args[1])

	}

	if err != nil {
		log.Fatalf("Error printing mismatch: %v", err)
	}
}

var MismatchCmd = &cobra.Command{
	Use:   "mismatch <seq1> <seq2>",
	Short: "Find the mismatches between two sequences",
	Run: func(cmd *cobra.Command, args []string) {
		printMismatchMainCmd(cmd, args)
	},
	Args:    cobra.ExactArgs(2),
	Example: "mismatch GGGCCGTTGGT GGACCGTTGAC",
}

func InitMismatchCmd() {
	MismatchCmd.Flags().BoolP("no-print", "P", false, "Do not print the mismatches")
}
