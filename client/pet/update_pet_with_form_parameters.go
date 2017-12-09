// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdatePetWithFormParams creates a new UpdatePetWithFormParams object
// with the default values initialized.
func NewUpdatePetWithFormParams() *UpdatePetWithFormParams {
	var ()
	return &UpdatePetWithFormParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePetWithFormParamsWithTimeout creates a new UpdatePetWithFormParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePetWithFormParamsWithTimeout(timeout time.Duration) *UpdatePetWithFormParams {
	var ()
	return &UpdatePetWithFormParams{

		timeout: timeout,
	}
}

// NewUpdatePetWithFormParamsWithContext creates a new UpdatePetWithFormParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePetWithFormParamsWithContext(ctx context.Context) *UpdatePetWithFormParams {
	var ()
	return &UpdatePetWithFormParams{

		Context: ctx,
	}
}

// NewUpdatePetWithFormParamsWithHTTPClient creates a new UpdatePetWithFormParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePetWithFormParamsWithHTTPClient(client *http.Client) *UpdatePetWithFormParams {
	var ()
	return &UpdatePetWithFormParams{
		HTTPClient: client,
	}
}

/*UpdatePetWithFormParams contains all the parameters to send to the API endpoint
for the update pet with form operation typically these are written to a http.Request
*/
type UpdatePetWithFormParams struct {

	/*Name
	  Updated name of the pet

	*/
	Name *string
	/*PetID
	  ID of pet that needs to be updated

	*/
	PetID int64
	/*Status
	  Updated status of the pet

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update pet with form params
func (o *UpdatePetWithFormParams) WithTimeout(timeout time.Duration) *UpdatePetWithFormParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pet with form params
func (o *UpdatePetWithFormParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pet with form params
func (o *UpdatePetWithFormParams) WithContext(ctx context.Context) *UpdatePetWithFormParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pet with form params
func (o *UpdatePetWithFormParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pet with form params
func (o *UpdatePetWithFormParams) WithHTTPClient(client *http.Client) *UpdatePetWithFormParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pet with form params
func (o *UpdatePetWithFormParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update pet with form params
func (o *UpdatePetWithFormParams) WithName(name *string) *UpdatePetWithFormParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update pet with form params
func (o *UpdatePetWithFormParams) SetName(name *string) {
	o.Name = name
}

// WithPetID adds the petID to the update pet with form params
func (o *UpdatePetWithFormParams) WithPetID(petID int64) *UpdatePetWithFormParams {
	o.SetPetID(petID)
	return o
}

// SetPetID adds the petId to the update pet with form params
func (o *UpdatePetWithFormParams) SetPetID(petID int64) {
	o.PetID = petID
}

// WithStatus adds the status to the update pet with form params
func (o *UpdatePetWithFormParams) WithStatus(status *string) *UpdatePetWithFormParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the update pet with form params
func (o *UpdatePetWithFormParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePetWithFormParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	// path param petId
	if err := r.SetPathParam("petId", swag.FormatInt64(o.PetID)); err != nil {
		return err
	}

	if o.Status != nil {

		// form param status
		var frStatus string
		if o.Status != nil {
			frStatus = *o.Status
		}
		fStatus := frStatus
		if fStatus != "" {
			if err := r.SetFormParam("status", fStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
