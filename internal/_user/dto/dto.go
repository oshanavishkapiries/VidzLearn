package dto

type RegisterUserDTO struct {
	Email     string `json:"email"`
	SignUpVia int    `json:"signup_via"`
	GoogleID  string `json:"google_id,omitempty"`
	AppleID   string `json:"apple_id,omitempty"`
	PushID    string `json:"push_id,omitempty"`
	IpAddress string `json:"ip_address,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
}

type RequestOtpDTO struct {
	Email string `json:"email"`
}

type VerifyOtpDTO struct {
	Email        string `json:"email"`
	Otp          string `json:"otp"`
	OtpReference string `json:"otp_reference"`
}

type SetPasswordDTO struct {
	UserId          string `json:"user_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CompleteProfileDTO struct {
	UserId            string `json:"user_id"`
	FullName          string `json:"full_name"`
	UserName          string `json:"username"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	CountryCode       string `json:"country_code"`
	DialCode          string `json:"dial_code"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
