// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetLeasesIDParams creates a new GetLeasesIDParams object
// with the default values initialized.
func NewGetLeasesIDParams() *GetLeasesIDParams {
	var ()
	return &GetLeasesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLeasesIDParamsWithTimeout creates a new GetLeasesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLeasesIDParamsWithTimeout(timeout time.Duration) *GetLeasesIDParams {
	var ()
	return &GetLeasesIDParams{

		timeout: timeout,
	}
}

// NewGetLeasesIDParamsWithContext creates a new GetLeasesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLeasesIDParamsWithContext(ctx context.Context) *GetLeasesIDParams {
	var ()
	return &GetLeasesIDParams{

		Context: ctx,
	}
}

// NewGetLeasesIDParamsWithHTTPClient creates a new GetLeasesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLeasesIDParamsWithHTTPClient(client *http.Client) *GetLeasesIDParams {
	var ()
	return &GetLeasesIDParams{
		HTTPClient: client,
	}
}

/*GetLeasesIDParams contains all the parameters to send to the API endpoint
for the get leases ID operation typically these are written to a http.Request
*/
type GetLeasesIDParams struct {

	/*ID
	  Id for lease

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get leases ID params
func (o *GetLeasesIDParams) WithTimeout(timeout time.Duration) *GetLeasesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get leases ID params
func (o *GetLeasesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get leases ID params
func (o *GetLeasesIDParams) WithContext(ctx context.Context) *GetLeasesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get leases ID params
func (o *GetLeasesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get leases ID params
func (o *GetLeasesIDParams) WithHTTPClient(client *http.Client) *GetLeasesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get leases ID params
func (o *GetLeasesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get leases ID params
func (o *GetLeasesIDParams) WithID(id string) *GetLeasesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get leases ID params
func (o *GetLeasesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLeasesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}