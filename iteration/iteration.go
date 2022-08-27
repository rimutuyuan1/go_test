package iteration

func Iteration(param string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += param
	}
	return repeated
}
