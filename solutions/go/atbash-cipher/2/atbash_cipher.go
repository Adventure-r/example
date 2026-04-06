package atbashcipher

func Atbash(s string) string {
	length := len(s)
	b := make([]byte, 0, length+length/5+1)
	counter := 0
	for i := 0; i < length; i++ {
		letter := s[i]
		var out byte
		switch {
		case letter >= '0' && letter <= '9':
			out = letter
		case letter >= 'a' && letter <= 'z':
			out = 'z' - letter + 'a'
		case letter >= 'A' && letter <= 'Z':
			out = 'z' - letter + 'A'
		default:
			continue
		}
		b = append(b, out)
		counter++
		if counter%5 == 0 {
			b = append(b, ' ')
		}
	}
	bLength := len(b)
	if b[bLength-1] == ' ' {
		b = b[:bLength-1]
	}
	return string(b)
}
