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

// TableOperations contains the methods for the Table group.
type TableOperations interface {
	// Create - Creates a new table with the specified table name, under the specified account.
	Create(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableCreateOptions) (*TableResponse, error)
	// Delete - Deletes the table with the specified table name, under the specified account if it exists.
	Delete(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableDeleteOptions) (*http.Response, error)
	// Get - Gets the table with the specified table name, under the specified account if it exists.
	Get(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableGetOptions) (*TableResponse, error)
	// List - Gets a list of all the tables under the specified storage account
	List(resourceGroupName string, accountName string, options *TableListOptions) ListTableResourcePager
	// Update - Creates a new table with the specified table name, under the specified account.
	Update(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableUpdateOptions) (*TableResponse, error)
}

// TableClient implements the TableOperations interface.
// Don't use this type directly, use NewTableClient() instead.
type TableClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewTableClient creates a new instance of TableClient with the specified values.
func NewTableClient(con *armcore.Connection, subscriptionID string) TableOperations {
	return &TableClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *TableClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Create - Creates a new table with the specified table name, under the specified account.
func (client *TableClient) Create(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableCreateOptions) (*TableResponse, error) {
	req, err := client.CreateCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateHandleError(resp)
	}
	result, err := client.CreateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCreateRequest creates the Create request.
func (client *TableClient) CreateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// CreateHandleResponse handles the Create response.
func (client *TableClient) CreateHandleResponse(resp *azcore.Response) (*TableResponse, error) {
	result := TableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Table)
}

// CreateHandleError handles the Create error response.
func (client *TableClient) CreateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes the table with the specified table name, under the specified account if it exists.
func (client *TableClient) Delete(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableDeleteOptions) (*http.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *TableClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// DeleteHandleError handles the Delete error response.
func (client *TableClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the table with the specified table name, under the specified account if it exists.
func (client *TableClient) Get(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableGetOptions) (*TableResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *TableClient) GetCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
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

// GetHandleResponse handles the Get response.
func (client *TableClient) GetHandleResponse(resp *azcore.Response) (*TableResponse, error) {
	result := TableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Table)
}

// GetHandleError handles the Get error response.
func (client *TableClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets a list of all the tables under the specified storage account
func (client *TableClient) List(resourceGroupName string, accountName string, options *TableListOptions) ListTableResourcePager {
	return &listTableResourcePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListTableResourceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListTableResource.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *TableClient) ListCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables"
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
func (client *TableClient) ListHandleResponse(resp *azcore.Response) (*ListTableResourceResponse, error) {
	result := ListTableResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListTableResource)
}

// ListHandleError handles the List error response.
func (client *TableClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Update - Creates a new table with the specified table name, under the specified account.
func (client *TableClient) Update(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableUpdateOptions) (*TableResponse, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	result, err := client.UpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCreateRequest creates the Update request.
func (client *TableClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// UpdateHandleResponse handles the Update response.
func (client *TableClient) UpdateHandleResponse(resp *azcore.Response) (*TableResponse, error) {
	result := TableResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Table)
}

// UpdateHandleError handles the Update error response.
func (client *TableClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
