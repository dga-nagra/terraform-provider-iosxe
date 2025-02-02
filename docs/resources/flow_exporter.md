---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxe_flow_exporter Resource - terraform-provider-iosxe"
subcategory: "Flow"
description: |-
  This resource can manage the Flow Exporter configuration.
---

# iosxe_flow_exporter (Resource)

This resource can manage the Flow Exporter configuration.

## Example Usage

```terraform
resource "iosxe_flow_exporter" "example" {
  name                  = "EXPORTER1"
  description           = "My exporter"
  destination_ip        = "1.1.1.1"
  source_loopback       = 123
  transport_udp         = 655
  template_data_timeout = 60
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `description` (String) Provide a description for this Flow Exporter
- `destination_ip` (String)
- `device` (String) A device name from the provider configuration.
- `source_loopback` (Number) Loopback interface
  - Range: `0`-`2147483647`
- `template_data_timeout` (Number) Resend data based on a timeout
  - Range: `1`-`86400`
- `transport_udp` (Number) UDP transport protocol
  - Range: `0`-`65535`

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxe_flow_exporter.example "Cisco-IOS-XE-native:native/Cisco-IOS-XE-flow:flow/exporter=EXPORTER1"
```
