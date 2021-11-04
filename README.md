# fhir-poc
 

## Docker Build and Testing

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

### Build

docker build --tag mindboard.com/fhir-api-poc .

### Run

docker run --publish 1500:1500 mindboard.com/fhir-api-poc:latest

### Test

curl -F 'payload=@test/in.xml' http://127.0.0.1:8080/hl72fhir


