package player_test

import (
	"blackjack/player"
	"testing"
)

func TestPlayer_Dealer(t *testing.T){
	var dealer player.Dealer
	var player player.Player
	dealer.Init()
	player = &dealer
	score := player.GetScore()
	if score != 0 {
		t.Error("Something wrong with getScore on the player interface from dealer")
	}
}

func TestPlayer_RegularPlayer(t *testing.T){
	var regularPlayer player.RegularPlayer
	var player player.Player
	regularPlayer.Init()
	player = &regularPlayer
	score := player.GetScore()
	if score != 0 {
		t.Error("Something wrong with getScore on the player interface from regular player")
	}
}