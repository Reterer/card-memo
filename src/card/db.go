package card

import "fmt"

var cards []Card

func init() {
	fmt.Println("init of db cards")
	cards = []Card{
		MakeCard("Hello World", 0),
		MakeCard("TestCard 1", 0),
		MakeCard("TestCard 2", 0),
	}
}

func AddCardIntoDB(card Card) bool {
	cards = append(cards, card)
	return true
}

func GetsCard() []Card {
	res := make([]Card, len(cards))
	copy(res, cards)
	return res
}

func ReplaceCard(i int, card Card) {
	cards[i] = card
}

func RemoveCard(i int) {
	cards[i], cards[len(cards)-1] = cards[len(cards)-1], cards[i]
	cards = cards[:len(cards)-1]
}
