/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateThingsReferencesCreateParams creates a new WeaviateThingsReferencesCreateParams object
// with the default values initialized.
func NewWeaviateThingsReferencesCreateParams() *WeaviateThingsReferencesCreateParams {
	var ()
	return &WeaviateThingsReferencesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsReferencesCreateParamsWithTimeout creates a new WeaviateThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsReferencesCreateParamsWithTimeout(timeout time.Duration) *WeaviateThingsReferencesCreateParams {
	var ()
	return &WeaviateThingsReferencesCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsReferencesCreateParamsWithContext creates a new WeaviateThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsReferencesCreateParamsWithContext(ctx context.Context) *WeaviateThingsReferencesCreateParams {
	var ()
	return &WeaviateThingsReferencesCreateParams{

		Context: ctx,
	}
}

// NewWeaviateThingsReferencesCreateParamsWithHTTPClient creates a new WeaviateThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsReferencesCreateParamsWithHTTPClient(client *http.Client) *WeaviateThingsReferencesCreateParams {
	var ()
	return &WeaviateThingsReferencesCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsReferencesCreateParams contains all the parameters to send to the API endpoint
for the weaviate things references create operation typically these are written to a http.Request
*/
type WeaviateThingsReferencesCreateParams struct {

	/*Body*/
	Body *models.SingleRef
	/*PropertyName
	  Unique name of the property related to the Thing.

	*/
	PropertyName string
	/*ThingID
	  Unique ID of the Thing.

	*/
	ThingID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithTimeout(timeout time.Duration) *WeaviateThingsReferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithContext(ctx context.Context) *WeaviateThingsReferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithHTTPClient(client *http.Client) *WeaviateThingsReferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithBody(body *models.SingleRef) *WeaviateThingsReferencesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithPropertyName adds the propertyName to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithPropertyName(propertyName string) *WeaviateThingsReferencesCreateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithThingID adds the thingID to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) WithThingID(thingID strfmt.UUID) *WeaviateThingsReferencesCreateParams {
	o.SetThingID(thingID)
	return o
}

// SetThingID adds the thingId to the weaviate things references create params
func (o *WeaviateThingsReferencesCreateParams) SetThingID(thingID strfmt.UUID) {
	o.ThingID = thingID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsReferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	// path param thingId
	if err := r.SetPathParam("thingId", o.ThingID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}