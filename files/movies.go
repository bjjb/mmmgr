package files

var moviePatterns = []string{}

var movieREs *reList

func init() {
	movieREs = compileREs(moviePatterns)
}
