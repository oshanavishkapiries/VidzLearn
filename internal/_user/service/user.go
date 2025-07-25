package service

import (
	"context"
	"strings"
	"time"

	"github.com/Cenzios/pf-backend/internal/_user/dto"
	"github.com/Cenzios/pf-backend/internal/models/mongo"
	"github.com/Cenzios/pf-backend/pkg/db"
	"github.com/Cenzios/pf-backend/pkg/response"
	"github.com/Cenzios/pf-backend/pkg/smtp"
	"github.com/Cenzios/pf-backend/pkg/utils"
)

func RegisterUserService(dto dto.RegisterUserDTO) (*mongo.User, error) {
	// 1. Check if user exists
	filter := map[string]interface{}{
		"email":     strings.ToLower(dto.Email),
		"is_active": true,
	}
	existingUser, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return nil, response.InternalServerException("Failed to check existing user", err)
	}
	if existingUser != nil {
		return nil, response.ConflictException("User already exists", nil)
	}

	// 2. Create new user
	now := time.Now().Format(time.RFC3339)

	user := &mongo.User{
		Email:             &dto.Email,
		SignUpVia:         dto.SignUpVia,
		GoogleID:          &dto.GoogleID,
		AppleID:           &dto.AppleID,
		PushID:            &dto.PushID,
		IpAddress:         &dto.IpAddress,
		DeviceID:          &dto.DeviceID,
		IsAccountVerified: true,
		IsActive:          true,
		CreatedAt:         &now,
		UpdatedAt:         &now,
	}
	if err := db.DB.InsertOne(context.Background(), "users", user); err != nil {
		return nil, response.InternalServerException("Failed to register user", err)
	}
	return user, nil
}

func RequestOtpService(dto dto.RequestOtpDTO) error {
	// 1. Check if user exists
	filter := map[string]interface{}{
		"email":     strings.ToLower(dto.Email),
		"is_active": true,
	}
	userDoc, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return response.InternalServerException("Failed to check existing user", err)
	}
	if userDoc == nil {
		return response.NotFoundException("User not found", nil)
	}

	// 2. Generate OTP and reference
	otp := utils.GenerateRandomOtp()
	otpRef := utils.GenerateOtpReference()
	now := time.Now().Format(time.RFC3339)

	// 2.1: Hash OTP and reference
	hashedOtp, err := utils.HashOtp(otp)
	if err != nil {
		return response.InternalServerException("Failed to hash OTP", err)
	}
	hashedRef, err := utils.HashOtpReference(otpRef)
	if err != nil {
		return response.InternalServerException("Failed to hash OTP reference", err)
	}

	// 3: Store in user table
	update := map[string]interface{}{
		"latest_otp":            hashedOtp,
		"latest_otp_reference":  hashedRef,
		"latest_otp_created_at": now,
		"updated_at":            now,
	}
	err = db.DB.UpdateOne(context.Background(), "users", filter, update)
	if err != nil {
		return response.InternalServerException("Failed to update user with OTP", err)
	}

	// 4. Send email (HTML)
	htmlBody, err := utils.SendOtpEmail(dto.Email, otp)
	if err != nil {
		return response.InternalServerException("Failed to generate OTP email body", err)
	}
	err = smtp.SendMail(dto.Email, "Your OTP Code", htmlBody)
	if err != nil {
		return response.InternalServerException("Failed to send OTP email", err)
	}

	return nil
}

func VerifyOtpService(dto dto.VerifyOtpDTO) error {
	// 1. Check if user exists
	filter := map[string]interface{}{
		"email":     strings.ToLower(dto.Email),
		"is_active": true,
	}
	userDoc, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return response.InternalServerException("Failed to check user", err)
	}
	if userDoc == nil {
		return response.NotFoundException("User not found", nil)
	}

	// 2. Extract stored OTP and reference hashes
	userMap, ok := userDoc.(map[string]interface{})
	if !ok {
		return response.InternalServerException("User data format error", nil)
	}
	storedOtpHash, _ := userMap["latest_otp"].(string)
	storedRefHash, _ := userMap["latest_otp_reference"].(string)
	if storedOtpHash == "" || storedRefHash == "" {
		return response.BadRequestException("No OTP or reference found for user", nil)
	}

	// 3. Compare OTP and reference
	if !utils.CompareOtpHash(dto.Otp, storedOtpHash) || !utils.CompareOtpReferenceHash(dto.OtpReference, storedRefHash) {
		return response.BadRequestException("Invalid OTP or reference", nil)
	}

	// 4. Remove OTP and reference from user record
	now := time.Now().Format(time.RFC3339)
	unset := map[string]interface{}{
		"latest_otp":           "",
		"latest_otp_reference": "",
		"updated_at":           now,
	}
	err = db.DB.UpdateOne(context.Background(), "users", filter, unset)
	if err != nil {
		return response.InternalServerException("Failed to remove OTP from user", err)
	}

	return nil
}

