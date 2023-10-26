package jsonAtaTransform

import (
	"encoding/json"
	"fmt"
	"log"
	jsonata "github.com/blues/jsonata-go"
)

const jsonAtaExpr = `
    $.data.{
        "materialDefinition":{
        "id":article
        }
    }
`

func TransformPLMArticleString(jsonString string)(interface{}) {
	var data interface{}
	// Decode JSON.
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err)
	}
	return TransformSAPOrder(data) 
}

func TransformPLMArticle(data interface{})(interface{}) {
	e := jsonata.MustCompile(jsonAtaExpr)
	// Evaluate.
	res, err := e.Eval(data)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func TransformSAPOrder_test() {
	res = TransformSAPOrderString(plmArticle_sample)
	fmt.Println(res)
}

const plmArticle_sample=`{
  "_id": "0A56948624F21EDE93C0B89EFE2D452C",
  "_flow": "0A56948624F21EDE93C0B89EFE2D452C",
  "_env": "p",
  "_tstamp": "2023-09-08T05:08:39Z",
  "_source": "urn:eda:ric:application:sap",
  "_namespace": "urn:eda:ric:ns:sap-productionorder-v1",
  "_event": "update",
  "_type": "productionorder",
  "_object": "000701944972",
  "_topic": "/emv1/p/e/sap/productionorder/update/v1/car/chm5/z041/000701944972",
  "_replyTo": "",
  "_correlationId": "",
  "data": {
    "article":"CRMLDMK"
  }
}
`
