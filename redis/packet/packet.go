package packet

const EL = "\r\n"

func OkLine(s string) string {
	return "+" + s + EL
}

func GetString(s string) string {
	return OkLine("\"" + s + "\"")
}

func ErrLine(s string) string {
	return "-" + s + EL
}
