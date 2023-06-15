package adapter

type Round interface {
	GetRadius() int
}

var _ Round = &RoundHole{}

type RoundHole struct {
	Radius int
}

func (r *RoundHole) GetRadius() int {
	return r.Radius
}

var _ Round = &RoundPeg{}

type RoundPeg struct {
	Radius int
}

func (r *RoundPeg) GetRadius() int {
	return r.Radius
}

type Square interface {
	GetWidth() int
}

var _ Square = &SquarePeg{}

type SquarePeg struct {
	Width int
}

func (s *SquarePeg) GetWidth() int {
	return s.Width
}

type roundAdapter struct {
	square Square
}

func (r roundAdapter) GetRadius() int {
	return r.square.GetWidth()
}

func NewRouadAdapater(square Square) Round {
	return &roundAdapter{square: square}
}
