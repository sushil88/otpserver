package otp

import "errors"

func verifyPhoneNumber(phoneNumber string) error {
	return nil
}

func generateOTP() (int, int, error)  {
  // insert one random 4 digit OTP in sql
  return 1, 4444, nil
}

func sendOTP(otp int, ph string) error {
	// sending otp here
	return nil
}
func GenerateOTP(phoneNumber string) (int, error)  {
	// Verify phone number
	err := verifyPhoneNumber(phoneNumber)
	if err != nil {
		return 0, errors.New("Invalid phone number")
	}
	// generate ID
	otpID, otp, err := generateOTP()
	if err != nil {
		return 0, errors.New("Failed to generate OTP")
	}
	// Send OTP
	err = sendOTP(otp, phoneNumber)
    if err != nil {
    	return 0, errors.New("Invalid send otp")
    }

	return otpID, err
}
