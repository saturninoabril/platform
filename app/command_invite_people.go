// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"strings"

	l4g "github.com/alecthomas/log4go"
	"github.com/mattermost/platform/model"
	"github.com/mattermost/platform/utils"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
)

type InvitePeopleProvider struct {
}

const (
	CMD_INVITE_PEOPLE = "invite_people"
)

func init() {
	RegisterCommandProvider(&InvitePeopleProvider{})
}

func (me *InvitePeopleProvider) GetTrigger() string {
	return CMD_INVITE_PEOPLE
}

func (me *InvitePeopleProvider) GetCommand(T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_INVITE_PEOPLE,
		AutoComplete:     true,
		AutoCompleteDesc: T("i18n.server.api.command.invite_people.desc"),
		AutoCompleteHint: T("i18n.server.api.command.invite_people.hint"),
		DisplayName:      T("i18n.server.api.command.invite_people.name"),
	}
}

func (me *InvitePeopleProvider) DoCommand(args *model.CommandArgs, message string) *model.CommandResponse {
	if !utils.Cfg.EmailSettings.SendEmailNotifications {
		return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL, Text: args.T("i18n.server.api.command.invite_people.email_off")}
	}

	emailList := strings.Fields(message)

	for i := len(emailList) - 1; i >= 0; i-- {
		emailList[i] = strings.Trim(emailList[i], ",")
		if !strings.Contains(emailList[i], "@") {
			emailList = append(emailList[:i], emailList[i+1:]...)
		}
	}

	if len(emailList) == 0 {
		return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL, Text: args.T("i18n.server.api.command.invite_people.no_email")}
	}

	if err := InviteNewUsersToTeam(emailList, args.TeamId, args.UserId); err != nil {
		l4g.Error(err.Error())
		return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL, Text: args.T("i18n.server.api.command.invite_people.fail")}
	}

	return &model.CommandResponse{ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL, Text: args.T("i18n.server.api.command.invite_people.sent")}
}
