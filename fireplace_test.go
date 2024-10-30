package esb_fireplace

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func shouldNotRun(input_data string, args []string) (interface{}, error) {
	return nil, errors.New("This should not be running")
}

func solve_pt1(input_data string, args []string) (interface{}, error) {
	return input_data, nil
}

var pt2_answer = "christmas gopher"

func solve_pt2(input_data string, args []string) (interface{}, error) {
	return pt2_answer, nil
}

func SuppressOutput(fn interface{}) interface{} {
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()

	return reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		oldStdout := os.Stdout
		oldStderr := os.Stderr
		os.Stdout, _ = os.Open(os.DevNull)
		os.Stderr, _ = os.Open(os.DevNull)
		defer func() { os.Stdout, os.Stderr = oldStdout, oldStderr }()
		results = fnVal.Call(args)
		return
	}).Interface()
}

func TestRun(t *testing.T) {
	input_data := "123"
	args := []string{}
	type RunTestCase = struct {
		fn1        AoCSolutionFn
		fn2        AoCSolutionFn
		part       AoCPart
		expected   string
		shouldFail bool
	}
	tests := []RunTestCase{
		{solve_pt1, shouldNotRun, AoCPart1, input_data, false},
		{shouldNotRun, solve_pt2, AoCPart2, pt2_answer, false},
		{shouldNotRun, shouldNotRun, AoCPart1, "", true},
		{shouldNotRun, shouldNotRun, AoCPart2, "", true},
	}

	for _, tt := range tests {
		silentRun := SuppressOutput(run).(func(AoCSolutionFn, AoCSolutionFn, string, []string, AoCPart) (interface{}, error))
		answer, err := silentRun(
			tt.fn1,
			tt.fn2,
			input_data,
			args,
			tt.part,
		)
		if tt.shouldFail {
			if err == nil {
				t.Errorf("Solution should return an error")
			}
		} else {
			if err != nil {
				t.Errorf("Solution shouldn't return an error")
			}
			if answer != tt.expected {
				t.Errorf("Solution didn't ran properly")
			}
		}

	}
}

func TestV1Run(t *testing.T) {
	input_data := "123"
	// args := []string{}
	type RunTestCase = struct {
		fn1        AoCSolutionFn
		fn2        AoCSolutionFn
		part       string
		input_data string
		expected   string
	}
	tests := []RunTestCase{
		{solve_pt1, shouldNotRun, "1", input_data, input_data},
		{shouldNotRun, solve_pt2, "2", input_data, pt2_answer},
	}

	for _, tt := range tests {
		originalArgs := os.Args
		os.Args = []string{"cmd", "--part", tt.part}
		defer func() { os.Args = originalArgs }() // Restore original arguments after test

		originalStdin := os.Stdin
		stdinReader, stdinWriter, err := os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe for stdin: %v", err)
		}
		os.Stdin = stdinReader
		defer func() { os.Stdin = originalStdin }()

		go func() {
			defer stdinWriter.Close()
			_, err := stdinWriter.Write([]byte(tt.input_data))
			if err != nil {
				panic(fmt.Sprintf("Error writing test stdin: %v", err))
			}
		}()

		originalStdout := os.Stdout
		stdoutReader, stdoutWriter, err := os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe for stdout: %v", err)
		}
		os.Stdout = stdoutWriter
		defer func() { os.Stdout = originalStdout }()

		var outputBuf bytes.Buffer
		outC := make(chan string)
		go func() {
			_, err := io.Copy(&outputBuf, stdoutReader)
			if err != nil {
				panic(fmt.Sprintf("Error copying stdout to buffer: %v", err))
			}
			outC <- outputBuf.String()
		}()

		V1Run(tt.fn1, tt.fn2)

		stdoutWriter.Close()
		output := <-outC

		if output != tt.expected {
			t.Errorf("Expected output %q, got %q", tt.expected, output)
		}
	}
}

func TestParser(t *testing.T) {
	type ParserTestCase = struct {
		test         []string
		expectedPart AoCPart
		expectedArgs []string
		shouldFail   bool
	}
	tests := []ParserTestCase{
		{[]string{}, AoCPart1, []string{}, true},
		{[]string{"-p", "1"}, AoCPart1, []string{}, false},
		{[]string{"-p", "2"}, AoCPart2, []string{}, false},
		{[]string{"--part", "1"}, AoCPart1, []string{}, false},
		{[]string{"-p", "3"}, AoCPart1, []string{}, true},
		{[]string{"-p", "1", "--args", "abc"}, AoCPart1, []string{"abc"}, false},
		{[]string{"-p", "2", "--args", "abc", "def"}, AoCPart2, []string{"abc", "def"}, false},
	}

	for _, tt := range tests {
		silentRun := SuppressOutput(parseArgs).(func(cmd_args []string) (AoCPart, []string, error))
		part, args, err := silentRun(tt.test)
		if tt.shouldFail {
			if err == nil {
				t.Errorf("Solution should return an error")
			}
		} else {

			if part != tt.expectedPart {
				t.Errorf("Parser did not return correct part %d", part)
			}
			if !reflect.DeepEqual(args, tt.expectedArgs) {
				t.Errorf("Parser did not return correct args %v != %v", args, tt.expectedArgs)
			}
		}
	}
}