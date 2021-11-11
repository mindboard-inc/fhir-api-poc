// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources


import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	m "mindboard.com/fhirpoc/pkg/model"
)


// Organization resource
type Organization struct {
	Domain
	Identifier    []d.Identifier              `json:"identifier,omitempty"`
	Active        *d.Boolean                  `json:"active,omitempty"`
	Name          *d.String              	  `json:"name,omitempty"`
	Telecom       []d.ContactPoint            `json:"telecom,omitempty"`
	Address       []d.Address                 `json:"address,omitempty"`
}


func NewOrganization() *Organization {

	org := Organization{}
	org.ResourceType = "Organization"

	// domain attributes
	b := NewDomain(org.ResourceType, d.ResourceOrganization)
	org.ID=b.ID
	org.Meta =b.Meta
	org.Text=b.Text
	org.Contained =b.Contained
	org.Extension=b.Extension

	// Domain Resource Attributes
	telecom := make([]d.ContactPoint, 0)
	address := make([]d.Address, 0)
	identifiers := make([]d.Identifier, 0)

	org.Telecom = telecom
	org.Address = address
	org.Identifier = identifiers

	// update model 
	m.AddModelEntry(org.ResourceType, *org.ID, org)
    return &org
}

