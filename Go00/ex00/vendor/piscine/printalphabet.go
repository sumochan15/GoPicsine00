package piscine

import "ft"

func PrintAlphabet() {
	for al := 'a'; al <= 'z'; al++ {
		ft.PrintRune(al)
	}
	ft.PrintRune('\n')
}
