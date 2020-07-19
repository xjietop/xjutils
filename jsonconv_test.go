package xjutils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonConv(t *testing.T) {
	type Person struct {
		HelloWold       string
		LightWeightBaby string
	}
	var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
	res, _ := json.Marshal(a)
	fmt.Printf("%s\n", res)
	res, _ = json.Marshal(JsonSnakeCase{a})
	fmt.Printf("%s\n", res)
	res, _ = json.Marshal(JsonCamelCase{a})
	fmt.Printf("%s\n", res)
	var b Person
	err := json.Unmarshal(res,&b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
