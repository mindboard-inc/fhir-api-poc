package model

import (
	"errors"

	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)
var Fhir_MAP = make(map[string]interface{})

func AddModelEntry(resourcetype string, id d.ID, v interface{}) {
	var idstr, key string 
	idstr = string(id)
	key = resourcetype + "/" + idstr

	Fhir_MAP[key]=v 
}

func UpdateModelEntry(resourcetype string, id d.ID, v interface{}) error {
	var idstr, key string 
	idstr = string(id)
	key = resourcetype + "/" + idstr

	_, ok := Fhir_MAP[key]
	if ok {
		Fhir_MAP[key]=v 
		return nil
	} else {
		return errors.New("No model object with key")
	}
}

func GetModelSize() int {
	return len(Fhir_MAP)
}

func ResetModel() {
	Fhir_MAP = make(map[string]interface{})
}

func GetModelEntry(resourcetype string, id d.ID) (interface{}, error) {
	var idstr, key string 
	idstr = string(id)
	key = resourcetype + "/" + idstr
	obj, ok := Fhir_MAP[key]
	if ok {
		return obj, nil
	} else {
		return "", errors.New("No model object with key")
	}
}
func GetModelEntryByReference(key string) (interface{}, error) {
	obj, ok := Fhir_MAP[key]
	if ok {
		return obj, nil
	} else {
		return "", errors.New("No model object with key")
	}
}
