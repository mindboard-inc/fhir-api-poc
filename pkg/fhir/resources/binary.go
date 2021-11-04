// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Binary resource
type Binary struct {
	Base
	ContentType     *d.Code         `json:"contentType,omitempty"`
	SecurityContext *d.Reference    `json:"securityContext,omitempty"`
	Content         *d.Base64Binary `json:"content,omitempty"`
}

