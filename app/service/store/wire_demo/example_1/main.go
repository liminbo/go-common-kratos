package example_1

type Message struct {
	msg string
}

type Greeter struct {
	message Message
}

type Event struct {
	greeter Greeter
}

func newMessage(msg string) Message{
	return Message{
		msg:msg,
	}
}

func newGreeter(message Message) Greeter{
	return Greeter{message:message}
}

func newEvent(greeter Greeter) Event{
	return Event{greeter:Greeter{}}
}

func main() {

}
