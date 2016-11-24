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
	"errors"
	"fmt"
	"strings"
)

const (
	// CompressionGzip will compress the squashfs with gzip
	CompressionGzip = "gzip"

	// CompressionXZ will compress the squashfs using xz
	CompressionXZ = "xz"
)

// SectionLiveOS is the Live ISO specific configuration
type SectionLiveOS struct {
	Compression  string `toml:"compression"`   // The type of compression to use on the LiveOS
	FileName     string `toml:"filename"`      // The resulting filename for this image spin
	RootfsSize   int    `toml:"rootfs_size"`   // Size of the image in megabytes (default 4000)
	RootfsFormat string `toml:"rootfs_format"` // Format of the rootfs, defaults to ext4
}

// ValidateSectionLiveOS will determine if the configuration is valid for a LiveOS
func ValidateSectionLiveOS(l *SectionLiveOS) error {
	switch l.Compression {
	case CompressionGzip, CompressionXZ:
	default:
		return fmt.Errorf("Unknown compression type: %v", l.Compression)
	}
	l.FileName = strings.TrimSpace(l.FileName)
	if l.FileName == "" {
		return errors.New("Invalid filename for livecd")
	}
	return nil
}
