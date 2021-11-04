// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// ServiceDefinition resource
type ServiceDefinition struct {
	Domain
	URL                 *d.URI                `json:"url,omitempty"`
	Identifier          []d.Identifier        `json:"identifier,omitempty"`
	Version             *d.String             `json:"version,omitempty"`
	Name                *d.String             `json:"name,omitempty"`
	Title               *d.String             `json:"title,omitempty"`
	Status              *d.Code               `json:"status,omitempty"`
	Experimental        *d.Boolean            `json:"experimental,omitempty"`
	Date                *d.DateTime           `json:"date,omitempty"`
	Publisher           *d.String             `json:"publisher,omitempty"`
	Description         *d.Markdown           `json:"description,omitempty"`
	Purpose             *d.Markdown           `json:"purpose,omitempty"`
	Usage               *d.String             `json:"usage,omitempty"`
	ApprovalDate        *d.Date               `json:"approvalDate,omitempty"`
	LastReviewDate      *d.Date               `json:"lastReviewDate,omitempty"`
	EffectivePeriod     *d.Period             `json:"effectivePeriod,omitempty"`
	UseContext          []d.UsageContext      `json:"useContext,omitempty"`
	Jurisdiction        []d.CodeableConcept   `json:"jurisdiction,omitempty"`
	Topic               []d.CodeableConcept   `json:"topic,omitempty"`
	Contributor         []d.Contributor       `json:"contributor,omitempty"`
	Contact             []d.ContactDetail     `json:"contact,omitempty"`
	Copyright           *d.Markdown           `json:"copyright,omitempty"`
	RelatedArtifact     []d.RelatedArtifact   `json:"relatedArtifact,omitempty"`
	Trigger             []d.TriggerDefinition `json:"trigger,omitempty"`
	DataRequirement     []d.DataRequirement   `json:"dataRequirement,omitempty"`
	OperationDefinition *d.Reference          `json:"operationDefinition,omitempty"`
}
