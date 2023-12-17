// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-iosxe/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPAddressFamilyIPv4 struct {
	Device                           types.String                                        `tfsdk:"device"`
	Id                               types.String                                        `tfsdk:"id"`
	DeleteMode                       types.String                                        `tfsdk:"delete_mode"`
	Asn                              types.String                                        `tfsdk:"asn"`
	AfName                           types.String                                        `tfsdk:"af_name"`
	Ipv4UnicastRedistributeConnected types.Bool                                          `tfsdk:"ipv4_unicast_redistribute_connected"`
	Ipv4UnicastRedistributeStatic    types.Bool                                          `tfsdk:"ipv4_unicast_redistribute_static"`
	Ipv4UnicastAggregateAddresses    []BGPAddressFamilyIPv4Ipv4UnicastAggregateAddresses `tfsdk:"ipv4_unicast_aggregate_addresses"`
	Ipv4UnicastNetworksMask          []BGPAddressFamilyIPv4Ipv4UnicastNetworksMask       `tfsdk:"ipv4_unicast_networks_mask"`
	Ipv4UnicastNetworks              []BGPAddressFamilyIPv4Ipv4UnicastNetworks           `tfsdk:"ipv4_unicast_networks"`
}

type BGPAddressFamilyIPv4Data struct {
	Device                           types.String                                        `tfsdk:"device"`
	Id                               types.String                                        `tfsdk:"id"`
	Asn                              types.String                                        `tfsdk:"asn"`
	AfName                           types.String                                        `tfsdk:"af_name"`
	Ipv4UnicastRedistributeConnected types.Bool                                          `tfsdk:"ipv4_unicast_redistribute_connected"`
	Ipv4UnicastRedistributeStatic    types.Bool                                          `tfsdk:"ipv4_unicast_redistribute_static"`
	Ipv4UnicastAggregateAddresses    []BGPAddressFamilyIPv4Ipv4UnicastAggregateAddresses `tfsdk:"ipv4_unicast_aggregate_addresses"`
	Ipv4UnicastNetworksMask          []BGPAddressFamilyIPv4Ipv4UnicastNetworksMask       `tfsdk:"ipv4_unicast_networks_mask"`
	Ipv4UnicastNetworks              []BGPAddressFamilyIPv4Ipv4UnicastNetworks           `tfsdk:"ipv4_unicast_networks"`
}
type BGPAddressFamilyIPv4Ipv4UnicastAggregateAddresses struct {
	Ipv4Address types.String `tfsdk:"ipv4_address"`
	Ipv4Mask    types.String `tfsdk:"ipv4_mask"`
}
type BGPAddressFamilyIPv4Ipv4UnicastNetworksMask struct {
	Network  types.String `tfsdk:"network"`
	Mask     types.String `tfsdk:"mask"`
	RouteMap types.String `tfsdk:"route_map"`
	Backdoor types.Bool   `tfsdk:"backdoor"`
}
type BGPAddressFamilyIPv4Ipv4UnicastNetworks struct {
	Network  types.String `tfsdk:"network"`
	RouteMap types.String `tfsdk:"route_map"`
	Backdoor types.Bool   `tfsdk:"backdoor"`
}

func (data BGPAddressFamilyIPv4) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv4=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

func (data BGPAddressFamilyIPv4Data) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp=%v/address-family/no-vrf/ipv4=%s", url.QueryEscape(fmt.Sprintf("%v", data.Asn.ValueString())), url.QueryEscape(fmt.Sprintf("%v", data.AfName.ValueString())))
}

