package event-transformer-sap

import (
	"encoding/json"
	"fmt"
	"log"

	jsonata "github.com/blues/jsonata-go"
)

const jsonString = `
    {
        "orders": [
            {"price": 10, "quantity": 3},
            {"price": 0.5, "quantity": 10},
            {"price": 100, "quantity": 1}
        ]
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
	e := jsonata.MustCompile("$sum(orders.(price*quantity))")

	// Evaluate.
	res, err := e.Eval(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	// Output: 135
}