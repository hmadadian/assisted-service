// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDownloadInfraEnvDiscoveryImageHeadersParams creates a new DownloadInfraEnvDiscoveryImageHeadersParams object
// with the default values initialized.
func NewDownloadInfraEnvDiscoveryImageHeadersParams() *DownloadInfraEnvDiscoveryImageHeadersParams {
	var ()
	return &DownloadInfraEnvDiscoveryImageHeadersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadInfraEnvDiscoveryImageHeadersParamsWithTimeout creates a new DownloadInfraEnvDiscoveryImageHeadersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadInfraEnvDiscoveryImageHeadersParamsWithTimeout(timeout time.Duration) *DownloadInfraEnvDiscoveryImageHeadersParams {
	var ()
	return &DownloadInfraEnvDiscoveryImageHeadersParams{

		timeout: timeout,
	}
}

// NewDownloadInfraEnvDiscoveryImageHeadersParamsWithContext creates a new DownloadInfraEnvDiscoveryImageHeadersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadInfraEnvDiscoveryImageHeadersParamsWithContext(ctx context.Context) *DownloadInfraEnvDiscoveryImageHeadersParams {
	var ()
	return &DownloadInfraEnvDiscoveryImageHeadersParams{

		Context: ctx,
	}
}

// NewDownloadInfraEnvDiscoveryImageHeadersParamsWithHTTPClient creates a new DownloadInfraEnvDiscoveryImageHeadersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadInfraEnvDiscoveryImageHeadersParamsWithHTTPClient(client *http.Client) *DownloadInfraEnvDiscoveryImageHeadersParams {
	var ()
	return &DownloadInfraEnvDiscoveryImageHeadersParams{
		HTTPClient: client,
	}
}

/*DownloadInfraEnvDiscoveryImageHeadersParams contains all the parameters to send to the API endpoint
for the download infra env discovery image headers operation typically these are written to a http.Request
*/
type DownloadInfraEnvDiscoveryImageHeadersParams struct {

	/*InfraEnvID
	  The InfraEnv whose image headers should be retrieved.

	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) WithTimeout(timeout time.Duration) *DownloadInfraEnvDiscoveryImageHeadersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) WithContext(ctx context.Context) *DownloadInfraEnvDiscoveryImageHeadersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) WithHTTPClient(client *http.Client) *DownloadInfraEnvDiscoveryImageHeadersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraEnvID adds the infraEnvID to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) WithInfraEnvID(infraEnvID strfmt.UUID) *DownloadInfraEnvDiscoveryImageHeadersParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the download infra env discovery image headers params
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadInfraEnvDiscoveryImageHeadersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
