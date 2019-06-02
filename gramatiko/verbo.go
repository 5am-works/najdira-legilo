package gramatiko

type Verbfinaĵo struct {
	V0 string
	V1 string
	V1S string
	A1 string
}

var be = Verbfinaĵo{
	V0: "be",
	V1: "pe",
	V1S: "ve",
	A1: "ne",
}

var vi = Verbfinaĵo{
	V0: "vi",
	V1: "fi",
	V1S: "pi",
	A1: "zi",
}

var sa = Verbfinaĵo{
	V0: "sa",
	V1: "za",
	V1S: "da",
	A1: "va",
}

var nai = Verbfinaĵo{
	V0: "nA",
	V1: "dA",
	V1S: "zA",
	A1: "sE",
}

var zɔ = Verbfinaĵo{
	V0: "zɔ",
	V1: "sɔ",
	V1S: "nɔ",
	A1: "ʃɔ",
}

var lu = Verbfinaĵo{
	V0: "lu",
	V1: "ru",
	V1S: "ʃu",
	A1: "zu",
}

var ko = Verbfinaĵo{
	V0: "ko",
	V1: "go",
	V1S: "wo",
	A1: "po",
}

var V0Finaĵoj = map[string]Verbfinaĵo{
	"be": be,
	"vi": vi,
	"sa": sa,
	"nA": nai,
	"zɔ": zɔ,
	"lu": lu,
	"ko": ko,
}