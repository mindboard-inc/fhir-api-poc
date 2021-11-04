// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Group resource
type Group struct {
	Domain
	Identifier     []d.Identifier        `json:"identifier,omitempty"`
	Active         *d.Boolean            `json:"active,omitempty"`
	Type           *d.Code               `json:"type,omitempty"`
	Actual         *d.Boolean            `json:"actual,omitempty"`
	Code           *d.CodeableConcept    `json:"code,omitempty"`
	Name           *d.String             `json:"name,omitempty"`
	Quantity       *d.UnsignedInt        `json:"quantity,omitempty"`
	Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`
	Member         []GroupMember         `json:"member,omitempty"`
}

// GroupCharacteristic sub-resource
type GroupCharacteristic struct {
	Code                 *d.CodeableConcept `json:"code,omitempty"`
	ValueCodeableConcept *d.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *d.Boolean         `json:"valueBoolean,omitempty"`
	ValueQuantity        *d.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *d.Range           `json:"valueRange,omitempty"`
	Exclude              *d.Boolean         `json:"exclude,omitempty"`
	Period               *d.Period          `json:"period,omitempty"`
}

// GroupMember sub-resource
type GroupMember struct {
	Entity   *d.Reference `json:"entity,omitempty"`
	Period   *d.Period    `json:"period,omitempty"`
	Inactive *d.Boolean   `json:"inactive,omitempty"`
}
