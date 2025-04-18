// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package brokerapi

import (
	"reflect"

	"code.cloudfoundry.org/brokerapi/v13/domain"
)

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type Service = domain.Service

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServiceDashboardClient = domain.ServiceDashboardClient

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServicePlan = domain.ServicePlan

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServiceSchemas = domain.ServiceSchemas

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServiceInstanceSchema = domain.ServiceInstanceSchema

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServiceBindingSchema = domain.ServiceBindingSchema

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type Schema = domain.Schema

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServicePlanMetadata = domain.ServicePlanMetadata

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServicePlanCost = domain.ServicePlanCost

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type ServiceMetadata = domain.ServiceMetadata

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
func FreeValue(v bool) *bool {
	return domain.FreeValue(v)
}

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
func BindableValue(v bool) *bool {
	return domain.BindableValue(v)
}

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
type RequiredPermission = domain.RequiredPermission

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
const (
	PermissionRouteForwarding = domain.PermissionRouteForwarding
	PermissionSyslogDrain     = domain.PermissionSyslogDrain
	PermissionVolumeMount     = domain.PermissionVolumeMount
)

// Deprecated: Use code.cloudfoundry.org/brokerapi/v13/domain
func GetJsonNames(s reflect.Value) (res []string) {
	return domain.GetJsonNames(s)
}
