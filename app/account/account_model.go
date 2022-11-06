package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountName  string `gorm:"account_name"`
	Email        string `gorm:"email"`
	Password     string `gorm:"password"`
	AccessToken  string `gorm:"access_token"`
	RefreshToken string `gorm:"refresh_token"`
}
