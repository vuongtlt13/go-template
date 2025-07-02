package response

// Standard API Response wrapper
type APIResponse[T any] struct {
	Success   bool   `json:"success" description:"Request success status"`
	Data      T      `json:"data,omitempty" description:"Response data"`
	Message   string `json:"message,omitempty" description:"Response message"`
	ErrorCode int    `json:"errorCode,omitempty" description:"Error code, only present if error"`
	Stack     string `json:"stack,omitempty" description:"Stack trace, only present if error"`
}
