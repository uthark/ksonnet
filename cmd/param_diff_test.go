// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"testing"

	"github.com/ksonnet/ksonnet/actions"
)

func Test_paramDiffCmd(t *testing.T) {
	cases := []cmdTestCase{
		{
			name:   "with a component",
			args:   []string{"param", "diff", "env1", "env2", "--component", "component-name"},
			action: actionParamDiff,
			expected: map[string]interface{}{
				actions.OptionApp:           ka,
				actions.OptionComponentName: "component-name",
				actions.OptionEnvName1:      "env1",
				actions.OptionEnvName2:      "env2",
			},
		},
		{
			name:   "without a component",
			args:   []string{"param", "diff", "env1", "env2"},
			action: actionParamDiff,
			expected: map[string]interface{}{
				actions.OptionApp:           ka,
				actions.OptionComponentName: "component-name",
				actions.OptionEnvName1:      "env1",
				actions.OptionEnvName2:      "env2",
			},
		},
	}

	runTestCmd(t, cases)
}
