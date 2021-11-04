package profiles

import (
	"errors"

	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	r "mindboard.com/fhirpoc/pkg/fhir/resources"
	m "mindboard.com/fhirpoc/pkg/model"
)


// CDA clinical document bundle 
type ClinicalDocument struct {
	r.Base
	Identifier *d.Identifier  `json:"identifier,omitempty"`
	Total      *d.UnsignedInt `json:"total,omitempty"`
	Link       []r.BundleLink   `json:"link,omitempty"`
	Entry      []r.BundleEntry  `json:"entry,omitempty"`
	Signature  *d.Signature   `json:"signature,omitempty"`
}


func NewClinicalDocument() *ClinicalDocument {

	cd := ClinicalDocument{}
	cd.ResourceType = "Bundle"
	entries := make([]r.BundleEntry, 0)

	b := r.NewBase(cd.ResourceType, d.ProfileClinicalDocument)
	// Base Attributes 
	cd.ID=b.ID
	// cd.ResourceType=b.ResourceType
	cd.Meta =b.Meta
	
	// Clinical Document Attributes
	cd.Entry = entries

	// update model 
	m.AddModelEntry(cd.ResourceType, *cd.ID, cd)
    return &cd
}

func SetClinicalDocumentComposition(cd *ClinicalDocument, comp *r.Composition)  error {

	entry := &r.BundleEntry{}
	entry.Resource = *comp
	//cd.Entry = append(cd.Entry, *entry)
	if (len(cd.Entry) > 0) {
		cd.Entry[0]=*entry
	} else {
		cd.Entry = append(cd.Entry, *entry)
	}

	return nil
}

func SetClinicalDocumentEncounter(cd *ClinicalDocument, enc *r.Encounter) {
	var comp r.Composition
	if (len(cd.Entry) > 0) {
		comp = cd.Entry[0].Resource.(r.Composition)
	}
	r.SetCompositionEncounter(&comp, enc)
}

func addClinicalDocumentBundleEntry(cd *ClinicalDocument, ref *d.String)  error {

	entry := &r.BundleEntry{}
	if v, error := m.GetModelEntryByReference(string(*ref)); error != nil {
		return errors.New("No model object with reference")
	} else {
		entry.Resource = v
		cd.Entry = append(cd.Entry, *entry)
	}

	return nil

}

func FinalizeClinicalDocumentBundle(cd *ClinicalDocument)  error {

	var comp r.Composition
	if (len(cd.Entry) > 0) {
		comp = cd.Entry[0].Resource.(r.Composition)
	} 

	// Patient
	subject := comp.Subject
	addClinicalDocumentBundleEntry(cd, subject.Reference)
	// Author / Organization
	if (len(comp.Author) > 0) {
		for _, org := range comp.Author {
			addClinicalDocumentBundleEntry(cd, org.Reference)
		}
	}

	// Encounter
	
	enc := comp.Encounter
	if (enc != nil) {
		addClinicalDocumentBundleEntry(cd, enc.Reference)
	}
	
	// make sure the model is updated
	if err:= m.UpdateModelEntry("Bundle", *cd.ID, cd); err != nil {
			return err
	}
	
	total := d.UnsignedInt(len(cd.Entry))
	cd.Total = &total
	return nil
}

