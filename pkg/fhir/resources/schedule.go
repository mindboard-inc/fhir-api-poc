// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Schedule resource
type Schedule struct {
	Domain
	Identifier      []d.Identifier      `json:"identifier,omitempty"`
	Active          *d.Boolean          `json:"active,omitempty"`
	ServiceCategory *d.CodeableConcept  `json:"serviceCategory,omitempty"`
	ServiceType     []d.CodeableConcept `json:"serviceType,omitempty"`
	Specialty       []d.CodeableConcept `json:"specialty,omitempty"`
	Actor           []d.Reference       `json:"actor,omitempty"`
	PlanningHorizon *d.Period           `json:"planningHorizon,omitempty"`
	Comment         *d.String           `json:"comment,omitempty"`
}
