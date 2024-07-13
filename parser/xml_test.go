package parser

import (
	"testing"
)

func TestAssertXml(t *testing.T) {
	var xmlTests = []struct {
		name   string
		input  string
		output bool
	}{
		{
			"normal case", "<Design><Code>hello world</Code></Design>", true,
		},
		{
			"no closing tag", "<Design><Code>hello world</Design><People>", false,
		},
		{
			"non corresponding tags", "<People><Design><Code>hello world</People></Code></Design>", false,
		},
		{
			"attribute is not supported", "<People age='1'>hello world</People>", false,
		},
	}

	for _, tt := range xmlTests {
		t.Run(tt.name, func(t *testing.T) {
			d := AssertXml(tt.input)
			if d != tt.output {
				t.Errorf("got: %v, want: %v\n", d, tt.output)
			}
		})
	}
}
