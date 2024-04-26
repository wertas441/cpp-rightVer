// Copyright © 2024 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { defineMessages } from 'react-intl'
import { useSelector } from 'react-redux'

import Icon from '@ttn-lw/components/icon'
import Status from '@ttn-lw/components/status'

import Message from '@ttn-lw/lib/components/message'

import LastSeen from '@console/components/last-seen'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import {
  selectBookmarksList,
  selectBookmarksTotalCount,
} from '@console/store/selectors/user-preferences'

import EntitiesList from '../list'

const m = defineMessages({
  noTopEntities: 'No top entities yet',
  noTopEntitiesDescription: 'Your most visited, and bookmarked entities will be listed here.',
})

const AllTopEntitiesList = () => {
  const allBookmarks = useSelector(state => selectBookmarksList(state))

  const headers = [
    {
      name: 'type',
      displayName: sharedMessages.type,
      width: '35px',
      render: icon => <Icon icon={icon} />,
    },
    {
      name: 'name',
      displayName: sharedMessages.name,
      align: 'left',
      render: (name, id) => <Message content={name === '' ? id : name} />,
    },
    {
      name: 'lastSeen',
      displayName: sharedMessages.lastSeen,
      render: lastSeen => {
        if (!lastSeen) {
          return (
            <Status
              status="mediocre"
              label={sharedMessages.noRecentActivity}
              className="d-flex j-end al-center"
            />
          )
        }
        if (lastSeen && typeof lastSeen === 'string') {
          return <LastSeen lastSeen={lastSeen} short statusClassName="j-end" />
        }

        return lastSeen.isDisconnected ? (
          <LastSeen
            status="bad"
            message={sharedMessages.disconnected}
            lastSeen={lastSeen.disconnectedAt}
            statusClassName="j-end"
            short
          />
        ) : Boolean(lastSeen.gatewayLastSeen) ? (
          <LastSeen lastSeen={lastSeen.gatewayLastSeen} statusClassName="j-end" short />
        ) : (
          <Status
            status="mediocre"
            label={sharedMessages.noRecentActivity}
            className="d-flex j-end al-center"
          />
        )
      },
    },
  ]

  return (
    <EntitiesList
      allBookmarks={allBookmarks}
      itemsCountSelector={selectBookmarksTotalCount}
      headers={headers}
      emptyMessage={m.noTopEntities}
      emptyDescription={m.noTopEntitiesDescription}
    />
  )
}

export default AllTopEntitiesList
