package service

import (
	"fmt"
	"strings"
	"testing"
)

// Create game and exchange sets of cards in successive hands
func TestExchange(t *testing.T) {
	game := NewGame(5, 200)
	hand := append([]string(nil), game.Deal()...)
	newHand, _, _ := game.Exchange([]int{0, 1, 2})
	for i := 0; i < 3; i++ {
		if hand[i] == newHand[i] {
			t.Log(fmt.Sprintf("Card %d not exchanged %s - %s", i+1,
				strings.Join(hand, " "),
				strings.Join(newHand, " ")))
			t.Fail()
		}
	}
	for i := 3; i < 5; i++ {
		if hand[i] != newHand[i] {
			t.Log(fmt.Sprintf("Card %d exchanged %s - %s", i+1,
				strings.Join(hand, " "),
				strings.Join(newHand, " ")))
			t.Fail()
		}
	}
	hand = append([]string(nil), game.Deal()...)
	newHand, _, _ = game.Exchange([]int{3, 4})
	for i := 0; i < 3; i++ {
		if hand[i] != newHand[i] {
			t.Log(fmt.Sprintf("Card %d exchanged %s - %s", i+1,
				strings.Join(hand, " "),
				strings.Join(newHand, " ")))
			t.Fail()
		}
	}
	for i := 3; i < 5; i++ {
		if hand[i] == newHand[i] {
			t.Log(fmt.Sprintf("Card %d not exchanged %s - %s", i+1,
				strings.Join(hand, " "),
				strings.Join(newHand, " ")))
			t.Fail()
		}
	}
}
