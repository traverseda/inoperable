package layout

type DrawObject interface {
	layout()
}

type BaseObject struct {
	realX    float64 //The actual size of the object, after all our calculation
	realY    float64
	minX     float32
	minY     float32
	maxX     float32
	maxY     float32
	depth    int
	parent   DrawObject
	children []DrawObject
}

func (self BaseObject) layout() {
	for _, drawObject := range f.children {
		e.depth = self.depth + 1
		e.layout()
	}
}

//Varients of drawObject
type Text struct {
	BaseObject
	content []rune
}

type Box struct {
	BaseObject
}
