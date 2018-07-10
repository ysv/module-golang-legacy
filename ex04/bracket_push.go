package brackets

//works with text, f.e. "{adasd[gg]asd(2+4)}asdf"

func skip_valid_pairs(runes []rune) bool {
	len := len(runes)
	for i := 0; i < len; {
		if is_bracket(runes[i]) {
			if is_opening_bracket(runes[i]) {
				skipped := skip_valid_pair(runes[i+1:], runes[i])
				if skipped < 0 {
					return false
				} else {
					i += skipped
				}
			} else {
				return false
			}
		}
		i += 1
	}
	return true
}

func is_bracket(symb rune) bool {
	switch symb {
	case '{', '}', '(', ')', '[', ']':
		return true
	}

	return false
}

func is_opening_bracket(symb rune) bool {
	switch symb {
	case '{', '(', '[':
		return true
	}

	return false
}

func is_closing_bracket(symb rune) bool {
	switch symb {
	case '}', ')', ']':
		return true
	}

	return false
}

func is_pair(os, cs rune) bool {
	switch {
	case os == '{' && cs == '}':
		return true
	case os == '(' && cs == ')':
		return true
	case os == '[' && cs == ']':
		return true
	}

	return false
}

func skip_valid_pair(runes []rune, symb rune) int {
	len := len(runes)
	for i := 0; i < len; {
		if is_bracket(runes[i]) {
			if is_opening_bracket(runes[i]) {
				skipped := skip_valid_pair(runes[i+1:], runes[i])
				if skipped < 0 {
					return -1
				} else {
					i += skipped
				}
			} else {
				if is_closing_bracket(runes[i]) {
					if is_pair(symb, runes[i]) {
						return i + 1
					} else {
						return -1
					}
				}
			}
		}
		i += 1
	}

	return -1
}

func Bracket(s string) (bool, error) {
	runes := []rune(s)

	b := skip_valid_pairs(runes)

	if !b {
		return false, nil
	}
	return true, nil
}
