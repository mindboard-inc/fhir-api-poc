// Package resources provides all resources implementation of FHIR v3.0.2
//
// Reference: http://hl7.org/fhir/STU3/resourcelist.html
package resources

import (
	"time"
	"errors"

	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	m "mindboard.com/fhirpoc/pkg/model"
)

// Section for  CDA clinical document bundle composition resource
type CompositionSection struct {
	Title *d.String `json:"title,omitempty"`
	Code  *d.Coding `json:"code,omitempty"`
	Text *d.Narrative `json:"text,omitempty"`
	Entry []d.Reference `json:"entry,omitempty"`
}

// CDA clinical document bundle composition resource
type Composition struct {
	Domain
	Identifier *d.Identifier  	`json:"identifier,omitempty"`
	// <xs:element name="date" minOccurs="1" maxOccurs="1" type="dateTime">
	// date		            .effectiveTime
	Date 	 *d.Date			 `json:"date,omitempty"`

	// <xs:element name="type" minOccurs="1" maxOccurs="1" type="CodeableConcept">
	Type 	 *d.Coding 			`json:"type,omitempty"`

	//<xs:element name="class" minOccurs="0" maxOccurs="1" type="CodeableConcept">
	Class   *d.String        	`json:"class,omitempty"`

	// <xs:element name="status" minOccurs="1" maxOccurs="1" type="CompositionStatus">
	Status   *d.String        	`json:"status,omitempty"`

	//<xs:element name="subject" minOccurs="1" maxOccurs="1" type="Reference">
	Subject  *d.Reference        `json:"subject,omitempty"`

	Encounter  *d.Reference      `json:"encounter,omitempty"`

	// <xs:element name="author" minOccurs="1" maxOccurs="unbounded" type="Reference">
	
	Author   []d.Reference       `json:"author,omitempty"`
	Title    *d.String       	 `json:"title,omitempty"`
	Confidentiality *d.String    `json:"confidentiality,omitempty"`
	Attester []d.Attester		 `json:"attester,omitempty"`
	Custodian *d.Reference       `json:"custodian,omitempty"`
	Event *d.ServiceEvent        `json:"event,omitempty"`
	Section []CompositionSection `json:"section,omitempty"`

}


func NewComposition() *Composition {

	comp := Composition{}
	comp.ResourceType = "Composition"

	// domain attributes
	b := NewDomain(comp.ResourceType, d.ResourceComposition)
	comp.ID=b.ID
	comp.Meta =b.Meta
	comp.Text=b.Text
	comp.Contained =b.Contained
	comp.Extension=b.Extension

	// Composition attributes
	// Default mandatory attributes
	// Date: bundle packaging timestamp
	t := time.Now()
	// Date date
	var effectiveTime d.Date
	effectiveTime = d.Date(t.Format("20060102150405-0000"))
	comp.Date = &effectiveTime 

	// Type: Generic / Interim Doc Type
	var codeSystem d.URI
	var displayName d.String
	var code d.Code

	codeType := &d.Coding{}

	codeSystem = d.CompositionTypeCodeSystemDefault
	code = d.CompositionTypeCodeDefault
	displayName=d.CompositionTypeDisplayNameDefault

	codeType.System = &codeSystem
	codeType.Code = &code
	codeType.Display = &displayName
	
	comp.Type = codeType

	// Status: Generic / Interim Status

	// Status
	var status d.String
	status = d.CompositionStatusPreliminary
	comp.Status = &status

	// Title: Default Title
	var title d.String
	title = d.CompositionTitleDefault
	comp.Title = &title 

	// Patient: Default Patient

	subject := NewPatient()
	var pref string
	var patientid string
	var prefstr d.String

	patientid = string(*subject.ID)
	pref = subject.ResourceType + "/" + string(patientid)
	prefstr = d.String(pref)
	patientref := &d.Reference {}
	patientref.Reference = &prefstr

	comp.Subject = patientref

	// Author: No Default Author
	authors := make([]d.Reference, 0)
	comp.Author = authors

	// Non Mandatory Attributes
	// Identifier
	var setIdRoot d.String
	setIdRoot = d.CompositionIdentifierDefault
	identifier := &d.Identifier{}
	identifier.Value = &setIdRoot
	comp.Identifier = identifier

	// Confidentiality 
	var confidentialityCode d.String
	confidentialityCode = d.CompositionConfidentialityCodeDefault
	comp.Confidentiality = &confidentialityCode 

	// Attester - Empty Array 
	attesters := make([]d.Attester, 0)
	comp.Attester = attesters
	// Section - Empty Array 
	sections := make([]CompositionSection, 0)
	comp.Section = sections
	// update model 
	m.AddModelEntry(comp.ResourceType, *comp.ID, comp)
    return &comp
}

func SetCompositionEncounter(comp *Composition, enc *Encounter) error {
	if (comp == nil || enc == nil) {
		return errors.New("Unable to Set encounter.")
	}
	var eref string
	var encid string
	var erefstr d.String
	
	encid = string(*enc.ID)
	eref = enc.ResourceType + "/" + string(encid)
	erefstr = d.String(eref)
	encref := &d.Reference {}
	encref.Reference = &erefstr
	comp.Encounter = encref
	return nil
}

func SetCompositionAuthorOrganization(comp *Composition, org *Organization) error {
	var author *Organization

	if (comp == nil) {
		return errors.New("Unable to add author/organization.")
	}
	if (org == nil) {
		author = NewOrganization()
	} else {
		author = org
	}
	var ref string
	var authorid string
	var refstr d.String

	authorid = string(*author.ID)
	ref = author.ResourceType + "/" + string(authorid)
	refstr = d.String(ref)
	authorref := &d.Reference {}
	authorref.Reference = &refstr

	comp.Author = append(comp.Author, *authorref)

	// make sure the model is updated
	if err:= m.UpdateModelEntry("Organization", *author.ID, author); err != nil {
			return err
	}
	
	return nil
}



