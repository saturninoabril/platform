// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package timezones

type SupportedTimeZone struct {
	Value  string `json:"value"`
	Abbr   string `json:"abbr"`
	Offset int    `json:"offset"`
	IsDST  bool   `json:"is_dst"`
	Text   string `json:"text"`
	UTC    string `json:"utc"`
}

type SupportedTimeZones struct {
	SupportedTimeZones []SupportedTimeZone `json:"supported_time_zones"`
}
