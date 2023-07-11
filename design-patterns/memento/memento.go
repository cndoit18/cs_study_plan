package memento

type Memento interface {
	GetState() any
}

type memento struct {
	state any
}

func (m *memento) GetState() any {
	return m.state
}

func NewMemento(state any) Memento {
	return &memento{
		state: state,
	}
}

type Originator interface {
	Save() Memento
	Restore(Memento)
}

var _ Originator = &Caretakers{}

type Caretakers struct {
	Name    string
	History []Memento
}

func (c *Caretakers) ChangeName(name string) {
	c.History = append(c.History, c.Save())
	c.Name = name
}

func (c *Caretakers) RollBack() {
	target := c.History[len(c.History)-1]
	c.History = c.History[:len(c.History)-1]
	c.Restore(target)
}

func (c *Caretakers) Save() Memento {
	return NewMemento(c.Name)
}

func (c *Caretakers) Restore(m Memento) {
	c.Name = m.GetState().(string)
}
