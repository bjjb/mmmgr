package files

var tvPatterns = []string{}

var tvREs *reList

func init() {
	tvREs = compileREs(tvPatterns)
}
