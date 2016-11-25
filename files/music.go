package files

var musicPatterns = []string{}

var musicREs *reList

func init() {
	musicREs = compileREs(musicPatterns)
}
