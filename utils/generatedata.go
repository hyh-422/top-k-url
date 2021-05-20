package utils

import (
	"os"
	"strconv"
)

// GenerateUrlData generate url file as data.txt
func GenerateUrlData(str string) error {
	f, err := os.OpenFile(str, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	} else {
		count := 0
		for count < 1000 {
			for i := count; i >= 0; i-- {
				_, err = f.Write([]byte("https://www.baidu.com/" + strconv.Itoa(count) + "\n"))
				if err != nil {
					return err
				}
			}
			count++
		}
	}
	return nil
}
