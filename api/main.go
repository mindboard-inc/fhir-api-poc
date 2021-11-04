package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	h "mindboard.com/fhirpoc/pkg/hl73"
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
	p "mindboard.com/fhirpoc/pkg/fhir/profiles"
	m "mindboard.com/fhirpoc/pkg/model"
	t "mindboard.com/fhirpoc/pkg/translator"
)

func map2Fhir (c echo.Context) error {
	m.ResetModel()

	xmlFile, err := c.FormFile("payload")
	if err != nil {
		return err
	}
	src, err := xmlFile.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	byteValue, _ := ioutil.ReadAll(src)
	
	var hl7cda h.ClinicalDocument
    xml.Unmarshal(byteValue, &hl7cda)
	// hl7cda.Reflect()

	//fmt.Printf("INXML: %+v\n", string(byteValue))

	//fmt.Printf("RecordTarget.PatientRole.ID[0]: %+v\n", hl7cda.RecordTarget.PatientRole.ID[0].Root)
	fmt.Printf("RecordTarget.Author.Time.Value: %+v\n", hl7cda.Author.Time.Value)
	fmt.Printf("RecordTarget.Author.Time.Text: %+v\n", hl7cda.Author.Time.Text)
	fmt.Printf("RecordTarget.RealmCode.Code: %+v\n", hl7cda.RealmCode.Code)

	

	cda := t.MappedClinicalDocument(hl7cda)
	if (cda == nil) {
		fmt.Printf("Error in mapping.")
	}

	if err := c.Bind(cda); err != nil {
		return err
	}
	fmt.Printf("Model counter:  %v\n", m.GetModelSize())

	if cdalt, error := m.GetModelEntry("Bundle", *cda.ID); error != nil {
		fmt.Printf("Error in retrieving Model Entry")
	} else {
		cdv := cdalt.(*p.ClinicalDocument)
		var total *d.UnsignedInt
		total = cdv.Total
		fmt.Printf("CD Entries Total:  %v\n", total)
	}
	
	return c.JSON(http.StatusOK, cda)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/hl72fhir", map2Fhir)

	// Start server
	e.Logger.Fatal(e.Start(":1500"))

}