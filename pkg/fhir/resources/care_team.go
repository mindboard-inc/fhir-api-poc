// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)


// CareTeam resource
type CareTeam struct {
	Domain
	Identifier           []d.Identifier        `json:"identifier,omitempty"`
	Status               *d.Code               `json:"status,omitempty"`
	Category             []d.CodeableConcept   `json:"category,omitempty"`
	Name                 *d.String             `json:"name,omitempty"`
	Subject              *d.Reference          `json:"subject,omitempty"`
	Context              *d.Reference          `json:"context,omitempty"`
	Period               *d.Period             `json:"period,omitempty"`
	Participant          []CareTeamParticipant `json:"participant,omitempty"`
	ReasonCode           []d.CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference      []d.Reference         `json:"reasonReference,omitempty"`
	ManagingOrganization []d.Reference         `json:"managingOrganization,omitempty"`
	Note                 []d.Annotation        `json:"note,omitempty"`
}

// CareTeamParticipant sub-resource
type CareTeamParticipant struct {
	Role       *d.CodeableConcept `json:"role,omitempty"`
	Member     *d.Reference       `json:"member,omitempty"`
	OnBehalfOf *d.Reference       `json:"onBehalfOf,omitempty"`
	Period     *d.Period          `json:"period,omitempty"`
}
