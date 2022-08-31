package main

import "fmt"

type Notifier interface {
	Send(string)
}

type NotifierFunc func(string)

func (f NotifierFunc) Send(msg string) {
	f(msg)
}

var QQNotifier = func(f NotifierFunc) NotifierFunc {
	return func(msg string) {
		if f != nil {
			f(msg)
		}
		fmt.Printf("%s: Alert by QQ!\n", msg)
	}
}

var WeChatNotifier = func(f NotifierFunc) NotifierFunc {
	return func(msg string) {
		if f != nil {
			f(msg)
		}
		fmt.Printf("%s: Alert by WeChat!\n", msg)
	}
}

var EmailNotifier = func(f NotifierFunc) NotifierFunc {
	return func(msg string) {
		if f != nil {
			f(msg)
		}
		fmt.Printf("%s: Alert by Email!\n", msg)
	}
}
