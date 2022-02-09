# \TaskApi

All URIs are relative to *https://api.ebay.com/sell/feed/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](TaskApi.md#CreateTask) | **Post** /task | 
[**GetInputFile**](TaskApi.md#GetInputFile) | **Get** /task/{task_id}/download_input_file | 
[**GetResultFile**](TaskApi.md#GetResultFile) | **Get** /task/{task_id}/download_result_file | 
[**GetTask**](TaskApi.md#GetTask) | **Get** /task/{task_id} | 
[**GetTasks**](TaskApi.md#GetTasks) | **Get** /task | 
[**UploadFile**](TaskApi.md#UploadFile) | **Post** /task/{task_id}/upload_file | 



## CreateTask

> CreateTask(ctx).CreateTaskRequest(createTaskRequest).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()





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
    createTaskRequest := *openapiclient.NewCreateTaskRequest() // CreateTaskRequest | description not needed
    xEBAYCMARKETPLACEID := "xEBAYCMARKETPLACEID_example" // string | The ID of the eBay marketplace where the item is hosted. Note: This value is case sensitive. For example: X-EBAY-C-MARKETPLACE-ID:EBAY_US This identifies the eBay marketplace that applies to this task. See MarketplaceIdEnum. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.CreateTask(context.Background()).CreateTaskRequest(createTaskRequest).XEBAYCMARKETPLACEID(xEBAYCMARKETPLACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) | description not needed | 
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


## GetInputFile

> map[string]interface{} GetInputFile(ctx, taskId).Execute()





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
    taskId := "taskId_example" // string | The task ID associated with the file to be downloaded.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.GetInputFile(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetInputFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInputFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetInputFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The task ID associated with the file to be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInputFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResultFile

> map[string]interface{} GetResultFile(ctx, taskId).Execute()





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
    taskId := "taskId_example" // string | The ID of the task associated with the file you want to download. This ID was generated when the task was created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.GetResultFile(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetResultFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetResultFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task associated with the file you want to download. This ID was generated when the task was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId).Execute()





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
    taskId := "taskId_example" // string | The ID of the task. This ID was generated when the task was created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.GetTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task. This ID was generated when the task was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> TaskCollection GetTasks(ctx).DateRange(dateRange).FeedType(feedType).Limit(limit).LookBackDays(lookBackDays).Offset(offset).ScheduleId(scheduleId).Execute()





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
    dateRange := "dateRange_example" // string | Specifies the range of task creation dates used to filter the results. The results are filtered to include only tasks with a creation date that is equal to this date or is within specified range. Only tasks that are less than 90 days can be retrieved. Note: Maximum date range window size is 90 days. Valid Format (UTC):yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks created on September 8, 2019 2019-09-08T00:00:00.000Z..2019-09-09T00:00:00.000Z (optional)
    feedType := "feedType_example" // string | The feed type associated with the tasks to be returned. Only use a feedType that is available for your API: Order Feeds: LMS_ORDER_ACK, LMS_ORDER_REPORT Large Merchant Services (LMS) Feeds: See Available FeedTypes Do not use with the schedule_id parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type. (optional)
    limit := "limit_example" // string | The maximum number of tasks that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 (optional)
    lookBackDays := "lookBackDays_example" // string | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter's default value is used. Default: 7 Range: 1-90 (inclusive) (optional)
    offset := "offset_example" // string | The number of tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 (optional)
    scheduleId := "scheduleId_example" // string | The schedule ID associated with the task. A schedule periodically generates a report for the feed type specified by the schedule template (see scheduleTemplateId in createSchedule). Do not use with the feed_type parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.GetTasks(context.Background()).DateRange(dateRange).FeedType(feedType).Limit(limit).LookBackDays(lookBackDays).Offset(offset).ScheduleId(scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: TaskCollection
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateRange** | **string** | Specifies the range of task creation dates used to filter the results. The results are filtered to include only tasks with a creation date that is equal to this date or is within specified range. Only tasks that are less than 90 days can be retrieved. Note: Maximum date range window size is 90 days. Valid Format (UTC):yyyy-MM-ddThh:mm:ss.SSSZ..yyyy-MM-ddThh:mm:ss.SSSZ For example: Tasks created on September 8, 2019 2019-09-08T00:00:00.000Z..2019-09-09T00:00:00.000Z | 
 **feedType** | **string** | The feed type associated with the tasks to be returned. Only use a feedType that is available for your API: Order Feeds: LMS_ORDER_ACK, LMS_ORDER_REPORT Large Merchant Services (LMS) Feeds: See Available FeedTypes Do not use with the schedule_id parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type. | 
 **limit** | **string** | The maximum number of tasks that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves tasks 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 | 
 **lookBackDays** | **string** | The number of previous days in which to search for tasks. Do not use with the date_range parameter. If both date_range and look_back_days are omitted, this parameter&#39;s default value is used. Default: 7 Range: 1-90 (inclusive) | 
 **offset** | **string** | The number of tasks to skip in the result set before returning the first task in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 | 
 **scheduleId** | **string** | The schedule ID associated with the task. A schedule periodically generates a report for the feed type specified by the schedule template (see scheduleTemplateId in createSchedule). Do not use with the feed_type parameter. Since schedules are based on feed types, you can specify a schedule (schedule_id) that returns the needed feed_type. | 

### Return type

[**TaskCollection**](TaskCollection.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> map[string]interface{} UploadFile(ctx, taskId).CreationDate(creationDate).FileName(fileName).ModificationDate(modificationDate).Name(name).ReadDate(readDate).Size(size).Type_(type_).Execute()





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
    taskId := "taskId_example" // string | The task_id associated with the file that will be uploaded. This ID was generated when the specified task was created.
    creationDate := "creationDate_example" // string | The file creation date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 8, 2019 2019-09-08T00:00:00.000Z (optional)
    fileName := "fileName_example" // string | The name of the file including its extension (for example, xml or csv) to be uploaded. (optional)
    modificationDate := "modificationDate_example" // string | The file modified date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 9, 2019 2019-09-09T00:00:00.000Z (optional)
    name := "name_example" // string | A content identifier. The only presently supported name is file. (optional)
    readDate := "readDate_example" // string | The date you read the file. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 10, 2019 2019-09-10T00:00:00.000Z (optional)
    size := int32(56) // int32 | The size of the file. (optional)
    type_ := "type__example" // string | The file type. The only presently supported type is form-data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskApi.UploadFile(context.Background(), taskId).CreationDate(creationDate).FileName(fileName).ModificationDate(modificationDate).Name(name).ReadDate(readDate).Size(size).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The task_id associated with the file that will be uploaded. This ID was generated when the specified task was created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creationDate** | **string** | The file creation date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 8, 2019 2019-09-08T00:00:00.000Z | 
 **fileName** | **string** | The name of the file including its extension (for example, xml or csv) to be uploaded. | 
 **modificationDate** | **string** | The file modified date. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 9, 2019 2019-09-09T00:00:00.000Z | 
 **name** | **string** | A content identifier. The only presently supported name is file. | 
 **readDate** | **string** | The date you read the file. Format: UTC yyyy-MM-ddThh:mm:ss.SSSZ For example: Created on September 10, 2019 2019-09-10T00:00:00.000Z | 
 **size** | **int32** | The size of the file. | 
 **type_** | **string** | The file type. The only presently supported type is form-data. | 

### Return type

**map[string]interface{}**

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

