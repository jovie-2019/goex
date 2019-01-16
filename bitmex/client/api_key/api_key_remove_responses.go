// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// APIKeyRemoveReader is a Reader for the APIKeyRemove structure.
type APIKeyRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIKeyRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAPIKeyRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAPIKeyRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAPIKeyRemoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAPIKeyRemoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIKeyRemoveOK creates a APIKeyRemoveOK with default headers values
func NewAPIKeyRemoveOK() *APIKeyRemoveOK {
	return &APIKeyRemoveOK{}
}

/*APIKeyRemoveOK handles this case with default header values.

Request was successful
*/
type APIKeyRemoveOK struct {
	Payload *APIKeyRemoveOKBody
}

func (o *APIKeyRemoveOK) Error() string {
	return fmt.Sprintf("[DELETE /apiKey][%d] apiKeyRemoveOK  %+v", 200, *o.Payload)
}

func (o *APIKeyRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(APIKeyRemoveOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyRemoveBadRequest creates a APIKeyRemoveBadRequest with default headers values
func NewAPIKeyRemoveBadRequest() *APIKeyRemoveBadRequest {
	return &APIKeyRemoveBadRequest{}
}

/*APIKeyRemoveBadRequest handles this case with default header values.

Parameter Error
*/
type APIKeyRemoveBadRequest struct {
	Payload *models.Error
}

func (o *APIKeyRemoveBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /apiKey][%d] apiKeyRemoveBadRequest  %+v", 400, *o.Payload)
}

func (o *APIKeyRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyRemoveUnauthorized creates a APIKeyRemoveUnauthorized with default headers values
func NewAPIKeyRemoveUnauthorized() *APIKeyRemoveUnauthorized {
	return &APIKeyRemoveUnauthorized{}
}

/*APIKeyRemoveUnauthorized handles this case with default header values.

Unauthorized
*/
type APIKeyRemoveUnauthorized struct {
	Payload *models.Error
}

func (o *APIKeyRemoveUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apiKey][%d] apiKeyRemoveUnauthorized  %+v", 401, o.Payload)
}

func (o *APIKeyRemoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIKeyRemoveNotFound creates a APIKeyRemoveNotFound with default headers values
func NewAPIKeyRemoveNotFound() *APIKeyRemoveNotFound {
	return &APIKeyRemoveNotFound{}
}

/*APIKeyRemoveNotFound handles this case with default header values.

Not Found
*/
type APIKeyRemoveNotFound struct {
	Payload *models.Error
}

func (o *APIKeyRemoveNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apiKey][%d] apiKeyRemoveNotFound  %+v", 404, o.Payload)
}

func (o *APIKeyRemoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*APIKeyRemoveOKBody API key remove o k body
swagger:model APIKeyRemoveOKBody
*/
type APIKeyRemoveOKBody struct {

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this API key remove o k body
func (o *APIKeyRemoveOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *APIKeyRemoveOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *APIKeyRemoveOKBody) UnmarshalBinary(b []byte) error {
	var res APIKeyRemoveOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
