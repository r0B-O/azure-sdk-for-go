// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// QueueServicesOperations contains the methods for the QueueServices group.
type QueueServicesOperations interface {
	// GetServiceProperties - Gets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
	// Sharing) rules.
	GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesGetServicePropertiesOptions) (*QueueServicePropertiesResponse, error)
	// List - List all queue services for the storage account
	List(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesListOptions) (*ListQueueServicesResponse, error)
	// SetServiceProperties - Sets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
	// Sharing) rules.
	SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties, options *QueueServicesSetServicePropertiesOptions) (*QueueServicePropertiesResponse, error)
}

// QueueServicesClient implements the QueueServicesOperations interface.
// Don't use this type directly, use NewQueueServicesClient() instead.
type QueueServicesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewQueueServicesClient creates a new instance of QueueServicesClient with the specified values.
func NewQueueServicesClient(con *armcore.Connection, subscriptionID string) QueueServicesOperations {
	return &QueueServicesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *QueueServicesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetServiceProperties - Gets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
func (client *QueueServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesGetServicePropertiesOptions) (*QueueServicePropertiesResponse, error) {
	req, err := client.GetServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetServicePropertiesHandleError(resp)
	}
	result, err := client.GetServicePropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *QueueServicesClient) GetServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesGetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *QueueServicesClient) GetServicePropertiesHandleResponse(resp *azcore.Response) (*QueueServicePropertiesResponse, error) {
	result := QueueServicePropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.QueueServiceProperties)
}

// GetServicePropertiesHandleError handles the GetServiceProperties error response.
func (client *QueueServicesClient) GetServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List all queue services for the storage account
func (client *QueueServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesListOptions) (*ListQueueServicesResponse, error) {
	req, err := client.ListCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result, err := client.ListHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCreateRequest creates the List request.
func (client *QueueServicesClient) ListCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *QueueServicesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *QueueServicesClient) ListHandleResponse(resp *azcore.Response) (*ListQueueServicesResponse, error) {
	result := ListQueueServicesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListQueueServices)
}

// ListHandleError handles the List error response.
func (client *QueueServicesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetServiceProperties - Sets the properties of a storage account’s Queue service, including properties for Storage Analytics and CORS (Cross-Origin Resource
// Sharing) rules.
func (client *QueueServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties, options *QueueServicesSetServicePropertiesOptions) (*QueueServicePropertiesResponse, error) {
	req, err := client.SetServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.SetServicePropertiesHandleError(resp)
	}
	result, err := client.SetServicePropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *QueueServicesClient) SetServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters QueueServiceProperties, options *QueueServicesSetServicePropertiesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{queueServiceName}", url.PathEscape("default"))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// SetServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *QueueServicesClient) SetServicePropertiesHandleResponse(resp *azcore.Response) (*QueueServicePropertiesResponse, error) {
	result := QueueServicePropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.QueueServiceProperties)
}

// SetServicePropertiesHandleError handles the SetServiceProperties error response.
func (client *QueueServicesClient) SetServicePropertiesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