// if last path element has a key -> remove it
func (data BGPAddressFamilyIPv4) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data BGPAddressFamilyIPv4) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"af-name", data.AfName.ValueString())
	}
	if !data.Ipv4UnicastRedistributeConnected.IsNull() && !data.Ipv4UnicastRedistributeConnected.IsUnknown() {
		if data.Ipv4UnicastRedistributeConnected.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.redistribute.connected", map[string]string{})
		}
	}
	if !data.Ipv4UnicastRedistributeStatic.IsNull() && !data.Ipv4UnicastRedistributeStatic.IsUnknown() {
		if data.Ipv4UnicastRedistributeStatic.ValueBool() {
			body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.redistribute.static", map[string]string{})
		}
	}
	if len(data.Ipv4UnicastAggregateAddresses) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.aggregate-address", []interface{}{})
		for index, item := range data.Ipv4UnicastAggregateAddresses {
			if !item.Ipv4Address.IsNull() && !item.Ipv4Address.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.aggregate-address"+"."+strconv.Itoa(index)+"."+"ipv4-address", item.Ipv4Address.ValueString())
			}
			if !item.Ipv4Mask.IsNull() && !item.Ipv4Mask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.aggregate-address"+"."+strconv.Itoa(index)+"."+"ipv4-mask", item.Ipv4Mask.ValueString())
			}
		}
	}
	if len(data.Ipv4UnicastNetworksMask) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.with-mask", []interface{}{})
		for index, item := range data.Ipv4UnicastNetworksMask {
			if !item.Network.IsNull() && !item.Network.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.with-mask"+"."+strconv.Itoa(index)+"."+"number", item.Network.ValueString())
			}
			if !item.Mask.IsNull() && !item.Mask.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.with-mask"+"."+strconv.Itoa(index)+"."+"mask", item.Mask.ValueString())
			}
			if !item.RouteMap.IsNull() && !item.RouteMap.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.with-mask"+"."+strconv.Itoa(index)+"."+"route-map", item.RouteMap.ValueString())
			}
			if !item.Backdoor.IsNull() && !item.Backdoor.IsUnknown() {
				if item.Backdoor.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.with-mask"+"."+strconv.Itoa(index)+"."+"backdoor", map[string]string{})
				}
			}
		}
	}
	if len(data.Ipv4UnicastNetworks) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.no-mask", []interface{}{})
		for index, item := range data.Ipv4UnicastNetworks {
			if !item.Network.IsNull() && !item.Network.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.no-mask"+"."+strconv.Itoa(index)+"."+"number", item.Network.ValueString())
			}
			if !item.RouteMap.IsNull() && !item.RouteMap.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.no-mask"+"."+strconv.Itoa(index)+"."+"route-map", item.RouteMap.ValueString())
			}
			if !item.Backdoor.IsNull() && !item.Backdoor.IsUnknown() {
				if item.Backdoor.ValueBool() {
					body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"ipv4-unicast.network.no-mask"+"."+strconv.Itoa(index)+"."+"backdoor", map[string]string{})
				}
			}
		}
	}
	return body
}

