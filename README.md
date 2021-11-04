# fhir-poc
 
https://forums.docker.com/t/strange-docker-output-or-help-me-please-im-very-noob/100788
disable docker buildkit

{
“features”: {
“buildkit”: false
},
“experimental”: false
}



docker build --tag mindboard.com/fhir-api-poc .

docker run --publish 1500:1500 mindboard.com/fhir-api-poc:latest
