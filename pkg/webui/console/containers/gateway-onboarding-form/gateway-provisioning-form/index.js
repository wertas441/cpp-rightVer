// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

import React, { useCallback, useEffect, useMemo, useState } from 'react'
import { ErrorMessage } from 'formik'
import { defineMessages } from 'react-intl'
import { useDispatch } from 'react-redux'

import Button from '@ttn-lw/components/button'
import Form, { useFormContext } from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'

import Message from '@ttn-lw/lib/components/message'

import OwnersSelect from '@console/containers/owners-select'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import useDebounce from '@ttn-lw/lib/hooks/use-debounce'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { getGatewayClaimInfoByEui } from '@console/store/actions/gateways'

import GatewayRegistrationFormSection from './gateway-registration-form-section'
import GatewayClaimFormSection from './gateway-claim-form-section'

const m = defineMessages({
  continue: 'To continue, please confirm the Gateway EUI so we can determine onboarding options',
  emtyEui: 'Continue without EUI',
  noEui: 'No gateway EUI',
})

// Save EUI in both fields.
const gatewayEuiEncoder = value => ({
  // Empty strings will be interpreted as zero value (00 fill) by the backend
  // For this reason, they need to be transformed to undefined instead,
  // so that the value will be properly stripped when sent to the backend.
  ids: { eui: value === '' ? undefined : value },
  authenticated_identifiers: {
    gateway_eui: value,
  },
})
const gatewayEuiDecoder = value => (value?.ids?.eui === undefined ? '' : value.ids.eui)
// The default merge of the value setter cannot be used here, since it would
// disregard `undefined` values.
const gatewayEuiValueSetter = ({ setValues }, { value }) =>
  setValues(values => ({
    ...values,
    ids: { ...values.ids, ...value.ids },
    authenticated_identifiers: {
      ...values.authenticated_identifiers,
      ...value.authenticated_identifiers,
    },
  }))

const GatewayProvisioningFormSection = () => {
  const [euiError, setEuiError] = useState(undefined)
  const dispatch = useDispatch()
  const {
    values: {
      _ownerId: ownerId,
      _inputMethod: inputMethod,
      _withQRdata: withQRdata,
      ids: { eui = '' },
    },
    initialValues,
    resetForm,
    setFieldValue,
    setFieldTouched,
    touched,
  } = useFormContext()
  const isEuiMac = useMemo(() => eui?.length === 12, [eui.length])
  const debouncedEui = useDebounce(eui, 3000)
  const isDebouncedEuiMac = useMemo(() => debouncedEui?.length === 12, [debouncedEui])
  const showMacConvert = isEuiMac && (isDebouncedEuiMac || touched.ids?.eui)

  useEffect(() => {
    if (showMacConvert) {
      setFieldTouched('ids.eui', true)
    } else if (isDebouncedEuiMac && !isEuiMac) {
      setFieldTouched('ids.eui', false)
    }
  }, [isDebouncedEuiMac, isEuiMac, setFieldTouched, showMacConvert])

  const hasEmptyEui = eui === ''
  const hasCompletedEui = eui.length === 16
  const hasInputMethod = Boolean(inputMethod)

  // Based on the EUI, we can determine whether the gateway is claimable
  // using the `getInfoByEUI` service. Based on the result, we can
  // toggle claim/manual registration section of the form.
  const handleGatewayEUI = useCallback(async () => {
    setEuiError(undefined)

    try {
      if (!hasEmptyEui) {
        const { supports_claiming, is_managed } = await dispatch(
          attachPromise(getGatewayClaimInfoByEui(eui)),
        )
        setFieldValue(
          '_inputMethod',
          is_managed ? 'managed' : supports_claiming ? 'claim' : 'register',
        )
      } else {
        // Gateways without Gateway EUI cannot be claimed.
        setFieldValue('_inputMethod', 'register')
      }
    } catch (error) {
      setEuiError(error)
    }
  }, [dispatch, eui, hasEmptyEui, setFieldValue])

  useEffect(() => {
    // Auto-confirm the join EUI when using QR code data.
    if (withQRdata) {
      handleGatewayEUI()
    }
  }, [withQRdata, handleGatewayEUI])

  const handleEuiReset = useCallback(async () => {
    setEuiError(undefined)
    resetForm({ values: { ...initialValues, _ownerId: ownerId } })
  }, [initialValues, ownerId, resetForm])

  const handleGatewayEUIKeydown = useCallback(
    event => {
      // Allow using "Enter"-key to confirm the EUI.
      if (hasCompletedEui) {
        if (event.keyCode === 13) {
          handleGatewayEUI()
        }
      }
    },
    [handleGatewayEUI, hasCompletedEui],
  )

  let statusElement = null
  if (euiError) {
    statusElement = <ErrorMessage error={euiError} />
  } else if (!hasInputMethod) {
    statusElement = <Message component="p" content={m.continue} />
  }
  const handleConvertMacToEui = useCallback(() => {
    const euiValue = `${eui.substring(0, 6)}FFFE${eui.substring(6)}`
    setFieldValue('authenticated_identifiers.gateway_eui', euiValue)
    setFieldValue('ids.eui', euiValue)
  }, [eui, setFieldValue])

  return (
    <>
      <OwnersSelect name="_ownerId" required />
      <Form.Field
        title={sharedMessages.gatewayEUI}
        name="ids.eui,authenticated_identifiers.gateway_eui"
        type="byte"
        min={8}
        max={8}
        placeholder={hasInputMethod && hasEmptyEui ? m.noEui : sharedMessages.gatewayEUI}
        component={Input}
        tooltipId={tooltipIds.GATEWAY_EUI}
        required={inputMethod !== 'register'}
        disabled={hasInputMethod || withQRdata}
        onKeyDown={handleGatewayEUIKeydown}
        encode={gatewayEuiEncoder}
        decode={gatewayEuiDecoder}
        valueSetter={gatewayEuiValueSetter}
        autoFocus
      >
        {showMacConvert ? (
          <Button
            className="ml-cs-xs"
            type="button"
            message={sharedMessages.convertMacToEui}
            onClick={handleConvertMacToEui}
            secondary
          />
        ) : hasInputMethod ? (
          <Button
            className="ml-cs-xs"
            type="button"
            message={sharedMessages.reset}
            onClick={handleEuiReset}
            secondary
            disabled={withQRdata}
          />
        ) : (
          <Button
            className="ml-cs-xs"
            type="button"
            message={hasEmptyEui ? m.emtyEui : sharedMessages.confirm}
            onClick={handleGatewayEUI}
            disabled={!hasCompletedEui && !hasEmptyEui}
            secondary
          />
        )}
      </Form.Field>
      {(inputMethod === 'claim' || inputMethod === 'managed') && <GatewayClaimFormSection />}
      {inputMethod === 'register' && <GatewayRegistrationFormSection />}
      {statusElement}
    </>
  )
}

export { GatewayProvisioningFormSection as default }
