// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

import ONLINE_STATUS from '@ttn-lw/constants/online-status'

export const selectStatusStore = state => state.status

export const selectOnlineStatus = state => selectStatusStore(state).onlineStatus

export const selectIsOnlineStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.ONLINE

export const selectIsOfflineStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.OFFLINE

export const selectIsCheckingStatus = state =>
  selectStatusStore(state).onlineStatus === ONLINE_STATUS.CHECKING

export const selectSummaryStore = state => selectStatusStore(state).summary
export const selectUpcomingMaintenancesStore = state =>
  selectSummaryStore(state).scheduled_maintenances ?? []
export const selectNetworkStatusIndicatorStore = state => selectSummaryStore(state).status.indicator
