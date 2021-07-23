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
		wantErr   string
	}{
		{
			"sort -c test",
			"|",
			"sort",
			[]string{"-c", "test"},
			"",
		},
		{
			"test",
			"|",
			"test",
			[]string{},
			"",
		},
		{
			"",
			"",
			"",
			[]string{},
			"command is not defined",
		},
	}

	for _, testInput := range testInputs {
		got, err := parseCmd(testInput.inputCmd, testInput.inputPipe)
		if err != nil {
			if err.Error() != testInput.wantErr {
				t.Errorf("parse command error: %v", err)
			}

			continue
		}

		if got.name != testInput.wantName {
			t.Errorf("name parsing failed: result = %v, want = %v", got.name, testInput.wantName)
		}

		if !reflect.DeepEqual(got.args, testInput.wantArgs) {
			t.Errorf("args parsing failed: result args = %v, want = %v", got.args, testInput.wantArgs)
		}
	}
}

func TestAllCmd(t *testing.T) {
	testInfos := []struct {
		input string
		want  []CMDInformation
		err   string
	}{
		{
			"sort -c",
			[]CMDInformation{{"sort", []string{"-c"}, ""}},
			"",
		},
		{
			"sort -c | awk abde | test -a -s",
			[]CMDInformation{{"sort", []string{"-c"}, "|"}, {"awk", []string{"abde"}, "|"}, {"test", []string{"-a", "-s"}, ""}},
			"",
		},
		{
			"grep",
			[]CMDInformation{{"grep", []string{}, ""}},
			"",
		},
		{
			"",
			[]CMDInformation{{"", []string{}, ""}},
			"parse input command: command is not defined",
		},
	}

	for _, testInfo := range testInfos {
		got, err := Parse(testInfo.input)
		if err != nil {
			if err.Error() != testInfo.err {
				t.Errorf("error from func: %v", err)
			}
			continue
		}

		if len(testInfo.want) != len(got) {
			t.Errorf("incorrect len of result. want: %v, got len: %v", len(testInfo.want), len(got))
			continue
		}

		for i := range got {
			if testInfo.want[i].name != got[i].name {
				t.Errorf("incorrect parsing of input string. name isn't eq. want: %v, got: %v", testInfo.want[i].name, got[i].name)
			}

			if !reflect.DeepEqual(testInfo.want[i].args, got[i].args) {
				t.Errorf("incorrect parsing of input string: args aren't eq. want: %v, got: %v", testInfo.want[i].args, got[i].args)
			}

			if testInfo.want[i].pipelineOperation != got[i].pipelineOperation {
				t.Errorf("incorrect parsing of input string: operation isn't eq. want: %v, got: %v",
					testInfo.want[i].pipelineOperation, got[i].pipelineOperation)
			}
		}
	}
}
