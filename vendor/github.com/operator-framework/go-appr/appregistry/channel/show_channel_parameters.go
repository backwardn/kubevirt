// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowChannelParams creates a new ShowChannelParams object
// with the default values initialized.
func NewShowChannelParams() *ShowChannelParams {
	var ()
	return &ShowChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowChannelParamsWithTimeout creates a new ShowChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowChannelParamsWithTimeout(timeout time.Duration) *ShowChannelParams {
	var ()
	return &ShowChannelParams{

		timeout: timeout,
	}
}

// NewShowChannelParamsWithContext creates a new ShowChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowChannelParamsWithContext(ctx context.Context) *ShowChannelParams {
	var ()
	return &ShowChannelParams{

		Context: ctx,
	}
}

// NewShowChannelParamsWithHTTPClient creates a new ShowChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowChannelParamsWithHTTPClient(client *http.Client) *ShowChannelParams {
	var ()
	return &ShowChannelParams{
		HTTPClient: client,
	}
}

/*ShowChannelParams contains all the parameters to send to the API endpoint
for the show channel operation typically these are written to a http.Request
*/
type ShowChannelParams struct {

	/*Channel
	  channel name

	*/
	Channel string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show channel params
func (o *ShowChannelParams) WithTimeout(timeout time.Duration) *ShowChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show channel params
func (o *ShowChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show channel params
func (o *ShowChannelParams) WithContext(ctx context.Context) *ShowChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show channel params
func (o *ShowChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show channel params
func (o *ShowChannelParams) WithHTTPClient(client *http.Client) *ShowChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show channel params
func (o *ShowChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannel adds the channel to the show channel params
func (o *ShowChannelParams) WithChannel(channel string) *ShowChannelParams {
	o.SetChannel(channel)
	return o
}

// SetChannel adds the channel to the show channel params
func (o *ShowChannelParams) SetChannel(channel string) {
	o.Channel = channel
}

// WithNamespace adds the namespace to the show channel params
func (o *ShowChannelParams) WithNamespace(namespace string) *ShowChannelParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the show channel params
func (o *ShowChannelParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the show channel params
func (o *ShowChannelParams) WithPackage(packageVar string) *ShowChannelParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the show channel params
func (o *ShowChannelParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WriteToRequest writes these params to a swagger request
func (o *ShowChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel
	if err := r.SetPathParam("channel", o.Channel); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}