package Models

import "time"

// FrontEndUser struct
type FrontEndUser struct {
	ID                        uint   `gorm:"primary_key"`
	FirstName                 string `gorm:"type:varchar(255)"`
	LastName                  string `gorm:"type:varchar(255)"`
	Email                     string `gorm:"type:varchar(255)"`
	EmailStatus               int
	EmailDatetimeStatus       time.Time
	PhoneNumber               string `gorm:"type:varchar(255)"`
	PhoneNumberStatus         int
	PhoneNumberDatetimeStatus string `gorm:"type:varchar(255)"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	EncryptedPassword         string `gorm:"type:varchar(255)"`
	AuthenticationToken       string `gorm:"type:varchar(255)"`
	PushToken                 string `gorm:"type:varchar(255)"`
	ResetPasswordToken        string `gorm:"type:varchar(255)"`
	ResetPasswordSentAt       time.Time
	PushTokenProd             bool
	APIKey                    string `gorm:"type:varchar(255)"`
	PhoneToken                string `gorm:"type:varchar(255)"`
	SaasCompanyID             int
	SecondEmail               string `gorm:"type:varchar(255)"`
	PhoneNumberTbc            string `gorm:"type:varchar(255)"`
	Status                    int
	IsAPIUser                 bool
	CreationChannel           int
	LandlineNumber            string
	JwtValidityToken          string
}
