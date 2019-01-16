// Code generated by go-swagger; DO NOT EDIT.

package quote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// QuoteGetReader is a Reader for the QuoteGet structure.
type QuoteGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuoteGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQuoteGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewQuoteGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewQuoteGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewQuoteGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQuoteGetOK creates a QuoteGetOK with default headers values
func NewQuoteGetOK() *QuoteGetOK {
	return &QuoteGetOK{}
}

/*QuoteGetOK handles this case with default header values.

Request was successful
*/
type QuoteGetOK struct {
	Payload []*models.Quote
}

func (o *QuoteGetOK) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteGetOK  %+v", 200, o.Payload)
}

func (o *QuoteGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetBadRequest creates a QuoteGetBadRequest with default headers values
func NewQuoteGetBadRequest() *QuoteGetBadRequest {
	return &QuoteGetBadRequest{}
}

/*QuoteGetBadRequest handles this case with default header values.

Parameter Error
*/
type QuoteGetBadRequest struct {
	Payload *models.Error
}

func (o *QuoteGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteGetBadRequest  %+v", 400, *o.Payload)
}

func (o *QuoteGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetUnauthorized creates a QuoteGetUnauthorized with default headers values
func NewQuoteGetUnauthorized() *QuoteGetUnauthorized {
	return &QuoteGetUnauthorized{}
}

/*QuoteGetUnauthorized handles this case with default header values.

Unauthorized
*/
type QuoteGetUnauthorized struct {
	Payload *models.Error
}

func (o *QuoteGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteGetUnauthorized  %+v", 401, o.Payload)
}

func (o *QuoteGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuoteGetNotFound creates a QuoteGetNotFound with default headers values
func NewQuoteGetNotFound() *QuoteGetNotFound {
	return &QuoteGetNotFound{}
}

/*QuoteGetNotFound handles this case with default header values.

Not Found
*/
type QuoteGetNotFound struct {
	Payload *models.Error
}

func (o *QuoteGetNotFound) Error() string {
	return fmt.Sprintf("[GET /quote][%d] quoteGetNotFound  %+v", 404, o.Payload)
}

func (o *QuoteGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
