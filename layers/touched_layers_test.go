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
	"path/filepath"
	"testing"

	"github.com/heroku/libhkbuildpack/layers"
	"github.com/heroku/libhkbuildpack/logger"
	"github.com/heroku/libhkbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestTouchedLayers(t *testing.T) {
	spec.Run(t, "TouchedLayers", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var (
			root    string
			touched layers.TouchedLayers
		)

		it.Before(func() {
			root = test.ScratchDir(t, "touched-layers")
			touched = layers.NewTouchedLayers(root, logger.Logger{})
		})

		it("does not remove touched layers", func() {
			test.TouchFile(t, root, "test-layer.toml")

			touched.Add(filepath.Join(root, "test-layer.toml"))
			g.Expect(touched.Cleanup()).To(Succeed())

			g.Expect(filepath.Join(root, "test-layer.toml")).To(BeARegularFile())
		})

		it("removes untouched layers", func() {
			test.TouchFile(t, root, "test-layer.toml")

			g.Expect(touched.Cleanup()).To(Succeed())

			g.Expect(filepath.Join(root, "test-layer.toml")).NotTo(BeARegularFile())
		})

		it("does not remove app.toml", func() {
			test.TouchFile(t, root, "app.toml")

			g.Expect(touched.Cleanup()).To(Succeed())

			g.Expect(filepath.Join(root, "app.toml")).To(BeARegularFile())
		})

		it("does not remove launch.toml", func() {
			test.TouchFile(t, root, "launch.toml")

			g.Expect(touched.Cleanup()).To(Succeed())

			g.Expect(filepath.Join(root, "launch.toml")).To(BeARegularFile())
		})

		it("does not remove store.toml", func() {
			test.TouchFile(t, root, "store.toml")

			g.Expect(touched.Cleanup()).To(Succeed())

			g.Expect(filepath.Join(root, "store.toml")).To(BeARegularFile())
		})
	}, spec.Report(report.Terminal{}))
}
