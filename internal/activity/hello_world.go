package activity

import "context"

func (a *Activity) HelloWorldActivity(ctx context.Context, name string) (string, error) {
	return a.SayHello(name), nil
}
