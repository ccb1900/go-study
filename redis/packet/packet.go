package packet

const EL = "\r\n"

func OkLine(s string) string {
	return "+" + s + EL
}

func ErrLine(s string) string {
	return "-" + s + EL
}
