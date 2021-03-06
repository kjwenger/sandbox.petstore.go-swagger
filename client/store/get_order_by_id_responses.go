// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kjwenger/sandbox.petstore.go-swagger/models"
)

// GetOrderByIDReader is a Reader for the GetOrderByID structure.
type GetOrderByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrderByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrderByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrderByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrderByIDOK creates a GetOrderByIDOK with default headers values
func NewGetOrderByIDOK() *GetOrderByIDOK {
	return &GetOrderByIDOK{}
}

/*GetOrderByIDOK handles this case with default header values.

successful operation
*/
type GetOrderByIDOK struct {
	Payload *models.Order
}

func (o *GetOrderByIDOK) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] getOrderByIdOK  %+v", 200, o.Payload)
}

func (o *GetOrderByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderByIDBadRequest creates a GetOrderByIDBadRequest with default headers values
func NewGetOrderByIDBadRequest() *GetOrderByIDBadRequest {
	return &GetOrderByIDBadRequest{}
}

/*GetOrderByIDBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetOrderByIDBadRequest struct {
}

func (o *GetOrderByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] getOrderByIdBadRequest ", 400)
}

func (o *GetOrderByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrderByIDNotFound creates a GetOrderByIDNotFound with default headers values
func NewGetOrderByIDNotFound() *GetOrderByIDNotFound {
	return &GetOrderByIDNotFound{}
}

/*GetOrderByIDNotFound handles this case with default header values.

Order not found
*/
type GetOrderByIDNotFound struct {
}

func (o *GetOrderByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] getOrderByIdNotFound ", 404)
}

func (o *GetOrderByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
