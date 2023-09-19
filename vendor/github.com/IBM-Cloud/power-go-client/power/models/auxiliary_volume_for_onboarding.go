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

// AuxiliaryVolumeForOnboarding auxiliary volume for onboarding
//
// swagger:model AuxiliaryVolumeForOnboarding
type AuxiliaryVolumeForOnboarding struct {

	// auxiliary volume name at storage host level
	// Required: true
	AuxVolumeName *string `json:"auxVolumeName"`

	// display name of auxVolumeName once onboarded,auxVolumeName will be set to display name if not provided.
	Name string `json:"name,omitempty"`
}

// Validate validates this auxiliary volume for onboarding
func (m *AuxiliaryVolumeForOnboarding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuxVolumeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuxiliaryVolumeForOnboarding) validateAuxVolumeName(formats strfmt.Registry) error {

	if err := validate.Required("auxVolumeName", "body", m.AuxVolumeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this auxiliary volume for onboarding based on context it is used
func (m *AuxiliaryVolumeForOnboarding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuxiliaryVolumeForOnboarding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuxiliaryVolumeForOnboarding) UnmarshalBinary(b []byte) error {
	var res AuxiliaryVolumeForOnboarding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}