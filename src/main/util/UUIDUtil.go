package util

import (
	"github.com/go-basic/uuid"
)

func GetUUID() (result string) {
	return uuid.New()
}