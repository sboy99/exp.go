package helpers

import (
	"fmt"
	"math/rand"

	"github.com/sboy99/exp.go/go-task-management/src/infra/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type CryptoHelper struct{}

func (h *CryptoHelper) GetId() int64 {
	return rand.Int63()
}

func (h *CryptoHelper) GetHash(str string) string {
	strBytes := []byte(str)
	hashedStr, err := bcrypt.GenerateFromPassword(strBytes, 10)
	exceptions.HandleHashStringException(err)
	return string(hashedStr)
}

func (h *CryptoHelper) GetRandStr(size int) string {
	randStr := fmt.Sprint(h.GetRandInt(0, size))
	randStrLen := len(randStr)
	sizeLen := len(fmt.Sprint(size))

	for l := 1; l < sizeLen-randStrLen; l++ {
		randStr = "0" + randStr
	}

	return randStr
}

func (h *CryptoHelper) GetRandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
