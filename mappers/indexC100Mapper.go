package mappers

type C100Entry struct {
	Term        string
	Description string
}

var C100Mapper = map[string]C100Entry{
	"Drehzahllimit":        {"G92=", "#1 Drehzahl begrenzen"},
	"Werkzeugwechselpunkt": {"GXZ73", "Revolver in X und Z-Achse auf Wechselposition"},
	"G40":                  {"G40", "Radiuskorrektur abwaehlen"},
	"G41":                  {"G41", "Radiuskorrektur links"},
	"G42":                  {"G42", "Radiuskorrektur rechts"},
}
