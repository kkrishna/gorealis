/**
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package realis

import (
	"github.com/rdelval/gorealis/gen-go/apache/aurora"
)

type Container interface {
	Build() *aurora.Container
}

// TODO(rdelvalle): Implement Mesos container builder
type MesosContainer struct {
	container *aurora.MesosContainer
}

type DockerContainer struct {
	container *aurora.DockerContainer
}

func NewDockerContainer() DockerContainer {
	return DockerContainer{container: aurora.NewDockerContainer()}
}

func (c DockerContainer) Build() *aurora.Container {
	return &aurora.Container{Docker: c.container}
}

func (c DockerContainer) Image(image string) DockerContainer {
	c.container.Image = image
	return c
}

func (c DockerContainer) AddParameter(name, value string) DockerContainer {
	c.container.Parameters = append(c.container.Parameters, &aurora.DockerParameter{name, value})
	return c
}