func (data *BGPAddressFamilyIPv4) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "af-name"); value.Exists() && !data.AfName.IsNull() {
		data.AfName = types.StringValue(value.String())
	} else {
		data.AfName = types.StringNull()
	}
	if value := res.Get(prefix + "ipv4-unicast.redistribute.connected"); !data.Ipv4UnicastRedistributeConnected.IsNull() {
		if value.Exists() {
			data.Ipv4UnicastRedistributeConnected = types.BoolValue(true)
		} else {
			data.Ipv4UnicastRedistributeConnected = types.BoolValue(false)
		}
	} else {
		data.Ipv4UnicastRedistributeConnected = types.BoolNull()
	}
	if value := res.Get(prefix + "ipv4-unicast.redistribute.static"); !data.Ipv4UnicastRedistributeStatic.IsNull() {
		if value.Exists() {
			data.Ipv4UnicastRedistributeStatic = types.BoolValue(true)
		} else {
			data.Ipv4UnicastRedistributeStatic = types.BoolValue(false)
		}
	} else {
		data.Ipv4UnicastRedistributeStatic = types.BoolNull()
	}
	for i := range data.Ipv4UnicastAggregateAddresses {
		keys := [...]string{"ipv4-address", "ipv4-mask"}
		keyValues := [...]string{data.Ipv4UnicastAggregateAddresses[i].Ipv4Address.ValueString(), data.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv4-unicast.aggregate-address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("ipv4-address"); value.Exists() && !data.Ipv4UnicastAggregateAddresses[i].Ipv4Address.IsNull() {
			data.Ipv4UnicastAggregateAddresses[i].Ipv4Address = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastAggregateAddresses[i].Ipv4Address = types.StringNull()
		}
		if value := r.Get("ipv4-mask"); value.Exists() && !data.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.IsNull() {
			data.Ipv4UnicastAggregateAddresses[i].Ipv4Mask = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastAggregateAddresses[i].Ipv4Mask = types.StringNull()
		}
	}
	for i := range data.Ipv4UnicastNetworksMask {
		keys := [...]string{"number", "mask"}
		keyValues := [...]string{data.Ipv4UnicastNetworksMask[i].Network.ValueString(), data.Ipv4UnicastNetworksMask[i].Mask.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv4-unicast.network.with-mask").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("number"); value.Exists() && !data.Ipv4UnicastNetworksMask[i].Network.IsNull() {
			data.Ipv4UnicastNetworksMask[i].Network = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastNetworksMask[i].Network = types.StringNull()
		}
		if value := r.Get("mask"); value.Exists() && !data.Ipv4UnicastNetworksMask[i].Mask.IsNull() {
			data.Ipv4UnicastNetworksMask[i].Mask = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastNetworksMask[i].Mask = types.StringNull()
		}
		if value := r.Get("route-map"); value.Exists() && !data.Ipv4UnicastNetworksMask[i].RouteMap.IsNull() {
			data.Ipv4UnicastNetworksMask[i].RouteMap = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastNetworksMask[i].RouteMap = types.StringNull()
		}
		if value := r.Get("backdoor"); !data.Ipv4UnicastNetworksMask[i].Backdoor.IsNull() {
			if value.Exists() {
				data.Ipv4UnicastNetworksMask[i].Backdoor = types.BoolValue(true)
			} else {
				data.Ipv4UnicastNetworksMask[i].Backdoor = types.BoolValue(false)
			}
		} else {
			data.Ipv4UnicastNetworksMask[i].Backdoor = types.BoolNull()
		}
	}
	for i := range data.Ipv4UnicastNetworks {
		keys := [...]string{"number"}
		keyValues := [...]string{data.Ipv4UnicastNetworks[i].Network.ValueString()}

		var r gjson.Result
		res.Get(prefix + "ipv4-unicast.network.no-mask").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("number"); value.Exists() && !data.Ipv4UnicastNetworks[i].Network.IsNull() {
			data.Ipv4UnicastNetworks[i].Network = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastNetworks[i].Network = types.StringNull()
		}
		if value := r.Get("route-map"); value.Exists() && !data.Ipv4UnicastNetworks[i].RouteMap.IsNull() {
			data.Ipv4UnicastNetworks[i].RouteMap = types.StringValue(value.String())
		} else {
			data.Ipv4UnicastNetworks[i].RouteMap = types.StringNull()
		}
		if value := r.Get("backdoor"); !data.Ipv4UnicastNetworks[i].Backdoor.IsNull() {
			if value.Exists() {
				data.Ipv4UnicastNetworks[i].Backdoor = types.BoolValue(true)
			} else {
				data.Ipv4UnicastNetworks[i].Backdoor = types.BoolValue(false)
			}
		} else {
			data.Ipv4UnicastNetworks[i].Backdoor = types.BoolNull()
		}
	}
}

