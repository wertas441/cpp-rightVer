// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *CreateRelayRequest) SetFields(src *CreateRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &RelaySettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayResponse) SetFields(src *CreateRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &RelaySettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayRequest) SetFields(src *GetRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayResponse) SetFields(src *GetRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &RelaySettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayRequest) SetFields(src *UpdateRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &RelaySettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayResponse) SetFields(src *UpdateRelayResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "settings":
			if len(subs) > 0 {
				var newDst, newSrc *RelaySettings
				if (src == nil || src.Settings == nil) && dst.Settings == nil {
					continue
				}
				if src != nil {
					newSrc = src.Settings
				}
				if dst.Settings != nil {
					newDst = dst.Settings
				} else {
					newDst = &RelaySettings{}
					dst.Settings = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Settings = src.Settings
				} else {
					dst.Settings = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayRequest) SetFields(src *DeleteRelayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayResponse) SetFields(src *DeleteRelayResponse, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message DeleteRelayResponse has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *CreateRelayUplinkForwardingRuleRequest) SetFields(src *CreateRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardingRule
				if (src == nil || src.Rule == nil) && dst.Rule == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rule
				}
				if dst.Rule != nil {
					newDst = dst.Rule
				} else {
					newDst = &RelayUplinkForwardingRule{}
					dst.Rule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rule = src.Rule
				} else {
					dst.Rule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CreateRelayUplinkForwardingRuleResponse) SetFields(src *CreateRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardingRule
				if (src == nil || src.Rule == nil) && dst.Rule == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rule
				}
				if dst.Rule != nil {
					newDst = dst.Rule
				} else {
					newDst = &RelayUplinkForwardingRule{}
					dst.Rule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rule = src.Rule
				} else {
					dst.Rule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayUplinkForwardingRuleRequest) SetFields(src *GetRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetRelayUplinkForwardingRuleResponse) SetFields(src *GetRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardingRule
				if (src == nil || src.Rule == nil) && dst.Rule == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rule
				}
				if dst.Rule != nil {
					newDst = dst.Rule
				} else {
					newDst = &RelayUplinkForwardingRule{}
					dst.Rule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rule = src.Rule
				} else {
					dst.Rule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListRelayUplinkForwardingRulesRequest) SetFields(src *ListRelayUplinkForwardingRulesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListRelayUplinkForwardingRulesResponse) SetFields(src *ListRelayUplinkForwardingRulesResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "rules":
			if len(subs) > 0 {
				return fmt.Errorf("'rules' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rules = src.Rules
			} else {
				dst.Rules = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayUplinkForwardingRuleRequest) SetFields(src *UpdateRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}
		case "rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardingRule
				if (src == nil || src.Rule == nil) && dst.Rule == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rule
				}
				if dst.Rule != nil {
					newDst = dst.Rule
				} else {
					newDst = &RelayUplinkForwardingRule{}
					dst.Rule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rule = src.Rule
				} else {
					dst.Rule = nil
				}
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UpdateRelayUplinkForwardingRuleResponse) SetFields(src *UpdateRelayUplinkForwardingRuleResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "rule":
			if len(subs) > 0 {
				var newDst, newSrc *RelayUplinkForwardingRule
				if (src == nil || src.Rule == nil) && dst.Rule == nil {
					continue
				}
				if src != nil {
					newSrc = src.Rule
				}
				if dst.Rule != nil {
					newDst = dst.Rule
				} else {
					newDst = &RelayUplinkForwardingRule{}
					dst.Rule = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Rule = src.Rule
				} else {
					dst.Rule = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayUplinkForwardingRuleRequest) SetFields(src *DeleteRelayUplinkForwardingRuleRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "index":
			if len(subs) > 0 {
				return fmt.Errorf("'index' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Index = src.Index
			} else {
				var zero uint32
				dst.Index = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *DeleteRelayUplinkForwardingRuleResponse) SetFields(src *DeleteRelayUplinkForwardingRuleResponse, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message DeleteRelayUplinkForwardingRuleResponse has no fields, but paths %s were specified", paths)
	}
	return nil
}
