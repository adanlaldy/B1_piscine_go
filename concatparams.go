package piscine

func ConcatParams(args []string) string {
	videuh := ""
	for i := 0; i < len(args); i++ {
		videuh += args[i]
		if i < len(args)-1 {
			videuh += "\n"
		}
	}
	return videuh
}
