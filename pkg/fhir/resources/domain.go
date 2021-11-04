// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	"gopkg.in/mgo.v2/bson"
	
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)

// Domain resource
type Domain struct {
	Base
	// <xs:element name="text" type="Narrative" minOccurs="0" maxOccurs="1">
	Text      *d.Narrative `json:"text,omitempty"`
	// <xs:element name="contained" type="ResourceContainer" minOccurs="0" maxOccurs="unbounded">
	Contained []Base       `json:"contained,omitempty"`
	// <xs:element name="extension" type="Extension" minOccurs="0" maxOccurs="unbounded">
	Extension []d.Extension       `json:"extension,omitempty"`
}
 
func NewDomain (resourcetype string,  profileuri string) *Domain {

	// base attributes
	b := Domain{}
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

	// Domain Resource Attributes
	containees := make([]Base, 0)
	extensions := make([]d.Extension, 0)
	text := &d.Narrative{}
	b.Contained = containees
	b.Extension = extensions
	b.Text = text

	return &b

}