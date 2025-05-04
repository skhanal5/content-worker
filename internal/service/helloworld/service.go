package helloworld

type HelloWorldManager interface {
	SayHello(name string) string
}

type HelloWorldService struct{}

func (h HelloWorldService) SayHello(name string) string {
	return "Hello, " + name
}
