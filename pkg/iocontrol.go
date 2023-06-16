package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CheckFileAccessibility(file_path string) error {
	//file accessiblity check
	file, err := os.Open(file_path)
	if err != nil {
		if os.IsPermission(err) { fmt.Println("File Permission Denied : ", err) ; return err
		} else if os.IsNotExist(err) { fmt.Println("File Not Exist : ", err) ; return err
		} else { fmt.Println("Undefined Error : ", err) ; return err }
	} //errors
	defer file.Close()
	return nil
}

func ReadFile(file_path string) ([]byte, error) {

	if err := CheckFileAccessibility(file_path); err != nil { return nil, err }

	content, err := ioutil.ReadFile(file_path)
	if err != nil { fmt.Println("File Reading Error : ", err) ; return nil, err	}

	return content, nil
}

func WriteFile(file_path string, data []byte) error {

	if err := ioutil.WriteFile(file_path, data, 0644); err != nil { fmt.Println("File Write Error : ", err) ; return err }

	return nil
}
