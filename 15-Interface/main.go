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
	fmt.Printf("the book title is %s and its of %f rupees\n", b.title, b.price)
}

func (g game) print() {
	fmt.Printf("the game title is %s and its of %f rupees\n", g.title, g.price)
}

type printer interface {
	print()
}

type list []printer

func main() {
	mobydick := book{
		title: "mobydick",
		price: 12.4,
	}
	minecraft := game{
		title: "minecraft",
		price: 25.6,
	}
	tetris := game{
		title: "tetris",
		price: 21.2,
	}
	var store list
	store = append(store, &minecraft, &tetris, mobydick)
	for _, item := range store {
		item.print()
	}
}
