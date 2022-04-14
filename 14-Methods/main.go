package main

import "fmt"

type book struct {
	title string
	price float64
}

type game struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Printf("the title is %s and price is %f", b.title, b.price)
}

func (g game) print() {
	fmt.Printf("the title is %s and price is %f", g.title, g.price)
}

func (g *game) discount(r float64) {
	g.price *= (1 - r)
}

func main() {
	myBook := book{
		title: "mobydick",
		price: 12.2,
	}

	minecraft := game{
		title: "minecraft",
		price: 15,
	}
	minecraft.print()
	minecraft.discount(.5)
	minecraft.print()
	myBook.print()
}
