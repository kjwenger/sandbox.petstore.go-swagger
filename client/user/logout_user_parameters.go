// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLogoutUserParams creates a new LogoutUserParams object
// with the default values initialized.
func NewLogoutUserParams() *LogoutUserParams {

	return &LogoutUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogoutUserParamsWithTimeout creates a new LogoutUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogoutUserParamsWithTimeout(timeout time.Duration) *LogoutUserParams {

	return &LogoutUserParams{

		timeout: timeout,
	}
}

// NewLogoutUserParamsWithContext creates a new LogoutUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogoutUserParamsWithContext(ctx context.Context) *LogoutUserParams {

	return &LogoutUserParams{

		Context: ctx,
	}
}

// NewLogoutUserParamsWithHTTPClient creates a new LogoutUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogoutUserParamsWithHTTPClient(client *http.Client) *LogoutUserParams {

	return &LogoutUserParams{
		HTTPClient: client,
	}
}

/*LogoutUserParams contains all the parameters to send to the API endpoint
for the logout user operation typically these are written to a http.Request
*/
type LogoutUserParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logout user params
func (o *LogoutUserParams) WithTimeout(timeout time.Duration) *LogoutUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logout user params
func (o *LogoutUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logout user params
func (o *LogoutUserParams) WithContext(ctx context.Context) *LogoutUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logout user params
func (o *LogoutUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logout user params
func (o *LogoutUserParams) WithHTTPClient(client *http.Client) *LogoutUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logout user params
func (o *LogoutUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogoutUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
