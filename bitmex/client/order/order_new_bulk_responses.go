// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// OrderNewBulkReader is a Reader for the OrderNewBulk structure.
type OrderNewBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderNewBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderNewBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderNewBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderNewBulkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderNewBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderNewBulkOK creates a OrderNewBulkOK with default headers values
func NewOrderNewBulkOK() *OrderNewBulkOK {
	return &OrderNewBulkOK{}
}

/*OrderNewBulkOK handles this case with default header values.

Request was successful
*/
type OrderNewBulkOK struct {
	Payload []*models.Order
}

func (o *OrderNewBulkOK) Error() string {
	return fmt.Sprintf("[POST /order/bulk][%d] orderNewBulkOK  %+v", 200, o.Payload)
}

func (o *OrderNewBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewBulkBadRequest creates a OrderNewBulkBadRequest with default headers values
func NewOrderNewBulkBadRequest() *OrderNewBulkBadRequest {
	return &OrderNewBulkBadRequest{}
}

/*OrderNewBulkBadRequest handles this case with default header values.

Parameter Error
*/
type OrderNewBulkBadRequest struct {
	Payload *models.Error
}

func (o *OrderNewBulkBadRequest) Error() string {
	return fmt.Sprintf("[POST /order/bulk][%d] orderNewBulkBadRequest  %+v", 400, *o.Payload)
}

func (o *OrderNewBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewBulkUnauthorized creates a OrderNewBulkUnauthorized with default headers values
func NewOrderNewBulkUnauthorized() *OrderNewBulkUnauthorized {
	return &OrderNewBulkUnauthorized{}
}

/*OrderNewBulkUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderNewBulkUnauthorized struct {
	Payload *models.Error
}

func (o *OrderNewBulkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /order/bulk][%d] orderNewBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderNewBulkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderNewBulkNotFound creates a OrderNewBulkNotFound with default headers values
func NewOrderNewBulkNotFound() *OrderNewBulkNotFound {
	return &OrderNewBulkNotFound{}
}

/*OrderNewBulkNotFound handles this case with default header values.

Not Found
*/
type OrderNewBulkNotFound struct {
	Payload *models.Error
}

func (o *OrderNewBulkNotFound) Error() string {
	return fmt.Sprintf("[POST /order/bulk][%d] orderNewBulkNotFound  %+v", 404, o.Payload)
}

func (o *OrderNewBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
