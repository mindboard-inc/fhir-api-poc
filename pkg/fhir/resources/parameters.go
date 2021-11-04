// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)


// Parameters resource
type Parameters struct {
	Base
	Parameter []ParametersParameter `json:"parameter,omitempty"`
}

// ParametersParameter subResource
type ParametersParameter struct {
	Name                 *d.String          `json:"name,omitempty"`
	ValueInteger         *d.Integer         `json:"valueInteger,omitempty"`
	ValueDecimal         *d.Decimal         `json:"valueDecimal,omitempty"`
	ValueDateTime        *d.DateTime        `json:"valueDateTime,omitempty"`
	ValueDate            *d.Date            `json:"valueDate,omitempty"`
	ValueInstant         *d.Instant         `json:"valueInstant,omitempty"`
	ValueString          *d.String          `json:"valueString,omitempty"`
	ValueURI             *d.URI             `json:"valueUri,omitempty"`
	ValueBoolean         *d.Boolean         `json:"valueBoolean,omitempty"`
	ValueCode            *d.Code            `json:"valueCode,omitempty"`
	ValueBase64Binary    *d.Base64Binary    `json:"valueBase64Binary,omitempty"`
	ValueCoding          *d.Coding          `json:"valueCoding,omitempty"`
	ValueCodeableConcept *d.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueAttachment      *d.Attachment      `json:"valueAttachment,omitempty"`
	ValueIdentifier      *d.Identifier      `json:"valueIdentifier,omitempty"`
	ValueQuantity        *d.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *d.Range           `json:"valueRange,omitempty"`
	ValuePeriod          *d.Period          `json:"valuePeriod,omitempty"`
	ValueRatio           *d.Ratio           `json:"valueRatio,omitempty"`
	ValueHumanName       *d.HumanName       `json:"valueHumanName,omitempty"`
	ValueAddress         *d.Address         `json:"valueAddress,omitempty"`
	ValueContactPoint    *d.ContactPoint    `json:"valueContactPoint,omitempty"`
	ValueSchedule        interface{}        `json:"valueSchedule,omitempty"`
	ValueReference       *d.Reference       `json:"valueReference,omitempty"`
	Resource             interface{}        `json:"resource,omitempty"`
	Part                 interface{}        `json:"part,omitempty"`
}
