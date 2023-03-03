// Package internal provides internal helper functions.
package internal

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	var mapper map[interface{}]interface{}
	var tests = []struct {
		name   string
		text   string
		config string
		want   string
	}{
		{
			"one",
			`FROM {{ .Base }}`,
			`Base: ubuntu`,
			`FROM ubuntu`,
		},
		{
			"two",
			`RUN cat {{ .mylongdirpath | upper }}`,
			`mylongdirpath: /tmp/folder/folder2/file.txt`,
			`RUN cat /TMP/FOLDER/FOLDER2/FILE.TXT`,
		},
		{
			"three",
			`RUN touch {{ .message | nospace }}`,
			`message: h e ll o wor ld`,
			`RUN touch helloworld`,
		},
		{
			"four",
			`RUN echo maximum is: {{ max .maxNumbers.a .maxNumbers.b .maxNumbers.c }}`,
			`
maxNumbers:
  a: 4
  b: 99
  c: 1`,
			`RUN echo maximum is: 99`,
		},
		{
			"five",
			`
FROM {{ .Base }}
RUN cat {{ .mylongdirpath | upper }}
RUN touch {{ .message | nospace }}
RUN echo maximum is: {{ max .maxNumbers.a .maxNumbers.b .maxNumbers.c }}
			`,
			`
Base: ubuntu
mylongdirpath: /tmp/folder/folder2/file.txt
message: h e ll o wor ld
maxNumbers:
  a: 4
  b: 99
  c: 1
`,
			`
FROM ubuntu
RUN cat /TMP/FOLDER/FOLDER2/FILE.TXT
RUN touch helloworld
RUN echo maximum is: 99
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			mapper, err = ReadConfigFromText(tt.config)
			if err != nil {
				t.Errorf("test %s's config couldn't be loaded! (%s)", tt.name, err)
			}

			ans, err := ParseFile(tt.name, tt.text, mapper)

			if err != nil {
				t.Errorf("test %s couldn't be parsed! (%s)", tt.name, err)
			}

			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
