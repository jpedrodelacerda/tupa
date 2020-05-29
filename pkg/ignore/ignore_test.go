package ignore_test

import (
	"bytes"
	"github.com/jpedrodelacerda/tupa/pkg/ignore"
	"testing"
)

func TestCreateGitIgnore(t *testing.T) {
	testCases := []struct {
		desc  string
		query string
		want  string
	}{
		{
			desc:  "try simple query",
			query: "go",
			want:  ignoreGo,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			want := []byte(tC.want)
			got, err := ignore.FetchGitIgnore(tC.query)
			if err != nil {
				t.Errorf("fail to fetch gitignore: - %s\n\tquery:\t%s\nwant:\t%s\n\tgot:\t%s", tC.desc, tC.query, tC.want, string(got))
			}

			if r := bytes.Compare(got, want); r != 0 {
				t.Errorf("got wrong gitignore: - %s\n\tquery:\t%s\nwant:\t%s\n\tgot:\t%s", tC.desc, tC.query, tC.want, string(got))
			}

		})
	}
}

// https://www.gitignore.io/api/go
var ignoreGo = `
# Created by https://www.gitignore.io/api/go
# Edit at https://www.gitignore.io/?templates=go

### Go ###
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + "`go test -c`" + `
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

### Go Patch ###
/vendor/
/Godeps/

# End of https://www.gitignore.io/api/go
`
