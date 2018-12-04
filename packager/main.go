/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	d := os.Args[1]

	packager, err := defaultPackager(d)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to initialize Packager: %s\n", err)
		os.Exit(101)
	}

	if err := packager.Create(); err != nil {
		packager.Logger.Info(err.Error())
		os.Exit(102)
	}
}
