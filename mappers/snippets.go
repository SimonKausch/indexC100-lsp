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
	"#rant": {
		Body:        "NAH FR THO 😤 ${1:topic} is actually WILD 💯 like how are people even ${2:action} fr⁉️ this is actually so ${3:adjective} i cant even- 💀 ${4:extra_thoughts}",
		Description: "Create a passionate rant template",
	},
}