func SetPasswordService(dto dto.SetPasswordDTO) error {
	// 1. Check if passwords match
	if dto.Password != dto.ConfirmPassword {
		return response.BadRequestException("Passwords do not match", nil)
	}

	// 2. Check if user exists
	filter := map[string]interface{}{
		"_id":       dto.UserId,
		"is_active": true,
	}
	userDoc, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return response.InternalServerException("Failed to check user", err)
	}
	if userDoc == nil {
		return response.NotFoundException("User not found", nil)
	}

	// 3. Ensure password is not already set
	userMap, ok := userDoc.(map[string]interface{})
	if !ok {
		return response.InternalServerException("User data format error", nil)
	}
	if userMap["latest_password"] != nil && userMap["latest_password"].(string) != "" {
		return response.BadRequestException("Password already set", nil)
	}

	// 4. Hash and set password
	hashedPassword, err := utils.HashOtp(dto.Password) // bcrypt is fine
	if err != nil {
		return response.InternalServerException("Failed to hash password", err)
	}
	now := time.Now().Format(time.RFC3339)
	update := map[string]interface{}{
		"latest_password":            hashedPassword,
		"latest_password_created_at": now,
		"last_password":              hashedPassword,
		"updated_at":                 now,
	}
	err = db.DB.UpdateOne(context.Background(), "users", filter, update)
	if err != nil {
		return response.InternalServerException("Failed to set password", err)
	}

	return nil
}

func CompleteProfileService(dto dto.CompleteProfileDTO) error {
	// 1. Check if user exists
	filter := map[string]interface{}{
		"_id":       dto.UserId,
		"is_active": true,
	}
	userDoc, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return response.InternalServerException("Failed to check user", err)
	}
	if userDoc == nil {
		return response.NotFoundException("User not found", nil)
	}

	// 2. Build update map only with non-empty fields
	update := map[string]interface{}{}
	if dto.FullName != "" {
		update["full_name"] = dto.FullName
	}
	if dto.UserName != "" {
		update["user_name"] = dto.UserName
	}
	if dto.FirstName != "" {
		update["first_name"] = dto.FirstName
	}
	if dto.LastName != "" {
		update["last_name"] = dto.LastName
	}
	if dto.PhoneNumber != "" {
		update["phone_number"] = dto.PhoneNumber
	}
	if dto.CountryCode != "" {
		update["country_code"] = dto.CountryCode
	}
	if dto.DialCode != "" {
		update["dial_code"] = dto.DialCode
	}
	if dto.ProfilePictureUrl != "" {
		update["profile_picture_url"] = dto.ProfilePictureUrl
	}

	if len(update) == 0 {
		return response.BadRequestException("No profile fields to update", nil)
	}

	now := time.Now().Format(time.RFC3339)
	update["updated_at"] = now
	update["is_profile_completed"] = true

	err = db.DB.UpdateOne(context.Background(), "users", filter, update)
	if err != nil {
		return response.InternalServerException("Failed to update profile", err)
	}

	return nil
}

func LoginUserService(dto dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	// 1. Find user by email
	filter := map[string]interface{}{
		"email":     strings.ToLower(dto.Email),
		"is_active": true,
	}
	userDoc, err := db.DB.FindOne(context.Background(), "users", filter)
	if err != nil {
		return nil, response.InternalServerException("Failed to check user", err)
	}
	if userDoc == nil {
		return nil, response.NotFoundException("User not found", nil)
	}

	userMap, ok := userDoc.(map[string]interface{})
	if !ok {
		return nil, response.InternalServerException("User data format error", nil)
	}
	storedPassword, _ := userMap["latest_password"].(string)
	if storedPassword == "" {
		return nil, response.BadRequestException("Password not set for user", nil)
	}

	// 2. Compare password (assuming utils.CompareOtpHash works for password)
	if !utils.CompareOtpHash(dto.Password, storedPassword) {
		return nil, response.BadRequestException("Invalid credentials", nil)
	}

	// 3. Issue JWT
	userID, _ := userMap["_id"].(string)
	email, _ := userMap["email"].(string)
	jwtPayload := utils.JWTPayload{
		Sub:   userID,
		Exp:   time.Now().Add(24 * time.Hour).Unix(), // 24h expiry
		Email: email,
	}
	token, err := utils.SignJWT(jwtPayload)
	if err != nil {
		return nil, response.InternalServerException("Failed to sign JWT", err)
	}

	return &dto.LoginResponseDTO{
		Token: token,
		User: dto.UserInfo{
			ID:    userID,
			Email: email,
		},
	}, nil

}
