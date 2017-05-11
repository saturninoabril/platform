// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import React, {PropTypes} from 'react';
import $ from 'jquery';
import {FormattedMessage} from 'react-intl';
import {browserHistory} from 'react-router/es6';

import TeamStore from 'stores/team_store.jsx';
import * as Utils from 'utils/utils.jsx';
import AppDispatcher from '../dispatcher/app_dispatcher.jsx';
import Constants from 'utils/constants.jsx';
const ActionTypes = Constants.ActionTypes;

export default function Jump(props) {
    function hideSidebar() {
        $('.sidebar--right').removeClass('move--left');
    }

    function shrinkSidebar() {
        setTimeout(() => {
            props.shrink();
        });
    }

    function handleClick() {
        if (Utils.isMobile()) {
            AppDispatcher.handleServerAction({
                type: ActionTypes.RECEIVED_SEARCH,
                results: null
            });

            AppDispatcher.handleServerAction({
                type: ActionTypes.RECEIVED_SEARCH_TERM,
                term: null,
                do_search: false,
                is_mention_search: false
            });

            AppDispatcher.handleServerAction({
                type: ActionTypes.RECEIVED_POST_SELECTED,
                postId: null
            });

            hideSidebar();
        }
        shrinkSidebar();
        browserHistory.push(TeamStore.getCurrentTeamRelativeUrl() + '/pl/' + props.postId);
    }

    return (
        <a
            onClick={handleClick}
            className='search-item__jump'
        >
            <FormattedMessage
                id='search_item.jump'
                defaultMessage='Jump'
            />
        </a>
    );
}

Jump.propTypes = {
    idPrefix: PropTypes.string.isRequired,
    idCount: PropTypes.number,
    postId: PropTypes.string.isRequired,
    shrink: PropTypes.func
};

Jump.defaultProps = {
    idCount: -1,
    postId: ''
};