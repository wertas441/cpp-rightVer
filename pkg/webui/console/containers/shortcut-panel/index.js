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

import Panel from '../../../components/panel'

import ShortcutItem from './shortcut-item'

const m = defineMessages({
  shortcuts: 'Quick actions',
  addApplication: 'New Application',
  addGateway: 'New Gateway',
  addOrganization: 'New Organization',
  addPersonalApiKey: 'New personal API key',
  registerDevice: 'Register a device',
})

const ShortcutPanel = () => (
  <Panel title={m.shortcuts} path="/edit-shortcuts" icon="bolt" divider className="h-full">
    <div className="grid">
      <ShortcutItem
        icon="application"
        title={m.addApplication}
        link="/applications/add"
        className="item-6"
      />
      <ShortcutItem icon="gateway" title={m.addGateway} link="/gateways/add" className="item-6" />
      <ShortcutItem
        icon="users-group"
        title={m.addOrganization}
        link="/organizations/add"
        className="item-4"
      />
      <ShortcutItem
        icon="key"
        title={m.addPersonalApiKey}
        link="/user/api-keys/add"
        className="item-4"
      />
      <ShortcutItem
        icon="device"
        title={m.registerDevice}
        link="/applications"
        className="item-4"
      />
    </div>
  </Panel>
)

export default ShortcutPanel
