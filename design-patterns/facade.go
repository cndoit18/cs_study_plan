package main

import "fmt"

type ImageFacade struct {
	a *Account
	p *PictureHost
}

func (i *ImageFacade) Push() {
	i.a.Login()
	i.p.PushImage()
}

type Account struct{}

func (a *Account) Login() {
	fmt.Println("login")
}

type PictureHost struct{}

func (p *PictureHost) PushImage() {
	fmt.Println("pushImage")
}
