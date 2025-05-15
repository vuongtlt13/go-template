// This file is under development for the new admin gRPC services.
// Remove this comment when ready for production.

package admin

import (
	"context"
	"yourapp/internal/model"
	"yourapp/internal/service"
	"yourapp/pb/admin"

	"connectrpc.com/connect"
)

// PermissionHandler implements the Connect gRPC interface for admin permission service
// It should satisfy adminconnect.PermissionServiceHandler

type PermissionHandler struct {
	svc service.PermissionService
}

// NewPermissionHandler creates a new permission handler using singleton instances
func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{
		svc: service.NewPermissionService(),
	}
}

// CreatePermission creates a new permission
func (h *PermissionHandler) CreatePermission(ctx context.Context, req *connect.Request[admin.CreatePermissionRequest]) (*connect.Response[admin.PermissionResponse], error) {
	permission := &model.Permission{
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Service:     req.Msg.Service,
		Method:      req.Msg.Method,
	}

	err := h.svc.CreatePermission(ctx, permission)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&admin.PermissionResponse{
		Permission: &admin.Permission{
			Id:          permission.ID,
			Code:        permission.Code,
			Name:        permission.Name,
			Description: permission.Description,
			Service:     permission.Service,
			Method:      permission.Method,
		},
	}), nil
}

// GetPermission retrieves a permission by ID
func (h *PermissionHandler) GetPermission(ctx context.Context, req *connect.Request[admin.GetPermissionRequest]) (*connect.Response[admin.PermissionResponse], error) {
	permission, err := h.svc.GetPermissionByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.PermissionResponse{
		Permission: &admin.Permission{
			Id:          permission.ID,
			Code:        permission.Code,
			Name:        permission.Name,
			Description: permission.Description,
			Service:     permission.Service,
			Method:      permission.Method,
		},
	}), nil
}

// UpdatePermission updates a permission
func (h *PermissionHandler) UpdatePermission(ctx context.Context, req *connect.Request[admin.UpdatePermissionRequest]) (*connect.Response[admin.PermissionResponse], error) {
	permission := &model.Permission{
		ID:          req.Msg.Id,
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Service:     req.Msg.Service,
		Method:      req.Msg.Method,
	}

	err := h.svc.UpdatePermission(ctx, permission)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.PermissionResponse{
		Permission: &admin.Permission{
			Id:          permission.ID,
			Code:        permission.Code,
			Name:        permission.Name,
			Description: permission.Description,
			Service:     permission.Service,
			Method:      permission.Method,
		},
	}), nil
}

// DeletePermission deletes a permission
func (h *PermissionHandler) DeletePermission(ctx context.Context, req *connect.Request[admin.DeletePermissionRequest]) (*connect.Response[admin.DeletePermissionResponse], error) {
	err := h.svc.DeletePermission(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.DeletePermissionResponse{
		Success: true,
	}), nil
}

// ListPermissions retrieves all permissions
func (h *PermissionHandler) ListPermissions(ctx context.Context, req *connect.Request[admin.ListPermissionsRequest]) (*connect.Response[admin.ListPermissionsResponse], error) {
	permissions, total, err := h.svc.ListPermissions(ctx, int(req.Msg.Page), int(req.Msg.PageSize))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var pbPermissions []*admin.Permission
	for _, p := range permissions {
		pbPermissions = append(pbPermissions, &admin.Permission{
			Id:          p.ID,
			Code:        p.Code,
			Name:        p.Name,
			Description: p.Description,
			Service:     p.Service,
			Method:      p.Method,
		})
	}

	return connect.NewResponse(&admin.ListPermissionsResponse{
		Permissions: pbPermissions,
		Total:       uint32(total),
	}), nil
}
