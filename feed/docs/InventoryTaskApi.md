# \InventoryTaskApi

All URIs are relative to *https://api.ebay.com/sell/feed/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInventoryTask**](InventoryTaskApi.md#CreateInventoryTask) | **Post** /inventory_task | 
[**GetInventoryTask**](InventoryTaskApi.md#GetInventoryTask) | **Get** /inventory_task/{task_id} | 
[**GetInventoryTasks**](InventoryTaskApi.md#GetInventoryTasks) | **Get** /inventory_task | 



## CreateInventoryTask

> CreateInventoryTask(ctx).CreateInventoryTaskRequest(createInventoryTaskRequest).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()





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
    createInventoryTaskRequest := *openapiclient.NewCreateInventoryTaskRequest() // CreateInventoryTaskRequest | The request payload containing the version, feedType, and optional filterCriteria.
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | The ID of the eBay marketplace where the item is hosted. Note: This value is case sensitive. For example: X-EBAY-C-MARKETPLACE-ID:EBAY_US This identifies the eBay marketplace that applies to this task. See MarketplaceIdEnum. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryTaskApi.CreateInventoryTask(context.Background()).CreateInventoryTaskRequest(createInventoryTaskRequest).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryTaskApi.CreateInventoryTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInventoryTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInventoryTaskRequest** | [**CreateInventoryTaskRequest**](CreateInventoryTaskRequest.md) | The request payload containing the version, feedType, and optional filterCriteria. | 
 **xEBAYCMARKETPLACEID** | **string** | The ID of the eBay marketplace where the item is hosted. Note: This value is case sensitive. For example: X-EBAY-C-MARKETPLACE-ID:EBAY_US This identifies the eBay marketplace that applies to this task. See MarketplaceIdEnum. | 

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


## GetInventoryTask

> InventoryTask GetInventoryTask(ctx, taskId).Execute()





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
    taskId := "taskId_example" // string | The ID of the task. This ID was generated when the task was created by the createInventoryTask method

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryTaskApi.GetInventoryTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryTaskApi.GetInventoryTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventoryTask`: InventoryTask
    fmt.Fprintf(os.Stdout, "Response from `InventoryTaskApi.GetInventoryTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task. This ID was generated when the task was created by the createInventoryTask method | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryTask**](InventoryTask.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoryTasks

> InventoryTaskCollection GetInventoryTasks(ctx).FeedType(feedType).ScheduleId(scheduleId).LookBackDays(lookBackDays).DateRange(dateRange).Limit(limit).Offset(offset).Execute()





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
    feedType := "feedType_example" // string | The feed type associated with the inventory task. Either feed_type or schedule_id is required. Do not use with the schedule_id parameter. Presently, only one feed type is available: LMS_ACTIVE_INVENTORY_REPORT (optional)
    scheduleId := "scheduleId_example" // string | The ID of the schedule for which to retrieve the latest result file. This ID is generated when the schedule was created by the createSchedule method. Schedules apply to downloaded reports (LMS_ACTIVE_INVENTORY_REPORT). Either schedule_id or feed_type is required. Do not use with the feed_type parameter. (optional)
    lookBackDays := "lookBackDays_example" // string | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter's default value is used. Default: 7 Range: 1-90 (inclusive) (optional)
    dateRange := "dateRange_example" // string | Specifies the range of task creation dates used to filter the results. The results are filtered to include only tasks with a creation date that is equal to this date or is within specified range. Note: Maximum date range window size is 90 days. Valid Format (UTC): yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks created on March 31, 2021 2021-03-31T00:00:00.000Z..2021-03-31T00:00:00.000Z (optional)
    limit := "limit_example" // string | The maximum number of tasks that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 (optional)
    offset := "offset_example" // string | The number of tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryTaskApi.GetInventoryTasks(context.Background()).FeedType(feedType).ScheduleId(scheduleId).LookBackDays(lookBackDays).DateRange(dateRange).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryTaskApi.GetInventoryTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInventoryTasks`: InventoryTaskCollection
    fmt.Fprintf(os.Stdout, "Response from `InventoryTaskApi.GetInventoryTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedType** | **string** | The feed type associated with the inventory task. Either feed_type or schedule_id is required. Do not use with the schedule_id parameter. Presently, only one feed type is available: LMS_ACTIVE_INVENTORY_REPORT | 
 **scheduleId** | **string** | The ID of the schedule for which to retrieve the latest result file. This ID is generated when the schedule was created by the createSchedule method. Schedules apply to downloaded reports (LMS_ACTIVE_INVENTORY_REPORT). Either schedule_id or feed_type is required. Do not use with the feed_type parameter. | 
 **lookBackDays** | **string** | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter&#39;s default value is used. Default: 7 Range: 1-90 (inclusive) | 
 **dateRange** | **string** | Specifies the range of task creation dates used to filter the results. The results are filtered to include only tasks with a creation date that is equal to this date or is within specified range. Note: Maximum date range window size is 90 days. Valid Format (UTC): yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks created on March 31, 2021 2021-03-31T00:00:00.000Z..2021-03-31T00:00:00.000Z | 
 **limit** | **string** | The maximum number of tasks that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 | 
 **offset** | **string** | The number of tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 | 

### Return type

[**InventoryTaskCollection**](InventoryTaskCollection.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

