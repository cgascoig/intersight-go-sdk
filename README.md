# (Unofficial) Intersight Go SDK

This is a Go module that contains the Intersight SDK. 

This module is generated automatically (using the [OpenAPITools OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator)) from the [Intersight OpenAPI specification](https://intersight.com/apidocs/downloads/) and updated automatically whenever there is a new version of the specification available. 

The canonical version of the SDK code is [here](https://github.com/CiscoDevNet/terraform-provider-intersight/tree/master/intersight_gosdk) but that version is not able to be imported directly. 

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