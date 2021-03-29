package iteration

const repeatCount = 5

func repeat(s string) string {
	ss := ""

	for i := 0; i < repeatCount; i++ {
		ss += s
	}
	return ss
}
