// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// DeregisterInfraEnvReader is a Reader for the DeregisterInfraEnv structure.
type DeregisterInfraEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeregisterInfraEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeregisterInfraEnvNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeregisterInfraEnvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeregisterInfraEnvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeregisterInfraEnvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeregisterInfraEnvMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeregisterInfraEnvConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeregisterInfraEnvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeregisterInfraEnvNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeregisterInfraEnvNoContent creates a DeregisterInfraEnvNoContent with default headers values
func NewDeregisterInfraEnvNoContent() *DeregisterInfraEnvNoContent {
	return &DeregisterInfraEnvNoContent{}
}

/*DeregisterInfraEnvNoContent handles this case with default header values.

Success.
*/
type DeregisterInfraEnvNoContent struct {
}

func (o *DeregisterInfraEnvNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvNoContent ", 204)
}

func (o *DeregisterInfraEnvNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeregisterInfraEnvUnauthorized creates a DeregisterInfraEnvUnauthorized with default headers values
func NewDeregisterInfraEnvUnauthorized() *DeregisterInfraEnvUnauthorized {
	return &DeregisterInfraEnvUnauthorized{}
}

/*DeregisterInfraEnvUnauthorized handles this case with default header values.

Unauthorized.
*/
type DeregisterInfraEnvUnauthorized struct {
	Payload *models.InfraError
}

func (o *DeregisterInfraEnvUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvUnauthorized  %+v", 401, o.Payload)
}

func (o *DeregisterInfraEnvUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DeregisterInfraEnvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvForbidden creates a DeregisterInfraEnvForbidden with default headers values
func NewDeregisterInfraEnvForbidden() *DeregisterInfraEnvForbidden {
	return &DeregisterInfraEnvForbidden{}
}

/*DeregisterInfraEnvForbidden handles this case with default header values.

Forbidden.
*/
type DeregisterInfraEnvForbidden struct {
	Payload *models.InfraError
}

func (o *DeregisterInfraEnvForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvForbidden  %+v", 403, o.Payload)
}

func (o *DeregisterInfraEnvForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *DeregisterInfraEnvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvNotFound creates a DeregisterInfraEnvNotFound with default headers values
func NewDeregisterInfraEnvNotFound() *DeregisterInfraEnvNotFound {
	return &DeregisterInfraEnvNotFound{}
}

/*DeregisterInfraEnvNotFound handles this case with default header values.

Error.
*/
type DeregisterInfraEnvNotFound struct {
	Payload *models.Error
}

func (o *DeregisterInfraEnvNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvNotFound  %+v", 404, o.Payload)
}

func (o *DeregisterInfraEnvNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterInfraEnvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvMethodNotAllowed creates a DeregisterInfraEnvMethodNotAllowed with default headers values
func NewDeregisterInfraEnvMethodNotAllowed() *DeregisterInfraEnvMethodNotAllowed {
	return &DeregisterInfraEnvMethodNotAllowed{}
}

/*DeregisterInfraEnvMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type DeregisterInfraEnvMethodNotAllowed struct {
	Payload *models.Error
}

func (o *DeregisterInfraEnvMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeregisterInfraEnvMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterInfraEnvMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvConflict creates a DeregisterInfraEnvConflict with default headers values
func NewDeregisterInfraEnvConflict() *DeregisterInfraEnvConflict {
	return &DeregisterInfraEnvConflict{}
}

/*DeregisterInfraEnvConflict handles this case with default header values.

Error.
*/
type DeregisterInfraEnvConflict struct {
	Payload *models.Error
}

func (o *DeregisterInfraEnvConflict) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvConflict  %+v", 409, o.Payload)
}

func (o *DeregisterInfraEnvConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterInfraEnvConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvInternalServerError creates a DeregisterInfraEnvInternalServerError with default headers values
func NewDeregisterInfraEnvInternalServerError() *DeregisterInfraEnvInternalServerError {
	return &DeregisterInfraEnvInternalServerError{}
}

/*DeregisterInfraEnvInternalServerError handles this case with default header values.

Error.
*/
type DeregisterInfraEnvInternalServerError struct {
	Payload *models.Error
}

func (o *DeregisterInfraEnvInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvInternalServerError  %+v", 500, o.Payload)
}

func (o *DeregisterInfraEnvInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterInfraEnvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeregisterInfraEnvNotImplemented creates a DeregisterInfraEnvNotImplemented with default headers values
func NewDeregisterInfraEnvNotImplemented() *DeregisterInfraEnvNotImplemented {
	return &DeregisterInfraEnvNotImplemented{}
}

/*DeregisterInfraEnvNotImplemented handles this case with default header values.

Not implemented.
*/
type DeregisterInfraEnvNotImplemented struct {
	Payload *models.Error
}

func (o *DeregisterInfraEnvNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /v2/infra-envs/{infra_env_id}][%d] deregisterInfraEnvNotImplemented  %+v", 501, o.Payload)
}

func (o *DeregisterInfraEnvNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeregisterInfraEnvNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
