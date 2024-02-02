package dom


// User model to stores user details
type User struct {
	UserName string
	Email    string
	Password string
	Role     string
}

// Login model for storing logging in credentials
type Login struct {
	Email    string
	Password string
}
