package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var TextLexer = internal.Register(MustNewLexer(
	&Config{
		Name:      "Text",
		Aliases:   []string{"txt"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
	},
	Rules{
		"root": {
			{`.+`, Text, nil},
		},
	},
))
