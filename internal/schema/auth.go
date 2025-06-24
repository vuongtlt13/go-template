package schema

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

// LoginResponse represents login response with JWT token
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterResponse represents registration response
type RegisterResponse struct {
	Message string `json:"message"`
}
