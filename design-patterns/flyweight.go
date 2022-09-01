package main

type ColorFactory struct {
	colors map[string]Color
}

type PtrColor struct {
	color string
}

func (p *PtrColor) String() string {
	return p.color
}

func (f *ColorFactory) GetColor(t string) Color {
	switch t {
	case "red":
		c, ok := f.colors["red"]
		if !ok {
			c = &PtrColor{color: "red"}
			f.colors["red"] = c
		}
		return c
	case "blue":
		c, ok := f.colors["blue"]
		if !ok {
			c = &PtrColor{color: "blue"}
			f.colors["blue"] = c
		}
		return c
	}
	return nil
}
