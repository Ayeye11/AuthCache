package types

// Types
type Permission struct {
	Category string `json:"category,omitempty"`
	Action   string `json:"action,omitempty"`
}

type Role struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
