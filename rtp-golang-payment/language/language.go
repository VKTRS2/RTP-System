package language

type Language uint

const (
	Invalid Language = iota
	Catalan
	English
	Spanish
)

var values = []string{
	"Indefinit",
	"Català",
	"Espanyol",
	"Anglès",
}

func (p Language) String() string {
	return values[p]
}
