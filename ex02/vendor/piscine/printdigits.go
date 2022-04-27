package piscine

import "ft"

func PrintDigits(){
	for number := '0'; number <= '9'; number++{
		ft.PrintRune(number)
	} 
	ft.PrintRune('\n')
}