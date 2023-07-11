package observer_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/observer"
)

func ExampleObserver() {
	publisher := observer.Publisher{Subscribers: []observer.Subscriber{}}

	a := observer.Subscriber{
		Name:   "a",
		Writer: os.Stdout,
	}

	b := observer.Subscriber{
		Name:   "b",
		Writer: os.Stdout,
	}

	publisher.AddSubscriber(a)
	publisher.Notify("hello")

	publisher.AddSubscriber(b)
	publisher.Notify("world")
	// Output:
	// [a]: hello
	// [a]: world
	// [b]: world
}
