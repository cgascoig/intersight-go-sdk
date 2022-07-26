# Deprecation Notice

This repository is no longer updated (manually or automatically)! Please use the supported [Intersight Go](https://github.com/ciscodevnet/intersight-go) SDK in the Cisco DevNet repository.

# (Unofficial) Intersight Go SDK

This is a Go module that contains the Intersight SDK. 

This module is generated automatically (using the [OpenAPITools OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)) from the [Intersight OpenAPI specification](https://intersight.com/apidocs/downloads/) and updated automatically whenever there is a new version of the specification available. 

# Usage

With each new specification release, a new tag is created in this repository for the SDK, for example. `v1.0.9-4870`. 

To add the SDK to your `go.mod` file, use:

```
go get github.com/cgascoig/intersight-go-sdk/intersight@v1.0.9-4870
```

Then in your code:

```
...
import "github.com/cgascoig/intersight-go-sdk/intersight"

func main() {
	config := intersight.NewConfiguration()
	client := intersight.NewAPIClient(config)
...
```

You can find a simple example using this module in the [intersight-go-example](https://github.com/cgascoig/intersight-go-example) repository. 
