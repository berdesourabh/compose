/*
   Copyright 2020 Docker Compose CLI authors

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

package compose

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"gotest.tools/v3/assert"
)

func TestApplyScaleOpt(t *testing.T) {
	p := types.Project{
		Services: types.Services{
			{
				Name: "foo",
			},
			{
				Name: "bar",
			},
		},
	}
	opt := createOptions{scale: []string{"foo=2"}}
	err := opt.Apply(&p)
	assert.NilError(t, err)
	foo, err := p.GetService("foo")
	assert.NilError(t, err)
	assert.Equal(t, *foo.Deploy.Replicas, uint64(2))
}
