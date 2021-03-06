// Code generated by go-swagger; DO NOT EDIT.

package c_o_r_s

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

// NewOptionsAccountsParams creates a new OptionsAccountsParams object
// with the default values initialized.
func NewOptionsAccountsParams() *OptionsAccountsParams {

	return &OptionsAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOptionsAccountsParamsWithTimeout creates a new OptionsAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOptionsAccountsParamsWithTimeout(timeout time.Duration) *OptionsAccountsParams {

	return &OptionsAccountsParams{

		timeout: timeout,
	}
}

// NewOptionsAccountsParamsWithContext creates a new OptionsAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewOptionsAccountsParamsWithContext(ctx context.Context) *OptionsAccountsParams {

	return &OptionsAccountsParams{

		Context: ctx,
	}
}

// NewOptionsAccountsParamsWithHTTPClient creates a new OptionsAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOptionsAccountsParamsWithHTTPClient(client *http.Client) *OptionsAccountsParams {

	return &OptionsAccountsParams{
		HTTPClient: client,
	}
}

/*OptionsAccountsParams contains all the parameters to send to the API endpoint
for the options accounts operation typically these are written to a http.Request
*/
type OptionsAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the options accounts params
func (o *OptionsAccountsParams) WithTimeout(timeout time.Duration) *OptionsAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the options accounts params
func (o *OptionsAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the options accounts params
func (o *OptionsAccountsParams) WithContext(ctx context.Context) *OptionsAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the options accounts params
func (o *OptionsAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the options accounts params
func (o *OptionsAccountsParams) WithHTTPClient(client *http.Client) *OptionsAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the options accounts params
func (o *OptionsAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OptionsAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
