// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	"gopkg.in/mgo.v2/bson"

	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Base Resource
type Base struct {
	ResourceType  string  `json:"resourceType"`
	ID            *d.ID   `json:"id,omitempty"`
	// <xs:element name="meta" type="Meta" minOccurs="0" maxOccurs="1">
	Meta          *d.Meta `json:"meta,omitempty"`
	ImplicitRules *d.URI  `json:"implicitRules,omitempty"`
	Language      *d.Code `json:"language,omitempty"`
}

func NewBase (resourcetype string,  profileuri string) *Base {
	b := Base{}
	var id d.ID
	id = d.ID(bson.NewObjectId().Hex())
	b.ID = &id
	b.ResourceType = resourcetype
	meta := &d.Meta{}
	// minimal profile
	// Profile     []URI `json:"profile,omitempty"`
	var profile d.URI
	profile = d.URI(profileuri)
	profiles := make([]d.URI, 0)
	profiles = append(profiles, profile)
	meta.Profile = profiles
	b.Meta = meta
	return &b

}




