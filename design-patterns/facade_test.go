package main

func ExampleImageFacade() {
	f := ImageFacade{
		a: &Account{},
		p: &PictureHost{},
	}

	// 将子系统封装起来
	f.Push()
	// Output:
	// login
	// pushImage
}
