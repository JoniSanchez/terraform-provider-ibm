// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VPNGatewayConnectionLocalCIDRs v p n gateway connection local c ID rs
// swagger:model VPNGatewayConnectionLocalCIDRs
type VPNGatewayConnectionLocalCIDRs struct {

	// A collection of local CIDRs for this resource
	LocalCidrs []string `json:"local_cidrs"`
}

// Validate validates this v p n gateway connection local c ID rs
func (m *VPNGatewayConnectionLocalCIDRs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VPNGatewayConnectionLocalCIDRs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNGatewayConnectionLocalCIDRs) UnmarshalBinary(b []byte) error {
	var res VPNGatewayConnectionLocalCIDRs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
