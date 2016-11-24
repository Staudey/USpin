//
// Copyright © 2016 Ikey Doherty <ikey@solus-project.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package config

import (
	"testing"
)

const (
	confTestPath = "../../../testdata/minimal.spin"
)

func TestConfig(t *testing.T) {
	c, err := New(confTestPath)
	if err != nil {
		t.Fatalf("Couldn't open good config: %v", err)
	}
	if c.Image.Packages != "minimal.packages" {
		t.Fatalf("Invalid packages file")
	}
	if c.Image.Type != "livecd" {
		t.Fatalf("Invalid type")
	}
	if c.LiveOS.Compression != "gzip" {
		t.Fatalf("Invalid compression: %v", c.LiveOS.Compression)
	}
}
