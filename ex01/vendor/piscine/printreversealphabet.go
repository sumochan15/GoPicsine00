package	piscine

import	"ft"

func PrintReverseAlphabet() {
	for za := 'z'; za >= 'a'; za--{
		ft.PrintRune(za)
	}
	ft.PrintRune('\n')
}