package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MetricsClient is the composite Swagger for Application Insights Data Client
type MetricsClient struct {
	BaseClient
}

// NewMetricsClient creates an instance of the MetricsClient client.
func NewMetricsClient() MetricsClient {
	return NewMetricsClientWithBaseURI(DefaultBaseURI)
}

// NewMetricsClientWithBaseURI creates an instance of the MetricsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMetricsClientWithBaseURI(baseURI string) MetricsClient {
	return MetricsClient{NewWithBaseURI(baseURI)}
}

// Get gets metric values for a single metric
// Parameters:
// appID - ID of the application. This is Application ID from the API Access settings blade in the Azure
// portal.
// metricID - ID of the metric. This is either a standard AI metric, or an application-specific custom metric.
// timespan - the timespan over which to retrieve metric values. This is an ISO8601 time period value. If
// timespan is omitted, a default time range of `PT12H` ("last 12 hours") is used. The actual timespan that is
// queried may be adjusted by the server based. In all cases, the actual time span used for the query is
// included in the response.
// interval - the time interval to use when retrieving metric values. This is an ISO8601 duration. If interval
// is omitted, the metric value is aggregated across the entire timespan. If interval is supplied, the server
// may adjust the interval to a more appropriate size based on the timespan used for the query. In all cases,
// the actual interval used for the query is included in the response.
// aggregation - the aggregation to use when computing the metric values. To retrieve more than one aggregation
// at a time, separate them with a comma. If no aggregation is specified, then the default aggregation for the
// metric is used.
// segment - the name of the dimension to segment the metric values by. This dimension must be applicable to
// the metric you are retrieving. To segment by more than one dimension at a time, separate them with a comma
// (,). In this case, the metric data will be segmented in the order the dimensions are listed in the
// parameter.
// top - the number of segments to return.  This value is only valid when segment is specified.
// orderby - the aggregation function and direction to sort the segments by.  This value is only valid when
// segment is specified.
// filter - an expression used to filter the results.  This value should be a valid OData filter expression
// where the keys of each clause should be applicable dimensions for the metric you are retrieving.
func (client MetricsClient) Get(ctx context.Context, appID string, metricID MetricID, timespan string, interval *string, aggregation []MetricsAggregation, segment []MetricsSegment, top *int32, orderby string, filter string) (result MetricsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: aggregation,
			Constraints: []validation.Constraint{{Target: "aggregation", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "aggregation", Name: validation.MinItems, Rule: 1, Chain: nil}}}}},
		{TargetValue: segment,
			Constraints: []validation.Constraint{{Target: "segment", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "segment", Name: validation.MinItems, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("insights.MetricsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, appID, metricID, timespan, interval, aggregation, segment, top, orderby, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MetricsClient) GetPreparer(ctx context.Context, appID string, metricID MetricID, timespan string, interval *string, aggregation []MetricsAggregation, segment []MetricsSegment, top *int32, orderby string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId":    autorest.Encode("path", appID),
		"metricId": autorest.Encode("path", metricID),
	}

	queryParameters := map[string]interface{}{}
	if len(timespan) > 0 {
		queryParameters["timespan"] = autorest.Encode("query", timespan)
	}
	if interval != nil {
		queryParameters["interval"] = autorest.Encode("query", *interval)
	}
	if aggregation != nil && len(aggregation) > 0 {
		queryParameters["aggregation"] = autorest.Encode("query", aggregation, ",")
	}
	if segment != nil && len(segment) > 0 {
		queryParameters["segment"] = autorest.Encode("query", segment, ",")
	}
	if top != nil {
		queryParameters["top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["orderby"] = autorest.Encode("query", orderby)
	}
	if len(filter) > 0 {
		queryParameters["filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/apps/{appId}/metrics/{metricId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MetricsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MetricsClient) GetResponder(resp *http.Response) (result MetricsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMetadata gets metadata describing the available metrics
// Parameters:
// appID - ID of the application. This is Application ID from the API Access settings blade in the Azure
// portal.
func (client MetricsClient) GetMetadata(ctx context.Context, appID string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricsClient.GetMetadata")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetMetadataPreparer(ctx, appID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMetadata", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMetadataSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMetadata", resp, "Failure sending request")
		return
	}

	result, err = client.GetMetadataResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMetadata", resp, "Failure responding to request")
	}

	return
}

// GetMetadataPreparer prepares the GetMetadata request.
func (client MetricsClient) GetMetadataPreparer(ctx context.Context, appID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/apps/{appId}/metrics/metadata", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMetadataSender sends the GetMetadata request. The method will close the
// http.Response Body if it receives an error.
func (client MetricsClient) GetMetadataSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetMetadataResponder handles the response to the GetMetadata request. The method always
// closes the http.Response Body.
func (client MetricsClient) GetMetadataResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMultiple gets metric values for multiple metrics
// Parameters:
// appID - ID of the application. This is Application ID from the API Access settings blade in the Azure
// portal.
// body - the batched metrics query.
func (client MetricsClient) GetMultiple(ctx context.Context, appID string, body []MetricsPostBodySchema) (result ListMetricsResultsItem, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MetricsClient.GetMultiple")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MetricsClient", "GetMultiple", err.Error())
	}

	req, err := client.GetMultiplePreparer(ctx, appID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMultiple", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMultipleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMultiple", resp, "Failure sending request")
		return
	}

	result, err = client.GetMultipleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MetricsClient", "GetMultiple", resp, "Failure responding to request")
	}

	return
}

// GetMultiplePreparer prepares the GetMultiple request.
func (client MetricsClient) GetMultiplePreparer(ctx context.Context, appID string, body []MetricsPostBodySchema) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/apps/{appId}/metrics", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMultipleSender sends the GetMultiple request. The method will close the
// http.Response Body if it receives an error.
func (client MetricsClient) GetMultipleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetMultipleResponder handles the response to the GetMultiple request. The method always
// closes the http.Response Body.
func (client MetricsClient) GetMultipleResponder(resp *http.Response) (result ListMetricsResultsItem, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
