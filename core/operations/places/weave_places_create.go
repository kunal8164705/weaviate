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
 package places


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeavePlacesCreateHandlerFunc turns a function with the right signature into a weave places create handler
type WeavePlacesCreateHandlerFunc func(WeavePlacesCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeavePlacesCreateHandlerFunc) Handle(params WeavePlacesCreateParams) middleware.Responder {
	return fn(params)
}

// WeavePlacesCreateHandler interface for that can handle valid weave places create params
type WeavePlacesCreateHandler interface {
	Handle(WeavePlacesCreateParams) middleware.Responder
}

// NewWeavePlacesCreate creates a new http.Handler for the weave places create operation
func NewWeavePlacesCreate(ctx *middleware.Context, handler WeavePlacesCreateHandler) *WeavePlacesCreate {
	return &WeavePlacesCreate{Context: ctx, Handler: handler}
}

/*WeavePlacesCreate swagger:route POST /places/create places weavePlacesCreate

Registers a new place.

*/
type WeavePlacesCreate struct {
	Context *middleware.Context
	Handler WeavePlacesCreateHandler
}

func (o *WeavePlacesCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeavePlacesCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
