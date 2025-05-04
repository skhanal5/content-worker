package helloworld

type HelloWorldManager interface {
	SayHello(name string) string
}

func SayHello(name string) (string) {
	return "Hello, " + name
}
