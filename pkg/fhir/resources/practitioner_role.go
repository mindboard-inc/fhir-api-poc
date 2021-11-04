// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)


// PractitionerRole resource
type PractitionerRole struct {
	Domain
	Identifier             []d.Identifier                  `json:"identifier,omitempty"`
	Active                 *d.Boolean                      `json:"active,omitempty"`
	Period                 *d.Period                       `json:"period,omitempty"`
	Practitioner           *d.Reference                    `json:"practitioner,omitempty"`
	Organization           *d.Reference                    `json:"organization,omitempty"`
	Code                   []d.CodeableConcept             `json:"code,omitempty"`
	Specialty              []d.CodeableConcept             `json:"specialty,omitempty"`
	Location               []d.Reference                   `json:"location,omitempty"`
	HealthcareService      []d.Reference                   `json:"healthcareService,omitempty"`
	Telecom                []d.ContactPoint                `json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *d.String                       `json:"availabilityExceptions,omitempty"`
	Endpoint               []d.Reference                   `json:"endpoint,omitempty"`
}

// PractitionerRoleAvailableTime sub-resource
type PractitionerRoleAvailableTime struct {
	DaysOfWeek         []d.Code   `json:"daysOfWeek,omitempty"`
	AllDay             *d.Boolean `json:"allDay,omitempty"`
	AvailableStartTime *d.Time    `json:"availableStartTime,omitempty"`
	AvailableEndTime   *d.Time    `json:"availableEndTime,omitempty"`
}

// PractitionerRoleNotAvailable sub-resource
type PractitionerRoleNotAvailable struct {
	Description *d.String `json:"description,omitempty"`
	During      *d.Period `json:"during,omitempty"`
}
