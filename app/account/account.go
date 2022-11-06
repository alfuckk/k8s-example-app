package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/horzions/pkg/config"
	"github.com/horzions/pkg/jwt"
	"github.com/horzions/pkg/password"
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

func (us *AccountService) Register(c *gin.Context) {
	var ra RegisterAccount
	err := c.ShouldBindQuery(&ra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	validate := validator.New()
	err = validate.Struct(ra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	hash, _ := password.HashPassword(ra.Password)

	var account Account
	user := us.DB.Where(&Account{AccountName: ra.AccountName, Email: ra.Email}).Limit(1).Find(&account)
	if user.RowsAffected > 0 {
		c.JSON(http.StatusAccepted, gin.H{"msg": "account exists."})
		return
	}
	result := us.DB.Create(&Account{AccountName: ra.AccountName, Email: ra.Email, Password: hash})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "created failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "account created success"})
}

func (us *AccountService) ResetPassword(c *gin.Context) {

}

func (us *AccountService) Login(c *gin.Context) {
	var la LoginAccount
	err := c.ShouldBindQuery(&la)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	validate := validator.New()
	err = validate.Struct(la)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var account Account
	result := us.DB.Where(&Account{Email: la.Email}).Limit(1).Find(&account)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusAccepted, gin.H{"msg": "account invalid."})
		return
	}
	isLogin := password.CheckPasswordHash(la.Password, account.Password)

	if !isLogin {
		c.JSON(http.StatusAccepted, gin.H{"msg": "password invalid."})
		return
	}
	token, _ := jwt.NewJwt(&us.Config.Server).GenerateJWT(account.Email, account.AccountName)
	c.JSON(http.StatusOK, gin.H{"msg": "login success.", "token": token})
}

func (us *AccountService) AddAccount(c *gin.Context) {

}

func (us *AccountService) DeleteAccount(c *gin.Context) {

}

func (us *AccountService) ModifyAccount(c *gin.Context) {

}

func (us *AccountService) GetAccounts(c *gin.Context) {

}

func (us *AccountService) AccountInfo(c *gin.Context) {

}
