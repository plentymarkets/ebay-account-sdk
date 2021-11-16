# \ProgramApi

All URIs are relative to *https://api.ebay.com/sell/account/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptedInPrograms**](ProgramApi.md#GetOptedInPrograms) | **Get** /program/get_opted_in_programs | 
[**OptInToProgram**](ProgramApi.md#OptInToProgram) | **Post** /program/opt_in | 
[**OptOutOfProgram**](ProgramApi.md#OptOutOfProgram) | **Post** /program/opt_out | 



## GetOptedInPrograms

> Programs GetOptedInPrograms(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProgramApi.GetOptedInPrograms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.GetOptedInPrograms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOptedInPrograms`: Programs
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.GetOptedInPrograms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptedInProgramsRequest struct via the builder pattern


### Return type

[**Programs**](Programs.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptInToProgram

> map[string]interface{} OptInToProgram(ctx).Program(program).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    program := *openapiclient.NewProgram() // Program | Program being opted-in to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProgramApi.OptInToProgram(context.Background()).Program(program).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.OptInToProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptInToProgram`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.OptInToProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptInToProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **program** | [**Program**](Program.md) | Program being opted-in to. | 

### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptOutOfProgram

> map[string]interface{} OptOutOfProgram(ctx).Program(program).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    program := *openapiclient.NewProgram() // Program | Program being opted-out of.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProgramApi.OptOutOfProgram(context.Background()).Program(program).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProgramApi.OptOutOfProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptOutOfProgram`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProgramApi.OptOutOfProgram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptOutOfProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **program** | [**Program**](Program.md) | Program being opted-out of. | 

### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