func (data *BGPAddressFamilyIPv4Data) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "ipv4-unicast.redistribute.connected"); value.Exists() {
		data.Ipv4UnicastRedistributeConnected = types.BoolValue(true)
	} else {
		data.Ipv4UnicastRedistributeConnected = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv4-unicast.redistribute.static"); value.Exists() {
		data.Ipv4UnicastRedistributeStatic = types.BoolValue(true)
	} else {
		data.Ipv4UnicastRedistributeStatic = types.BoolValue(false)
	}
	if value := res.Get(prefix + "ipv4-unicast.aggregate-address"); value.Exists() {
		data.Ipv4UnicastAggregateAddresses = make([]BGPAddressFamilyIPv4Ipv4UnicastAggregateAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv4Ipv4UnicastAggregateAddresses{}
			if cValue := v.Get("ipv4-address"); cValue.Exists() {
				item.Ipv4Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv4-mask"); cValue.Exists() {
				item.Ipv4Mask = types.StringValue(cValue.String())
			}
			data.Ipv4UnicastAggregateAddresses = append(data.Ipv4UnicastAggregateAddresses, item)
			return true
		})
	}
	if value := res.Get(prefix + "ipv4-unicast.network.with-mask"); value.Exists() {
		data.Ipv4UnicastNetworksMask = make([]BGPAddressFamilyIPv4Ipv4UnicastNetworksMask, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv4Ipv4UnicastNetworksMask{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Network = types.StringValue(cValue.String())
			}
			if cValue := v.Get("mask"); cValue.Exists() {
				item.Mask = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map"); cValue.Exists() {
				item.RouteMap = types.StringValue(cValue.String())
			}
			if cValue := v.Get("backdoor"); cValue.Exists() {
				item.Backdoor = types.BoolValue(true)
			} else {
				item.Backdoor = types.BoolValue(false)
			}
			data.Ipv4UnicastNetworksMask = append(data.Ipv4UnicastNetworksMask, item)
			return true
		})
	}
	if value := res.Get(prefix + "ipv4-unicast.network.no-mask"); value.Exists() {
		data.Ipv4UnicastNetworks = make([]BGPAddressFamilyIPv4Ipv4UnicastNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := BGPAddressFamilyIPv4Ipv4UnicastNetworks{}
			if cValue := v.Get("number"); cValue.Exists() {
				item.Network = types.StringValue(cValue.String())
			}
			if cValue := v.Get("route-map"); cValue.Exists() {
				item.RouteMap = types.StringValue(cValue.String())
			}
			if cValue := v.Get("backdoor"); cValue.Exists() {
				item.Backdoor = types.BoolValue(true)
			} else {
				item.Backdoor = types.BoolValue(false)
			}
			data.Ipv4UnicastNetworks = append(data.Ipv4UnicastNetworks, item)
			return true
		})
	}
}

