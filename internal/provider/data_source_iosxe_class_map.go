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
	_ datasource.DataSource              = &ClassMapDataSource{}
	_ datasource.DataSourceWithConfigure = &ClassMapDataSource{}
)

func NewClassMapDataSource() datasource.DataSource {
	return &ClassMapDataSource{}
}

type ClassMapDataSource struct {
	clients map[string]*restconf.Client
}

func (d *ClassMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_class_map"
}

func (d *ClassMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Class Map configuration.",

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
				MarkdownDescription: "name of the class map",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "type of the class-map",
				Computed:            true,
			},
			"subscriber": schema.BoolAttribute{
				MarkdownDescription: "Domain name of the class map",
				Computed:            true,
			},
			"prematch": schema.StringAttribute{
				MarkdownDescription: "Logical-AND/Logical-OR of all matching statements under this class map",
				Computed:            true,
			},
			"match_authorization_status_authorized": schema.BoolAttribute{
				MarkdownDescription: "authorized",
				Computed:            true,
			},
			"match_result_type_aaa_timeout": schema.BoolAttribute{
				MarkdownDescription: "aaa timeout type",
				Computed:            true,
			},
			"match_authorization_status_unauthorized": schema.BoolAttribute{
				MarkdownDescription: "unauthorized",
				Computed:            true,
			},
			"match_activated_service_templates": schema.ListNestedAttribute{
				MarkdownDescription: "match name of service template activated on session",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_name": schema.StringAttribute{
							MarkdownDescription: "Enter service name",
							Computed:            true,
						},
					},
				},
			},
			"match_authorizing_method_priority_greater_than": schema.ListAttribute{
				MarkdownDescription: "greater than",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"match_method_dot1x": schema.BoolAttribute{
				MarkdownDescription: "dot1x",
				Computed:            true,
			},
			"match_result_type_method_dot1x_authoritative": schema.BoolAttribute{
				MarkdownDescription: "failure type",
				Computed:            true,
			},
			"match_result_type_method_dot1x_agent_not_found": schema.BoolAttribute{
				MarkdownDescription: "agent not found type",
				Computed:            true,
			},
			"match_result_type_method_dot1x_method_timeout": schema.BoolAttribute{
				MarkdownDescription: "method_timeout type",
				Computed:            true,
			},
			"match_method_mab": schema.BoolAttribute{
				MarkdownDescription: "mab",
				Computed:            true,
			},
			"match_result_type_method_mab_authoritative": schema.BoolAttribute{
				MarkdownDescription: "failure type",
				Computed:            true,
			},
			"match_dscp": schema.ListAttribute{
				MarkdownDescription: "Match DSCP in IP(v4) and IPv6 packets",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Class-Map description",
				Computed:            true,
			},
		},
	}
}

func (d *ClassMapDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*restconf.Client)
}

func (d *ClassMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ClassMapData

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
		config = ClassMapData{Device: config.Device}
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
