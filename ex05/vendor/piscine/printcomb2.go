package piscine

import "ft"

func PrintComb2Set(n int){
	ft.PrintRune(rune(n / 10 + '0'))
	ft.PrintRune(rune(n % 10 + '0'))
}

func PrintComb2() {
	for i := 0; i < 100; i++{
		for j := i + 1; j < 100; j++{
			PrintComb2Set(i)
			ft.PrintRune(' ')
			PrintComb2Set(j)
			if i != 98 || j != 99{
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}