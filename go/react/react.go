package react

// Define reactor, cell and canceler types here.
// These types will implement the Reactor, Cell and Canceler interfaces, respectively.
type reactor struct {
	cells []cell
}

type cell struct {
	dep1     Cell
	dep2     Cell
	compute  func(int, int) int
	callback func(int)
	value    int
}

type canceler struct {
	cell *cell
}

func (c *canceler) Cancel() {
	c.cell.callback = nil
}

func (c *cell) Value() int {
	if c.compute == nil {
		return c.value
	}
	return c.compute(c.dep1.Value(), c.dep2.Value())
}

func (c *cell) SetValue(value int) {
	c.value = value
	if c.callback != nil {
		c.callback(value)
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callback = callback
	return &canceler{cell: c}
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	cell := cell{value: initial}
	return &cell
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	c := cell{dep1: dep, dep2: dep, compute: func(v int, _ int) int { return compute(v) }}
	r.cells = append(r.cells, c)
	return &c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	c := cell{dep1: dep1, dep2: dep2, compute: compute}
	r.cells = append(r.cells, c)
	return &c
}
