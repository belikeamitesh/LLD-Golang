package util
import (
	model "snakeandladder/internal/model"
)
type WinChecker interface {
	IsWinMove(*model.Player) bool
}

type SimpleWinChecker struct {}

func NewWinChecker() WinChecker {
	return &SimpleWinChecker{}
}

func (wc *SimpleWinChecker) IsWinMove(player *model.Player) bool{
	return player.Position == 100
}
