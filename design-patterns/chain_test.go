package main

func ExampleChain() {
	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)

	// Output:
	// Reception registering patient
	// Doctor checking patient
	// Medical giving medicine to patient
	// Cashier getting money from patient patient
}
