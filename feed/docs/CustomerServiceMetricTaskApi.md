# \CustomerServiceMetricTaskApi

All URIs are relative to *https://api.ebay.com/sell/feed/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerServiceMetricTask**](CustomerServiceMetricTaskApi.md#CreateCustomerServiceMetricTask) | **Post** /customer_service_metric_task | 
[**GetCustomerServiceMetricTask**](CustomerServiceMetricTaskApi.md#GetCustomerServiceMetricTask) | **Get** /customer_service_metric_task/{task_id} | 
[**GetCustomerServiceMetricTasks**](CustomerServiceMetricTaskApi.md#GetCustomerServiceMetricTasks) | **Get** /customer_service_metric_task | 



## CreateCustomerServiceMetricTask

> CreateCustomerServiceMetricTask(ctx).AcceptLanguage(acceptLanguage).CreateServiceMetricsTaskRequest(createServiceMetricsTaskRequest).Execute()





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
    acceptLanguage := "acceptLanguage_example" // string | Use this header to specify the natural language in which the authenticated user desires the response.
    createServiceMetricsTaskRequest := *openapiclient.NewCreateServiceMetricsTaskRequest() // CreateServiceMetricsTaskRequest | Request payload containing version, feedType, and optional filterCriteria.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerServiceMetricTaskApi.CreateCustomerServiceMetricTask(context.Background()).AcceptLanguage(acceptLanguage).CreateServiceMetricsTaskRequest(createServiceMetricsTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceMetricTaskApi.CreateCustomerServiceMetricTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerServiceMetricTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use this header to specify the natural language in which the authenticated user desires the response. | 
 **createServiceMetricsTaskRequest** | [**CreateServiceMetricsTaskRequest**](CreateServiceMetricsTaskRequest.md) | Request payload containing version, feedType, and optional filterCriteria. | 

### Return type

 (empty response body)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerServiceMetricTask

> ServiceMetricsTask GetCustomerServiceMetricTask(ctx, taskId).Execute()





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
    taskId := "taskId_example" // string | Use this path parameter to specify the task ID value for the customer service metric task to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerServiceMetricTaskApi.GetCustomerServiceMetricTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceMetricTaskApi.GetCustomerServiceMetricTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerServiceMetricTask`: ServiceMetricsTask
    fmt.Fprintf(os.Stdout, "Response from `CustomerServiceMetricTaskApi.GetCustomerServiceMetricTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Use this path parameter to specify the task ID value for the customer service metric task to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerServiceMetricTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceMetricsTask**](ServiceMetricsTask.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerServiceMetricTasks

> CustomerServiceMetricTaskCollection GetCustomerServiceMetricTasks(ctx).DateRange(dateRange).FeedType(feedType).Limit(limit).LookBackDays(lookBackDays).Offset(offset).Execute()





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
    dateRange := "dateRange_example" // string | The task creation date range. The results are filtered to include only tasks with a creation date that is equal to the dates specified or is within the specified range. Do not use with the look_back_days parameter. Format: UTC For example, tasks within a range: yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ Tasks created on March 8, 2020 2020-03-08T00:00.00.000Z..2020-03-09T00:00:00.000Z Maximum: 90 days (optional)
    feedType := "feedType_example" // string | The feed type associated with the task. The only presently supported value is CUSTOMER_SERVICE_METRICS_REPORT. (optional)
    limit := "limit_example" // string | The number of customer service metric tasks to return per page of the result set. Use this parameter in conjunction with the offset parameter to control the pagination of the output. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Note:This feature employs a zero-based list, where the first item in the list has an offset of 0. Default: 10 Maximum: 500 (optional)
    lookBackDays := "lookBackDays_example" // string | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter's default value is used. Default value: 7 Range: 1-90 (inclusive) (optional)
    offset := "offset_example" // string | The number of customer service metric tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. Default: 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerServiceMetricTaskApi.GetCustomerServiceMetricTasks(context.Background()).DateRange(dateRange).FeedType(feedType).Limit(limit).LookBackDays(lookBackDays).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerServiceMetricTaskApi.GetCustomerServiceMetricTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerServiceMetricTasks`: CustomerServiceMetricTaskCollection
    fmt.Fprintf(os.Stdout, "Response from `CustomerServiceMetricTaskApi.GetCustomerServiceMetricTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerServiceMetricTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateRange** | **string** | The task creation date range. The results are filtered to include only tasks with a creation date that is equal to the dates specified or is within the specified range. Do not use with the look_back_days parameter. Format: UTC For example, tasks within a range: yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ Tasks created on March 8, 2020 2020-03-08T00:00.00.000Z..2020-03-09T00:00:00.000Z Maximum: 90 days | 
 **feedType** | **string** | The feed type associated with the task. The only presently supported value is CUSTOMER_SERVICE_METRICS_REPORT. | 
 **limit** | **string** | The number of customer service metric tasks to return per page of the result set. Use this parameter in conjunction with the offset parameter to control the pagination of the output. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Note:This feature employs a zero-based list, where the first item in the list has an offset of 0. Default: 10 Maximum: 500 | 
 **lookBackDays** | **string** | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter&#39;s default value is used. Default value: 7 Range: 1-90 (inclusive) | 
 **offset** | **string** | The number of customer service metric tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. Default: 0 | 

### Return type

[**CustomerServiceMetricTaskCollection**](CustomerServiceMetricTaskCollection.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

