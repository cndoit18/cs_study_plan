package builder

import "fmt"

func ExampleCarBuilder() {
	builder := CarBuilder{}
	fmt.Println(builder.SetEngine("honda").Build())
	fmt.Println(builder.SetEngine("honda").Reset().Build())
	fmt.Println(builder.SetEngine("honda").SetSeats(1).Build())
	fmt.Println(builder.SetEngine("mercedes").SetGPS("beidou").SetSeats(1).Build())
	// Output:
	// This is a car with honda's engine.
	// This is a car.
	// This is a car with 1 seats and honda's engine.
	// This is a car with 1 seats and mercedes's engine and beidou's GPS.
}
