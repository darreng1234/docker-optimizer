package analyze

import (
	"bufio"
	"os"

	"github.com/pickme-go/log"
)

func ReadManifest(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		log.Error("Cannot open file", path, err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil

}
