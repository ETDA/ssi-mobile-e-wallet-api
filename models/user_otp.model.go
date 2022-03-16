package models

import "time"

type UserOTP struct {
	ID         string     `json:"id" gorm:"column:id"`
	OTPNumber  string     `json:"otp_number" gorm:"column:otp_number"`
	UserID     string     `json:"user_id" gorm:"column:user_id"`
	VerifiedAt *time.Time `json:"verified_at" gorm:"column:verified_at"`
	ExpiredAt  *time.Time `json:"expired_at" gorm:"column:expired_at"`
	RevokedAt  *time.Time `json:"revoked_at" gorm:"column:revoked_at"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (r UserOTP) TableName() string {
	return "user_otps"
}
