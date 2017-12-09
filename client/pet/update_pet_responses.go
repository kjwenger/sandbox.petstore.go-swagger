// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdatePetReader is a Reader for the UpdatePet structure.
type UpdatePetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewUpdatePetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdatePetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewUpdatePetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePetBadRequest creates a UpdatePetBadRequest with default headers values
func NewUpdatePetBadRequest() *UpdatePetBadRequest {
	return &UpdatePetBadRequest{}
}

/*UpdatePetBadRequest handles this case with default header values.

Invalid ID supplied
*/
type UpdatePetBadRequest struct {
}

func (o *UpdatePetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] updatePetBadRequest ", 400)
}

func (o *UpdatePetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePetNotFound creates a UpdatePetNotFound with default headers values
func NewUpdatePetNotFound() *UpdatePetNotFound {
	return &UpdatePetNotFound{}
}

/*UpdatePetNotFound handles this case with default header values.

Pet not found
*/
type UpdatePetNotFound struct {
}

func (o *UpdatePetNotFound) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] updatePetNotFound ", 404)
}

func (o *UpdatePetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePetMethodNotAllowed creates a UpdatePetMethodNotAllowed with default headers values
func NewUpdatePetMethodNotAllowed() *UpdatePetMethodNotAllowed {
	return &UpdatePetMethodNotAllowed{}
}

/*UpdatePetMethodNotAllowed handles this case with default header values.

Validation exception
*/
type UpdatePetMethodNotAllowed struct {
}

func (o *UpdatePetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /pet][%d] updatePetMethodNotAllowed ", 405)
}

func (o *UpdatePetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
