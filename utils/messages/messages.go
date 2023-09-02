package messages

/*
*
Error messages
*/
const (
	FailToParseUserID    = "Failed to parse user ID"
	InvalidRequestData   = "Failed to parse request data"
	FailToCreateUser     = "Failed to create new user"
	FailToDeleteUser     = "Failed to delete a user"
	FailToUpdateUser     = "Failed to update user"
	UnsuccessfulLogin    = "Login attempt failed"
	FailToRegister       = "Registration attempt failed"
	JWTClaimError        = "JWT claim error"
	AllFieldMustBeFilled = "All field must be filled"
	AlreadyLoggedIn      = "Already logged in"
	TooManyRequest       = "Too many request"
)

/*
*
Success messages
*/
const (
	SuccessfulUserObtain   = "Successfully obtained user data"
	SuccessfulUserCreation = "Successfully created a new user"
	SuccessfulUserDeletion = "Successfully deleted a new user"
	SuccessfulUserUpdate   = "Successfully updated a user"
	SuccessfulLogin        = "Successfully logged in"
	SuccessfulRegister     = "Successfully registered a user"
)
