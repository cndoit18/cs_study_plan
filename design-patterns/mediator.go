package main

import "reflect"

type Mediator interface {
	Notify(component interface{}, evnet string)
}

type CheckBox struct {
	m Mediator
}

func (c *CheckBox) Click() {
	c.m.Notify(c, "Click")
}

type Dialog struct {
	m Mediator
}

func (d *Dialog) Show() {
	d.m.Notify(d, "Show")
}

type mediator struct {
	components map[reflect.Type]map[string]func(mediator)
}

func (m *mediator) register(component interface{}, evnent string, evnentFunc func(m mediator)) {
	key := reflect.TypeOf(component)
	eventFuncs := map[string]func(mediator){}
	if v, ok := m.components[key]; ok {
		eventFuncs = v
	}
	eventFuncs[evnent] = evnentFunc
	m.components[key] = eventFuncs
}

func (m *mediator) Notify(component interface{}, evnet string) {
	f := m.components[reflect.TypeOf(component)][evnet]
	f(*m)
}
