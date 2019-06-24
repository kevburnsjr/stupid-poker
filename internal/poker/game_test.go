package poker

import (
	"fmt"
	"testing"
)

// Create game and exchange sets of cards in successive hands
func TestExchange(t *testing.T) {
	game := NewGame(5, 200)
	h, _, _ := game.Deal()
	hcopy := h.clone()
	hfinal, _ := game.Exchange([]int{0, 1, 2})
	for i := 0; i < 3; i++ {
		if hcopy[i] == hfinal[i] {
			t.Log(fmt.Sprintf("Card %d not exchanged %s - %s", i+1, hcopy, hfinal))
			t.Fail()
		}
	}
	for i := 3; i < 5; i++ {
		if hcopy[i] != hfinal[i] {
			t.Log(fmt.Sprintf("Card %d exchanged %s - %s", i+1, hcopy, hfinal))
			t.Fail()
		}
	}
	h, _, _ = game.Deal()
	hcopy = h.clone()
	hfinal, _ = game.Exchange([]int{3, 4})
	for i := 0; i < 3; i++ {
		if hcopy[i] != hfinal[i] {
			t.Log(fmt.Sprintf("Card %d exchanged %s - %s", i+1, hcopy, hfinal))
			t.Fail()
		}
	}
	for i := 3; i < 5; i++ {
		if hcopy[i] == hfinal[i] {
			t.Log(fmt.Sprintf("Card %d not exchanged %s - %s", i+1, hcopy, hfinal))
			t.Fail()
		}
	}
}

// Create game with 0 balance and ensure it can't be played
func TestBroke(t *testing.T) {
	game := NewGame(5, 0)
	_, _, err := game.Deal()
	if err == nil {
		t.Log("Game dealt hand to player with 0 balance")
		t.Fail()
	}
}
