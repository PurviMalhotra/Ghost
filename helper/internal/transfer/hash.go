package transfer

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func GenerateHash(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha256.New()

	_, err = io.Copy(hash, file)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func HasHash(targetHash string) bool {

	files, err := os.ReadDir("./data")

	if err != nil {
		return false
	}

	for _, file := range files {

		path := "./data/" + file.Name()

		hash, err := GenerateHash(path)

		if err != nil {
			continue
		}

		if hash == targetHash {
			return true
		}
	}

	return false
}
