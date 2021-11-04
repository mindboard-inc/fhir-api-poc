// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Specimen resource
type Specimen struct {
	Domain
	Identifier          []d.Identifier       `json:"identifier,omitempty"`
	AccessionIdentifier *d.Identifier        `json:"accessionIdentifier,omitempty"`
	Status              *d.Code              `json:"status,omitempty"`
	Type                *d.CodeableConcept   `json:"type,omitempty"`
	Subject             *d.Reference         `json:"subject,omitempty"`
	ReceivedTime        *d.DateTime          `json:"receivedTime,omitempty"`
	Parent              []d.Reference        `json:"parent,omitempty"`
	Request             []d.Reference        `json:"request,omitempty"`
	Collection          *SpecimenCollection  `json:"collection,omitempty"`
	Processing          []SpecimenProcessing `json:"processing,omitempty"`
	Container           []SpecimenContainer  `json:"container,omitempty"`
	Note                []d.Annotation       `json:"note,omitempty"`
}

// SpecimenCollection subResource
type SpecimenCollection struct {
	Collector         *d.Reference       `json:"collector,omitempty"`
	CollectedDateTime *d.DateTime        `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *d.Period          `json:"collectedPeriod,omitempty"`
	Quantity          *d.SimpleQuantity  `json:"quantity,omitempty"`
	Method            *d.CodeableConcept `json:"method,omitempty"`
	BodySite          *d.CodeableConcept `json:"bodySite,omitempty"`
}

// SpecimenProcessing subResource
type SpecimenProcessing struct {
	Description  *d.String          `json:"description,omitempty"`
	Procedure    *d.CodeableConcept `json:"procedure,omitempty"`
	Additive     []d.Reference      `json:"additive,omitempty"`
	TimeDateTime *d.DateTime        `json:"timeDateTime,omitempty"`
	TimePeriod   *d.Period          `json:"timePeriod,omitempty"`
}

// SpecimenContainer subResource
type SpecimenContainer struct {
	Identifier              []d.Identifier     `json:"identifier,omitempty"`
	Description             *d.String          `json:"description,omitempty"`
	Type                    *d.CodeableConcept `json:"type,omitempty"`
	Capacity                *d.SimpleQuantity  `json:"capacity,omitempty"`
	SpecimenQuantity        *d.SimpleQuantity  `json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *d.CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *d.Reference       `json:"additiveReference,omitempty"`
}
