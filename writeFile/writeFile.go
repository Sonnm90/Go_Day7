package writeFile

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(data []byte) {

	//// Phương thức WriteFile trả về lỗi nêú không thành công
	//err := os.WriteFile("data.csv", data, 0777)
	//// xử lý lỗi này
	//if err != nil {
	//	// in ra thông tin lỗi
	//	fmt.Println(err)
	//}
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		os.Create("data.csv")
	}
	defer file.Close()

	content := string(data)

	_, err = fmt.Fprintln(file, content)
	if err != nil {
		log.Fatal(err)
	}

}
