// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

import { mergeWith } from 'lodash'

import { EVENT_END_DEVICE_HEARTBEAT_FILTERS_REGEXP } from '@console/constants/event-filters'

import { getCombinedDeviceId, combineDeviceIds } from '@ttn-lw/lib/selectors/id'
import getByPath from '@ttn-lw/lib/get-by-path'

import { parseLorawanMacVersion } from '@console/lib/device-utils'

import { GET_INFO_BY_JOIN_EUI_SUCCESS } from '@console/store/actions/claim'
import {
  GET_DEV,
  GET_DEVICES_LIST_SUCCESS,
  GET_DEV_SUCCESS,
  UPDATE_DEV_SUCCESS,
  RESET_DEV_SUCCESS,
  GET_DEVICE_EVENT_MESSAGE_SUCCESS,
  DELETE_DEV_SUCCESS,
  FETCH_DEVICES_LIST_SUCCESS,
} from '@console/store/actions/devices'
import { GET_APP_EVENT_MESSAGE_SUCCESS } from '@console/store/actions/applications'

const defaultState = {
  entities: {},
  derived: {},
  selectedDevice: undefined,
  selectedDeviceClaimable: undefined,
}
const defaultDerived = {
  uplinkFrameCount: undefined,
  downlinkAppFrameCount: undefined,
  downlinkNwkFrameCount: undefined,
}

const heartbeatFilterRegExp = new RegExp(EVENT_END_DEVICE_HEARTBEAT_FILTERS_REGEXP)
const uplinkFrameCountEvent = 'ns.up.data.process'
const downlinkFrameCountEvent = 'ns.down.data.schedule.attempt'

const mergeDerived = (state, id, derived) =>
  Object.keys(derived).length > 0
    ? {
        ...state,
        derived: {
          ...state.derived,
          [id]: {
            ...state.derived[id],
            ...derived,
          },
        },
      }
    : state

