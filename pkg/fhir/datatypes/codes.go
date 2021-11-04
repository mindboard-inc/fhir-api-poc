    // DAF Profile URLs (based on FHIR DSTU2)
	package datatypes

	const ( 
		ProfileAllergyIntolerance="http://hl7.org/fhir/StructureDefinition/allergyintolerance"
		ProfileAlergyList= "http://hl7.org/fhir/StructureDefinition/allergylist"
		ProfileCondition= "http://hl7.org/fhir/StructureDefinition/condition"
		ProfileDiagnosticOrder= "http://hl7.org/fhir/StructureDefinition/diagnosticorder"
		ProfileDiagnosticReport= "http://hl7.org/fhir/StructureDefinition/diagnosticreport"
		ProfileEncounter= "http://hl7.org/fhir/StructureDefinition/encounter"
		ProfileEncounterList= "http://hl7.org/fhir/StructureDefinition/encounterlist"
		ProfileFamilyMemberHistory= "http://hl7.org/fhir/StructureDefinition/familymemberhistory"
		ProfileImmunization= "http://hl7.org/fhir/StructureDefinition/immunization"
		ProfileImmunizationList= "http://hl7.org/fhir/StructureDefinition/immunizationlist"
		ProfileLocation= "http://hl7.org/fhir/StructureDefinition/location"
		ProfileMedication= "http://hl7.org/fhir/StructureDefinition/medication"
		ProfileMedicationAdministration= "http://hl7.org/fhir/StructureDefinition/medicationadministration"
		ProfileMedicationDispense= "http://hl7.org/fhir/StructureDefinition/medicationdispense"
		ProfileMedicationList= "http://hl7.org/fhir/StructureDefinition/medicationlist"
		ProfileMedicationOrder= "http://hl7.org/fhir/StructureDefinition/medicationorder"
		ProfileMedicationStatement= "http://hl7.org/fhir/StructureDefinition/medicationstatement"
		ProfileOrganization= "http://hl7.org/fhir/StructureDefinition/organization"
		ProfilePatient= "http://hl7.org/fhir/StructureDefinition/patient"
		ProfilePractitioner= "http://hl7.org/fhir/StructureDefinition/pract"
		ProfileProblemList= "http://hl7.org/fhir/StructureDefinition/problemlist"
		ProfileProcedure= "http://hl7.org/fhir/StructureDefinition/procedure"
		ProfileProcedureList= "http://hl7.org/fhir/StructureDefinition/procedurelist"
		ProfileRelatedPerson= "http://hl7.org/fhir/StructureDefinition/relatedperson"
		ProfileResultList= "http://hl7.org/fhir/StructureDefinition/resultlist"
		ProfileResultJobs= "http://hl7.org/fhir/StructureDefinition/resultobs"
		ProfileSmokingStatus= "http://hl7.org/fhir/StructureDefinition/smokingstatus"
		ProfileSpecimen= "http://hl7.org/fhir/StructureDefinition/spec"
		ProfileSubstance= "http://hl7.org/fhir/StructureDefinition/substance"
		ProfileVitalSigns= "http://hl7.org/fhir/StructureDefinition/vitalsigns"
		// Extension URLs
		ProfileExtensionBirthPlace= "http://hl7.org/fhir/StructureDefinition/birthPlace"
		ProfileExtensionDataAbsentReason= "http://hl7.org/fhir/StructureDefinition/data-absent-reason"
		ProfileExtensionEthnicity= "http://hl7.org/fhir/StructureDefinition/us-core-ethnicity"
		ProfileExtensionRace= "http://hl7.org/fhir/StructureDefinition/us-core-race"
		ProfileExtensionReligion= "http://hl7.org/fhir/StructureDefinition/us-core-religion"

		ProfileClinicalDocument="http://hl7.org/fhir/StructureDefinition/bundle"
		
	)

	// Mandatory Composition Status Enumeration
	const (
		CompositionStatusPreliminary="preliminary" 	// 	This is a preliminary composition or document (also known as initial or interim). The content may be incomplete or unverified.
		CompositionStatusFinal="final" 			// This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition.
	    CompositionStatusAmended="amended" 		//	The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as "final" and the composition is complete and verified by an authorized person.
		ComposotionStatusEnteredInError="entered-in-error" // sThe composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid.
	)

	// Defaults Composition Type LOINC
	const (
		CompositionTypeCodeSystemDefault="http://loinc.org"
		CompositionTypeCodeDefault="69376-2"
		CompositionTypeDisplayNameDefault="Current medical information"
	)

	// Other Composition Defaults
	  const (
			CompositionIdentifierDefault="0000000000000" 	
			CompositionTitleDefault="New Composition Document" 	
			CompositionConfidentialityCodeDefault="N"
	  )

		
	// Resource 
	const (
		ResourceComposition="https://www.hl7.org/fhir/composition"
		ResourceOrganization="https://www.hl7.org/fhir/organization"
		ResourcePatient="https://www.hl7.org/fhir/patient"
		ResourceEncounter="https://www.hl7.org/fhir/encounter"
		ResourcePractitioner="https://www.hl7.org/fhir/practitioner"
	)
