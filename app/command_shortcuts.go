// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"bytes"
	"strings"

	"github.com/mattermost/platform/model"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
)

type ShortcutsProvider struct {
}

const (
	CMD_SHORTCUTS = "shortcuts"
)

func init() {
	RegisterCommandProvider(&ShortcutsProvider{})
}

func (me *ShortcutsProvider) GetTrigger() string {
	return CMD_SHORTCUTS
}

func (me *ShortcutsProvider) GetCommand(T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_SHORTCUTS,
		AutoComplete:     true,
		AutoCompleteDesc: T("i18n.server.api.command_shortcuts.desc"),
		AutoCompleteHint: "",
		DisplayName:      T("i18n.server.api.command_shortcuts.name"),
	}
}

func (me *ShortcutsProvider) DoCommand(args *model.CommandArgs, message string) *model.CommandResponse {
	shortcutIds := [28]string{
		"i18n.server.api.command_shortcuts.header",
		// Nav shortcuts
		"i18n.server.api.command_shortcuts.nav.header",
		"i18n.server.api.command_shortcuts.nav.prev",
		"i18n.server.api.command_shortcuts.nav.next",
		"i18n.server.api.command_shortcuts.nav.unread_prev",
		"i18n.server.api.command_shortcuts.nav.unread_next",
		"i18n.server.api.command_shortcuts.nav.switcher",
		"i18n.server.api.command_shortcuts.nav.settings",
		"i18n.server.api.command_shortcuts.nav.recent_mentions",
		// Files shortcuts
		"i18n.server.api.command_shortcuts.files.header",
		"i18n.server.api.command_shortcuts.files.upload",
		// Msg shortcuts
		"i18n.server.api.command_shortcuts.msgs.header",
		"i18n.server.api.command_shortcuts.msgs.mark_as_read",
		"i18n.server.api.command_shortcuts.msgs.reprint_prev",
		"i18n.server.api.command_shortcuts.msgs.reprint_next",
		"i18n.server.api.command_shortcuts.msgs.edit",
		"i18n.server.api.command_shortcuts.msgs.comp_username",
		"i18n.server.api.command_shortcuts.msgs.comp_channel",
		"i18n.server.api.command_shortcuts.msgs.comp_emoji",
		// Browser shortcuts
		"i18n.server.api.command_shortcuts.browser.header",
		"i18n.server.api.command_shortcuts.browser.channel_prev",
		"i18n.server.api.command_shortcuts.browser.channel_next",
		"i18n.server.api.command_shortcuts.browser.font_increase",
		"i18n.server.api.command_shortcuts.browser.font_decrease",
		"i18n.server.api.command_shortcuts.browser.highlight_prev",
		"i18n.server.api.command_shortcuts.browser.highlight_next",
		"i18n.server.api.command_shortcuts.browser.newline",
	}

	var osDependentWords map[string]interface{}
	if strings.Contains(message, "mac") {
		osDependentWords = map[string]interface{}{
			"CmdOrCtrl":      args.T("i18n.server.api.command_shortcuts.cmd"),
			"ChannelPrevCmd": args.T("i18n.server.api.command_shortcuts.browser.channel_prev.cmd_mac"),
			"ChannelNextCmd": args.T("i18n.server.api.command_shortcuts.browser.channel_next.cmd_mac"),
		}
	} else {
		osDependentWords = map[string]interface{}{
			"CmdOrCtrl":      args.T("i18n.server.api.command_shortcuts.ctrl"),
			"ChannelPrevCmd": args.T("i18n.server.api.command_shortcuts.browser.channel_prev.cmd"),
			"ChannelNextCmd": args.T("i18n.server.api.command_shortcuts.browser.channel_next.cmd"),
		}
	}

	var buffer bytes.Buffer
	for _, element := range shortcutIds {
		buffer.WriteString(args.T(element, osDependentWords))
	}

	return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL, Text: buffer.String()}
}
