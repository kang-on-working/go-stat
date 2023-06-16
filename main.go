package main

import (
	"fmt"
	"pkg"
)

func main() {
	target := "./example"

	err := pkg.WriteFile(target, []byte("example text test"))

	fmt.Println("main")
	FILEDATA, err := pkg.ReadFile(target)
	if err != nil { return } //에러는 함수 단에서 이미 출력될 것

	fmt.Println(string(FILEDATA))

}
