// FHIR GO - Golang Implementation of FHIR Data Types and Resources
// Via --
// [![Go Report Card](https://goreportcard.com/badge/github.com/monarko/fhirgo)](https://goreportcard.com/report/github.com/monarko/fhirgo)

package resources

import (
	d "mindboard.com/fhirpoc/pkg/fhir/datatypes"
)


// Bundle resource
type Bundle struct {
	Base
	Identifier *d.Identifier  `json:"identifier,omitempty"`
	Type       *d.Code        `json:"type,omitempty"`
	Total      *d.UnsignedInt `json:"total,omitempty"`
	Link       []BundleLink   `json:"link,omitempty"`
	Entry      []BundleEntry  `json:"entry,omitempty"`
	Signature  *d.Signature   `json:"signature,omitempty"`
}


// BundleLink subResource
type BundleLink struct {
	Relation *d.String `json:"relation,omitempty"`
	URL      *d.URI    `json:"url,omitempty"`
}

// BundleEntrySearch subResource
type BundleEntrySearch struct {
	Mode  *d.Code    `json:"mode,omitempty"`
	Score *d.Decimal `json:"score,omitempty"`
}

// BundleEntryRequest subResource
type BundleEntryRequest struct {
	Method          *d.Code    `json:"method,omitempty"`
	URL             *d.URI     `json:"url,omitempty"`
	IfNoneMatch     *d.String  `json:"ifNoneMatch,omitempty"`
	IfModifiedSince *d.Instant `json:"ifModifiedSince,omitempty"`
	IfMatch         *d.String  `json:"ifMatch,omitempty"`
	IfNoneExist     *d.String  `json:"ifNoneExist,omitempty"`
}

// BundleEntryResponse subResource
type BundleEntryResponse struct {
	Status       *d.String   `json:"status,omitempty"`
	Location     *d.URI      `json:"location,omitempty"`
	Etag         *d.String   `json:"etag,omitempty"`
	LastModified *d.Instant  `json:"lastModified,omitempty"`
	Outcome      interface{} `json:"outcome,omitempty"` // Resource
}

// BundleEntry subResource
type BundleEntry struct {
	Link     []BundleLink         `json:"link,omitempty"`
	FullURL  *d.URI               `json:"fullUrl,omitempty"`
	Resource interface{}          `json:"resource,omitempty"` // A resource in the bundle
	Search   *BundleEntrySearch   `json:"search,omitempty"`
	Request  *BundleEntryRequest  `json:"request,omitempty"`
	Response *BundleEntryResponse `json:"response,omitempty"`
}


