package shortener

import (
	"fmt"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink := "https://github.com/practical-tutorials/project-based-learning#go"
	shortLink := GenerateShortLink(initialLink, UserId)

	fmt.Println(shortLink)
}
