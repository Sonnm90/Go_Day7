package inputData

import (
	"bufio"
	"demo_day_7/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func InputData() {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var person json.Person
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Nhập dữ liệu (Nhập 'exit' để thoát):")
	for {
		fmt.Println("nhập tên")
		for scanner.Scan() {
			person.Name = scanner.Text()
			if person.Name != "" {
				break
			}
		}
		fmt.Println("Nhap tuoi")
		for scanner.Scan() {
			temp := scanner.Text()
			value, err := strconv.ParseInt(temp, 10, 0)
			if err != nil {
				fmt.Println("convert failed")
			} else {
				person.Age = int(value)
			}
			if person.Name != "" {
				break
			}
		}
		fmt.Println("Nhap email")
		for scanner.Scan() {
			person.Email = scanner.Text()
			if person.Name != "" {
				break
			}
		}

		_, err = fmt.Fprintln(file, string(json.Marshal(person)))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(person)
		fmt.Println("Dữ liệu đã được ghi vào file.")

		fmt.Println("Nữa không ?")
		var confirm string
		fmt.Scan(&confirm)
		if strings.EqualFold(confirm, "không") {
			break
		}
	}

}
