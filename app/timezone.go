// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

type Timezone struct {
	Value  string `json:"value"`
	Abbr   string `json:"abbr"`
	Offset int    `json:"offset"`
	IsDST  bool   `json:"is_dst"`
	Text   string `json:"text"`
	UTC    string `json:"utc"`
}

type Timezones struct {
	supportedZones []Timezone
}

// func New() *Timezones {
// 	timezones := Timezones{}

// 	supportedZones, err := GetSupportedTimezone()
// 	if err != nil {
// 		mlog.Warn("Failed to get supported timezones")
// 	}

// 	timezones.supportedZones = supportedZones

// 	return &timezones
// }

// func (t *Timezones) GetSupported() []Timezone {
// 	return t.supportedZones
// }

// // GetSupportedTimeZones returns all supported time zones
// func (a *App) GetSupportedTimezone() []Timezone {
// 	return DefaultSupportedTimezone
// }

// func (a *App) DefaultUserTimezone() map[string]string {
// 	defaultTimezone := make(map[string]string)
// 	defaultTimezone["useAutomaticTimezone"] = "true"
// 	defaultTimezone["automaticTimezone"] = ""
// 	defaultTimezone["manualTimezone"] = ""
// 	defaultTimezone["manualTimezoneLabel"] = ""

// 	return defaultTimezone
// }
