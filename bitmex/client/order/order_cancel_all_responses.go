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

// OrderCancelAllReader is a Reader for the OrderCancelAll structure.
type OrderCancelAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderCancelAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderCancelAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderCancelAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewOrderCancelAllUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderCancelAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderCancelAllOK creates a OrderCancelAllOK with default headers values
func NewOrderCancelAllOK() *OrderCancelAllOK {
	return &OrderCancelAllOK{}
}

/*OrderCancelAllOK handles this case with default header values.

Request was successful
*/
type OrderCancelAllOK struct {
	Payload []*models.Order
}

func (o *OrderCancelAllOK) Error() string {
	return fmt.Sprintf("[DELETE /order/all][%d] orderCancelAllOK  %+v", 200, o.Payload)
}

func (o *OrderCancelAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllBadRequest creates a OrderCancelAllBadRequest with default headers values
func NewOrderCancelAllBadRequest() *OrderCancelAllBadRequest {
	return &OrderCancelAllBadRequest{}
}

/*OrderCancelAllBadRequest handles this case with default header values.

Parameter Error
*/
type OrderCancelAllBadRequest struct {
	Payload *models.Error
}

func (o *OrderCancelAllBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /order/all][%d] orderCancelAllBadRequest  %+v", 400, *o.Payload)
}

func (o *OrderCancelAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllUnauthorized creates a OrderCancelAllUnauthorized with default headers values
func NewOrderCancelAllUnauthorized() *OrderCancelAllUnauthorized {
	return &OrderCancelAllUnauthorized{}
}

/*OrderCancelAllUnauthorized handles this case with default header values.

Unauthorized
*/
type OrderCancelAllUnauthorized struct {
	Payload *models.Error
}

func (o *OrderCancelAllUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /order/all][%d] orderCancelAllUnauthorized  %+v", 401, o.Payload)
}

func (o *OrderCancelAllUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderCancelAllNotFound creates a OrderCancelAllNotFound with default headers values
func NewOrderCancelAllNotFound() *OrderCancelAllNotFound {
	return &OrderCancelAllNotFound{}
}

/*OrderCancelAllNotFound handles this case with default header values.

Not Found
*/
type OrderCancelAllNotFound struct {
	Payload *models.Error
}

func (o *OrderCancelAllNotFound) Error() string {
	return fmt.Sprintf("[DELETE /order/all][%d] orderCancelAllNotFound  %+v", 404, o.Payload)
}

func (o *OrderCancelAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
