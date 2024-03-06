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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-restconf"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FlowRecordDataSource{}
	_ datasource.DataSourceWithConfigure = &FlowRecordDataSource{}
)

func NewFlowRecordDataSource() datasource.DataSource {
	return &FlowRecordDataSource{}
}

type FlowRecordDataSource struct {
	clients map[string]*restconf.Client
}

func (d *FlowRecordDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flow_record"
}

func (d *FlowRecordDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Flow Record configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the retrieved object.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Provide a description for this Flow Record",
				Computed:            true,
			},
			"match_ipv4_source_address": schema.BoolAttribute{
				MarkdownDescription: "IPv4 source address",
				Computed:            true,
			},
			"match_ipv4_destination_address": schema.BoolAttribute{
				MarkdownDescription: "IPv4 destination address",
				Computed:            true,
			},
			"match_ipv4_protocol": schema.BoolAttribute{
				MarkdownDescription: "IPv4 protocol",
				Computed:            true,
			},
			"match_ipv4_tos": schema.BoolAttribute{
				MarkdownDescription: "IPv4 type of service",
				Computed:            true,
			},
			"match_transport_source_port": schema.BoolAttribute{
				MarkdownDescription: "Transport source port",
				Computed:            true,
			},
			"match_transport_destination_port": schema.BoolAttribute{
				MarkdownDescription: "Transport destination port",
				Computed:            true,
			},
			"match_interface_input": schema.BoolAttribute{
				MarkdownDescription: "The input interface",
				Computed:            true,
			},
			"match_flow_direction": schema.BoolAttribute{
				MarkdownDescription: "Direction the flow was monitored in",
				Computed:            true,
			},
			"collect_interface_output": schema.BoolAttribute{
				MarkdownDescription: "The output interface",
				Computed:            true,
			},
			"collect_counter_bytes_long": schema.BoolAttribute{
				MarkdownDescription: "Total number of bytes (64 bit counter)",
				Computed:            true,
			},
			"collect_counter_packets_long": schema.BoolAttribute{
				MarkdownDescription: "Total number of packets (64 bit counter)",
				Computed:            true,
			},
			"collect_transport_tcp_flags": schema.BoolAttribute{
				MarkdownDescription: "TCP flags",
				Computed:            true,
			},
			"collect_timestamp_absolute_first": schema.BoolAttribute{
				MarkdownDescription: "Absolute time the first packet was seen (milliseconds)",
				Computed:            true,
			},
			"collect_timestamp_absolute_last": schema.BoolAttribute{
				MarkdownDescription: "Absolute time the most recent packet was seen (milliseconds)",
				Computed:            true,
			},
		},
	}
}

func (d *FlowRecordDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *FlowRecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FlowRecordData

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, ok := d.clients[config.Device.ValueString()]; !ok {
		resp.Diagnostics.AddAttributeError(path.Root("device"), "Invalid device", fmt.Sprintf("Device '%s' does not exist in provider configuration.", config.Device.ValueString()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	res, err := d.clients[config.Device.ValueString()].GetData(config.getPath())
	if res.StatusCode == 404 {
		config = FlowRecordData{Device: config.Device}
	} else {
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		config.fromBody(ctx, res.Res)
	}

	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