func (data *BGPAddressFamilyIPv4) getDeletedItems(ctx context.Context, state BGPAddressFamilyIPv4) []string {
	deletedItems := make([]string, 0)
	if !state.Ipv4UnicastRedistributeConnected.IsNull() && data.Ipv4UnicastRedistributeConnected.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/redistribute/connected", state.getPath()))
	}
	if !state.Ipv4UnicastRedistributeStatic.IsNull() && data.Ipv4UnicastRedistributeStatic.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/redistribute/static", state.getPath()))
	}
	for i := range state.Ipv4UnicastAggregateAddresses {
		stateKeyValues := [...]string{state.Ipv4UnicastAggregateAddresses[i].Ipv4Address.ValueString(), state.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4UnicastAggregateAddresses[i].Ipv4Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4UnicastAggregateAddresses {
			found = true
			if state.Ipv4UnicastAggregateAddresses[i].Ipv4Address.ValueString() != data.Ipv4UnicastAggregateAddresses[j].Ipv4Address.ValueString() {
				found = false
			}
			if state.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.ValueString() != data.Ipv4UnicastAggregateAddresses[j].Ipv4Mask.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/aggregate-address=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv4UnicastNetworksMask {
		stateKeyValues := [...]string{state.Ipv4UnicastNetworksMask[i].Network.ValueString(), state.Ipv4UnicastNetworksMask[i].Mask.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4UnicastNetworksMask[i].Network.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv4UnicastNetworksMask[i].Mask.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4UnicastNetworksMask {
			found = true
			if state.Ipv4UnicastNetworksMask[i].Network.ValueString() != data.Ipv4UnicastNetworksMask[j].Network.ValueString() {
				found = false
			}
			if state.Ipv4UnicastNetworksMask[i].Mask.ValueString() != data.Ipv4UnicastNetworksMask[j].Mask.ValueString() {
				found = false
			}
			if found {
				if !state.Ipv4UnicastNetworksMask[i].RouteMap.IsNull() && data.Ipv4UnicastNetworksMask[j].RouteMap.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/with-mask=%v/route-map", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Ipv4UnicastNetworksMask[i].Backdoor.IsNull() && data.Ipv4UnicastNetworksMask[j].Backdoor.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/with-mask=%v/backdoor", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/with-mask=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	for i := range state.Ipv4UnicastNetworks {
		stateKeyValues := [...]string{state.Ipv4UnicastNetworks[i].Network.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4UnicastNetworks[i].Network.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4UnicastNetworks {
			found = true
			if state.Ipv4UnicastNetworks[i].Network.ValueString() != data.Ipv4UnicastNetworks[j].Network.ValueString() {
				found = false
			}
			if found {
				if !state.Ipv4UnicastNetworks[i].RouteMap.IsNull() && data.Ipv4UnicastNetworks[j].RouteMap.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/no-mask=%v/route-map", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				if !state.Ipv4UnicastNetworks[i].Backdoor.IsNull() && data.Ipv4UnicastNetworks[j].Backdoor.IsNull() {
					deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/no-mask=%v/backdoor", state.getPath(), strings.Join(stateKeyValues[:], ",")))
				}
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/ipv4-unicast/network/no-mask=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	return deletedItems
}

func (data *BGPAddressFamilyIPv4) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.Ipv4UnicastRedistributeConnected.IsNull() && !data.Ipv4UnicastRedistributeConnected.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv4-unicast/redistribute/connected", data.getPath()))
	}
	if !data.Ipv4UnicastRedistributeStatic.IsNull() && !data.Ipv4UnicastRedistributeStatic.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv4-unicast/redistribute/static", data.getPath()))
	}

	for i := range data.Ipv4UnicastNetworksMask {
		keyValues := [...]string{data.Ipv4UnicastNetworksMask[i].Network.ValueString(), data.Ipv4UnicastNetworksMask[i].Mask.ValueString()}
		if !data.Ipv4UnicastNetworksMask[i].Backdoor.IsNull() && !data.Ipv4UnicastNetworksMask[i].Backdoor.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv4-unicast/network/with-mask=%v/backdoor", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}

	for i := range data.Ipv4UnicastNetworks {
		keyValues := [...]string{data.Ipv4UnicastNetworks[i].Network.ValueString()}
		if !data.Ipv4UnicastNetworks[i].Backdoor.IsNull() && !data.Ipv4UnicastNetworks[i].Backdoor.ValueBool() {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/ipv4-unicast/network/no-mask=%v/backdoor", data.getPath(), strings.Join(keyValues[:], ",")))
		}
	}
	return emptyLeafsDelete
}

func (data *BGPAddressFamilyIPv4) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Ipv4UnicastRedistributeConnected.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4-unicast/redistribute/connected", data.getPath()))
	}
	if !data.Ipv4UnicastRedistributeStatic.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4-unicast/redistribute/static", data.getPath()))
	}
	for i := range data.Ipv4UnicastAggregateAddresses {
		keyValues := [...]string{data.Ipv4UnicastAggregateAddresses[i].Ipv4Address.ValueString(), data.Ipv4UnicastAggregateAddresses[i].Ipv4Mask.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4-unicast/aggregate-address=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Ipv4UnicastNetworksMask {
		keyValues := [...]string{data.Ipv4UnicastNetworksMask[i].Network.ValueString(), data.Ipv4UnicastNetworksMask[i].Mask.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4-unicast/network/with-mask=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	for i := range data.Ipv4UnicastNetworks {
		keyValues := [...]string{data.Ipv4UnicastNetworks[i].Network.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4-unicast/network/no-mask=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	return deletePaths
}
