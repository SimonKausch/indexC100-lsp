package mappers

type CodeSnippet struct {
	Body        string
	Description string
}

var CodeSnippets = map[string]CodeSnippet{
	"T": {
		Body:        "T${1:Station}D${2:Schneidennr.}",
		Description: "Werkzeug anwaehlen",
	},
}
