package structure

type Vertexer interface {
	GetX() int
	GetY() int
}

type Vertex2 struct {
	X int
	Y int
}

type Vertex struct {
	X int
	Y int
}

func (v Vertex2) GetX() int {
	return v.X * 20
}

func (v Vertex2) GetY() int {
	return v.Y * 20
}

func (v Vertex) GetX() int {
	return v.X
}

func (v Vertex) GetY() int {
	return v.Y
}

func (v *Vertex) SetX(x int) {
	v.X = x
}

func (v *Vertex) SetY(y int) {
	v.Y = y
}
