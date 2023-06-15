package decorator

import (
	"fmt"
	"io"
)

type Notifiter interface {
	Send(out io.Writer, meesage string)
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(out io.Writer, message string) {
	fmt.Fprintf(out, "From Email: %s\n", message)
}

type QQDecorator struct {
	Notifiter Notifiter
}

func (q *QQDecorator) Send(out io.Writer, message string) {
	fmt.Fprintf(out, "From QQ: %s\n", message)
	q.Notifiter.Send(out, message)
}

type WeChatDecorator struct {
	Notifiter Notifiter
}

func (w *WeChatDecorator) Send(out io.Writer, message string) {
	fmt.Fprintf(out, "From WeChat: %s\n", message)
	w.Notifiter.Send(out, message)
}