const devices = (state = defaultState, { type, payload, event, meta }) => {
  switch (type) {
    case GET_INFO_BY_JOIN_EUI_SUCCESS:
      return {
        ...state,
        selectedDeviceClaimable: payload.supports_claiming ?? false,
      }
    case GET_DEV:
      if (meta.options.noSelect) {
        return state
      }
      return {
        ...state,
        selectedDevice: combineDeviceIds(payload.appId, payload.deviceId),
      }
    case DELETE_DEV_SUCCESS: {
      const id = getCombinedDeviceId(payload)
      const { [id]: deleted, ...entities } = state.entities
      const { [id]: deleted2, ...derived } = state.derived

      return {
        ...defaultState,
        entities,
        derived,
      }
    }
    case UPDATE_DEV_SUCCESS:
    case GET_DEV_SUCCESS:
      const id = getCombinedDeviceId(payload)
      const lorawanVersion = payload.lorawan_version

      const mergedDevice = mergeWith({}, state.entities[id], payload, (_, __, key, ___, source) => {
        // Always set location from the payload.
        if (source === payload && key === 'locations') {
          return source.locations
        }

        if (source === payload && key === 'attributes') {
          return source.attributes
        }
      })

      const updatedEntities = {
        ...state.entities,
        [id]: mergedDevice,
      }

      // Update uplink and downlink frame counts if possible.
      const derived = {}
      const { session } = payload
      if (session) {
        derived.uplinkFrameCount = session.last_f_cnt_up
        if (parseLorawanMacVersion(lorawanVersion) < 110) {
          derived.downlinkNwkFrameCount = session.last_n_f_cnt_down
        } else {
          // For 1.1+ end devices there are two frame counters.
          derived.downlinkAppFrameCount = session.last_a_f_cnt_down
          derived.downlinkNwkFrameCount = session.last_n_f_cnt_down
        }
      }

      return mergeDerived({ ...state, entities: updatedEntities }, id, derived)
    case RESET_DEV_SUCCESS:
      const combinedId = getCombinedDeviceId(payload)
      const device = state.entities[combinedId]
      const isOTAA = Boolean(device.supports_join)

      if (isOTAA && device.session && device.session.keys) {
        const { session } = device
        // Reset NS session keys and last seen information for joined OTAA end devices.
        const resetDevice = {
          ...device,
          session: {
            ...session,
            dev_addr: session.dev_addr,
            keys: { app_s_key: session.keys.app_s_key },
          },
        }

        return mergeDerived(
          {
            ...state,
            entities: {
              ...state.entities,
              [combinedId]: resetDevice,
            },
          },
          combinedId,
          defaultDerived,
        )
      }

      return mergeDerived(state, combinedId, defaultDerived)
    case FETCH_DEVICES_LIST_SUCCESS:
    case GET_DEVICES_LIST_SUCCESS:
      const newEntities = payload.entities.reduce(
        (acc, dev) => {
          const id = getCombinedDeviceId(dev)
          acc[id] = { ...state.entities[id], ...dev }
          return acc
        },
        { ...state.entities },
      )

      return { ...state, entities: newEntities }
    case GET_DEVICE_EVENT_MESSAGE_SUCCESS:
    case GET_APP_EVENT_MESSAGE_SUCCESS:
      let newState = state
      // Detect heartbeat events to update last seen state.
      if (heartbeatFilterRegExp.test(event.name)) {
        const id = getCombinedDeviceId(event.identifiers[0].device_ids)
        const receivedAt = getByPath(event, 'data.received_at')
        if (receivedAt) {
          const currentEndDevice = { ...state.entities[id] }
          if (currentEndDevice) {
            // Only update if the event was actually more recent than the current value.
            if (
              (currentEndDevice.last_seen_at && currentEndDevice.last_seen_at < receivedAt) ||
              !currentEndDevice.last_seen_at
            ) {
              const updatedEndDevice = {
                ...currentEndDevice,
                last_seen_at: receivedAt,
              }

              newState = {
                ...state,
                entities: {
                  ...state.entities,
                  [id]: updatedEndDevice,
                },
              }
            }
          }
        }
      }
      // Detect uplink/downlink process events to update uplink/downlink frame count state.
      if (event.name === uplinkFrameCountEvent) {
        const uplinkFrameCount = getByPath(event, 'data.payload.mac_payload.full_f_cnt')
        if (typeof uplinkFrameCount === 'number') {
          return mergeDerived(newState, getCombinedDeviceId(event.identifiers[0].device_ids), {
            uplinkFrameCount,
          })
        }
      } else if (event.name === downlinkFrameCountEvent) {
        const combinedDeviceId = getCombinedDeviceId(event.identifiers[0].device_ids)
        const lorawanVersion = getByPath(state.entities, `${combinedDeviceId}.lorawan_version`)

        if (parseLorawanMacVersion(lorawanVersion) < 110) {
          if (getByPath(event, 'data.payload.mac_payload.f_port') === undefined) {
            const downlinkNwkFrameCount = getByPath(event, 'data.payload.mac_payload.full_f_cnt')
            if (typeof downlinkNwkFrameCount === 'number') {
              return mergeDerived(newState, combinedDeviceId, {
                downlinkNwkFrameCount,
              })
            }
          } else if (getByPath(event, 'data.payload.mac_payload.f_port') > 0) {
            const downlinkAppFrameCount = getByPath(event, 'data.payload.mac_payload.full_f_cnt')
            if (typeof downlinkAppFrameCount === 'number') {
              return mergeDerived(newState, combinedDeviceId, {
                downlinkAppFrameCount,
              })
            }
          }
        }

        // For 1.1+ end devices there are two frame counters. If `f_port` is equal to 0 - then it's the "network" frame counter,
        // otherwise it's the "application" frame counter.
        if (getByPath(event, 'data.payload.mac_payload.f_port') === 0) {
          const downlinkFrameCount = getByPath(event, 'data.payload.mac_payload.f_cnt')
          if (typeof downlinkFrameCount === 'number') {
            return mergeDerived(newState, combinedDeviceId, {
              downlinkNwkFrameCount: downlinkFrameCount,
            })
          }
        } else {
          const downlinkFrameCount = getByPath(event, 'data.payload.mac_payload.full_f_cnt')
          if (typeof downlinkFrameCount === 'number') {
            return mergeDerived(newState, combinedDeviceId, {
              downlinkAppFrameCount: downlinkFrameCount,
            })
          }
        }
      }

      return newState
    default:
      return state
  }
}

export default devices
