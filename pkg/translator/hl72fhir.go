package translator

import (
	//"errors"
	"fmt"
	"strconv"

	h "mindboard.com/fhirpoc/pkg/hl73"
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	r "mindboard.com/fhirpoc/pkg/fhir/resources"
	p "mindboard.com/fhirpoc/pkg/fhir/profiles"
)

func MappedComposition(hl7cd h.ClinicalDocument) *r.Composition {

	comp:= r.NewComposition()
	enc:=r.NewEncounter()
	err := r.SetCompositionEncounter(comp, enc)
	if (err != nil) {
		return nil
	}
	fmt.Printf("extension.ValueInteger:  %v\n", comp.Status)

	/*
	fmt.Printf("RecordTarget.PatientRole.ID[0]: %+v\n", cd.RecordTarget.PatientRole.ID[0].Root)
	fmt.Printf("RecordTarget.Author.Time.Value: %+v\n", cd.Author.Time.Value)
	fmt.Printf("RecordTarget.Author.Time.Text: %+v\n", cd.Author.Time.Text)
	fmt.Printf("RecordTarget.RealmCode.Code: %+v\n", cd.RealmCode.Code)
	*/

	// Extension
	// Extension []d.Extension       `json:"extension,omitempty"`
	// MAP 	extension	-> versionNumber	
	extensions := make([]d.Extension, 0)
	extension := &d.Extension {}
	var versionNumber d.Integer
	versionnumerint, _ := strconv.Atoi(hl7cd.VersionNumber.Value)

	versionNumber = d.Integer(versionnumerint)
	extension.ValueInteger = &versionNumber
	//fmt.Printf("hl7cd.VersionNumber.Value:  %v\n", hl7cd.VersionNumber.Value)
	//fmt.Printf("extension.ValueInteger:  %v\n", extension.ValueInteger)
	extensions = append(extensions, *extension)
	comp.Extension = extensions

	// identifier		    .setId
	// Identifier
	var setIdRoot d.String
	setIdRoot = d.String(hl7cd.SetId.Root)
	identifier := &d.Identifier{}
	identifier.Value = &setIdRoot
	comp.Identifier = identifier



	// Type
	// type		            .code
	//<code codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" code="75619-7" displayName="National Ambulatory Medical Care Survey" />
	var codeSystem d.URI
	var displayName d.String
	var code d.Code

	codeType := &d.Coding{}

	codeSystem = d.URI(hl7cd.Code.CodeSystem) 
	code = d.Code(hl7cd.Code.Code)
	displayName= d.String(hl7cd.Code.DisplayName)

	codeType.System = &codeSystem
	codeType.Code = &code
	codeType.Display = &displayName
	
	comp.Type = codeType

	// Title
	// title		        .title
	var title d.String
	title = d.String(hl7cd.Title)
	comp.Title = &title 

	
	// Confidentiality 
	// confidentiality		.confidentialityCode
	var confidentialityCode d.String
	confidentialityCode = d.String(hl7cd.ConfidentialityCode.Code)
	comp.Confidentiality = &confidentialityCode 

	// Date date	
	// date		            .effectiveTime

	var effectiveTime d.Date
	effectiveTime = d.Date(hl7cd.EffectiveTime.Value)
	comp.Date = &effectiveTime 

/*
	subject		        .recordTarget -> <recordTarget><patientRole>
	encounter		    .componentOf.encompassingEncounter
	author		        .author.assignedAuthor

	// Subject
	// Encounter
	// Authors
	authors := make([]d.Reference, 0)
	author := dummyOrganization()
	var ref string
	var authorid string
	var refstr d.String
	authorid = string(*author.ID)
	ref = author.ResourceType + "/" + string(authorid)
	refstr = d.String(ref)
	authorref := &d.Reference {}
	authorref.Reference = &refstr
	authors = append(authors, *authorref)
	comp.Author = authors
	// Attester		        
	// Event 
	// Sections
*/
	
	return comp
}

func MappedClinicalDocument(hl7cd h.ClinicalDocument) *p.ClinicalDocument {
	// Prepare a dummy clinical document
	cda := p.NewClinicalDocument()
	comp := MappedComposition(hl7cd)

	p.SetClinicalDocumentComposition(cda, comp)
	if err:=p.FinalizeClinicalDocumentBundle(cda); err !=nil {
		return nil
	}
	return cda
}

