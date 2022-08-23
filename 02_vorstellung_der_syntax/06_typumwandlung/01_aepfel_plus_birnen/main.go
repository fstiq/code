package main

import (
	"fmt"
)

/*
Code kompiliert nicht, da wir Äpfel mit Birnen
vergleichen.
*/

func main() {
	type äpfel int
	type birnen float32
	a := äpfel(10)
	b := birnen(5.1)
	fmt.Println(a + äpfel(b))
}
