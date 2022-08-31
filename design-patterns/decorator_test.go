package main

func ExampleNotifier() {
	stack := EmailNotifier(nil)
	stack = QQNotifier(stack)
	stack = WeChatNotifier(stack)

	stack.Send("test")

	stack = QQNotifier(nil)
	stack.Send("test1")

	// Output:
	// test: Alert by Email!
	// test: Alert by QQ!
	// test: Alert by WeChat!
	// test1: Alert by QQ!
}
