// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources


import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	m "mindboard.com/fhirpoc/pkg/model"
)


// Patient resource
type Organization struct {
	Domain
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

	// update model 
	m.AddModelEntry(org.ResourceType, *org.ID, org)
    return &org
}

