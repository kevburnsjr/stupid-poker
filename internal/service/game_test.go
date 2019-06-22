package service

import (
	"fmt"
	"strings"
	"testing"
)

// Create game and exchange sets of cards in successive hands
func TestExchange(t *testing.T) {
	game := NewGame(5, 200)
	cards, _ := game.Deal()
	hand := append([]string(nil), cards...)
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
	cards, _ = game.Deal()
	hand = append([]string(nil), cards...)
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

// Create game with 0 balance and ensure it can't be played
func TestBroke(t *testing.T) {
	game := NewGame(5, 0)
	_, err := game.Deal()
	if err == nil {
		t.Log("Game dealt hand to player with 0 balance")
		t.Fail()
	}
}
