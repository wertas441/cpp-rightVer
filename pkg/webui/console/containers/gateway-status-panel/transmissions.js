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
import { FormattedNumber, defineMessages } from 'react-intl'

import Icon, {
  IconDownlink,
  IconUplink,
  IconActivity,
  IconX,
  IconClockCheck,
  IconClockCancel,
  IconProtocol,
} from '@ttn-lw/components/icon'

import DateTime from '@ttn-lw/lib/components/date-time'
import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import style from './gateway-status-panel.styl'

const m = defineMessages({
  noUplinks: 'No uplinks yet',
  noDownlinks: 'No downlinks yet',
  noStatus: 'No status',
  noDataYet: 'No data yet',
  established: 'Connection started {days}',
  notEstablished: 'Connection not started',
})

const Transmissions = ({ gatewayStats, isDisconnected }) => {
  const showUplink =
    Boolean(gatewayStats?.uplink_count) ||
    (typeof gatewayStats?.uplink_count === 'string' && gatewayStats?.uplink_count !== 0)
  const showUplinkTime = Boolean(gatewayStats?.last_uplink_received_at)
  const showDownlink =
    Boolean(gatewayStats?.downlink_count) ||
    (typeof gatewayStats?.downlink_count === 'string' && gatewayStats?.downlink_count !== 0)
  const showDownlinkTime = Boolean(gatewayStats?.last_downlink_received_at)
  const showStatus = Boolean(gatewayStats?.last_status_received_at)
  const showConnectionEstablished = Boolean(gatewayStats?.connected_at)
  const showProtocol = Boolean(gatewayStats?.protocol)

  return (
    <div>
      <div className={style.gtwStatusPanelTransmissions}>
        <div className="d-flex al-center gap-cs-xxs">
          <Icon icon={IconUplink} />
          {!showUplink ? (
            <Message content={m.noUplinks} className="fw-bold" />
          ) : (
            <FormattedNumber value={gatewayStats.uplink_count}>
              {parts => <b>{parts}</b>}
            </FormattedNumber>
          )}
        </div>
        {showUplinkTime && (
          <DateTime.Relative
            value={gatewayStats.last_uplink_received_at}
            relativeTimeStyle="short"
            className="sm-md:d-none"
          />
        )}
      </div>
      <div className={style.gtwStatusPanelTransmissions}>
        <div className="d-flex al-center gap-cs-xxs">
          <Icon icon={IconDownlink} />
          {!showDownlink ? (
            <Message content={m.noDownlinks} className="fw-bold" />
          ) : (
            <>
              <FormattedNumber value={gatewayStats.downlink_count}>
                {parts => <b>{parts}</b>}
              </FormattedNumber>
              {gatewayStats?.tx_acknowledgment_count && (
                <>
                  {'('}
                  <FormattedNumber value={gatewayStats.tx_acknowledgment_count} />
                  {' Ack’d)'}
                </>
              )}
            </>
          )}
        </div>
        {showDownlinkTime && (
          <DateTime.Relative
            value={gatewayStats.last_downlink_received_at}
            relativeTimeStyle="short"
            className="sm-md:d-none"
          />
        )}
      </div>
      <div className={style.gtwStatusPanelTransmissions}>
        <div className="d-flex al-center gap-cs-xxs fw-bold">
          {isDisconnected ? (
            <>
              <Icon icon={IconX} className="c-text-error-normal" />
              <Message content={sharedMessages.disconnected} className="c-text-error-normal" />
            </>
          ) : showStatus ? (
            <>
              <Icon icon={IconActivity} />
              <Message content={sharedMessages.received} />
            </>
          ) : (
            <>
              <Icon icon={IconX} className="c-text-error-normal" />
              <Message content={m.noStatus} className="c-text-error-normal" />
            </>
          )}
        </div>
        {showStatus && !isDisconnected && (
          <DateTime.Relative
            value={gatewayStats.last_status_received_at}
            relativeTimeStyle="short"
            className="fw-bold"
          />
        )}
      </div>
      <div className={style.gtwStatusPanelTransmissions}>
        <div className="d-flex al-center gap-cs-xxs">
          <Icon icon={showConnectionEstablished ? IconClockCheck : IconClockCancel} />
          {showConnectionEstablished ? (
            <Message
              content={m.established}
              className="fw-bold"
              values={{ days: <DateTime.Relative value={gatewayStats?.connected_at} /> }}
            />
          ) : (
            <Message content={m.notEstablished} className="fw-bold" />
          )}
        </div>
      </div>
      {showProtocol && (
        <div className={style.gtwStatusPanelTransmissions}>
          <div className="d-flex al-center gap-cs-xxs fw-bold">
            <Icon icon={IconProtocol} />
            <span>{gatewayStats.protocol.toUpperCase()}</span>
          </div>
        </div>
      )}
    </div>
  )
}

Transmissions.propTypes = {
  gatewayStats: PropTypes.shape({
    downlink_count: PropTypes.string,
    last_downlink_received_at: PropTypes.string,
    last_status_received_at: PropTypes.string,
    last_uplink_received_at: PropTypes.string,
    tx_acknowledgment_count: PropTypes.string,
    uplink_count: PropTypes.string,
    connected_at: PropTypes.string,
    protocol: PropTypes.string,
  }),
  isDisconnected: PropTypes.bool.isRequired,
}

Transmissions.defaultProps = {
  gatewayStats: undefined,
}

export default Transmissions
