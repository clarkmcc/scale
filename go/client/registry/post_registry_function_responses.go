// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/go/client/models"
)

// PostRegistryFunctionReader is a Reader for the PostRegistryFunction structure.
type PostRegistryFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRegistryFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRegistryFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRegistryFunctionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRegistryFunctionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRegistryFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRegistryFunctionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRegistryFunctionOK creates a PostRegistryFunctionOK with default headers values
func NewPostRegistryFunctionOK() *PostRegistryFunctionOK {
	return &PostRegistryFunctionOK{}
}

/*
PostRegistryFunctionOK describes a response with status code 200, with default header values.

OK
*/
type PostRegistryFunctionOK struct {
	Payload *models.ModelsCreateFunctionResponse
}

// IsSuccess returns true when this post registry function o k response has a 2xx status code
func (o *PostRegistryFunctionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post registry function o k response has a 3xx status code
func (o *PostRegistryFunctionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry function o k response has a 4xx status code
func (o *PostRegistryFunctionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post registry function o k response has a 5xx status code
func (o *PostRegistryFunctionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry function o k response a status code equal to that given
func (o *PostRegistryFunctionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post registry function o k response
func (o *PostRegistryFunctionOK) Code() int {
	return 200
}

func (o *PostRegistryFunctionOK) Error() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionOK  %+v", 200, o.Payload)
}

func (o *PostRegistryFunctionOK) String() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionOK  %+v", 200, o.Payload)
}

func (o *PostRegistryFunctionOK) GetPayload() *models.ModelsCreateFunctionResponse {
	return o.Payload
}

func (o *PostRegistryFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsCreateFunctionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistryFunctionBadRequest creates a PostRegistryFunctionBadRequest with default headers values
func NewPostRegistryFunctionBadRequest() *PostRegistryFunctionBadRequest {
	return &PostRegistryFunctionBadRequest{}
}

/*
PostRegistryFunctionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostRegistryFunctionBadRequest struct {
	Payload string
}

// IsSuccess returns true when this post registry function bad request response has a 2xx status code
func (o *PostRegistryFunctionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry function bad request response has a 3xx status code
func (o *PostRegistryFunctionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry function bad request response has a 4xx status code
func (o *PostRegistryFunctionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry function bad request response has a 5xx status code
func (o *PostRegistryFunctionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry function bad request response a status code equal to that given
func (o *PostRegistryFunctionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post registry function bad request response
func (o *PostRegistryFunctionBadRequest) Code() int {
	return 400
}

func (o *PostRegistryFunctionBadRequest) Error() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *PostRegistryFunctionBadRequest) String() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionBadRequest  %+v", 400, o.Payload)
}

func (o *PostRegistryFunctionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PostRegistryFunctionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistryFunctionUnauthorized creates a PostRegistryFunctionUnauthorized with default headers values
func NewPostRegistryFunctionUnauthorized() *PostRegistryFunctionUnauthorized {
	return &PostRegistryFunctionUnauthorized{}
}

/*
PostRegistryFunctionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostRegistryFunctionUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this post registry function unauthorized response has a 2xx status code
func (o *PostRegistryFunctionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry function unauthorized response has a 3xx status code
func (o *PostRegistryFunctionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry function unauthorized response has a 4xx status code
func (o *PostRegistryFunctionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry function unauthorized response has a 5xx status code
func (o *PostRegistryFunctionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry function unauthorized response a status code equal to that given
func (o *PostRegistryFunctionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post registry function unauthorized response
func (o *PostRegistryFunctionUnauthorized) Code() int {
	return 401
}

func (o *PostRegistryFunctionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRegistryFunctionUnauthorized) String() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRegistryFunctionUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *PostRegistryFunctionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistryFunctionNotFound creates a PostRegistryFunctionNotFound with default headers values
func NewPostRegistryFunctionNotFound() *PostRegistryFunctionNotFound {
	return &PostRegistryFunctionNotFound{}
}

/*
PostRegistryFunctionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostRegistryFunctionNotFound struct {
	Payload string
}

// IsSuccess returns true when this post registry function not found response has a 2xx status code
func (o *PostRegistryFunctionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry function not found response has a 3xx status code
func (o *PostRegistryFunctionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry function not found response has a 4xx status code
func (o *PostRegistryFunctionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post registry function not found response has a 5xx status code
func (o *PostRegistryFunctionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post registry function not found response a status code equal to that given
func (o *PostRegistryFunctionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post registry function not found response
func (o *PostRegistryFunctionNotFound) Code() int {
	return 404
}

func (o *PostRegistryFunctionNotFound) Error() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionNotFound  %+v", 404, o.Payload)
}

func (o *PostRegistryFunctionNotFound) String() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionNotFound  %+v", 404, o.Payload)
}

func (o *PostRegistryFunctionNotFound) GetPayload() string {
	return o.Payload
}

func (o *PostRegistryFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRegistryFunctionInternalServerError creates a PostRegistryFunctionInternalServerError with default headers values
func NewPostRegistryFunctionInternalServerError() *PostRegistryFunctionInternalServerError {
	return &PostRegistryFunctionInternalServerError{}
}

/*
PostRegistryFunctionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostRegistryFunctionInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this post registry function internal server error response has a 2xx status code
func (o *PostRegistryFunctionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post registry function internal server error response has a 3xx status code
func (o *PostRegistryFunctionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post registry function internal server error response has a 4xx status code
func (o *PostRegistryFunctionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post registry function internal server error response has a 5xx status code
func (o *PostRegistryFunctionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post registry function internal server error response a status code equal to that given
func (o *PostRegistryFunctionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post registry function internal server error response
func (o *PostRegistryFunctionInternalServerError) Code() int {
	return 500
}

func (o *PostRegistryFunctionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRegistryFunctionInternalServerError) String() string {
	return fmt.Sprintf("[POST /registry/function][%d] postRegistryFunctionInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRegistryFunctionInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PostRegistryFunctionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
