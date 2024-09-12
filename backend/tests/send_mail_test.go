package tests

import (
	accountcontroller "api/internal/controller/account_controller"
	"api/internal/models"
	"api/pkg/tools"
	"testing"

	"github.com/joho/godotenv"
)

var ac accountcontroller.AccountController = &accountcontroller.AccountControllerImpl{}

var mail = "depic98098@apifan.com"
var username = "test"

var user = models.User{
	Email:    mail,
	Username: username,
}

func TestSendWelcomeEmail(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}

	if err := ac.SendWelcomeEmail(user); err != nil {
		t.Errorf("Error sending welcome email: %v", err)
	}

	t.Logf("Welcome email sent to %s", user.Email)
}

func TestSendCode(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}

	if err := ac.SendCode(&user); err != nil {
		t.Errorf("Error sending code: %v", err)
	}

}

func TestParseHTMLToString(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}

	body, err := tools.ParseHTMLToString("welcome.html", user)
	if err != nil {
		t.Errorf("Error parsing HTML to string: %v", err)
	}

	t.Logf("Parsed HTML to string: %s", body)
}
