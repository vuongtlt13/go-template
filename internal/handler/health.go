package handler

import (
	"connectrpc.com/connect"
	"context"
	"yourapp/pb/health"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(ctx context.Context, c *connect.Request[health.HealthCheckRequest]) (*connect.Response[health.HealthCheckResponse], error) {
	//TODO implement me
	panic("implement me")
}
