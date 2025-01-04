package esb_fireplace

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/pflag"
)

type (
	AoCPart       int
	AoCSolutionFn func(string, []string) (interface{}, error)
)

const (
	AoCPartNone AoCPart = 0
	AoCPart1    AoCPart = 1
	AoCPart2    AoCPart = 2
)

func parseArgs(cmd_args []string) (AoCPart, []string, error) {
	parser := pflag.NewFlagSet("fireplace", pflag.ExitOnError)

	part := parser.IntP("part", "p", 0, "Run solution part 1 or part 2")
	var args []string
	parser.StringSliceVarP(
		&args,
		"args",
		"a",
		[]string{},
		"Additional arguments for running the solutions",
	)

	parser.Usage = func() {
		fmt.Fprintf(os.Stderr, "Elf Script Brigade Go solution runner\n")
		parser.PrintDefaults()
	}

	err := parser.Parse(cmd_args)
	if err != nil {
		return AoCPartNone, []string{}, errors.New("Error parsing arguments")
	}

	remainingArgs := parser.Args()
	args = append(args, remainingArgs...)

	if *part != 1 && *part != 2 {
		return AoCPartNone, []string{}, errors.New(
			"Invalid part, please use 1 or 2 as argument for --part flag.",
		)
	}
	return AoCPart(*part), args, nil
}

func readInput() (*string, error) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error reading input: %w", err))
	}
	inputData := string(stdin)
	return &inputData, nil
}

func run(
	SolvePt1 AoCSolutionFn,
	SolvePt2 AoCSolutionFn,
	inputData string,
	args []string,
	part AoCPart,
) (answer interface{}, elapsed int64, err error) {
	start := time.Now()
	switch part {
	case AoCPart1:
		answer, err = SolvePt1(inputData, args)
	case AoCPart2:
		start = time.Now()
		answer, err = SolvePt2(inputData, args)
	default:
		panic("Should not get here!")
	}
	elapsed = time.Since(start).Nanoseconds()
	return answer, elapsed, err
}

func V1Run(SolvePt1 AoCSolutionFn, SolvePt2 AoCSolutionFn) {
	part, args, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	inputData, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	answer, elapsed, err := run(SolvePt1, SolvePt2, *inputData, args, part)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(answer)
	fmt.Printf("RT %d ns\n", elapsed)
}
