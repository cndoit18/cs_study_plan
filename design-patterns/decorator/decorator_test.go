package decorator_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/decorator"
)

func ExampleDecorator() {
	var notifier decorator.Notifiter = &decorator.EmailNotifier{}
	notifier = &decorator.QQDecorator{Notifiter: notifier}
	notifier.Send(os.Stdout, "hello")
	notifier = &decorator.WeChatDecorator{Notifiter: notifier}
	notifier.Send(os.Stdout, "world")
	// Output:
	// From QQ: hello
	// From Email: hello
	// From WeChat: world
	// From QQ: world
	// From Email: world
}
