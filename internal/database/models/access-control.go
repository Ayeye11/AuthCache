package models

type Permission struct {
	Role     string `gorm:"column:role"`
	Category string `gorm:"column:category"`
	Action   string `gorm:"column:action"`
}

// TABLE: ac_roles
type AcRole struct {
	ID   uint   `gorm:"column:id"`
	Role string `gorm:"column:role"`
}

func (AcRole) TableName() string {
	return "ac_roles"
}

// TABLE: ac_categories
type AcCategory struct {
	ID       uint   `gorm:"column:id"`
	Category string `gorm:"column:category"`
}

func (AcCategory) TableName() string {
	return "ac_categories"
}

// TABLE: ac_actions
type AcAction struct {
	ID     uint   `gorm:"column:id"`
	Action string `gorm:"column:action"`
}

func (AcAction) TableName() string {
	return "ac_actions"
}

// TABLE: ac_relations
type AcRelation struct {
	RoleID     uint `gorm:"column:role_id"`
	CategoryID uint `gorm:"column:category_id"`
	ActionID   uint `gorm:"column:action_id"`
}

func (AcRelation) TableName() string {
	return "ac_relations"
}
