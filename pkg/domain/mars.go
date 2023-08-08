package domain

type Map struct {
	width     int
	height    int
	obstacles []Coordinate
}

func NewMap(width int, height int) *Map {
	return &Map{
		width:     width,
		height:    height,
		obstacles: nil,
	}
}

func NewMapWithObstacles(width int, height int, obstacles []Coordinate) *Map {
	return &Map{
		width:     width,
		height:    height,
		obstacles: obstacles,
	}
}

func (m *Map) Width() int {
	return m.width
}

func (m *Map) Height() int {
	return m.height
}

func (m *Map) HasObstacle(coord Coordinate) bool {
	for _, obstacle := range m.obstacles {
		if obstacle == coord {
			return true
		}
	}

	return false
}
