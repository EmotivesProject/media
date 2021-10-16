package test

import (
	"os"

	"github.com/TomBowyerResearchProject/common/logger"
)

func SetUpIntegrationTest() {
	logger.InitLogger("media")

	os.Setenv("FILE_LOCATION", "./images/")
}
