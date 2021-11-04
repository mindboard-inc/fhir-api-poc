// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Location resource
type Location struct {
	Domain
	Identifier           []d.Identifier     `json:"identifier,omitempty"`
	Status               *d.Code            `json:"status,omitempty"`
	OperationalStatus    *d.Coding          `json:"operationalStatus,omitempty"`
	Name                 *d.String          `json:"name,omitempty"`
	Alias                []d.String         `json:"alias,omitempty"`
	Description          *d.String          `json:"description,omitempty"`
	Mode                 *d.Code            `json:"mode,omitempty"`
	Type                 *d.CodeableConcept `json:"type,omitempty"`
	Telecom              []d.ContactPoint   `json:"telecom,omitempty"`
	Address              *d.Address         `json:"address,omitempty"`
	PhysicalType         *d.CodeableConcept `json:"physicalType,omitempty"`
	Position             *LocationPosition  `json:"position,omitempty"`
	ManagingOrganization *d.Reference       `json:"managingOrganization,omitempty"`
	PartOf               *d.Reference       `json:"partOf,omitempty"`
	Endpoint             []d.Reference      `json:"endpoint,omitempty"`
}

// LocationPosition sub-resource
type LocationPosition struct {
	Longitude *d.Decimal `json:"longitude,omitempty"`
	Latitude  *d.Decimal `json:"latitude,omitempty"`
	Altitude  *d.Decimal `json:"altitude,omitempty"`
}
