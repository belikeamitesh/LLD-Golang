package model

type Entity interface {
	GetStartPosition() int
	GetEndPosition() int
}