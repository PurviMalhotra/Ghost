package cache

import (
	"fmt"
	"os"

	"ghost/helper/internal/transfer"
)

var AssetIndex []transfer.Asset

func BuildIndex() error {

	fmt.Println("Building asset index...")

	files, err := os.ReadDir("./data")

	if err != nil {
		return err
	}

	AssetIndex = nil

	for _, file := range files {

		info, err := file.Info()

		if err != nil {
			continue
		}

		path := "./data/" + info.Name()

		hash, err := transfer.GenerateHash(path)

		if err != nil {
			continue
		}

		AssetIndex = append(
			AssetIndex,
			transfer.Asset{
				Name: info.Name(),
				Size: info.Size(),
				Hash: hash,
			},
		)
	}

	return nil
}

func RefreshIndex() error {
	return BuildIndex()
}