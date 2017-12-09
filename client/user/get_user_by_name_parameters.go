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

// NewGetUserByNameParams creates a new GetUserByNameParams object
// with the default values initialized.
func NewGetUserByNameParams() *GetUserByNameParams {
	var ()
	return &GetUserByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserByNameParamsWithTimeout creates a new GetUserByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserByNameParamsWithTimeout(timeout time.Duration) *GetUserByNameParams {
	var ()
	return &GetUserByNameParams{

		timeout: timeout,
	}
}

// NewGetUserByNameParamsWithContext creates a new GetUserByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserByNameParamsWithContext(ctx context.Context) *GetUserByNameParams {
	var ()
	return &GetUserByNameParams{

		Context: ctx,
	}
}

// NewGetUserByNameParamsWithHTTPClient creates a new GetUserByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserByNameParamsWithHTTPClient(client *http.Client) *GetUserByNameParams {
	var ()
	return &GetUserByNameParams{
		HTTPClient: client,
	}
}

/*GetUserByNameParams contains all the parameters to send to the API endpoint
for the get user by name operation typically these are written to a http.Request
*/
type GetUserByNameParams struct {

	/*Username
	  The name that needs to be fetched. Use user1 for testing.

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user by name params
func (o *GetUserByNameParams) WithTimeout(timeout time.Duration) *GetUserByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user by name params
func (o *GetUserByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user by name params
func (o *GetUserByNameParams) WithContext(ctx context.Context) *GetUserByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user by name params
func (o *GetUserByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user by name params
func (o *GetUserByNameParams) WithHTTPClient(client *http.Client) *GetUserByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user by name params
func (o *GetUserByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get user by name params
func (o *GetUserByNameParams) WithUsername(username string) *GetUserByNameParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get user by name params
func (o *GetUserByNameParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
