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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeFlowMonitor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "description", "My monitor"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "flow_monitor_exporter.0.name", "EXPORTER1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "cache_timeout_active", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxe_flow_monitor.test", "record", "FNF1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeFlowMonitorConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxeFlowMonitorConfig() string {
	config := `resource "iosxe_flow_monitor" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	name = "MON1"` + "\n"
	config += `	description = "My monitor"` + "\n"
	config += `	flow_monitor_exporter = [{` + "\n"
	config += `		name = "EXPORTER1"` + "\n"
	config += `	}]` + "\n"
	config += `	cache_timeout_active = 60` + "\n"
	config += `	record = "FNF1"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxe_flow_monitor" "test" {
			name = "MON1"
			depends_on = [iosxe_flow_monitor.test]
		}
	`
	return config
}
