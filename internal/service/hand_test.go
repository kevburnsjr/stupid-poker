package service

import (
	"strings"
	"testing"
)

// Table based test to ensure correct scoring of hands
func TestScoreHand(t *testing.T) {
	var tests = map[string]string{
		"A♠ K♠ Q♠ J♠ T♠": Hand_RoyalFlush,
		"J♦ T♦ A♦ Q♦ K♦": Hand_RoyalFlush,
		"K♠ Q♠ J♠ T♠ 9♠": Hand_StraightFlush,
		"K♠ K♥ K♦ K♣ 9♠": Hand_Quads,
		"K♠ 9♠ K♥ K♦ K♣": Hand_Quads,
		"K♠ 9♠ K♥ K♦ 9♣": Hand_FullHouse,
		"T♠ 8♠ Q♠ J♠ 4♠": Hand_Flush,
		"A♠ K♦ Q♠ J♠ T♠": Hand_Straight,
		"A♠ A♥ A♦ J♠ T♠": Hand_Trips,
		"A♠ A♥ J♠ A♦ 4♠": Hand_Trips,
		"A♠ A♥ J♠ J♦ 4♠": Hand_TwoPair,
		"A♠ J♦ 4♠ A♥ J♠": Hand_TwoPair,
		"A♠ J♦ A♥ J♠ 4♠": Hand_TwoPair,
		"A♠ A♦ 8♥ J♠ 4♠": Hand_Jacks,
		"J♠ J♦ 8♥ 9♠ 4♠": Hand_Jacks,
		"T♠ T♦ 8♥ 9♠ 4♠": Hand_Nothing,
	}
	for h, exp := range tests {
		res := scoreHand(strings.Split(h, " "))
		if exp != res {
			t.Log("Unexpected Result", h, res, "!=", exp)
			t.Fail()
		}
	}
}
