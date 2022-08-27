package main

func ExampleCreateSimpleTransport() {
	t := CreateSimpleTransport("road")
	t.Deliver()

	t = CreateSimpleTransport("ship")
	t.Deliver()

	f := CreateAbstractFactory("road")
	f.CreateTransport().Deliver()

	f = CreateAbstractFactory("ship")
	f.CreateTransport().Deliver()

	// Output:
	// truck deliver
	// ship deliver
	// truck deliver
	// ship deliver
}
