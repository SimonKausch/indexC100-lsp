package mappers

import protocol "github.com/tliron/glsp/protocol_3_16"

type C100Entry struct {
	Term          string
	Description   string
	Documentation protocol.MarkupContent
}

func createMarkup(markdown string) protocol.MarkupContent {
	return protocol.MarkupContent{Kind: protocol.MarkupKindMarkdown, Value: markdown}
}

var C100Mapper = map[string]C100Entry{
	"Drehzahllimit":        {"G92=", "#1 Drehzahl begrenzen", createMarkup("")},
	"Werkzeugwechselpunkt": {"GXZ73", "Revolver in X und Z-Achse auf Wechselposition", createMarkup("")},
	"G40":                  {"G40", "Radiuskorrektur abwaehlen", createMarkup("")},
	"G41":                  {"G41", "Radiuskorrektur links", createMarkup("")},
	"G42":                  {"G42", "Radiuskorrektur rechts", createMarkup("# Radiuskorrektur\nText")},
}
