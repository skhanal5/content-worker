package activity

import "context"

func HelloWorldActivity(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}