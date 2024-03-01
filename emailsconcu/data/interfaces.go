package data

// UserInterface is the interface for the user type. In order
// to satisfy this interface, all specified methods must be implemented.
// We do this so we can test things easily. Both data.User and data.UserTest
// implement this interface.
type UserInterface interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(id int) (*User, error)
	Update(user User) error
	// Delete() error
	DeleteByID(id int) error
	Insert(user User) (int, error)
	ResetPassword(password string) error
	PasswordMatches(plainText string) (bool, error)
}

// PlanInterface is the type for the plan type. Both data.Plan and data.PlanTest
// implement this interface.
type PlanInterface interface {
	GetAll() ([]*Plan, error)
	GetOne(id int) (*Plan, error)
	SubscribeUserToPlan(user User, plan Plan) error
	AmountForDisplay() string
}