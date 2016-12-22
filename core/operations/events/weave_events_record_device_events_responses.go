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
 package events




import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*WeaveEventsRecordDeviceEventsOK Successful response

swagger:response weaveEventsRecordDeviceEventsOK
*/
type WeaveEventsRecordDeviceEventsOK struct {
}

// NewWeaveEventsRecordDeviceEventsOK creates WeaveEventsRecordDeviceEventsOK with default headers values
func NewWeaveEventsRecordDeviceEventsOK() *WeaveEventsRecordDeviceEventsOK {
	return &WeaveEventsRecordDeviceEventsOK{}
}

// WriteResponse to the client
func (o *WeaveEventsRecordDeviceEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}
