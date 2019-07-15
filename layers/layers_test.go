/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package layers_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"

	layersBp "github.com/buildpack/libbuildpack/layers"
	"github.com/fatih/color"
	"github.com/heroku/libhkbuildpack/buildpack"
	"github.com/heroku/libhkbuildpack/layers"
	"github.com/heroku/libhkbuildpack/logger"
	"github.com/heroku/libhkbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestLayers(t *testing.T) {
	spec.Run(t, "Layers", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var (
			root string
			info bytes.Buffer
			l    layers.Layers
		)

		it.Before(func() {
			root = test.ScratchDir(t, "layers")
			logger := logger.NewFromWriters(nil, &info)
			l = layers.NewLayers(layersBp.Layers{Root: root}, layersBp.Layers{}, buildpack.Buildpack{}, logger)
		})

		it("logs process types", func() {
			g.Expect(l.WriteApplicationMetadata(layers.Metadata{
				Processes: []layers.Process{
					{"short", "test-command-1"},
					{"a-very-long-type", "test-command-2"},
				},
			})).To(Succeed())
			g.Expect(info.String()).To(Equal("-----> Process types:\n      a-very-long-type: test-command-2\n      short:            test-command-1\n"))
		})

		it("logs number of slices", func() {
			g.Expect(l.WriteApplicationMetadata(layers.Metadata{
				Slices: layers.Slices{
					layers.Slice{},
					layers.Slice{},
				},
			})).To(Succeed())

			g.Expect(info.String()).To(Equal(fmt.Sprintf(`%s 2 application slices
`, color.New(color.FgRed, color.Bold).Sprint("----->"))))
		})

		it("registers touched layers", func() {
			test.TouchFile(t, l.Root, "test-layer-1.toml")
			test.TouchFile(t, l.Root, "test-layer-2.toml")

			g.Expect(l.Layer("test-layer-1").Contribute(nil, func(layer layers.Layer) error {
				return nil
			})).To(Succeed())

			g.Expect(l.TouchedLayers.Cleanup()).To(Succeed())
			g.Expect(filepath.Join(l.Root, "test-layer-1.toml")).To(BeAnExistingFile())
			g.Expect(filepath.Join(l.Root, "test-layer-2.toml")).NotTo(BeAnExistingFile())
		})
	}, spec.Report(report.Terminal{}))
}
