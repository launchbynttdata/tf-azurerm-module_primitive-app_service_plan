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

variable "name" {
  description = "Name of the service plan to create"
  type        = string
}

variable "resource_group_name" {
  description = "name of the resource group where the function app will be created"
  type        = string
}

variable "location" {
  description = "Location where the function app will be created"
  type        = string
}

variable "os_type" {
  description = "The operating system type of the service plan"
  type        = string

  validation {
    condition     = contains(["Windows", "Linux", "WindowsContainer"], var.os_type)
    error_message = "os_type must be one of 'Windows', 'WindowsContainer', or 'Linux'"
  }
}

variable "sku_name" {
  description = "The SKU of the service plan"
  type        = string
}

variable "maximum_elastic_worker_count" {
  description = "The maximum number of workers that the plan can scale out to"
  type        = number
  default     = null
}

variable "worker_count" {
  description = "The number of workers to be allocated"
  type        = number
  default     = null
}

variable "per_site_scaling_enabled" {
  description = "Toggle for per-site scaling. Disabled by default"
  type        = bool
  default     = null
}

variable "zone_balancing_enabled" {
  description = "Toggle for zone balancing. Disabled by default"
  type        = bool
  default     = null
}

variable "tags" {
  description = "A mapping of tags to assign to the resource"
  type        = map(string)
  default     = {}
}
