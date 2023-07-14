package readFile

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() []string {
	// Đọc nội dung trong localfile.data
	data, err := os.ReadFile("data.csv")
	// Nếu chương trình không thể đọc file
	// in ra nguyên nhân tại sao
	if err != nil {
		fmt.Println(err)
	}

	// Nếu đọc file thành công thì
	// in ra nội dung như một string
	str := string(data)
	fmt.Println(str)
	result := strings.Split(str, "\n")
	return result

}
