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

package helper_test

import (
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/libhkbuildpack/helper"
	"github.com/cloudfoundry/libhkbuildpack/test"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestExtractZip(t *testing.T) {
	spec.Run(t, "ExtractZip", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var root string

		it.Before(func() {
			root = test.ScratchDir(t, "extract-zip")
		})

		it("extracts the archive", func() {
			g.Expect(helper.ExtractZip(filepath.Join("testdata", "test-archive.zip"), root, 0)).To(Succeed())
			g.Expect(filepath.Join(root, "fileA.txt")).To(BeARegularFile())
			g.Expect(filepath.Join(root, "dirA", "fileB.txt")).To(BeARegularFile())
			g.Expect(filepath.Join(root, "dirA", "fileC.txt")).To(BeARegularFile())
		})

		it("skips stripped components", func() {
			g.Expect(helper.ExtractZip(filepath.Join("testdata", "test-archive.zip"), root, 1)).To(Succeed())
			g.Expect(filepath.Join(root, "fileB.txt")).To(BeARegularFile())
			g.Expect(filepath.Join(root, "fileC.txt")).To(BeARegularFile())

		})

	}, spec.Report(report.Terminal{}))
}
