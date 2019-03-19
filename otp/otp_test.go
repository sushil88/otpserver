package otp

import "testing"

func Test_generateOTP(t *testing.T) {
	tests := []struct {
		PhoneNumber    string
		OtpID    int
		OTP   int
		wantErr bool
	}{
		struct {
			PhoneNumber string
			OtpID       int
			OTP         int
			wantErr     bool
		}{PhoneNumber: "+919191991919", OtpID: 1, OTP: 4444, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.PhoneNumber, func(t *testing.T) {
			optID, OTP, err := generateOTP()
			if (err != nil) != tt.wantErr {
				t.Errorf("generateOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if optID != tt.OtpID {
				t.Errorf("generateOTP() got optID = %v, want %v", optID, tt.OtpID)
			}
			if OTP != tt.OTP {
				t.Errorf("generateOTP() got OTP = %v, want %v", OTP, tt.OTP)
			}
			if err != nil && tt.wantErr {
				t.Errorf("generateOTP() got err = %v, want %v", err, nil)
			}
		})
	}
}
