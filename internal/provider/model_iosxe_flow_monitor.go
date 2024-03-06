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

type FlowMonitor struct {
	Device             types.String           `tfsdk:"device"`
	Id                 types.String           `tfsdk:"id"`
	DeleteMode         types.String           `tfsdk:"delete_mode"`
	Name               types.String           `tfsdk:"name"`
	Description        types.String           `tfsdk:"description"`
	Exporters          []FlowMonitorExporters `tfsdk:"exporters"`
	CacheTimeoutActive types.Int64            `tfsdk:"cache_timeout_active"`
	Record             types.String           `tfsdk:"record"`
}

type FlowMonitorData struct {
	Device             types.String           `tfsdk:"device"`
	Id                 types.String           `tfsdk:"id"`
	Name               types.String           `tfsdk:"name"`
	Description        types.String           `tfsdk:"description"`
	Exporters          []FlowMonitorExporters `tfsdk:"exporters"`
	CacheTimeoutActive types.Int64            `tfsdk:"cache_timeout_active"`
	Record             types.String           `tfsdk:"record"`
}
type FlowMonitorExporters struct {
	Name types.String `tfsdk:"name"`
}

func (data FlowMonitor) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/Cisco-IOS-XE-flow:flow/monitor=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

func (data FlowMonitorData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XE-native:native/Cisco-IOS-XE-flow:flow/monitor=%v", url.QueryEscape(fmt.Sprintf("%v", data.Name.ValueString())))
}

// if last path element has a key -> remove it
func (data FlowMonitor) getPathShort() string {
	path := data.getPath()
	re := regexp.MustCompile(`(.*)=[^\/]*$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) <= 1 {
		return path
	}
	return matches[1]
}

func (data FlowMonitor) toBody(ctx context.Context) string {
	body := `{"` + helpers.LastElement(data.getPath()) + `":{}}`
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"name", data.Name.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"description", data.Description.ValueString())
	}
	if !data.CacheTimeoutActive.IsNull() && !data.CacheTimeoutActive.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"cache.timeout.active", strconv.FormatInt(data.CacheTimeoutActive.ValueInt64(), 10))
	}
	if !data.Record.IsNull() && !data.Record.IsUnknown() {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"record.type", data.Record.ValueString())
	}
	if len(data.Exporters) > 0 {
		body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exporter", []interface{}{})
		for index, item := range data.Exporters {
			if !item.Name.IsNull() && !item.Name.IsUnknown() {
				body, _ = sjson.Set(body, helpers.LastElement(data.getPath())+"."+"exporter"+"."+strconv.Itoa(index)+"."+"name", item.Name.ValueString())
			}
		}
	}
	return body
}

func (data *FlowMonitor) updateFromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(prefix + "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Exporters {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Exporters[i].Name.ValueString()}

		var r gjson.Result
		res.Get(prefix + "exporter").ForEach(
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
		if value := r.Get("name"); value.Exists() && !data.Exporters[i].Name.IsNull() {
			data.Exporters[i].Name = types.StringValue(value.String())
		} else {
			data.Exporters[i].Name = types.StringNull()
		}
	}
	if value := res.Get(prefix + "cache.timeout.active"); value.Exists() && !data.CacheTimeoutActive.IsNull() {
		data.CacheTimeoutActive = types.Int64Value(value.Int())
	} else {
		data.CacheTimeoutActive = types.Int64Null()
	}
	if value := res.Get(prefix + "record.type"); value.Exists() && !data.Record.IsNull() {
		data.Record = types.StringValue(value.String())
	} else {
		data.Record = types.StringNull()
	}
}

func (data *FlowMonitorData) fromBody(ctx context.Context, res gjson.Result) {
	prefix := helpers.LastElement(data.getPath()) + "."
	if res.Get(helpers.LastElement(data.getPath())).IsArray() {
		prefix += "0."
	}
	if value := res.Get(prefix + "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := res.Get(prefix + "exporter"); value.Exists() {
		data.Exporters = make([]FlowMonitorExporters, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FlowMonitorExporters{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			}
			data.Exporters = append(data.Exporters, item)
			return true
		})
	}
	if value := res.Get(prefix + "cache.timeout.active"); value.Exists() {
		data.CacheTimeoutActive = types.Int64Value(value.Int())
	}
	if value := res.Get(prefix + "record.type"); value.Exists() {
		data.Record = types.StringValue(value.String())
	}
}

func (data *FlowMonitor) getDeletedItems(ctx context.Context, state FlowMonitor) []string {
	deletedItems := make([]string, 0)
	if !state.Description.IsNull() && data.Description.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/description", state.getPath()))
	}
	for i := range state.Exporters {
		stateKeyValues := [...]string{state.Exporters[i].Name.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Exporters[i].Name.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Exporters {
			found = true
			if state.Exporters[i].Name.ValueString() != data.Exporters[j].Name.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			deletedItems = append(deletedItems, fmt.Sprintf("%v/exporter=%v", state.getPath(), strings.Join(stateKeyValues[:], ",")))
		}
	}
	if !state.CacheTimeoutActive.IsNull() && data.CacheTimeoutActive.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/cache/timeout/active", state.getPath()))
	}
	if !state.Record.IsNull() && data.Record.IsNull() {
		deletedItems = append(deletedItems, fmt.Sprintf("%v/record/type", state.getPath()))
	}
	return deletedItems
}

func (data *FlowMonitor) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)

	return emptyLeafsDelete
}

func (data *FlowMonitor) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	for i := range data.Exporters {
		keyValues := [...]string{data.Exporters[i].Name.ValueString()}

		deletePaths = append(deletePaths, fmt.Sprintf("%v/exporter=%v", data.getPath(), strings.Join(keyValues[:], ",")))
	}
	if !data.CacheTimeoutActive.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/cache/timeout/active", data.getPath()))
	}
	if !data.Record.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/record/type", data.getPath()))
	}
	return deletePaths
}
