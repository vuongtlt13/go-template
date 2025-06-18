package admin

// DashboardResponse represents the admin dashboard data
type DashboardResponse struct {
	TotalUsers     int64 `json:"total_users"`
	ActiveUsers    int64 `json:"active_users"`
	TotalSessions  int64 `json:"total_sessions"`
	ActiveSessions int64 `json:"active_sessions"`
}

// AdminUsersResponse represents a paginated list of users
type AdminUsersResponse struct {
	Users      []AdminUserResponse `json:"users"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	Limit      int                 `json:"limit"`
	TotalPages int                 `json:"total_pages"`
}

// AdminUserResponse represents a user in admin view
type AdminUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Password string `json:"password" validate:"omitempty,min=6"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
	Status   string `json:"status" validate:"omitempty,oneof=active inactive"`
}

// SettingsResponse represents the system settings
type SettingsResponse struct {
	SiteName         string `json:"site_name"`
	SiteDescription  string `json:"site_description"`
	MaintenanceMode  bool   `json:"maintenance_mode"`
	MaxLoginAttempts int    `json:"max_login_attempts"`
	SessionTimeout   int    `json:"session_timeout"`
}

// UpdateSettingsRequest represents the request to update system settings
type UpdateSettingsRequest struct {
	SiteName         string `json:"site_name" validate:"required"`
	SiteDescription  string `json:"site_description"`
	MaintenanceMode  bool   `json:"maintenance_mode"`
	MaxLoginAttempts int    `json:"max_login_attempts" validate:"required,min=1"`
	SessionTimeout   int    `json:"session_timeout" validate:"required,min=5"`
}

// LogsResponse represents a paginated list of logs
type LogsResponse struct {
	Logs       []LogResponse `json:"logs"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
	TotalPages int           `json:"total_pages"`
}

// LogResponse represents a log entry
type LogResponse struct {
	ID        string `json:"id"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
}
