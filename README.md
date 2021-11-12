# Hl7 v3 to Fhir v4 Clinical Document Object Mapping POC

Extends Golang structs as defined in  https://goreportcard.com/badge/github.com/monarko/fhirgo to include ClinicalDocument profile, and Organization and Composition resources.

Implements mapping per Fhir v4 specifications http://hl7.org/fhir/composition-mappings.html

This code base is heavily under construction. Current version is the baseline with mappings implemented for Compositional and Provider [Organization] as Author inside the ClinicalDocument Bundle with the required Subject: Patient and Encounter references and place holders.

### Build & Run

`make build`

`make run`

## Container Image

### Docker Build

`docker build --tag mindboard.com/fhir-api-poc .`

### Docker Run

`docker run --publish 1500:1500 mindboard.com/fhir-api-poc:latest`

### Docker Test

`curl -F 'payload=@test/in.xml' http://localhost:1500/hl72fhir`

