package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenFile(file_path string) ([]byte, error) {
	file, err := os.Open(file_path)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("File Permission Denied : ", err)
			return nil, err
		} else if os.IsNotExist(err) {
			fmt.Println("File Not Exist : ", err)
			return nil, err
		} else {
			fmt.Println("Undefined Error : ", err)
			return nil, err 
		}
	}
	
	defer file.Close()

	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println("File Reading Error : ", err)
		return nil, err
	}
	return content, nil
}
