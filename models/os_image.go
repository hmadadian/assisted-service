// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OsImage os image
//
// swagger:model os-image
type OsImage struct {

	// The CPU architecture of the image (x86_64/arm64/etc).
	// Required: true
	CPUArchitecture *string `json:"cpu_architecture" gorm:"default:'x86_64'"`

	// Version of the operating system image
	// Example: 4.12
	// Required: true
	OpenshiftVersion *string `json:"openshift_version"`

	// The base OS image used for the discovery iso.
	// Required: true
	URL *string `json:"url"`

	// Build ID of the OS image.
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this os image
func (m *OsImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenshiftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsImage) validateCPUArchitecture(formats strfmt.Registry) error {

	if err := validate.Required("cpu_architecture", "body", m.CPUArchitecture); err != nil {
		return err
	}

	return nil
}

func (m *OsImage) validateOpenshiftVersion(formats strfmt.Registry) error {

	if err := validate.Required("openshift_version", "body", m.OpenshiftVersion); err != nil {
		return err
	}

	return nil
}

func (m *OsImage) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *OsImage) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this os image based on context it is used
func (m *OsImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OsImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsImage) UnmarshalBinary(b []byte) error {
	var res OsImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
