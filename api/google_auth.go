package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:1323/callback",
		ClientID:     getEnvVar("GOOGLE_CLIENT_ID"),
		ClientSecret: getEnvVar("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func LoginHandler(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("random_state")
	return c.Redirect(http.StatusSeeOther, url)
}

func CallbackHandler(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "random_state" {
		return c.String(http.StatusInternalServerError, "Google auth callback states mismatch")
	}

	code := c.QueryParam("code")
	token, err := googleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to exchange Code-Token")
	}

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to fetch user data")
	}

	userData, err := io.ReadAll(res.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to parse JSON")
	}

	return c.String(http.StatusOK, string(userData))

}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("failed to read .env file:", err)
	}

	return os.Getenv(key)
}
