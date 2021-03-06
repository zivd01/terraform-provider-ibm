// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudCloudinstancesGetParams creates a new PcloudCloudinstancesGetParams object
// with the default values initialized.
func NewPcloudCloudinstancesGetParams() *PcloudCloudinstancesGetParams {
	var ()
	return &PcloudCloudinstancesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesGetParamsWithTimeout creates a new PcloudCloudinstancesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesGetParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesGetParams {
	var ()
	return &PcloudCloudinstancesGetParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesGetParamsWithContext creates a new PcloudCloudinstancesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesGetParamsWithContext(ctx context.Context) *PcloudCloudinstancesGetParams {
	var ()
	return &PcloudCloudinstancesGetParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesGetParamsWithHTTPClient creates a new PcloudCloudinstancesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesGetParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesGetParams {
	var ()
	return &PcloudCloudinstancesGetParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesGetParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances get operation typically these are written to a http.Request
*/
type PcloudCloudinstancesGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) WithContext(ctx context.Context) *PcloudCloudinstancesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances get params
func (o *PcloudCloudinstancesGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
