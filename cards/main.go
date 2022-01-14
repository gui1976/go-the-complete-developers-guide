package main

func main() {
	cards := newDeck()

	cards.saveToFile("myCards.txt")

	anotherCards := newDeckFromFile("myCards.txt")

	anotherCards.print()
	anotherCards.shuffle()
	anotherCards.print()
}
