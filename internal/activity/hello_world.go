package activity

import (
	"clip-farmer-workflow/internal/service/helloworld"
	"context"
)

func (a *Activity) HelloWorldActivity(ctx context.Context, name string) (string, error) {
	result, err := helloworld.SayHello(name)
	if err != nil {
		return "", err
	}
	return result, nil
}
