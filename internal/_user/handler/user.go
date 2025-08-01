package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Cenzios/pf-backend/internal/_user/dto"
	userService "github.com/Cenzios/pf-backend/internal/_user/service"
	"github.com/Cenzios/pf-backend/pkg/response"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var req dto.RegisterUserDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	// Validation
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		response.BadRequest(w, "Email is required")
		return
	}
	if req.SignUpVia != 1 && req.SignUpVia != 2 && req.SignUpVia != 3 {
		response.BadRequest(w, "signup_via must be 1 (App), 2 (Google), or 3 (Apple)")
		return
	}
	if req.SignUpVia == 2 && strings.TrimSpace(req.GoogleID) == "" {
		response.BadRequest(w, "google_id is required for Google signup")
		return
	}
	if req.SignUpVia == 3 && strings.TrimSpace(req.AppleID) == "" {
		response.BadRequest(w, "apple_id is required for Apple signup")
		return
	}

	user, err := userService.RegisterUserService(dto.RegisterUserDTO{
		Email:     req.Email,
		SignUpVia: req.SignUpVia,
		GoogleID:  req.GoogleID,
		AppleID:   req.AppleID,
		PushID:    req.PushID,
		IpAddress: req.IpAddress,
		DeviceID:  req.DeviceID,
	})
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, user, "User registered successfully")
}

func RequestOtp(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestOtpDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" {
		response.BadRequest(w, "Email is required")
		return
	}

	err := userService.RequestOtpService(req)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, nil, "OTP sent successfully")
}

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyOtpDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Otp = strings.TrimSpace(req.Otp)
	req.OtpReference = strings.TrimSpace(req.OtpReference)
	if req.Email == "" || req.Otp == "" || req.OtpReference == "" {
		response.BadRequest(w, "Email, OTP, and OTP reference are required")
		return
	}

	err := userService.VerifyOtpService(req)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, nil, "OTP verified successfully")
}

func SetPassword(w http.ResponseWriter, r *http.Request) {
	var req dto.SetPasswordDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	req.UserId = strings.TrimSpace(req.UserId)
	req.Password = strings.TrimSpace(req.Password)
	req.ConfirmPassword = strings.TrimSpace(req.ConfirmPassword)
	if req.UserId == "" || req.Password == "" || req.ConfirmPassword == "" {
		response.BadRequest(w, "User ID, password, and confirm password are required")
		return
	}

	err := userService.SetPasswordService(req)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, nil, "Password set successfully")
}

func CompleteProfile(w http.ResponseWriter, r *http.Request) {
	var req dto.CompleteProfileDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	if req.UserId == "" {
		response.BadRequest(w, "User ID is required")
		return
	}

	err := userService.CompleteProfileService(req)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, nil, "Profile updated successfully")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" || req.Password == "" {
		response.BadRequest(w, "Email and password are required")
		return
	}

	loginResp, err := userService.LoginUserService(req)
	if err != nil {
		response.HandleError(w, err)
		return
	}
	response.Success(w, loginResp, "Login successful")
}
