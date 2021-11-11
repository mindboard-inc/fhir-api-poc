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

	// Authors
	// 	author		        .author.assignedAuthor
	// Provider Organization requested in the sample as opposed to Practitioner
	
	author := MappedOrganization(hl7cd)
	errorg := r.SetCompositionAuthorOrganization(comp, author)
	if (errorg != nil) {
		return nil
	}

/*
	subject		        .recordTarget -> <recordTarget><patientRole>
	encounter		    .componentOf.encompassingEncounter
	// Subject
	// Encounter
*/
	// Attester		        
	// Event 
	// Sections

	
	return comp
}

func MappedClinicalDocument(hl7cd h.ClinicalDocument) *p.ClinicalDocument {
	// Prepare a new/mapped clinical document
	cda := p.NewClinicalDocument()
	comp := MappedComposition(hl7cd)

	p.SetClinicalDocumentComposition(cda, comp)
	if err:=p.FinalizeClinicalDocumentBundle(cda); err !=nil {
		return nil
	}
	return cda
}



/*	8.6.9.3 RIM Mapping (http://hl7.org/v3 ) 

	Organization	Organization(classCode=ORG, determinerCode=INST)
    identifier	.scopes[Role](classCode=IDENT)
    active	.status
    type	.code
    name	.name
    alias	.name
    telecom	.telecom
    address	.address
    partOf	.playedBy[classCode=Part].scoper
    contact	.contactParty
        purpose	./type
        name	./name
        telecom	./telecom
        address	./addr
    endpoint	n/a
*/

/* SAMPLE CDA
	<author>
        <time value="20050329224411+0500" />
        <assignedAuthor>
            <id extension="99999999" root="2.16.840.1.113883.4.6" />
            <addr>
                <streetAddressLine>1002 Healthcare Drive </streetAddressLine>
                <city>Portland</city>
                <state>OR</state>
                <postalCode>99123</postalCode>
                <country>US</country>
            </addr>
            <telecom use="WP" value="tel:555-555-1002" />
            <assignedAuthoringDevice>
                <!-- Software Originator -->
                <manufacturerModelName>Good Health Medical Device </manufacturerModelName>
                <!-- Software Version number -->
                <softwareName>Good Health Document Generator version 2.3 </softwareName>
            </assignedAuthoringDevice>
        </assignedAuthor>
    </author>
*/

func MappedOrganization(hl7cd h.ClinicalDocument) *r.Organization {

	// Prepare a new/mapped organization
	org := r.NewOrganization()
	
	// ID
	//fmt.Printf("Author.AssignedAuthor.ID.Root: %+v\n", hl7cd.Author.AssignedAuthor.ID.Root)
	//fmt.Printf("Author.AssignedAuthor.ID.Extension: %+v\n", hl7cd.Author.AssignedAuthor.ID.Extension)

	// Identifier -- Skipping extension, only adding root as value
	var idRoot d.String
	idRoot = d.String(hl7cd.Author.AssignedAuthor.ID.Root)
	identifier := &d.Identifier{}
	identifier.Value = &idRoot
	org.Identifier = append(org.Identifier , *identifier)


	// Address

	//fmt.Printf("hl7cd.Author.AssignedAuthor.Addr.StreetAddressLine: %+v\n", hl7cd.Author.AssignedAuthor.Addr.StreetAddressLine)
	//fmt.Printf("hl7cd.Author.AssignedAuthor.Addr.City: %+v\n", hl7cd.Author.AssignedAuthor.Addr.City)
	//fmt.Printf("hl7cd.Author.AssignedAuthor.Addr.State: %+v\n", hl7cd.Author.AssignedAuthor.Addr.State)
	//fmt.Printf("hl7cd.Author.AssignedAuthor.Addr.PostalCode: %+v\n", hl7cd.Author.AssignedAuthor.Addr.PostalCode)
	//fmt.Printf("hl7cd.Author.AssignedAuthor.Addr.Country: %+v\n", hl7cd.Author.AssignedAuthor.Addr.Country)

	address := &d.Address{}

	lines := make([]d.String, 0)
	var line d.String 
	line = d.String(hl7cd.Author.AssignedAuthor.Addr.StreetAddressLine)
	lines = append(lines, line)
	address.Line = lines

	var city d.String 
	city = d.String(hl7cd.Author.AssignedAuthor.Addr.City)
	address.City = &city

	var state d.String 
	state = d.String(hl7cd.Author.AssignedAuthor.Addr.State)
	address.State = &state

	var postalcode d.String 
	postalcode = d.String(hl7cd.Author.AssignedAuthor.Addr.PostalCode)
	address.PostalCode = &postalcode

	var country d.String 
	country = d.String(hl7cd.Author.AssignedAuthor.Addr.Country)
	address.City = &country

	org.Address = append(org.Address , *address)

	// Telecom -> ContactPoint
	// 					Value  *String      `json:"value,omitempty"`
	//					Use    *Code        `json:"use,omitempty"`
	//fmt.Printf("hl7cd.Author.AssignedAuthor.Telecom.Use: %+v\n", hl7cd.Author.AssignedAuthor.Telecom.Use)
	// fmt.Printf("hl7cd.Author.AssignedAuthor.Telecom.Value: %+v\n", hl7cd.Author.AssignedAuthor.Telecom.Value)
	telecom := &d.ContactPoint{}
	var useCode d.Code
	useCode = d.Code(hl7cd.Author.AssignedAuthor.Telecom.Use)
	telecom.Use = &useCode
	var telvalue d.String
	telvalue = d.String(hl7cd.Author.AssignedAuthor.Telecom.Value)
	telecom.Value = &telvalue

	org.Telecom = append(org.Telecom , *telecom)
	return org
	
}