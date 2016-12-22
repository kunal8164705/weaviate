/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package devices


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveDevicesUpdateParentHandlerFunc turns a function with the right signature into a weave devices update parent handler
type WeaveDevicesUpdateParentHandlerFunc func(WeaveDevicesUpdateParentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesUpdateParentHandlerFunc) Handle(params WeaveDevicesUpdateParentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesUpdateParentHandler interface for that can handle valid weave devices update parent params
type WeaveDevicesUpdateParentHandler interface {
	Handle(WeaveDevicesUpdateParentParams, interface{}) middleware.Responder
}

// NewWeaveDevicesUpdateParent creates a new http.Handler for the weave devices update parent operation
func NewWeaveDevicesUpdateParent(ctx *middleware.Context, handler WeaveDevicesUpdateParentHandler) *WeaveDevicesUpdateParent {
	return &WeaveDevicesUpdateParent{Context: ctx, Handler: handler}
}

/*WeaveDevicesUpdateParent swagger:route POST /devices/{deviceId}/updateParent devices weaveDevicesUpdateParent

Updates parent of the child device. Only managers can use this method.

*/
type WeaveDevicesUpdateParent struct {
	Context *middleware.Context
	Handler WeaveDevicesUpdateParentHandler
}

func (o *WeaveDevicesUpdateParent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesUpdateParentParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
