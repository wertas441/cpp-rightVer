// Copyright © 2023 The Things Network Foundation, The Things Industries B.V.
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

import { defineMessage } from 'react-intl'

import tts from '@console/api/tts'

import toast from '@ttn-lw/components/toast'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as notifications from '@console/store/actions/notifications'

import { selectUserId } from '@console/store/selectors/logout'

const m = defineMessage({
  newNotifications: 'You have new notifications',
})

const getInboxNotificationsLogic = createRequestLogic({
  type: notifications.GET_INBOX_NOTIFICATIONS,
  process: async ({ action, getState }) => {
    const {
      payload: { page, limit },
    } = action
    const filter = ['NOTIFICATION_STATUS_UNSEEN', 'NOTIFICATION_STATUS_SEEN']
    const userId = selectUserId(getState())
    const result = await tts.Notifications.getAllNotifications(userId, filter, page, limit)

    return { notifications: result.notifications, totalCount: result.totalCount, page, limit }
  },
})

const getArchivedNotificationsLogic = createRequestLogic({
  type: notifications.GET_ARCHIVED_NOTIFICATIONS,
  process: async ({ action, getState }) => {
    const {
      payload: { page, limit },
    } = action
    const filter = ['NOTIFICATION_STATUS_ARCHIVED']
    const userId = selectUserId(getState())
    const result = await tts.Notifications.getAllNotifications(userId, filter, page, limit)

    return { notifications: result.notifications, totalCount: result.totalCount, page, limit }
  },
})

const getUnseenNotificationsPeriodicallyLogic = createRequestLogic({
  type: notifications.GET_UNSEEN_NOTIFICATIONS,
  process: async ({ getState }) => {
    clearInterval()
    const id = selectUserId(getState())
    const result = await tts.Notifications.getAllNotifications(id, ['NOTIFICATION_STATUS_UNSEEN'])
    let totalCount = result.totalCount
    setInterval(async () => {
      const newResult = await tts.Notifications.getAllNotifications(id, [
        'NOTIFICATION_STATUS_UNSEEN',
      ])
      if (newResult.totalCount > totalCount) {
        toast({ message: m.newNotifications, type: toast.types.INFO })
      }
      totalCount = newResult.totalCount
    }, [1000 * 60 * 5]) // 5 minutes
    return { notifications: result.notifications, totalCount: result.totalCount }
  },
})

const updateNotificationStatusLogic = createRequestLogic({
  type: notifications.UPDATE_NOTIFICATION_STATUS,
  process: async ({ action, getState }) => {
    const {
      payload: { notificationIds, newStatus },
    } = action
    const id = selectUserId(getState())
    const result = await tts.Notifications.updateNotificationStatus(id, notificationIds, newStatus)

    return { ...result, ids: notificationIds }
  },
})

export default [
  getInboxNotificationsLogic,
  getArchivedNotificationsLogic,
  getUnseenNotificationsPeriodicallyLogic,
  updateNotificationStatusLogic,
]
