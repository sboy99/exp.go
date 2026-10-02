package helpers

import (
	"os"
	"path/filepath"

	"github.com/sboy99/exp.go/go-book-store/infra/exceptions"
)

func SaveFile(fileName string, content []byte) {
	filePath := filepath.Join("../../storage", fileName)
	err := os.WriteFile(filePath, content, os.FileMode(0o644))
	exceptions.HandleFileWriteException(err)
}

func LoadFile(fileName string) []byte {
	filePath := filepath.Join("../../storage", fileName)
	content, err := os.ReadFile(filePath)
	exceptions.HandleFileReadException(err)
	return content
}
