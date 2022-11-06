package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/horzions/pkg/config"
	"github.com/horzions/pkg/helper"
	"github.com/horzions/pkg/jwt"

	"gorm.io/gorm"
)

type AccountService struct {
	DB     *gorm.DB
	Config *config.Config
	Engine *gin.Engine
}

func NewAccount(c *config.Config, e *gin.Engine, db *gorm.DB) *AccountService {
	db.AutoMigrate(&Account{})
	return &AccountService{
		Config: c,
		Engine: e,
		DB:     db,
	}
}

func (as *AccountService) Register(c *gin.Context) {
	var ra RegisterAccount
	err := c.ShouldBindJSON(&ra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	hash, _ := helper.HashPassword(ra.Password)

	var account Account
	user := as.DB.Where(&Account{AccountName: ra.AccountName, Email: ra.Email}).Limit(1).Find(&account)
	if user.RowsAffected > 0 {
		c.JSON(http.StatusAccepted, gin.H{"msg": "account exists."})
		return
	}
	result := as.DB.Create(&Account{AccountName: ra.AccountName, Email: ra.Email, Password: hash})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "created failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "account created success."})
}

func (us *AccountService) ResetPassword(c *gin.Context) {

}

func (as *AccountService) Login(c *gin.Context) {
	var la LoginAccount
	err := c.ShouldBindJSON(&la)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var account Account
	result := as.DB.Where(&Account{Email: la.Email}).Limit(1).Find(&account)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusAccepted, gin.H{"msg": "account invalid."})
		return
	}
	isLogin := helper.CheckPasswordHash(la.Password, account.Password)

	if !isLogin {
		c.JSON(http.StatusAccepted, gin.H{"msg": "password invalid."})
		return
	}
	token, _ := jwt.NewJwt(&as.Config.Server).GenerateJWT(account.Email, account.AccountName)
	c.JSON(http.StatusOK, gin.H{"msg": "login success.", "token": token})
}

func (as *AccountService) AddAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "add account."})
}

func (as *AccountService) DeleteAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "delete account."})
}

func (as *AccountService) ModifyAccount(c *gin.Context) {

}

func (as *AccountService) GetAccounts(c *gin.Context) {

}

func (as *AccountService) AccountInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "account info."})
}
