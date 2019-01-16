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

// InstrumentGetActiveAndIndicesReader is a Reader for the InstrumentGetActiveAndIndices structure.
type InstrumentGetActiveAndIndicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstrumentGetActiveAndIndicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewInstrumentGetActiveAndIndicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewInstrumentGetActiveAndIndicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewInstrumentGetActiveAndIndicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewInstrumentGetActiveAndIndicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInstrumentGetActiveAndIndicesOK creates a InstrumentGetActiveAndIndicesOK with default headers values
func NewInstrumentGetActiveAndIndicesOK() *InstrumentGetActiveAndIndicesOK {
	return &InstrumentGetActiveAndIndicesOK{}
}

/*InstrumentGetActiveAndIndicesOK handles this case with default header values.

Request was successful
*/
type InstrumentGetActiveAndIndicesOK struct {
	Payload []*models.Instrument
}

func (o *InstrumentGetActiveAndIndicesOK) Error() string {
	return fmt.Sprintf("[GET /instrument/activeAndIndices][%d] instrumentGetActiveAndIndicesOK  %+v", 200, o.Payload)
}

func (o *InstrumentGetActiveAndIndicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetActiveAndIndicesBadRequest creates a InstrumentGetActiveAndIndicesBadRequest with default headers values
func NewInstrumentGetActiveAndIndicesBadRequest() *InstrumentGetActiveAndIndicesBadRequest {
	return &InstrumentGetActiveAndIndicesBadRequest{}
}

/*InstrumentGetActiveAndIndicesBadRequest handles this case with default header values.

Parameter Error
*/
type InstrumentGetActiveAndIndicesBadRequest struct {
	Payload *models.Error
}

func (o *InstrumentGetActiveAndIndicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /instrument/activeAndIndices][%d] instrumentGetActiveAndIndicesBadRequest  %+v", 400, *o.Payload)
}

func (o *InstrumentGetActiveAndIndicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetActiveAndIndicesUnauthorized creates a InstrumentGetActiveAndIndicesUnauthorized with default headers values
func NewInstrumentGetActiveAndIndicesUnauthorized() *InstrumentGetActiveAndIndicesUnauthorized {
	return &InstrumentGetActiveAndIndicesUnauthorized{}
}

/*InstrumentGetActiveAndIndicesUnauthorized handles this case with default header values.

Unauthorized
*/
type InstrumentGetActiveAndIndicesUnauthorized struct {
	Payload *models.Error
}

func (o *InstrumentGetActiveAndIndicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /instrument/activeAndIndices][%d] instrumentGetActiveAndIndicesUnauthorized  %+v", 401, o.Payload)
}

func (o *InstrumentGetActiveAndIndicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstrumentGetActiveAndIndicesNotFound creates a InstrumentGetActiveAndIndicesNotFound with default headers values
func NewInstrumentGetActiveAndIndicesNotFound() *InstrumentGetActiveAndIndicesNotFound {
	return &InstrumentGetActiveAndIndicesNotFound{}
}

/*InstrumentGetActiveAndIndicesNotFound handles this case with default header values.

Not Found
*/
type InstrumentGetActiveAndIndicesNotFound struct {
	Payload *models.Error
}

func (o *InstrumentGetActiveAndIndicesNotFound) Error() string {
	return fmt.Sprintf("[GET /instrument/activeAndIndices][%d] instrumentGetActiveAndIndicesNotFound  %+v", 404, o.Payload)
}

func (o *InstrumentGetActiveAndIndicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
