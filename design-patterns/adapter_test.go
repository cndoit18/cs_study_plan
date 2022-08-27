package main

func ExampleComputer() {
	c := &Client{}
	var mac Computer = &Mac{}
	c.InsertLightningConnectorIntoComputer(mac)

	// glogr就是一个很典型的适配器模式
	var win Computer = &WindowsAdapter{&Windows{}}
	c.InsertLightningConnectorIntoComputer(win)

	// Output:
	// Client inserts Lightning connector into computer.
	// Lightning connector is plugged into mac machine.
	// Client inserts Lightning connector into computer.
	// Adapter converts Lightning signal to USB.
	// USB connector is plugged into windows machine.
}
