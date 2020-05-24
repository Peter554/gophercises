package lib

func CamelCase(s string) int {
	if len(s) == 0 {
		return 0
	}
	o := 1
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			o++
		}
	}
	return o
}
