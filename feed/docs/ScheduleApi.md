# \ScheduleApi

All URIs are relative to *https://api.ebay.com/sell/feed/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchedule**](ScheduleApi.md#CreateSchedule) | **Post** /schedule | 
[**DeleteSchedule**](ScheduleApi.md#DeleteSchedule) | **Delete** /schedule/{schedule_id} | 
[**GetLatestResultFile**](ScheduleApi.md#GetLatestResultFile) | **Get** /schedule/{schedule_id}/download_result_file | 
[**GetSchedule**](ScheduleApi.md#GetSchedule) | **Get** /schedule/{schedule_id} | 
[**GetScheduleTemplate**](ScheduleApi.md#GetScheduleTemplate) | **Get** /schedule_template/{schedule_template_id} | 
[**GetScheduleTemplates**](ScheduleApi.md#GetScheduleTemplates) | **Get** /schedule_template | 
[**GetSchedules**](ScheduleApi.md#GetSchedules) | **Get** /schedule | 
[**UpdateSchedule**](ScheduleApi.md#UpdateSchedule) | **Put** /schedule/{schedule_id} | 



## CreateSchedule

> map[string]interface{} CreateSchedule(ctx).CreateUserScheduleRequest(createUserScheduleRequest).Execute()





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
    createUserScheduleRequest := *openapiclient.NewCreateUserScheduleRequest() // CreateUserScheduleRequest | In the request payload: feedType and scheduleTemplateId are required; scheduleName is optional; preferredTriggerHour, preferredTriggerDayOfWeek, preferredTriggerDayOfMonth, scheduleStartDate, scheduleEndDate, and schemaVersion are conditional.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.CreateSchedule(context.Background()).CreateUserScheduleRequest(createUserScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.CreateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.CreateSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserScheduleRequest** | [**CreateUserScheduleRequest**](CreateUserScheduleRequest.md) | In the request payload: feedType and scheduleTemplateId are required; scheduleName is optional; preferredTriggerHour, preferredTriggerDayOfWeek, preferredTriggerDayOfMonth, scheduleStartDate, scheduleEndDate, and schemaVersion are conditional. | 

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


## DeleteSchedule

> DeleteSchedule(ctx, scheduleId).Execute()





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
    scheduleId := "scheduleId_example" // string | The schedule_id of the schedule to delete. This ID was generated when the task was created. If you do not know the schedule_id, use the getSchedules method to return all schedules based on a specified feed_type and find the schedule_id of the schedule to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.DeleteSchedule(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.DeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The schedule_id of the schedule to delete. This ID was generated when the task was created. If you do not know the schedule_id, use the getSchedules method to return all schedules based on a specified feed_type and find the schedule_id of the schedule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestResultFile

> map[string]interface{} GetLatestResultFile(ctx, scheduleId).Execute()





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
    scheduleId := "scheduleId_example" // string | The ID of the schedule for which to retrieve the latest result file. This ID is generated when the schedule was created by the createSchedule method.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.GetLatestResultFile(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetLatestResultFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestResultFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetLatestResultFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The ID of the schedule for which to retrieve the latest result file. This ID is generated when the schedule was created by the createSchedule method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestResultFileRequest struct via the builder pattern


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


## GetSchedule

> UserScheduleResponse GetSchedule(ctx, scheduleId).Execute()





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
    scheduleId := "scheduleId_example" // string | The ID of the schedule for which to retrieve the details. This ID is generated when the schedule was created by the createSchedule method.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.GetSchedule(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedule`: UserScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The ID of the schedule for which to retrieve the details. This ID is generated when the schedule was created by the createSchedule method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserScheduleResponse**](UserScheduleResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduleTemplate

> ScheduleTemplateResponse GetScheduleTemplate(ctx, scheduleTemplateId).Execute()





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
    scheduleTemplateId := "scheduleTemplateId_example" // string | The ID of the template to retrieve. If you do not know the schedule_template_id, refer to the documentation or use the getScheduleTemplates method to find the available schedule templates.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.GetScheduleTemplate(context.Background(), scheduleTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetScheduleTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduleTemplate`: ScheduleTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetScheduleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleTemplateId** | **string** | The ID of the template to retrieve. If you do not know the schedule_template_id, refer to the documentation or use the getScheduleTemplates method to find the available schedule templates. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduleTemplateResponse**](ScheduleTemplateResponse.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduleTemplates

> ScheduleTemplateCollection GetScheduleTemplates(ctx).FeedType(feedType).Limit(limit).Offset(offset).Execute()





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
    feedType := "feedType_example" // string | The feed type of the schedule templates to retrieve.
    limit := "limit_example" // string | The maximum number of schedule templates that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves schedule templates 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 (optional)
    offset := "offset_example" // string | The number of schedule templates to skip in the result set before returning the first template in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.GetScheduleTemplates(context.Background()).FeedType(feedType).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetScheduleTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduleTemplates`: ScheduleTemplateCollection
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetScheduleTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedType** | **string** | The feed type of the schedule templates to retrieve. | 
 **limit** | **string** | The maximum number of schedule templates that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves schedule templates 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 | 
 **offset** | **string** | The number of schedule templates to skip in the result set before returning the first template in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 | 

### Return type

[**ScheduleTemplateCollection**](ScheduleTemplateCollection.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedules

> UserScheduleCollection GetSchedules(ctx).FeedType(feedType).Limit(limit).Offset(offset).Execute()





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
    feedType := "feedType_example" // string | The feedType associated with the schedule.
    limit := "limit_example" // string | The maximum number of schedules that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves schedules 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 (optional)
    offset := "offset_example" // string | The number of schedules to skip in the result set before returning the first schedule in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.GetSchedules(context.Background()).FeedType(feedType).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.GetSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedules`: UserScheduleCollection
    fmt.Fprintf(os.Stdout, "Response from `ScheduleApi.GetSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedType** | **string** | The feedType associated with the schedule. | 
 **limit** | **string** | The maximum number of schedules that can be returned on each page of the paginated response. Use this parameter in conjunction with the offset parameter to control the pagination of the output. Note: This feature employs a zero-based list, where the first item in the list has an offset of 0. For example, if offset is set to 10 and limit is set to 10, the call retrieves schedules 11 thru 20 from the result set. If this parameter is omitted, the default value is used. Default: 10 Maximum: 500 | 
 **offset** | **string** | The number of schedules to skip in the result set before returning the first schedule in the paginated response. Combine offset with the limit query parameter to control the items returned in the response. For example, if you supply an offset of 0 and a limit of 10, the first page of the response contains the first 10 items from the complete list of items retrieved by the call. If offset is 10 and limit is 20, the first page of the response contains items 11-30 from the complete result set. If this query parameter is not set, the default value is used and the first page of records is returned. Default: 0 | 

### Return type

[**UserScheduleCollection**](UserScheduleCollection.md)

### Authorization

[api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchedule

> UpdateSchedule(ctx, scheduleId).UpdateUserScheduleRequest(updateUserScheduleRequest).Execute()





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
    scheduleId := "scheduleId_example" // string | The ID of the schedule to update. This ID is generated when the schedule was created by the createSchedule method.
    updateUserScheduleRequest := *openapiclient.NewUpdateUserScheduleRequest() // UpdateUserScheduleRequest | In the request payload: scheduleName is optional; preferredTriggerHour, preferredTriggerDayOfWeek, preferredTriggerDayOfMonth, scheduleStartDate, scheduleEndDate, and schemaVersion are conditional.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleApi.UpdateSchedule(context.Background(), scheduleId).UpdateUserScheduleRequest(updateUserScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleApi.UpdateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | The ID of the schedule to update. This ID is generated when the schedule was created by the createSchedule method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserScheduleRequest** | [**UpdateUserScheduleRequest**](UpdateUserScheduleRequest.md) | In the request payload: scheduleName is optional; preferredTriggerHour, preferredTriggerDayOfWeek, preferredTriggerDayOfMonth, scheduleStartDate, scheduleEndDate, and schemaVersion are conditional. | 

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

