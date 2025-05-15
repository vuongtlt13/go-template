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

// RoleHandler implements the Connect gRPC interface for admin role service
// It should satisfy adminconnect.RoleServiceHandler

type RoleHandler struct {
	roleService service.RoleService
}

// NewRoleHandler creates a new role handler using singleton instances
func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// CreateRole creates a new role
func (h *RoleHandler) CreateRole(ctx context.Context, req *connect.Request[admin.CreateRoleRequest]) (*connect.Response[admin.RoleResponse], error) {
	role := &model.Role{
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
	}

	err := h.roleService.CreateRole(ctx, role)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&admin.RoleResponse{
		Role: &admin.Role{
			Id:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Description: role.Description,
		},
	}), nil
}

// GetRole retrieves a role by ID
func (h *RoleHandler) GetRole(ctx context.Context, req *connect.Request[admin.GetRoleRequest]) (*connect.Response[admin.RoleResponse], error) {
	role, err := h.roleService.GetRoleByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.RoleResponse{
		Role: &admin.Role{
			Id:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Description: role.Description,
		},
	}), nil
}

// UpdateRole updates a role
func (h *RoleHandler) UpdateRole(ctx context.Context, req *connect.Request[admin.UpdateRoleRequest]) (*connect.Response[admin.RoleResponse], error) {
	role := &model.Role{
		ID:          req.Msg.Id,
		Code:        req.Msg.Code,
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
	}

	err := h.roleService.UpdateRole(ctx, role)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.RoleResponse{
		Role: &admin.Role{
			Id:          role.ID,
			Code:        role.Code,
			Name:        role.Name,
			Description: role.Description,
		},
	}), nil
}

// DeleteRole deletes a role
func (h *RoleHandler) DeleteRole(ctx context.Context, req *connect.Request[admin.DeleteRoleRequest]) (*connect.Response[admin.DeleteRoleResponse], error) {
	err := h.roleService.DeleteRole(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}

	return connect.NewResponse(&admin.DeleteRoleResponse{
		Success: true,
	}), nil
}

// ListRoles retrieves all roles
func (h *RoleHandler) ListRoles(ctx context.Context, req *connect.Request[admin.ListRolesRequest]) (*connect.Response[admin.ListRolesResponse], error) {
	roles, total, err := h.roleService.ListRoles(ctx, int(req.Msg.Page), int(req.Msg.PageSize))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var pbRoles []*admin.Role
	for _, r := range roles {
		pbRoles = append(pbRoles, &admin.Role{
			Id:          r.ID,
			Code:        r.Code,
			Name:        r.Name,
			Description: r.Description,
		})
	}

	return connect.NewResponse(&admin.ListRolesResponse{
		Roles: pbRoles,
		Total: uint32(total),
	}), nil
}
