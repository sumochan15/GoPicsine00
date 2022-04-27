package piscine

import "ft"

func IsNegative(check int) {
	if check < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
