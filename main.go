package main

func main() {
	cards := newDeckFromFile("mycards.csv")
	cards.print()
}
