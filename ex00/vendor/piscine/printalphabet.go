package piscine

import "ft"

func PrintAlphabet() {
	for az := 'a'; az <= 'z'; az++ {
		ft.PrintRune(az)
	}
	ft.PrintRune('\n')
}
