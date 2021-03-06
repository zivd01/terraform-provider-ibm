// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceBrokerTestTimeoutParams creates a new ServiceBrokerTestTimeoutParams object
// with the default values initialized.
func NewServiceBrokerTestTimeoutParams() *ServiceBrokerTestTimeoutParams {
	var ()
	return &ServiceBrokerTestTimeoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerTestTimeoutParamsWithTimeout creates a new ServiceBrokerTestTimeoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServiceBrokerTestTimeoutParamsWithTimeout(timeout time.Duration) *ServiceBrokerTestTimeoutParams {
	var ()
	return &ServiceBrokerTestTimeoutParams{

		timeout: timeout,
	}
}

// NewServiceBrokerTestTimeoutParamsWithContext creates a new ServiceBrokerTestTimeoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewServiceBrokerTestTimeoutParamsWithContext(ctx context.Context) *ServiceBrokerTestTimeoutParams {
	var ()
	return &ServiceBrokerTestTimeoutParams{

		Context: ctx,
	}
}

// NewServiceBrokerTestTimeoutParamsWithHTTPClient creates a new ServiceBrokerTestTimeoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServiceBrokerTestTimeoutParamsWithHTTPClient(client *http.Client) *ServiceBrokerTestTimeoutParams {
	var ()
	return &ServiceBrokerTestTimeoutParams{
		HTTPClient: client,
	}
}

/*ServiceBrokerTestTimeoutParams contains all the parameters to send to the API endpoint
for the service broker test timeout operation typically these are written to a http.Request
*/
type ServiceBrokerTestTimeoutParams struct {

	/*T
	  seconds

	*/
	T int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) WithTimeout(timeout time.Duration) *ServiceBrokerTestTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) WithContext(ctx context.Context) *ServiceBrokerTestTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) WithHTTPClient(client *http.Client) *ServiceBrokerTestTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithT adds the t to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) WithT(t int64) *ServiceBrokerTestTimeoutParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the service broker test timeout params
func (o *ServiceBrokerTestTimeoutParams) SetT(t int64) {
	o.T = t
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerTestTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param t
	qrT := o.T
	qT := swag.FormatInt64(qrT)
	if qT != "" {
		if err := r.SetQueryParam("t", qT); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
