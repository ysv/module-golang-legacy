package downcase

func Downcase(s string) (string, error) {
	r := []rune{}
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			v = (v - 'A') + 'a'
		}
		r = append(r, v)
	}
	return string(r), nil
}
