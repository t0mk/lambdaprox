package main

import (
	"github.com/t0mk/lambdaprox"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)


func main() {
	// Register the function with the Functions Framework
	funcframework.RegisterHTTPFunction("/", lambdaprox.LambdaProx)

	// Start the Functions Framework
	if err := funcframework.Start("8085"); err != nil {
		panic(err)
	}
}