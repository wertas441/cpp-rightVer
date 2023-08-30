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

import { defineMessages } from 'react-intl'

export default defineMessages({
  // Keep these sorted alphabetically.
  '16Bit': '16 bit',
  '32Bit': '32 bit',
  abp: 'Activation by personalization (ABP)',
  absenceBeCreative: 'You make it up, so be creative!',
  absenceContactManufacturer: 'Contact your device manufacturer or reseller.',
  absenceProvidedByManufacturer: 'This should be provided by the device manufacturer.',
  accessTokens: 'Access tokens',
  accountDeleteConfirmation: 'Are you sure you want to delete this account?',
  accountDeleted: 'Account deleted',
  accuracy: 'Accuracy',
  actions: 'Actions',
  activationMode: 'Activation mode',
  add: 'Add',
  addApiKey: 'Add API key',
  addAttributes: 'Add attributes',
  addCollaborator: 'Add collaborator',
  addDeviceBulk: 'End device bulk creation',
  addHeaderEntry: 'Add header entry',
  addOAuthClient: 'Add OAuth client',
  addPubsub: 'Add Pub/Sub',
  addWebhook: 'Add webhook',
  address: 'Address',
  addressPlaceholder: 'host',
  admin: 'Admin',
  adminContact: 'Administrative contact',
  adminDescription:
    'Admin status enables overarching rights such as managing other users or modifying entities regardless of collaboration status',
  adminPanel: 'Admin panel',
  all: 'All',
  allAdmin: 'All (Admin)',
  altitude: 'Altitude',
  altitudeDesc: 'The altitude in meters, where 0 means sea level',
  antennas: 'Antennas',
  apiKey: 'API key',
  apiKeyCounted: '{count, plural, one {API key} other {API keys}}',
  apiKeyId: '<b>API Key ID:</b> <code>{apiKeyId}</code>',
  apiKeyNamePlaceholder: 'My new API key',
  apiKeys: 'API keys',
  appData: 'Application data',
  appOverview: 'Application overview',
  appEUI: 'AppEUI',
  appEUIDescription:
    'The AppEUI uniquely identifies the owner of the end device. If no AppEUI is provided by the device manufacturer (usually for development), it can be filled with zeros.',
  appEUIJoinEUI: 'AppEUI/JoinEUI',
  appId: 'Application ID',
  appKey: 'AppKey',
  appKeyAbsence:
    'Contact the manufacturer or your reseller. If they cannot provide an AppKey, and your end device is programmable, it is okay to generate one.',
  appSKey: 'AppSKey',
  application: 'Application',
  applicationServerAddress: 'Application Server address',
  applications: 'Applications',
  approve: 'Approve',
  asServerID: 'Application Server ID',
  asServerIDDescription: 'The AS-ID of the Application Server to use',
  asServerKekLabel: 'Application Server KEK label',
  asServerKekLabelDescription:
    'The KEK label of the Application Server to use for wrapping the application session key',
  attributeDescription:
    'Attributes can be used to set arbitrary information about the entity, to be used by scripts, or simply for your own organization',
  attributeKeyValidateTooLong: 'Attribute keys must have less than 32 characters',
  attributeKeyValidateTooShort:
    'Attribute keys must have at least 3 characters and contain no special characters',
  attributeValueValidateTooLong: 'Attribute values must have less than 200 characters',
  attributes: 'Attributes',
  attributesValidateRequired:
    'All attribute entry values are required. Please remove empty entries.',
  attributesValidateTooMany: '{field} must be 10 items or fewer',
  authorization: 'Authorization',
  authorizationCode: 'Authorization code',
  authorizations: 'Authorizations',
  autoUpdateDescription: 'Gateway can be updated automatically',
  automaticUpdates: 'Automatic updates',
  backTo: 'Back to {siteTitle}',
  backToLogin: 'Back to login',
  backToOverview: 'Back to overview',
  beaconFrequency: 'Beacon frequency',
  bearerMyAuthToken: 'Bearer my-auth-token',
  brand: 'Brand',
  cancel: 'Cancel',
  changeLocation: 'Change location settings',
  changePassword: 'Change password',
  channel: 'Channel',
  claimAuthCode: 'Claim authentication code',
  claiming: 'Claiming',
  classBTimeout: 'Class B timeout',
  classCTimeout: 'Class C timeout',
  clear: 'Clear',
  client: 'Client',
  clientId: 'Client ID',
  collaborator: 'Collaborator',
  collaboratorCounted: '{count, plural, one {Collaborator} other {Collaborators}}',
  collaboratorDeleteSuccess: 'Collaborator removed',
  collaboratorEdit: 'Edit {collaboratorId}',
  collaboratorEditRights: 'Edit rights of {collaboratorId}',
  collaboratorId: 'Collaborator ID',
  collaboratorIdPlaceholder: 'collaborator-id',
  collaboratorModalWarning: 'Are you sure you want to remove {collaboratorId} as a collaborator?',
  collaboratorModalWarningSelf:
    'Are you sure you want to remove yourself as a collaborator? Access to this entity will be lost until someone else adds you as a collaborator again.',
  collaboratorRemove: 'Collaborator remove',
  collaboratorUpdateSuccess: 'Collaborator rights updated',
  collaboratorWarningAdmin:
    'This user is an administrator that will retain all rights to all entities regardless of collaborator status',
  collaboratorWarningAdminSelf:
    'As an administrator, you always retain all rights to all entities regardless of collaborator status',
  collaboratorWarningSelf: 'Changing your own rights could result in loss of access',
  collaborators: 'Collaborators',
  componentAs: 'Application Server',
  componentDcs: 'Device Claiming Server',
  componentEdtc: 'End Device Template Converter',
  componentGcs: 'Gateway Claiming Server',
  componentGs: 'Gateway Server',
  componentIs: 'Identity Server',
  componentJs: 'Join Server',
  componentNs: 'Network Server',
  componentQrg: 'QR Code Generator',
  confirm: 'Confirm',
  confirmPassword: 'Confirm password',
  confirmedDownlink: 'Confirmed downlink',
  connected: 'Connected',
  connecting: 'Connecting',
  connectionIssues: 'Connection issues',
  contactFieldPlaceholder: 'Type or choose a collaborator',
  contactInformation: 'Contact information',
  convertMacToEui: 'Convert MAC to EUI',
  copiedToClipboard: 'Copied to clipboard!',
  copyToClipboard: 'Copy to clipboard',
  createApiKey: 'Create API key',
  createApplication: 'Create application',
  createOrganization: 'Create organization',
  created: 'Created',
  createdAt: 'Created at',
  currentCollaborators: 'Current collaborators',
  currentUserIndicator: '(This is you)',
  dashboard: 'Dashboard',
  data: 'Data',
  defineRights: 'Define rights',
  delayWarning:
    'Delay too short. The lower bound ({minimumValue}ms) will be used by the Gateway Server.',
  deleteModalConfirmDeletion: 'Confirm deletion',
  deleteModalConfirmMessage: 'Please enter <pre>{entityId}</pre> to confirm the deletion.',
  deleteModalDefaultMessage:
    'This will <strong>PERMANENTLY DELETE THE ENTITY ITSELF AND ALL ASSOCIATED ENTITIES</strong>, including collaborator associations. It will also <strong>NOT BE POSSIBLE TO REUSE THE ENTITY ID</strong>.',
  deleteModalPurgeMessage:
    'This will <strong>PERMANENTLY DELETE THE ENTITY ITSELF AND ALL ASSOCIATED ENTITIES</strong>, including collaborator associations.',
  deleteModalPurgeWarning:
    'Releasing the entity IDs will make it possible to register a new entity with the same ID. Note that this <strong>irreversible</strong> and may lead to <strong>other users gaining access to historical data of the entity if they register an entity with the same ID</strong> . Please make sure you understand the implications of purging as described <DocLink>here</DocLink>.',
  deleteModalReleaseIdLabel: 'Also release entity IDs (purge)',
  deleteModalReleaseIdTitle: 'Entity purge (admin only)',
  deleteModalTitle: 'Are you sure you want to delete <pre>{entityName}</pre>?',
  deleted: 'Deleted (Admin)',
  description: 'Description',
  devAddr: 'Device address',
  devDesc: 'End device description',
  devEUI: 'DevEUI',
  devEUIBlockLimitReached: 'DevEUI generation limit reached',
  devID: 'End device ID',
  devName: 'End device name',
  device: 'End device',
  deviceCounted: '{count, plural, one {End device} other {End devices}}',
  deviceDescDescription:
    'Optional end device description; can also be used to save notes about the end device',
  deviceDescPlaceholder: 'Description for my new end device',
  deviceHardwareVersionAbsence:
    'Contact the manufacturer or reseller of your device. Providing an incorrect hardware version can result in unwanted device behavior.',
  deviceIdPlaceholder: 'my-new-device',
  deviceNamePlaceholder: 'My new end device',
  deviceSimulationDisabledWarning: 'Simulation is disabled for devices that skip payload crypto',
  devices: 'End devices',
  disabled: 'Disabled',
  disconnected: 'Disconnected',
  documentation: 'Documentation',
  downlink: 'Downlink',
  downlinkAck: 'Downlink ack',
  downlinkFailed: 'Downlink failed',
  downlinkFrameCount: 'Downlink frame count',
  downlinkNack: 'Downlink nack',
  downlinkPush: 'Downlink push',
  downlinkQueueInvalidated: 'Downlink queue invalidated',
  downlinkQueued: 'Downlink queued',
  downlinkReplace: 'Downlink replace',
  downlinkSent: 'Downlink sent',
  downlinksScheduled: 'Downlinks (re)scheduled',
  endDeviceOverview: 'End device overview',
  edit: 'Edit',
  editWebhook: 'Edit webhook',
  email: 'Email',
  emailAddress: 'Email address',
  emailAddressDescription:
    'Primary email address used for logging in; this address is not publicly visible',
  emailAddressValidation: 'Treat email address as validated',
  emailAddressValidationDescription:
    'Enable this option if you do not need this user to validate the email address',
  emailPlaceholder: 'mail@example.com',
  empty: 'Empty',
  enabled: 'Enabled',
  endDeviceModelsUnavailable: 'End device models unavailable',
  enforceDutyCycle: 'Enforce duty cycle',
  enforceDutyCycleDescription:
    'Recommended for all gateways in order to respect spectrum regulations',
  entityId: 'Entity ID',
  eventDownlinkAckDesc: 'A confirmed downlink is acknowledged by an end device',
  eventDownlinkFailedDesc: 'A downlink cannot be sent',
  eventDownlinkNackDesc: 'A sent confirmed downlink fails confirmation by the end device',
  eventDownlinkPushDesc: 'A downlink is pushed to the downlink queue',
  eventDownlinkQueueInvalidatedDesc: 'The downlink queue is reset due to frame counter mismatch',
  eventDownlinkQueuedDesc: 'A downlink is added to the downlink queue',
  eventDownlinkReplaceDesc: 'A downlink is used to replace the downlink queue',
  eventDownlinkSentDesc: 'A downlink is sent to an end device or multicast group',
  eventEnabledTypes: 'Enabled event types',
  eventJoinAcceptDesc: 'An end device successfully joins the network and starts a session',
  eventLocationSolvedDesc: 'An integration succeeded locating the end device',
  eventServiceDataDesc: 'An integration emits an event',
  eventUplinkMessageDesc: 'An uplink message is received by the application',
  eventUplinkNormalizedDesc: 'A normalized uplink payload',
  eventsCannotShow: 'Cannot show events',
  expiry: 'Expiry',
  exportJson: 'Export as JSON',
  external: 'External',
  externalJoinServer: 'External Join Server',
  fNwkSIntKey: 'FNwkSIntKey',
  factoryPresetFrequencies: 'Factory preset frequencies',
  fetching: 'Fetching data…',
  firmwareVersion: 'Firmware version',
  format: 'Format',
  fpNotFoundError:
    'The LoRaWAN version <code>{lorawanVersion}</code> does not support the <code>{freqPlan}</code> frequency plan. Please choose a different MAC version or frequency plan.',
  frameCounterWidth: 'Frame counter width',
  freqAdd: 'Add Frequency',
  frequencyPlaceholder: 'e.g. 869525000 for 869,525 MHz',
  frequencyPlan: 'Frequency plan',
  frequencyPlanWarning:
    'Without choosing a frequency plan, packets from the gateway will not be correctly processed',
  furtherResources: 'Further resources',
  gateway: 'Gateway',
  gatewayOverview: 'Gateway overview',
  gatewayDescDescription:
    'Optional gateway description; can also be used to save notes about the gateway',
  gatewayDescPlaceholder: 'Description for my new gateway',
  gatewayDescription: 'Gateway description',
  gatewayEUI: 'Gateway EUI',
  gatewayFanoutNotificationsTitle: 'Fan-out notifications',
  gatewayFanoutNotificationsLabel: 'Enable fan-out notifications',
  gatewayFanoutNotificationsDescription:
    'Notifications sent to this organization will be propagated to all collaborators instead of the respective administrative or technical contact.',
  gatewayID: 'Gateway ID',
  gatewayIdPlaceholder: 'my-new-gateway',
  gatewayLocation: 'Gateway location',
  gatewayLocationPublic: 'Share location within network',
  gatewayName: 'Gateway name',
  gatewayNamePlaceholder: 'My new gateway',
  gatewayScheduleDownlinkLate: 'Schedule downlink late',
  gatewayServerAddress: 'Gateway Server address',
  gatewayStatus: 'Gateway status',
  gatewayStatusPublic: 'Share status within network',
  gatewayUpdateOptions: 'Gateway updates',
  gateways: 'Gateways',
  general: 'General',
  generalInformation: 'General information',
  generalSettings: 'General settings',
  generateAPIKeyCups: 'Generate API key for CUPS',
  generateAPIKeyLNS: 'Generate API key for LNS',
  getSupport: 'Get support',
  grantAdminStatus: 'Grant this user admin status',
  grpcService: 'GRPC service',
  gsServerAddressDescription: 'The address of the Gateway Server to connect to',
  hardware: 'Hardware',
  hardwareVersion: 'Hardware version',
  homeNetID: 'Home NetID',
  homeNetIDDescription: 'ID to identify the LoRaWAN network',
  hours: 'hours',
  id: 'ID',
  idAlreadyExists: 'ID already exists',
  import: 'Import',
  importDevices: 'Import end devices',
  inputMethod: 'Input method',
  insufficientAppKeyRights: 'Insufficient rights to set an AppKey',
  insufficientNwkKeyRights: 'Insufficient rights to set a NwkKey',
  integrations: 'Integrations',
  invite: 'Invite',
  joinAccept: 'Join accept',
  joinEUI: 'JoinEUI',
  joinServerAddress: 'Join Server address',
  key: 'key',
  keyEdit: 'Edit API key',
  keyId: 'Key ID',
  language: 'Language',
  lastSeen: 'Last activity',
  latitude: 'Latitude',
  latitudeDesc: 'The north-south position in degrees, where 0 is the equator',
  lbsLNSSecret: 'LoRa Basics Station LNS Authentication Key',
  lbsLNSSecretDescription:
    'The Authentication Key for Lora Basics Station LNS connections. This field is ignored for other gateways.',
  link: 'Link',
  linked: 'Linked',
  list: 'List',
  liveData: 'Live data',
  location: 'Location',
  locationDescription:
    'When set to public, the gateway location may be visible to other users of the network',
  locationMarkerDescriptionNonUser:
    'This location has been set automatically from incoming (status) messages',
  locationMarkerDescriptionUntrusted:
    'This location was determined via an untrusted status message and may be inaccurate',
  locationMarkerDescriptionUser:
    'This location has been set manually (e.g. by using the "Location"-tab)',
  locationSolved: 'Location solved',
  locationSourceBtRssi: 'Bluetooth RSSI geolocation',
  locationSourceCombined: 'Combined geolocation',
  locationSourceGps: 'GPS-based location',
  locationSourceIpGeolocation: 'IP-based geolocation',
  locationSourceLoraRssi: 'LoRa RSSI geolocation',
  locationSourceLoraTdoa: 'LoRa TDOA geolocation',
  locationSourceRegistry: 'Manually set location',
  locationSourceWifiRssi: 'Wifi RSSI geolocation',
  login: 'Login',
  loginFailed: 'Login failed',
  logout: 'Logout',
  longitude: 'Longitude',
  longitudeDesc: 'The east-west position in degrees, where 0 is the prime meridian (Greenwich)',
  loraCloud: 'LoRa Cloud',
  loraCloudServerUrlDescription: 'LoRa Cloud Modem and Geolocation Services Server URL',
  lorawanClassCapabilities: 'LoRaWAN class capabilities',
  lorawanInformation: 'LoRaWAN information',
  lorawanOptions: 'LoRaWAN options',
  lorawanPhyVersionDescription: 'The LoRaWAN PHY version of the end device',
  macData: 'MAC data',
  macSettingsError:
    'There was an error and the default MAC settings for the <code>{freqPlan}</code> frequency plan could not be loaded',
  macVersion: 'LoRaWAN version',
  messageTypes: 'Message types',
  messages: 'messages',
  messaging: 'Messaging',
  milliseconds: 'milliseconds',
  minutes: 'minutes',
  model: 'Model',
  moreInformation: 'More information',
  mqtt: 'MQTT',
  multicast: 'Define multicast group (ABP & Multicast)',
  name: 'Name',
  netId: 'Net ID',
  networkInformation: 'Network information',
  networkServerAddress: 'Network Server address',
  networks: 'Networks',
  never: 'Never',
  next: 'Next',
  noActivityYet: 'No activity yet',
  noDesc: 'This end device has no description',
  noEvents: 'Waiting for events from <pre>{entityId}</pre>…',
  noLocation: 'No location information available',
  noMatch: 'No items found',
  noMatchingUserFound: 'No matching user or organization was found',
  noRecentActivity: 'No recent activity',
  none: 'None',
  normalizedPayloadAir: 'Air',
  normalizedPayloadSoil: 'Soil',
  normalizedPayloadWind: 'Wind',
  notAvailable: 'n/a',
  notLinked: 'Not linked',
  notifications: 'Notifications',
  notSet: 'Not set',
  nsEmptyDefault: 'Leave empty to link to the Network Server in the same cluster',
  nsServerKekLabel: 'Network Server KEK label',
  nsServerKekLabelDescription:
    'The KEK label of the Network Server to use for wrapping the network session key',
  nwkKey: 'NwkKey',
  nwkSEncKey: 'NwkSEncKey',
  nwkSEncKeyDescription: 'Network session encryption key',
  nwkSKey: 'NwkSKey',
  oauthClientAuthorizations: 'OAuth client authorizations',
  oauthClientId: 'OAuth client ID',
  oauthClients: 'OAuth clients',
  offline: 'Offline',
  ok: 'Ok',
  online: 'Online',
  options: 'Options',
  organization: 'Organization',
  organizationId: 'Organization ID',
  organizations: 'Organizations',
  otaa: 'Over the air activation (OTAA)',
  otherCluster: 'Other cluster',
  otherOption: 'Other…',
  overview: 'Overview',
  packetBroker: 'Packet Broker',
  password: 'Password',
  passwordChanged: 'Password changed',
  pause: 'Pause',
  payload: 'Payload',
  payloadFormatters: 'Payload formatters',
  payloadFormattersDownlink: 'Downlink payload formatters',
  payloadFormattersUpdateFailure: 'There was an error updating the payload formatter',
  payloadFormattersUpdateSuccess: 'Payload formatter updated',
  payloadFormattersUplink: 'Uplink payload formatters',
  personalApiKeys: 'Personal API keys',
  phyVersion: 'Regional Parameters version',
  phyVersionDescription:
    'The Regional Parameters version (PHY), as provided by the device manufacturer',
  pingSlotFrequency: 'Ping slot frequency',
  pingSlotPeriodicity: 'Ping slot periodicity',
  port: 'Port',
  privacyPolicy: 'Privacy policy',
  profileSettings: 'Profile settings',
  provider: 'Provider',
  provisionedOnExternalJoinServer: 'Provisioned on external Join Server',
  pubsubBaseTopic: 'Base topic',
  pubsubFormat: 'Pub/Sub format',
  pubsubId: 'Pub/Sub ID',
  pubsubs: 'Pub/Subs',
  purge: 'Purge',
  redirecting: 'Redirecting…',
  refresh: 'Refresh',
  registerEndDevice: 'Register end device',
  registerGateway: 'Register gateway',
  remove: 'Remove',
  removeCollaborator: 'Remove collaborator',
  removeCollaboratorLast: 'Cannot remove last collaborator',
  removeCollaboratorSelf: 'Remove yourself as collaborator',
  replaceWebhook: 'Replace webhook',
  requireAuthenticatedConnection: 'Require authenticated connection',
  requireAuthenticatedConnectionDescription:
    'Controls whether this gateway may only connect if it uses an authenticated Basic Station or MQTT connection',
  reset: 'Reset',
  resetWarning: 'Reseting is insecure and makes your end device susceptible for replay attacks',
  resetsFCnt: 'Resets frame counters',
  resetsJoinNonces: 'Resets join nonces',
  restartStream: 'Restart stream',
  restore: 'Restore',
  restrictedUser:
    'You can only set yourself as a contact. If you would like another collaborator set as contact, please contact this collaborator to self-assign as contact.',
  resume: 'Resume',
  rights: 'Rights',
  rightsList: 'Rights:',
  rootKeys: 'Root keys',
  rx1DataRateOffset: 'Rx1 data rate offset',
  rx1Delay: 'Rx1 delay',
  rx2Frequency: 'Rx2 frequency',
  sNwkSIKey: 'SNwkSIntKey',
  sNwkSIKeyDescription: 'Serving network session integrity key',
  saveChanges: 'Save changes',
  scanEndDevice: 'Scan end device QR code',
  scheduleAnyTimeDelay: 'Schedule any time delay',
  scheduleAnyTimeDescription:
    'Configure gateway delay (minimum: {minimumValue}ms, default: {defaultValue}ms)',
  scheduleDownlinkLateDescription: 'Enable server-side buffer of downlink messages',
  search: 'Search',
  secondInterval: '{count, plural, one {every second} other {every {count} seconds}}',
  seconds: 'seconds',
  secondsAbbreviated: 'sec',
  secret: 'Secret',
  secure: 'Secure',
  sendInvitation: 'Send invitation',
  serverUrl: 'Server URL',
  serviceData: 'Service data',
  sessions: 'Sessions',
  setLoRaCloudToken: 'Set LoRa Cloud token',
  settings: 'Settings',
  setup: 'Setup',
  shareGatewayInfo: 'Share gateway information',
  skipCryptoDescription: 'Skip decryption of uplink payloads and encryption of downlink payloads',
  skipCryptoPlaceholder: 'Encryption/decryption disabled',
  skipCryptoTitle: 'Skip payload encryption and decryption',
  source: 'Source',
  stable: 'Stable',
  state: 'State',
  stateApproved: 'Approved',
  stateDescription: 'State description',
  stateFlagged: 'Flagged',
  stateRejected: 'Rejected',
  stateRequested: 'Requested',
  stateSuspended: 'Suspended',
  status: 'Status',
  statusDescription: 'The status of this gateway may be visible to other users',
  statusPage: 'Status page',
  statusUnknown: 'Status unknown',
  success: 'Success',
  suggestions: 'Suggestions',
  support: 'Support',
  supportsClassB: 'Supports class B',
  supportsClassC: 'Supports class C',
  takeMeBack: 'Take me back',
  technicalContact: 'Technical contact',
  tenantId: 'Tenant ID',
  termsAndCondition: 'Terms and conditions',
  time: 'Time',
  token: 'Token',
  tokenDelete: 'Token delete',
  tokenDeleted: 'Token deleted',
  topApplications: 'Top applications',
  topGateways: 'Top gateways',
  topEntities: 'Top entities',
  tokenSet: 'Set token',
  tokenUpdated: 'Token updated',
  traffic: 'Traffic',
  troubleshooting: 'Troubleshooting',
  type: 'Type',
  typeToSearch: 'Type to search…',
  unexposed: 'Unexposed',
  unknown: 'Unknown',
  unknownError: 'Unknown error',
  unknownHwOption: 'Unknown ver.',
  updateChannelDescription: 'Channel for gateway automatic updates',
  updatedAt: 'Last updated at',
  uplink: 'Uplink',
  uplinkFrameCount: 'Uplink frame count',
  uplinkMessage: 'Uplink message',
  uplinkNormalized: 'Normalized uplink',
  uplinksReceived: 'Uplinks received',
  uploadAnImage: 'Upload an image',
  used: '{currentValue}/{maxValue} used',
  useDefaultPolicy: 'Use default routing policy for this network',
  user: 'User',
  userAdd: 'Add user',
  userDelete: 'Delete user',
  userDescDescription: 'Optional user description; can also be used to save notes about the user',
  userDescription: 'Description for my new user',
  userEdit: 'Edit user',
  userId: 'User ID',
  userIdPlaceholder: 'jane-doe',
  userInvitations: 'User invitations',
  userManagement: 'User management',
  userNamePlaceholder: 'Jane Doe',
  userOrgId: 'User / Organization ID',
  username: 'Username',
  users: 'Users',
  validFrom: 'Valid from',
  validTo: 'Valid to',
  validateAddressFormat: '{field} must be in the format "host" or "host:port"',
  validateApiKey: 'API keys must follow the format "NNSXS.[…].[…]"',
  validateDateInPast: '{field} must be a date in the future',
  validateDelayFormat: '{field} must be a positive, whole number',
  validateDigit: '{field} must have at least {digit} {digit, plural, one {digit} other {digits}}',
  validateEmail: 'An email address must use exactly one "@", one "." and use no special characters',
  validateFreqDynamic: '{field} must be 0 for dynamic frequencies or greater than 100000Hz',
  validateFreqNumeric: 'All frequency values must be positive integers',
  validateFreqRequired: 'All frequency values are required. Please remove empty entries.',
  validateHexLength: '{field} must be a complete hex value',
  validateIdFormat: '{field} must contain only lowercase letters, numbers and dashes (-)',
  validateInt32: '{field} must be a whole number, negative or positive',
  validateJson: '{field} must be a valid JSON object',
  validateLatitude: 'Latitude must be a whole or decimal number between -90 and 90',
  validateLength: '{field} must be exactly {length} characters long',
  validateLongitude: 'Longitude must be a whole or decimal number between -180 and 180',
  validateMacAddressEntered:
    '{field}s are 8 bytes. If you have entered a MAC address instead, use the button to convert it.',
  validateMqttPassword: '{field} must be empty or have at least 2 characters',
  validateMqttUrl:
    'MQTT URLs must have the format "mqtt[s]://[username][:password]@host.domain[:port]"',
  validateNoSpaces: '{field} must contain no spaces',
  validateNumberGte: '{field} must be at least {min} or higher',
  validateNumberLte: '{field} must be {max} or lower',
  validatePasswordMatch: 'Passwords must match',
  validateRequired: '{field} is required',
  validateRights: 'At least one right must be selected',
  validateSpecial:
    '{field} must have at least {special} special {special, plural, one {character} other {characters}}',
  validateTooLong: '{field} must have less than {max} characters',
  validateTooShort: '{field} must have at least {min} characters',
  validateUppercase:
    '{field} must have at least {upper} uppercase {upper, plural, one {character} other {characters}}',
  validateUrl: 'Must be a valid URL format, contain no spaces or special characters',
  value: 'value',
  webhookActivated: 'Webhook activated',
  viewLink: 'You can view and edit this API key <Link>here</Link>.',
  webhookAlreadyExistsModalMessage:
    'A Webhook with the ID "{id}" already exists. Do you wish to replace this webhook?',
  webhookBaseUrl: 'Base URL',
  webhookDeleted: 'Webhook deleted',
  webhookFormat: 'Webhook format',
  webhookId: 'Webhook ID',
  webhookUpdated: 'Webhook updated',
  webhooks: 'Webhooks',
})
