package schema

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `parse:"body:email,required" validate:"required,email" description:"User email address"`
	Password string `parse:"body:password,required" validate:"required,min=8" description:"User password"`
}

// RegisterRequest represents registration data
type RegisterRequest struct {
	Email    string `parse:"body:email,required" validate:"required,email" description:"User email address"`
	Password string `parse:"body:password,required" validate:"required,min=8" description:"User password"`
}

// LoginData represents login response data
type LoginData struct {
	Token string `json:"token" description:"JWT authentication token"`
}

// RegisterData represents registration response data
type RegisterData struct {
	Message string `json:"message" description:"Registration confirmation message"`
}
