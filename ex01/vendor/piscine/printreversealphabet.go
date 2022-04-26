package	piscine

import	"ft"

func PrintReverseAlphabet() {
	for zyx := 'z'; zyx >= 'a'; zyx--{
		ft.PrintRune(zyx)
	}
	ft.PrintRune('\n')
}