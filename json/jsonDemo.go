package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func Marshal(p Person) []byte {
	//p := Person{Name: "John Doe", Age: 30}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(jsonBytes))
	return jsonBytes
}
func Unmarshal(data []string) {
	//jsonStr := string(Marshal())
	for _, v := range data {
		var p Person
		err := json.Unmarshal([]byte(v), &p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p.Name, p.Age)
	}

}
