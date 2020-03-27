// 예제 13 JSON 스키마들 실행하는 코드

package main

import (
	"fmt"
	"gifhub.com/xeipuuv/gojsonchema"
	"io/ioutil"
)

func main() {
	// 스키마를 파일에서 읽어온다
	schema, err := iotuil.ReadFile("schema.json")
	if err != nil {
		panic(err)
	}
	schemaLoader := gojsonschema.NewBytesLoader(schema)

	// 검증 대상 파일을 파일에서 읽어온다
	document, err := ioutil.ReadFile("document.json")
	if err != nil {
		panic(err)
	}
	documentLoader := gojsonschema.NewBytesLoader(document)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("-%s\n", desc)
		}
	}
}
