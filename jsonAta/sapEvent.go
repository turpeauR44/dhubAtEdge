package jsonAtaTransform

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"

	jsonata "github.com/blues/jsonata-go"
)

const jsonAtaExpr = `
    $.{
        "operationRequest":{
        "id":manufacturingOrder
	}
    }
`
type Event struct{
	Topic string
	Wait int
	Qos int
	Data map[string]interface{}
}

func TransformSAPOrderString(jsonString string) {

	var data interface{}

	// Decode JSON.
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err)
	}

	// Create expression.
	e := jsonata.MustCompile(jsonAtaExpr)

	// Evaluate.
	res, err := e.Eval(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
	// Output: 135
}

func TransformSAPOrder(data interface{}) {
	e := jsonata.MustCompile(jsonAtaExpr)

	// Evaluate.
	res, err := e.Eval(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
	// Output: 135
}

func TransformSAPOrder_test() {
	content,err := ioutil.ReadFile("./samples/sap_order.json"+file.Name())
	if err != nil{
		log.Fatal("Error when opening file : ", err)
	}
	var file_event map[string]interface{}
	err = json.Unmarshal(content,&file_event)
	if err != nil{
		log.Fatal("Error when UnWarshall: ", err)
	}
	TransformSAPOrder(file_event)
}