package common

import (
	"os"
)

type adventPart func(fp *os.File) (int, error)

func RunPart(part adventPart, fileLoc string) (int, error) {
	fp, err := os.Open(fileLoc)
	if err != nil {
		return -1, err
	}
	defer fp.Close()
	answer, err := part(fp)
	if err != nil {
		return -1, nil
	}
	return answer, nil
}
