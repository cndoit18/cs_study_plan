package prototype

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type People struct {
	Name   string
	Age    uint
	Gender Gender
}

func (p *People) Clone() *People {
	o := &People{}
	*o = *p
	return o
}
