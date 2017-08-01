package main

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle Func
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return nil, nil
}

func main() {}
