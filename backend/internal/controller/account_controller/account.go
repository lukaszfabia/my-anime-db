package accountcontroller

import (
	"api/internal/controller"
	"api/internal/models"
	"api/internal/store"
	"api/pkg/db"
	"api/pkg/tools"
	"api/pkg/utils"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mail/mail"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var storage = store.NewVerificationStore()

type AccountController interface {
	CreateAccount(c *gin.Context) error
	GenerateToken(loginParams models.LoginForm) (string, error)
	EditAccount(c *gin.Context, user *models.User) error
	DeleteAccount(user *models.User) error
	VerifyAccount(c *gin.Context, user *models.User) error
	SendCode(user *models.User) error
	Get(user *models.User) (models.User, error)
	SendWelcomeEmail(user models.User) error
}

type AccountControllerImpl struct {
}

func (ac *AccountControllerImpl) SendWelcomeEmail(user models.User) error {
	senderMail := os.Getenv("GOOGLE_MAIL")
	senderPassword := os.Getenv("GOOGLE_PASSWORD")
	m := mail.NewMessage()

	m.SetHeader("From", senderMail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Reply-To", "no-reply@example.com")
	m.SetAddressHeader("Cc", user.Email, user.Username)
	m.SetHeader("Subject", "Welcome to myanime.db!")

	body, err := tools.ParseHTMLToString("welcome.html", user)

	if err != nil {
		log.Println(err)
		return err
	}

	m.SetBody("text/html", body)

	d := mail.NewDialer("smtp.gmail.com", 587, senderMail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v", err)
		return errors.New("failed to send email")
	}

	return nil
}

func (ac *AccountControllerImpl) CreateAccount(c *gin.Context) error {

	var signupForm = models.SignupForm{
		LoginForm: models.LoginForm{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		},
		Email: c.PostForm("email"),
	}

	hashed, err := utils.HashPassword(signupForm.Password)

	if err != nil {
		return errors.New("failed to hash password")
	}

	picUrl := utils.SaveImage(c, utils.Avatar, "pic")

	user := models.User{
		Username: signupForm.Username,
		Email:    signupForm.Email,
		Password: hashed,
		PicUrl:   picUrl,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		utils.RemoveImage(user.PicUrl)
		return errors.New("failed to create user")
	}

	user.UserStats = &models.UserStat{
		UserID: user.ID,
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return errors.New("failed to set user stats")
	}

	return nil
}

func (ac *AccountControllerImpl) GenerateToken(loginParams models.LoginForm) (string, error) {
	var user models.User

	if err := db.DB.First(&user, "username = ?", loginParams.Username).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password)); err != nil {
		return "", errors.New("credentials are invalid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenStr, nil
}

func (ac *AccountControllerImpl) EditAccount(c *gin.Context, user *models.User) error {
	var userToUpdate models.User

	if err := db.DB.First(&userToUpdate, user.ID).Error; err != nil {
		return errors.New("user not found")
	}

	userToUpdate.Username = controller.GetOrDefault(c.PostForm("username"), userToUpdate.Username).(string)

	newPassword, err := utils.HashPassword(c.PostForm("password"))

	if err == nil {
		user.Password = newPassword
	}

	userToUpdate.PicUrl = utils.UpdateImage(c, userToUpdate.PicUrl, utils.Avatar, "pic")

	newEmail := controller.GetOrDefault(c.PostForm("email"), userToUpdate.Email).(string)

	if newEmail != userToUpdate.Email {
		userToUpdate.IsVerified = false
	}

	userToUpdate.Email = newEmail

	userToUpdate.Bio = controller.GetOrDefault(c.PostForm("bio"), userToUpdate.Bio).(string)
	userToUpdate.Website = controller.GetOrDefault(c.PostForm("website"), userToUpdate.Website).(string)

	if err := db.DB.Save(&userToUpdate).Error; err != nil {
		log.Println(err)
		return errors.New("failed to update user")
	}

	return nil
}

func (ac *AccountControllerImpl) DeleteAccount(user *models.User) error {
	return db.Delete(&models.User{}, strconv.Itoa(int(user.ID)),
		db.Association{Model: "UserAnimes"},
		db.Association{Model: "Posts"},
		db.Association{Model: "Friends"},
	)
}

func (ac *AccountControllerImpl) VerifyAccount(c *gin.Context, user *models.User) error {
	code := c.PostForm("code")
	log.Println(code)
	if err := storage.Compare(code, user.Email); err != nil {
		log.Println(err)
		return err
	}

	user.IsVerified = true

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ac *AccountControllerImpl) SendCode(user *models.User) error {
	senderMail := os.Getenv("GOOGLE_MAIL")
	senderPassword := os.Getenv("GOOGLE_PASSWORD")

	code := utils.GenerateCode()

	m := mail.NewMessage()

	m.SetHeader("From", senderMail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Reply-To", "no-reply@example.com")
	m.SetAddressHeader("Cc", user.Email, user.Username)
	m.SetHeader("Subject", "Verify your account in myanime.db!")

	type dataForTempl struct {
		Username string
		Email    string
		Code     string
	}
	data := dataForTempl{
		Username: user.Username,
		Email:    user.Email,
		Code:     code,
	}

	body, err := tools.ParseHTMLToString("send_code.html", data)
	if err != nil {
		log.Println(err)
		return err
	}

	m.SetBody("text/html", body)

	d := mail.NewDialer("smtp.gmail.com", 587, senderMail, senderPassword)

	if err := d.DialAndSend(m); err != nil {
		return errors.New("failed to send email")
	}

	storage.Set(user.Email, code)

	return nil
}

func (ac *AccountControllerImpl) Get(user *models.User) (models.User, error) {
	var userDetails models.User
	if err := db.DB.
		Preload("Friends", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "pic_url", "is_verified", "created_at", "bio", "website")
		}).
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Preload("Reviews", func(db *gorm.DB) *gorm.DB {
			return db.Preload("UserAnime", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Anime", func(db *gorm.DB) *gorm.DB {
					return db.Preload("Genres").Preload("Studio").Preload("AlternativeTitles")
				}).Preload("User").Order("created_at DESC")
			})
		}).
		Preload("UserStats").
		First(&userDetails, user.ID).Error; err != nil {
		return models.User{}, err
	}

	return userDetails, nil
}
