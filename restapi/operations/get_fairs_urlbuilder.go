// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetFairsURL generates an URL for the get fairs operation
type GetFairsURL struct {
	DistrictName *string
	FairName     *string
	Neighborhood *string
	Region5      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFairsURL) WithBasePath(bp string) *GetFairsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFairsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetFairsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/fairs"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var districtNameQ string
	if o.DistrictName != nil {
		districtNameQ = *o.DistrictName
	}
	if districtNameQ != "" {
		qs.Set("districtName", districtNameQ)
	}

	var fairNameQ string
	if o.FairName != nil {
		fairNameQ = *o.FairName
	}
	if fairNameQ != "" {
		qs.Set("fairName", fairNameQ)
	}

	var neighborhoodQ string
	if o.Neighborhood != nil {
		neighborhoodQ = *o.Neighborhood
	}
	if neighborhoodQ != "" {
		qs.Set("neighborhood", neighborhoodQ)
	}

	var region5Q string
	if o.Region5 != nil {
		region5Q = *o.Region5
	}
	if region5Q != "" {
		qs.Set("region5", region5Q)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetFairsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetFairsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetFairsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetFairsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetFairsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetFairsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
