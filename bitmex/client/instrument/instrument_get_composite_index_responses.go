// Code generated by go-swagger; DO NOT EDIT.

package instrument

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// InstrumentGetCompositeIndexReader is a Reader for the InstrumentGetCompositeIndex structure.
type InstrumentGetCompositeIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstrumentGetCompositeIndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInstrumentGetCompositeIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInstrumentGetCompositeIndexBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInstrumentGetCompositeIndexUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInstrumentGetCompositeIndexNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstrumentGetCompositeIndexOK creates a InstrumentGetCompositeIndexOK with default headers values
func NewInstrumentGetCompositeIndexOK() *InstrumentGetCompositeIndexOK {
	return &InstrumentGetCompositeIndexOK{}
}

/*InstrumentGetCompositeIndexOK handles this case with default header values.

Request was successful
*/
type InstrumentGetCompositeIndexOK struct {
	Payload []*models.IndexComposite
}

func (o *InstrumentGetCompositeIndexOK) Error() string {
	return fmt.Sprintf("[GET /instrument/compositeIndex][%d] instrumentGetCompositeIndexOK  %+v", 200, o.Payload)
}

func (o *InstrumentGetCompositeIndexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetCompositeIndexBadRequest creates a InstrumentGetCompositeIndexBadRequest with default headers values
func NewInstrumentGetCompositeIndexBadRequest() *InstrumentGetCompositeIndexBadRequest {
	return &InstrumentGetCompositeIndexBadRequest{}
}

/*InstrumentGetCompositeIndexBadRequest handles this case with default header values.

Parameter Error
*/
type InstrumentGetCompositeIndexBadRequest struct {
	Payload *models.Error
}

func (o *InstrumentGetCompositeIndexBadRequest) Error() string {
	return fmt.Sprintf("[GET /instrument/compositeIndex][%d] instrumentGetCompositeIndexBadRequest  %+v", 400, *o.Payload)
}

func (o *InstrumentGetCompositeIndexBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetCompositeIndexUnauthorized creates a InstrumentGetCompositeIndexUnauthorized with default headers values
func NewInstrumentGetCompositeIndexUnauthorized() *InstrumentGetCompositeIndexUnauthorized {
	return &InstrumentGetCompositeIndexUnauthorized{}
}

/*InstrumentGetCompositeIndexUnauthorized handles this case with default header values.

Unauthorized
*/
type InstrumentGetCompositeIndexUnauthorized struct {
	Payload *models.Error
}

func (o *InstrumentGetCompositeIndexUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instrument/compositeIndex][%d] instrumentGetCompositeIndexUnauthorized  %+v", 401, o.Payload)
}

func (o *InstrumentGetCompositeIndexUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetCompositeIndexNotFound creates a InstrumentGetCompositeIndexNotFound with default headers values
func NewInstrumentGetCompositeIndexNotFound() *InstrumentGetCompositeIndexNotFound {
	return &InstrumentGetCompositeIndexNotFound{}
}

/*InstrumentGetCompositeIndexNotFound handles this case with default header values.

Not Found
*/
type InstrumentGetCompositeIndexNotFound struct {
	Payload *models.Error
}

func (o *InstrumentGetCompositeIndexNotFound) Error() string {
	return fmt.Sprintf("[GET /instrument/compositeIndex][%d] instrumentGetCompositeIndexNotFound  %+v", 404, o.Payload)
}

func (o *InstrumentGetCompositeIndexNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
