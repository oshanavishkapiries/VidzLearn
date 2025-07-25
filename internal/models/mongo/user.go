package mongo

type User struct {
	ID                      string  `bson:"_id,omitempty" json:"id,omitempty"`
	UserCode                string  `bson:"user_code" json:"user_code"`
	FullName                *string `bson:"full_name" json:"full_name"`
	UserName                *string `bson:"user_name" json:"user_name"`
	FirstName               *string `bson:"first_name" json:"first_name"`
	LastName                *string `bson:"last_name" json:"last_name"`
	Email                   *string `bson:"email" json:"email"`
	PhoneNumber             *string `bson:"phone_number" json:"phone_number"`
	CountryCode             *string `bson:"country_code" json:"country_code"`
	DialCode                *string `bson:"dial_code" json:"dial_code"`
	IsEmailVerified         bool    `bson:"is_email_verified" json:"is_email_verified"`
	IsPhoneNumberVerified   bool    `bson:"is_phone_number_verified" json:"is_phone_number_verified"`
	IsAccountVerified       bool    `bson:"is_account_verified" json:"is_account_verified"`
	IsProfileCompleted      bool    `bson:"is_profile_completed" json:"is_profile_completed"`
	LatestPassword          *string `bson:"latest_password" json:"latest_password"`
	LastPassword            *string `bson:"last_password" json:"last_password"`
	LatestPasswordCreatedAt *string `bson:"latest_password_created_at" json:"latest_password_created_at"`
	LatestOtp               *string `bson:"latest_otp" json:"latest_otp"`
	LatestOtpReference      *string `bson:"latest_otp_reference" json:"latest_otp_reference"`
	LatestOtpCreatedAt      *string `bson:"latest_otp_created_at" json:"latest_otp_created_at"`
	SignUpVia               int     `bson:"sign_up_via" json:"sign_up_via"`
	AppleID                 *string `bson:"apple_id" json:"apple_id"`
	GoogleID                *string `bson:"google_id" json:"google_id"`
	PushID                  *string `bson:"push_id" json:"push_id"`
	Token                   *string `bson:"token" json:"token"`
	DeviceID                *string `bson:"device_id" json:"device_id"`
	IpAddress               *string `bson:"ip_address" json:"ip_address"`
	LastLoginAt             *string `bson:"last_login_at" json:"last_login_at"`
	ProfilePictureUrl       *string `bson:"profile_picture_url" json:"profile_picture_url"`
	IsActive                bool    `bson:"is_active" json:"is_active"`
	CreatedAt               *string `bson:"created_at,omitempty" json:"created_at,omitempty"` // set automatically on insert
	UpdatedAt               *string `bson:"updated_at,omitempty" json:"updated_at,omitempty"` // set automatically on insert/update
	CreatedBy               *string `bson:"created_by" json:"created_by"`
	UpdatedBy               *string `bson:"updated_by" json:"updated_by"`
	RoleID                  *string `bson:"role_id" json:"role_id"`
}
