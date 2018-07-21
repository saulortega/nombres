package nombres

import (
	"strings"
)

func Corregir(n string) string {
	var pdzs = []string{}

	for _, nn := range strings.Split(n, " ") {
		pdzs = append(pdzs, corregir(nn))
	}

	return strings.Join(pdzs, " ")
}

func corregir(n string) string {
	n = strings.TrimSpace(n)
	mi := strings.ToLower(n)
	ma := strings.ToUpper(mi)
	in := strings.Title(mi)
	if n != mi && n != ma && n != in {
		return n
	}

	N, e := mapaCorrecciones[mi]
	if !e {
		return n
	}

	switch n {
	case mi:
		n = N
	case ma:
		n = strings.ToUpper(N)
	case in:
		n = strings.Title(N)
	}

	return n
}

var mapaCorrecciones = map[string]string{
	"saul":       "saúl",
	"raul":       "raúl",
	"ines":       "inés",
	"erika":      "érika",
	"angel":      "ángel",
	"maria":      "maría",
	"rocio":      "rocío",
	"fabian":     "fabián",
	"andres":     "andrés",
	"alexander":  "alexánder",
	"ramirez":    "ramírez",
	"diaz":       "díaz",
	"cordoba":    "córdoba",
	"garzon":     "garzón",
	"montañez":   "montáñez",
	"suarez":     "suárez",
	"garcia":     "garcía",
	"rodriguez":  "rodríguez",
	"gonzalez":   "gonzález",
	"martinez":   "martínez",
	"cardenas":   "cárdenas",
	"guarin":     "guarín",
	"farias":     "farías",
	"leguizamo":  "leguízamo",
	"hernandez":  "hernández",
	"estefania":  "estefanía",
	"nohemi":     "nohemí",
	"guzman":     "guzmán",
	"gomez":      "gómez",
	"perez":      "pérez",
	"fuquene":    "fúquene",
	"gutierrez":  "gutiérrez",
	"narvaez":    "narváez",
	"cuellar":    "cuéllar",
	"rincon":     "rincón",
	"cortes":     "cortés",
	"oscar":      "óscar",
	"sofia":      "sofía",
	"sanchez":    "sánchez",
	"bermudez":   "bermúdez",
	"pinzon":     "pinzón",
	"sebastian":  "sebastián",
	"lopez":      "lópez",
	"hortua":     "hortúa",
	"leon":       "león",
	"sepulveda":  "sepúlveda",
	"victor":     "víctor",
	"vasquez":    "vásquez",
	"ingrid":     "íngrid",
	"raquira":    "ráquira",
	"cesar":      "césar",
	"calderon":   "calderón",
	"giron":      "girón",
	"rios":       "ríos",
	"martin":     "martín",
	"matias":     "matías",
	"albarracin": "albarracín",
}
