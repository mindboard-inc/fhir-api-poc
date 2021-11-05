# Hl7 v3 to Fhir v4 Clinical Document Object Mapping POC

Extends Golang structs as defined in  https://goreportcard.com/badge/github.com/monarko/fhirgo to include ClinicalDocument and Composition objects.

Implements mapping per Fhir v4 specifications http://hl7.org/fhir/composition-mappings.html

This code base is heavily under construction. Current version is the baseline with Compositional level mappings with ClinicalDocument Bundle containing required Subject: Patient, Author: Organization and Encounter 

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


## Docker Build and Testing Tips

Note for in WSL2 Desktop Integration: Disable docker buildkit  under engine config prior to image build:

```
{
“features”: {
“buildkit”: false
},
“experimental”: false
}
```
*per https://forums.docker.com/t/strange-docker-output-or-help-me-please-im-very-noob/100788*