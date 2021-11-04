// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// RiskAssessment resource
type RiskAssessment struct {
	Domain
	Identifier            *d.Identifier              `json:"identifier,omitempty"`
	BasedOn               *d.Reference               `json:"basedOn,omitempty"`
	Parent                *d.Reference               `json:"parent,omitempty"`
	Status                *d.Code                    `json:"status,omitempty"`
	Method                *d.CodeableConcept         `json:"method,omitempty"`
	Code                  *d.CodeableConcept         `json:"code,omitempty"`
	Subject               *d.Reference               `json:"subject,omitempty"`
	Context               *d.Reference               `json:"context,omitempty"`
	OccurrenceDateTime    *d.DateTime                `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *d.Period                  `json:"occurrencePeriod,omitempty"`
	Condition             *d.Reference               `json:"condition,omitempty"`
	Performer             *d.Reference               `json:"performer,omitempty"`
	ReasonCodeableConcept *d.CodeableConcept         `json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *d.Reference               `json:"reasonReference,omitempty"`
	Basis                 []d.Reference              `json:"basis,omitempty"`
	Prediction            []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation            *d.String                  `json:"mitigation,omitempty"`
	Comment               *d.String                  `json:"comment,omitempty"`
}

// RiskAssessmentPrediction subResource
type RiskAssessmentPrediction struct {
	Outcome            *d.CodeableConcept `json:"outcome,omitempty"`
	ProbabilityDecimal *d.Decimal         `json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *d.Range           `json:"probabilityRange,omitempty"`
	QualitativeRisk    *d.CodeableConcept `json:"qualitativeRisk,omitempty"`
	RelativeRisk       *d.Decimal         `json:"relativeRisk,omitempty"`
	WhenPeriod         *d.Period          `json:"whenPeriod,omitempty"`
	WhenRange          *d.Range           `json:"whenRange,omitempty"`
	Rationale          *d.String          `json:"rationale,omitempty"`
}
