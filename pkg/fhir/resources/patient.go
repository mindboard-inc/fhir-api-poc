// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (

	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	m "mindboard.com/fhirpoc/pkg/model"
)

// Patient resource
type Patient struct {
	Domain

	Identifier           []d.Identifier         `json:"identifier,omitempty"`
	Active               *d.Boolean             `json:"active,omitempty"`
	Name                 []d.HumanName          `json:"name,omitempty"`
	Telecom              []d.ContactPoint       `json:"telecom,omitempty"`
	Gender               *d.Code                `json:"gender,omitempty"`
	BirthDate            *d.Date                `json:"birthDate,omitempty"`
	DeceasedBoolean      *d.Boolean             `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *d.DateTime            `json:"deceasedDateTime,omitempty"`
	Address              []d.Address            `json:"address,omitempty"`
	MaritalStatus        *d.CodeableConcept     `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *d.Boolean             `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *d.Integer             `json:"multipleBirthInteger,omitempty"`
	Photo                []d.Attachment         `json:"photo,omitempty"`
	Contact              []PatientContact       `json:"contact,omitempty"`
	Animal               *PatientAnimal         `json:"animal,omitempty"`
	Communication        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner  []d.Reference          `json:"generalPractitioner,omitempty"`
	ManagingOrganization *d.Reference           `json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `json:"link,omitempty"`
}


// PatientContact subResource
type PatientContact struct {
	Relationship []d.CodeableConcept `json:"relationship,omitempty"`
	Name         *d.HumanName        `json:"name,omitempty"`
	Telecom      []d.ContactPoint    `json:"telecom,omitempty"`
	Address      *d.Address          `json:"address,omitempty"`
	Gender       *d.Code             `json:"gender,omitempty"`
	Organization *d.Reference        `json:"patanization,omitempty"`
	Period       *d.Period           `json:"period,omitempty"`
}

// PatientAnimal subResource
type PatientAnimal struct {
	Species      *d.CodeableConcept `json:"species,omitempty"`
	Breed        *d.CodeableConcept `json:"breed,omitempty"`
	GenderStatus *d.CodeableConcept `json:"genderStatus,omitempty"`
}

// PatientCommunication subResource
type PatientCommunication struct {
	Language  *d.CodeableConcept `json:"language,omitempty"`
	Preferred *d.Boolean         `json:"preferred,omitempty"`
}

// PatientLink subResource
type PatientLink struct {
	Other *d.Reference `json:"other,omitempty"`
	Type  *d.Code      `json:"type,omitempty"`
}

func NewPatient() *Patient{

	pat := Patient{}
	pat.ResourceType = "Patient"

	// domain attributes
	b := NewDomain(pat.ResourceType, d.ResourcePatient)
	pat.ID=b.ID
	pat.Meta =b.Meta
	pat.Text=b.Text
	pat.Contained =b.Contained
	pat.Extension=b.Extension

	// update model 
	m.AddModelEntry(pat.ResourceType, *pat.ID, pat)
    return &pat
}

