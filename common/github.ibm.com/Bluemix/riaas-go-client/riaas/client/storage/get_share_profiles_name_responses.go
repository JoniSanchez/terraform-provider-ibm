// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetShareProfilesNameReader is a Reader for the GetShareProfilesName structure.
type GetShareProfilesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShareProfilesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetShareProfilesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetShareProfilesNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetShareProfilesNameOK creates a GetShareProfilesNameOK with default headers values
func NewGetShareProfilesNameOK() *GetShareProfilesNameOK {
	return &GetShareProfilesNameOK{}
}

/*GetShareProfilesNameOK handles this case with default header values.

dummy
*/
type GetShareProfilesNameOK struct {
	Payload *models.ShareProfile
}

func (o *GetShareProfilesNameOK) Error() string {
	return fmt.Sprintf("[GET /share/profiles/{name}][%d] getShareProfilesNameOK  %+v", 200, o.Payload)
}

func (o *GetShareProfilesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShareProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShareProfilesNameNotFound creates a GetShareProfilesNameNotFound with default headers values
func NewGetShareProfilesNameNotFound() *GetShareProfilesNameNotFound {
	return &GetShareProfilesNameNotFound{}
}

/*GetShareProfilesNameNotFound handles this case with default header values.

error
*/
type GetShareProfilesNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetShareProfilesNameNotFound) Error() string {
	return fmt.Sprintf("[GET /share/profiles/{name}][%d] getShareProfilesNameNotFound  %+v", 404, o.Payload)
}

func (o *GetShareProfilesNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
