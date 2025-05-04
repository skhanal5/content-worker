package activity

import "context"

func (a *Activity) HelloWorldActivity(ctx context.Context, name string) (string, error) {
	result, err := a.SayHello(name)
	if err != nil {
		return "", err
	}
	return result, nil
}
