// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nntaoli-project/GoEx/bitmex/models"
)

// UserSavePreferencesReader is a Reader for the UserSavePreferences structure.
type UserSavePreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSavePreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserSavePreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserSavePreferencesOK creates a UserSavePreferencesOK with default headers values
func NewUserSavePreferencesOK() *UserSavePreferencesOK {
	return &UserSavePreferencesOK{}
}

/*UserSavePreferencesOK handles this case with default header values.

Request was successful
*/
type UserSavePreferencesOK struct {
	Payload *models.User
}

func (o *UserSavePreferencesOK) Error() string {
	return fmt.Sprintf("[POST /user/preferences][%d] userSavePreferencesOK  %+v", 200, *o.Payload)
}

func (o *UserSavePreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
