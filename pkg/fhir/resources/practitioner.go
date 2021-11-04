// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Practitioner resource
type Practitioner struct {
	Domain
	Identifier    []d.Identifier              `json:"identifier,omitempty"`
	Active        *d.Boolean                  `json:"active,omitempty"`
	Name          []d.HumanName               `json:"name,omitempty"`
	Telecom       []d.ContactPoint            `json:"telecom,omitempty"`
	Address       []d.Address                 `json:"address,omitempty"`
	Gender        *d.Code                     `json:"gender,omitempty"`
	BirthDate     *d.Date                     `json:"birthDate,omitempty"`
	Photo         []d.Attachment              `json:"photo,omitempty"`
	Qualification []PractitionerQualification `json:"qualification,omitempty"`
	Communication []d.CodeableConcept         `json:"communication,omitempty"`
}


// PractitionerQualification sub-resource
type PractitionerQualification struct {
	Identifier []d.Identifier     `json:"identifier,omitempty"`
	Code       *d.CodeableConcept `json:"code,omitempty"`
	Period     *d.Period          `json:"period,omitempty"`
	Issuer     *d.Reference       `json:"issuer,omitempty"`
}
