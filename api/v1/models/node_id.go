// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeID Node ID with associated node IP addresses
//
// swagger:model NodeID
type NodeID struct {

	// ID allocated by the agent for the node
	// Required: true
	ID *int64 `json:"id"`

	// IP addresses of the node associated with the ID in the agent
	// Required: true
	Ips []string `json:"ips"`
}

// Validate validates this node ID
func (m *NodeID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeID) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NodeID) validateIps(formats strfmt.Registry) error {

	if err := validate.Required("ips", "body", m.Ips); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeID) UnmarshalBinary(b []byte) error {
	var res NodeID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
