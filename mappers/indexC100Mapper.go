package mappers

type C100Entry struct {
	Term        string
	Description string
}

var C100Mapper = map[string]C100Entry{
	"Drezahllimit":         {"G92=", "Drezahllimit auf Hauptspindel"},
	"Werkzeugwechselpunkt": {"GXZ73", "Kanal auf Werkzeugwechselpunkt bewegen"},
}
