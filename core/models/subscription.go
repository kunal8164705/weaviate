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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscription subscription
// swagger:model Subscription
type Subscription struct {

	// channel
	Channel *SubscriptionChannel `json:"channel,omitempty"`

	// Subscription target device IDs.
	DeviceID []string `json:"deviceId"`

	// Timestamp in milliseconds since epoch when the subscription expires and new notifications stop being sent.
	ExpirationTimeMs int64 `json:"expirationTimeMs,omitempty"`

	// Unique subscription ID.
	ID string `json:"id,omitempty"`

	// Identifies what kind of resource this is. Value: the fixed string "weave#subscription".
	Kind *string `json:"kind,omitempty"`

	// Subscription topics.
	//
	// Acceptable values are:
	// - "deviceInfo"
	// - "connectionStatus"
	// - "states"
	// - "myCommands"
	// - "deviceAccess"
	Topic []string `json:"topic"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {

		if err := m.Channel.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Subscription) validateDeviceID(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceID) { // not required
		return nil
	}

	return nil
}

var subscriptionTopicItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connectionStatus","deviceAccess","deviceInfo","myCommands","states"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTopicItemsEnum = append(subscriptionTopicItemsEnum, v)
	}
}

func (m *Subscription) validateTopicItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionTopicItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateTopic(formats strfmt.Registry) error {

	if swag.IsZero(m.Topic) { // not required
		return nil
	}

	for i := 0; i < len(m.Topic); i++ {

		// value enum
		if err := m.validateTopicItemsEnum("topic"+"."+strconv.Itoa(i), "body", m.Topic[i]); err != nil {
			return err
		}

	}

	return nil
}

// SubscriptionChannel Notification channel description.
// swagger:model SubscriptionChannel
type SubscriptionChannel struct {

	// gcm info
	GcmInfo *SubscriptionChannelGcmInfo `json:"gcmInfo,omitempty"`

	// Channel type supported by device.
	Type string `json:"type,omitempty"`
}

// Validate validates this subscription channel
func (m *SubscriptionChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcmInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionChannel) validateGcmInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.GcmInfo) { // not required
		return nil
	}

	if m.GcmInfo != nil {

		if err := m.GcmInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var subscriptionChannelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["gcm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionChannelTypeTypePropEnum = append(subscriptionChannelTypeTypePropEnum, v)
	}
}

const (
	subscriptionChannelTypeGcm string = "gcm"
)

// prop value enum
func (m *SubscriptionChannel) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionChannelTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SubscriptionChannel) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("channel"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// SubscriptionChannelGcmInfo GCM channel specific information.
// swagger:model SubscriptionChannelGcmInfo
type SubscriptionChannelGcmInfo struct {

	// GCM registration ID.
	RegistrationID string `json:"registrationId,omitempty"`
}

// Validate validates this subscription channel gcm info
func (m *SubscriptionChannelGcmInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
