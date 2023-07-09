package mediator

import (
	"fmt"
	"io"
)

type User struct {
	room    ChatRoom
	Name    string
	wrapper func(string) string
}

func (u *User) SendMessage(msg string) {
	u.room.SendMessage(u.wrapper(msg))
}

func NewUser(name string, room ChatRoom) *User {
	user := &User{
		Name: name,
		room: room,
	}
	f := func(msg string) string {
		return fmt.Sprintf("[%s]: %s", user.Name, msg)
	}
	user.wrapper = f
	return user
}

type ChatRoom interface {
	SendMessage(string)
}

type Room struct {
	io.Writer
}

func (r *Room) SendMessage(msg string) {
	fmt.Fprintln(r.Writer, msg)
}
