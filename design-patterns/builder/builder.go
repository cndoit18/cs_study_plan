package builder

import (
	"fmt"
	"strings"
)

type CarBuilder struct {
	seats  int
	engine string
	gps    string
}

func (c *CarBuilder) Reset() *CarBuilder {
	*c = CarBuilder{}
	return c
}

func (c *CarBuilder) SetSeats(n int) *CarBuilder {
	c.seats = n
	return c
}

func (c *CarBuilder) SetGPS(make string) *CarBuilder {
	c.gps = make
	return c
}
func (c *CarBuilder) SetEngine(make string) *CarBuilder {
	c.engine = make
	return c
}

func (h *CarBuilder) Build() string {
	b := strings.Builder{}

	_, _ = b.WriteString("This is a car")

	additional := []string{}
	if !isZero(h.seats) {
		additional = append(additional, fmt.Sprintf("%d seats", h.seats))
	}
	if !isZero(h.engine) {
		additional = append(additional, fmt.Sprintf("%s's engine", h.engine))
	}
	if !isZero(h.gps) {
		additional = append(additional, fmt.Sprintf("%s's GPS", h.gps))
	}
	if len(additional) > 0 {
		_, _ = b.WriteString(fmt.Sprintf(" with %s", strings.Join(additional, " and ")))
	}
	_, _ = b.WriteRune('.')
	return b.String()
}

func isZero[T string | int](i T) bool {
	var zreo T
	return i == zreo
}
