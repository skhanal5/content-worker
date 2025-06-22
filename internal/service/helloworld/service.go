package helloworld

import "fmt"

func SayHello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Name cannot be empty")
	}
	return "Hello, " + name, nil
}
