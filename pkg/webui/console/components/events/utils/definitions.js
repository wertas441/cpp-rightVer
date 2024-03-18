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

import DownlinkMessage from '../previews/downlink-message'
import GatewayUplinkMessage from '../previews/gateway-uplink-message'
import ApplicationDownlink from '../previews/application-downlink'
import ApplicationUplink from '../previews/application-uplink'
import ApplicationUplinkNormalized from '../previews/application-uplink-normalized'
import ApplicationUp from '../previews/application-up'
import ApplicationLocation from '../previews/application-location'
import JoinRequest from '../previews/join-request'
import JoinResponse from '../previews/join-response'
import GatewayStatus from '../previews/gateway-status'
import ErrorDetails from '../previews/error-details'
import Value from '../previews/value'

export const eventIconMap = [
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+create$/,
    icon: 'event-create',
  },
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+update$/,
    icon: 'event-update',
  },
  {
    test: /^(ns\.|as\.|js\.)?([a-z0-9](?:[-_]?[a-z0-9]){2,}\.)+delete$/,
    icon: 'event-delete',
  },
  {
    test: /^(ns|as)\.up(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-uplink',
  },
  {
    test: /^(ns|as)\.down(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-downlink',
  },
  {
    test: /^(js|ns|as)(\.up|\.down)?\.(join|rejoin)(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-join',
  },
  {
    test: /^gs\.up(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-uplink',
  },
  {
    test: /^gs\.down(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-downlink',
  },
  {
    test: /^gs.gateway.connect$/,
    icon: 'event-gateway-connect',
  },
  {
    test: /^gs.gateway.disconnect$/,
    icon: 'event-gateway-disconnect',
  },
  {
    test: /^ns\.mac\.rekey\..*$/,
    icon: 'event-rekey',
  },
  {
    test: /^ns\.mac\.device_mode\..*$/,
    icon: 'event-mode',
  },
  {
    test: /^ns\..*\.switch\..*$/,
    icon: 'event-switch',
  },
  {
    test: /^ns(\.[a-z0-9](?:[-_]?[a-z0-9]){2,})+$/,
    icon: 'event-connection',
  },
  {
    test: /^gs\.status\..*$/,
    icon: 'event-status',
  },
  {
    test: /^synthetic\.status\.cleared$/,
    icon: 'event-clear-all',
  },
  {
    test: /^synthetic\.error\..*$/,
    icon: 'event-error',
  },
]

export const dataTypeMap = {
  ApplicationDownlink,
  ApplicationUplink,
  ApplicationUplinkNormalized,
  ApplicationUp,
  ApplicationLocation,
  DownlinkMessage,
  GatewayUplinkMessage,
  JoinRequest,
  JoinResponse,
  ErrorDetails,
  GatewayStatus,
  Value,
}

export const applicationUpMessages = [
  'uplink_message',
  'uplink_normalized',
  'join_accept',
  'downlink_ack',
  'downlink_nack',
  'downlink_sent',
  'downlink_failed',
  'downlink_queued',
  'downlink_queue_invalidated',
  'location_solved',
  'service_data',
]
