// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)


// OperationOutcome resource
type OperationOutcome struct {
	Domain
	Issue []OperationOutcomeIssue `json:"issue,omitempty"`
}

// OperationOutcomeIssue sub-resource
type OperationOutcomeIssue struct {
	Severity    *d.Code            `json:"severity,omitempty"`
	Code        *d.Code            `json:"code,omitempty"`
	Details     *d.CodeableConcept `json:"details,omitempty"`
	Diagnostics *d.String          `json:"diagnostics,omitempty"`
	Location    []d.String         `json:"location,omitempty"`
	Expression  []d.String         `json:"expression,omitempty"`
}
