/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package support

import (
	"os"
)

const (
	// The environment variable for namespace where ODH is installed to.
	odhNamespaceEnvVar = "ODH_NAMESPACE"
)

func GetOpenDataHubNamespace() string {
	return lookupEnvOrDefault(odhNamespaceEnvVar, "opendatahub")
}

func lookupEnvOrDefault(key, value string) string {
	if v, yes := os.LookupEnv(key); yes {
		return v
	}
	return value
}
