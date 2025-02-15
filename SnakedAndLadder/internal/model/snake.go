package model

type Snake struct{
	start int
	end int
}

//constructor
func NewSnake(start,end int) Entity {
	return &Snake{
		start: start,
		end: end,
	}
}

func (s *Snake) GetStartPosition() int{
	return s.start
}

func (s *Snake) GetEndPosition() int{
	return s.end
}

