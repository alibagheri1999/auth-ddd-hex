package utils

const (
	PasswordRegex      = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$`
	PasswordValidation = "Enter a valid password: At least one lowercase letter ([a-z]), At least one uppercase letter ([A-Z]), At least one digit ([0-9]), At least one special character (symbols like !@#$%^&*), Minimum length of 8 characters"
)
