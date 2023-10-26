package jsonAtaTransform

import (
	"encoding/json"
	"fmt"
	"log"

	jsonata "github.com/blues/jsonata-go"
)
const jsonString = `
	{"manufacturingOrder":"000432"}
`
const jsonAtaExpr = `
    $.{
        "operationRequest":{
        "id":manufacturingOrder
	}
    }
`

func TransformSAPOrder() {

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
