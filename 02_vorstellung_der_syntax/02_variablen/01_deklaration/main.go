package main

import "fmt"

func main() {

	/*
		Dieser Code behandelt alle Beispiele zur Deklaration
		von Variablen
		Playground:
		https://play.golang.org/p/hzo7dE0NY8M
	*/

	var nummer int               // Definition ohne Wert
	var name string = "Rob Pike" // Definition mit Wert und Typ
	var anzahl = 10              // Definition ohne Typ

	fmt.Println(nummer, name, anzahl)

	/*
		Nullwerte bei der Deklaration mit var
	*/
	var nummer2 int // 0

	var txt string   // ""
	var checked bool // false

	type user struct{}
	var meinUser *user // nil - Pointer

	var liste []string // nil - Slice

	fmt.Println("nummer2", nummer2)
	fmt.Println("txt", txt)
	fmt.Println("checked", checked)
	fmt.Println("liste", liste)
	fmt.Println("meinUser", meinUser)

	nummer3 := 10      //  int
	name3 := getName() // string

	fmt.Println(nummer3, name3)

	a, b := 12, 34
	a, b = b, a // Werte tauschen
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// my test
	fmt.Println(add(2, 4))

	var test []int
	test = append(append(test, 12), 23)
	test = append(test, 123, 122, 333)
	fmt.Println(test)

	type test2 struct {
		id   int
		name string
	}
	var tUser *test2
	fmt.Println(tUser)

}

func getName() string {
	return "Andreas"
}

func add(a int, b int) int {
	return a + b
}
