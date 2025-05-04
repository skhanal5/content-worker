package helloworld

import "fmt"

type HelloWorldManager interface {
	SayHello(name string) (string, error)
}

type HelloWorldService struct{}

func (h HelloWorldService) SayHello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Name cannot be empty")
	}
	return "Hello, " + name, nil
}
