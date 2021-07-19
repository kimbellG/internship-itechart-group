package cmdparser

import (
	"reflect"
	"testing"
)

func TestParseOneCmd(t *testing.T) {
	testInputs := []struct {
		inputCmd  string
		inputPipe string
		wantName  string
		wantArgs  []string
	}{
		{
			"sort -c test",
			"|",
			"sort",
			[]string{"-c", "test"},
		},
		{
			"test",
			"&",
			"test",
			[]string{},
		},
	}

	for _, testInput := range testInputs {
		result, err := parseCmd(testInput.inputCmd, testInput.inputPipe)
		if err != nil {
			t.Errorf("parse command error: %v", err)
			continue
		}

		if result.name != testInput.wantName {
			t.Errorf("name parsing failed: result = %v, want = %v", result.name, testInput.wantName)
		}

		if !reflect.DeepEqual(result.args, testInput.wantArgs) {
			t.Errorf("args parsing failed: result args = %v, want = %v", result.args, testInput.wantArgs)
		}
	}
}
