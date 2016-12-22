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

package models

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DevicesPatchStateRequest devices patch state request
// swagger:model DevicesPatchStateRequest
type DevicesPatchStateRequest struct {

	// The list of state patches with corresponding timestamps.
	Patches []*DeviceStatePatchesStatePatch `json:"patches"`

	// Timestamp of a request. Local time, UNIX timestamp or time since last boot can be used.
	RequestTimeMs int64 `json:"requestTimeMs,omitempty"`
}

// Validate validates this devices patch state request
func (m *DevicesPatchStateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatches(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevicesPatchStateRequest) validatePatches(formats strfmt.Registry) error {

	if swag.IsZero(m.Patches) { // not required
		return nil
	}

	for i := 0; i < len(m.Patches); i++ {

		if swag.IsZero(m.Patches[i]) { // not required
			continue
		}

		if m.Patches[i] != nil {

			if err := m.Patches[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
