package controller

import (
	"context"
	model "snakeandladder/internal/model"
	util "snakeandladder/internal/util"

	"github.com/rs/zerolog"
)

type Game interface {
	Start()
	GetWinner() *model.Player
}

type gameImpl struct {
	dices []model.Dice
	board  model.Board
	players []*model.Player
	winner  *model.Player
	winChecker  util.WinChecker
}

func NewGame(dices []model.Dice,board  model.Board,players []*model.Player,winChecker  util.WinChecker) Game {
	game := &gameImpl{
		dices: make([]model.Dice, 0),
		board: board,
		players: make([]*model.Player, 0),
		winner: nil,
		winChecker: winChecker,
	}
	game.dices = append(game.dices, dices...)
	game.players = append(game.players, players...)
	return game
}

func (game *gameImpl) Start() {
	logger := zerolog.Ctx(context.Background()).With().Str("module", "game").Logger()
	i:=0
	for game.winner == nil {
		i = i% len(game.players)
		currentPlayer := game.players[i]
		initialPosition := currentPlayer.Position
		steps := 0
		for _,dice := range game.dices {
			steps+= dice.RollDice()
		}
		game.board.MakeMove(currentPlayer,steps)
		finalPosition := currentPlayer.Position
		finised := game.winChecker.IsWinMove(currentPlayer)
		logger.Info().Msgf("%s rolled a %d and moved from %d to %d ",currentPlayer.Name,steps,initialPosition,finalPosition)
		if finised {
			game.winner = currentPlayer
			return
		}
		i+=1
	}
}

func (game *gameImpl) GetWinner() *model.Player {
	return game.winner
}