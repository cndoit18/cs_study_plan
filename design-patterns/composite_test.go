package main

func ExampleComponent() {
	c := &Folder{name: "root"}
	c.add(&File{name: "text.txt"})
	subFolder := &Folder{name: "tmp"}
	subFolder.add(&File{name: "tmp.txt"})
	c.add(subFolder)

	c.Search("hello")
	// Output:
	// Searching for keyword hello in Folder root
	// Searching for keyword hello in file text.txt
	// Searching for keyword hello in Folder tmp
	// Searching for keyword hello in file tmp.txt
}
