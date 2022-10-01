package main

import (
	"fmt"
	"reflect"
)

func ExampleMediator() {
	m := &mediator{
		components: make(map[reflect.Type]map[string]func(mediator)),
	}
	c := &CheckBox{
		m: m,
	}
	m.register(c, "Click", func(m mediator) {
		m.Notify(&Dialog{}, "Show")
	})

	d := &Dialog{
		m: m,
	}
	m.register(d, "Show", func(m mediator) {
		fmt.Println("Show")
	})

	c.Click()
	// Output:
	// Show
}
