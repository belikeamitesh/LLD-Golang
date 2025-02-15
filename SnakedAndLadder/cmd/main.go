package main

import (
	"context"
	"fmt"
	"os"
	"snakeandladder/internal/controller"
	"snakeandladder/internal/model"
	"snakeandladder/internal/util"

	"github.com/rs/zerolog"
)

func init(){
    logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
    zerolog.DefaultContextLogger = &logger
}
func main() {
    ctx := context.Background()
    logger := zerolog.Ctx(ctx)
    logger.Info().Msg("Logger Initialized")

    players := []*model.Player {
        model.NewPlayer("Amitesh"),
        model.NewPlayer("Swastik"),
    }

    entities := []model.Entity {
        model.NewSnake(50, 10),
		model.NewSnake(30, 20),
		model.NewSnake(65, 15),
        model.NewLadder(3,27),
        model.NewLadder(54,79),
        model.NewLadder(24,91),
    }

    dice1,err := model.NewDice(6)
    if err != nil {
		logger.Err(err).Msg("Error creating dice")
		return
	}

	dice2, err := model.NewDice(6)
	if err != nil {
		logger.Err(err).Msg("Error creating dice")
		return
	}

    dices := []model.Dice{dice1, dice2}
    board := model.NewBoard(entities...)
    winChecker := util.NewWinChecker()

    game := controller.NewGame(dices,board,players,winChecker)
    game.Start()
    winner := game.GetWinner()
    fmt.Println("Winner of the Game is: ",winner.Name)
}