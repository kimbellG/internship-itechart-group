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

func TestAllCmd(t *testing.T) {
	testInfos := []struct {
		input string
		want  []cmdInformation
	}{
		{
			"sort -c",
			[]cmdInformation{{"sort", []string{"-c"}, ""}},
		},
		{
			"sort -c | awk abde | test -a -s",
			[]cmdInformation{{"sort", []string{"-c"}, "|"}, {"awk", []string{"abde"}, "|"}, {"test", []string{"-a", "-s"}, ""}},
		},
	}

	for _, testInfo := range testInfos {
		result, err := Parse(testInfo.input)
		if err != nil {
			t.Errorf("error from func: %v", err)
			continue
		}

		if len(testInfo.want) != len(result) {
			t.Errorf("incorrect len of result. want: %v, result len: %v", len(testInfo.want), len(result))
			continue
		}

		for i := range result {
			if testInfo.want[i].name != result[i].name {
				t.Errorf("incorrect parsing of input string. name isn't eq. want: %v, result: %v", testInfo.want[i].name, result[i].name)
			}

			if !reflect.DeepEqual(testInfo.want[i].args, result[i].args) {
				t.Errorf("incorrect parsing of input string: args aren't eq. want: %v, result: %v", testInfo.want[i].args, result[i].args)
			}

			if testInfo.want[i].pipelineOperation != result[i].pipelineOperation {
				t.Errorf("incorrect parsing of input string: operation isn't eq. want: %v, result: %v",
					testInfo.want[i].pipelineOperation, result[i].pipelineOperation)
			}
		}
	}
}
