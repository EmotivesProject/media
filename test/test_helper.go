package test

import (
	"os"

	"github.com/EmotivesProject/common/logger"
)

func SetUpIntegrationTest() {
	logger.InitLogger("media", logger.EmailConfig{
		From:     os.Getenv("EMAIL_FROM"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		Level:    os.Getenv("EMAIL_LEVEL"),
	})

	os.Setenv("FILE_LOCATION", "./images/")
}
