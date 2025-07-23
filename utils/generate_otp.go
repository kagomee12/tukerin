package utils

import(
	"math/rand"
	"strconv"
)

// GenerateOTP generates a one-time password (OTP) and sends it via email
func GenerateOTP(email string) (string, error) {
	otp := strconv.Itoa(rand.Intn(900000) + 100000) // Generate a 6-digit OTP

	return otp, nil
}