package observer

import (
	"fmt"
	"io"
)

type Publisher struct {
	Subscribers []Subscriber
}

func (p *Publisher) Notify(s string) {
	for _, v := range p.Subscribers {
		v.Notify(s)
	}
}

func (p *Publisher) AddSubscriber(s Subscriber) {
	for _, v := range p.Subscribers {
		if v.Name == s.Name {
			return
		}
	}
	p.Subscribers = append(p.Subscribers, s)
}

func (p *Publisher) RemoveSubscriber(s Subscriber) {
	replace := make([]Subscriber, 0, len(p.Subscribers))
	for _, v := range p.Subscribers {
		if v.Name == s.Name {
			continue
		}
		replace = append(replace, v)
	}
	p.Subscribers = replace
}

type Subscriber struct {
	Name   string
	Writer io.Writer
}

func (s *Subscriber) Notify(out string) {
	fmt.Fprintf(s.Writer, "[%s]: %s\n", s.Name, out)
}
