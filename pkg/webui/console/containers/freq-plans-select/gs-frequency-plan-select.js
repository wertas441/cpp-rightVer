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

import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'

import frequencyPlans from '@console/constants/frequency-plans'

import KeyValueMap from '@ttn-lw/components/key-value-map'
import Select from '@ttn-lw/components/select'
import Form, { useFormContext } from '@ttn-lw/components/form'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'

import { getGsFrequencyPlans } from '@console/store/actions/configuration'

import { selectGsFrequencyPlans } from '@console/store/selectors/configuration'

import { formatOptions, m } from './utils'

const GatewayFrequencyPlansSelect = () => {
  const { values } = useFormContext()
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getGsFrequencyPlans())
  }, [dispatch])

  const freqPlanOptions = [
    ...formatOptions(useSelector(selectGsFrequencyPlans)),
    { value: 'no-frequency-plan', label: m.none },
  ]

  return (
    <Form.Field
      name="frequency_plan_ids"
      title={sharedMessages.frequencyPlan}
      description={m.frequencyPlanDescription}
      valuePlaceholder={m.selectFrequencyPlan}
      component={KeyValueMap}
      tooltipId={tooltipIds.FREQUENCY_PLAN}
      additionalInputProps={{ options: freqPlanOptions }}
      inputElement={Select}
      warning={
        values.frequency_plan_id === frequencyPlans.EMPTY_FREQ_PLAN
          ? sharedMessages.frequencyPlanWarning
          : undefined
      }
      addMessage={m.addFrequencyPlan}
      removeMessage={m.removeFrequencyPlan}
      icon="remove"
      indexAsKey
      className="w-60"
      required
    />
  )
}

export default GatewayFrequencyPlansSelect
