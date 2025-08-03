
package main

import "fmt"

type Printable interface {
	doPrint()

	// this would break old code that use print()
	//fancyPrint()
}

func print (p Printable) {
	// old version
	// p.doPrint()

	// new update
	if fp, ok := p.(FancyPrintable); ok {
		fp.PrintDecor()
		p.doPrint()
		fp.PrintDecor()
	} else {
		p.doPrint()
	}
}

// added in new update
type FancyPrintable interface {
	// print checks if printable imp this
	// if yes use it
	// if no treat as legacy
	PrintDecor()

	Printable
}


// old legacy client code
type myMap map[string]string

func (mm myMap) doPrint() {
	for k, v := range mm {
		fmt.Printf("%s: %s\n", k, v)
	}
}

// new client code
type myFancyMap myMap

func (mm myFancyMap) PrintDecor() {
	fmt.Print("**********\n")
}


// front end uses both legacy and new imp
func main() {
	var nickNames myMap = map[string]string{
		"mike": "the aviator",
		"pual": "the jew",
		"hanna": "wonder cat",
	}

	print(nickNames)
	print(myFancyMap(nickNames))
}
