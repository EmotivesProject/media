// +build integration

package save_test

import (
	"image/jpeg"
	"media/internal/save"
	"media/test"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveImage(t *testing.T) {
	test.SetUpIntegrationTest()

	file, _ := os.Open("../../images/blank.jpg")
	image, _ := jpeg.Decode(file)

	_, err := save.File(
		&image,
		"nothing.jgp",
		"test",
		true,
	)
	assert.Nil(t, err)
}
