package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

//////////////CAESAR
func NewCaesar() Cipher {
	return NewShift(3)
}

//////////////SHIFT
func NewShift(sh int) Cipher {
	switch {
	case sh == 0:
		return nil
	case sh > 25:
		return nil
	case sh < -25:
		return nil
	default:
		return NewVigenere(string('a' + (26+rune(sh))%26))
	}
}

//////////////VIGENERE
type Vigenere struct {
	key string
}

func (vi Vigenere) Encode(s string) string {
	s, _ = CipherDowncase(s)
	r := []rune{}
	klen := len(vi.key)

	for i, v := range s {
		k := vi.key[i%klen]
		shift := rune(k - 'a')
		v = (v-'a'+rune(shift))%26 + 'a'
		r = append(r, v)
	}

	return string(r)
}

func (vi Vigenere) Decode(s string) string {
	s, _ = CipherDowncase(s)
	r := []rune{}
	klen := len(vi.key)

	for i, v := range s {
		k := vi.key[i%klen]
		shift := rune(k - 'a')
		v = (v-'a'+(26-rune(shift))%26)%26 + 'a'
		r = append(r, v)
	}

	return string(r)
}

func NewVigenere(key string) Cipher {
	if key_not_valid(key) {
		return nil
	}
	return Vigenere{key}
}

////////////////Other tools
//Skips spaces and other non-latin-characters
func CipherDowncase(s string) (string, error) {
	r := []rune{}
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			if v >= 'A' && v <= 'Z' {
				v = (v - 'A') + 'a'
			}
			r = append(r, v)
		}
	}
	return string(r), nil
}

func key_not_valid(key string) bool {
	a_flag := true

	if len(key) < 1 {
		return true
	}

	for _, v := range key {
		if v != 'a' {
			a_flag = false
		}
		if v < 'a' || v > 'z' {
			return true
		}
	}

	if a_flag {
		return true
	}

	return false
}
