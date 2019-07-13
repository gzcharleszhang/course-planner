package services

import "context"

type serviceHandler = func(ctx context.Context, req map[interface{}]interface{}) error

func newServiceHandler(h serviceHandler) serviceHandler {
	return serviceHandler(func)
}
