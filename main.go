package main

import (
  "log"

  "github.com/kjwenger/sandbox.petstore.go-swagger/client/operations"

  apiclient "github.com/kjwenger/sandbox.petstore.go-swagger/client"
  httptransport "github.com/go-openapi/runtime/client"

  "fmt"
)

func main() {

  // make the request to get all items
  resp, err := apiclient.Default.Operations.All(operations.AllParams{})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
}