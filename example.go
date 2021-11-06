package main

import (
	"fmt"
	
	"github.com/cgascoig/intersight-go-sdk/intersight"
)

func main() {
	config := intersight.NewConfiguration()
	fmt.Printf("Config: %v", config)
}