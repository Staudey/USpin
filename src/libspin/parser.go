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

package libspin

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ImageSpecParser does the heavy lifting of parsing a .spin file to pull all
// relevant stack operations from it.
type ImageSpecParser struct {
	CommentCharacter   string
	RepoSplitCharacter string
}

// NewParser will return a new parser for the image specification file
func NewParser() *ImageSpecParser {
	return &ImageSpecParser{
		CommentCharacter:   "#",
		RepoSplitCharacter: "=",
	}
}

// Parse will attempt to parse the given image speicifcation file at the given
// path, and will return an error if this fails.
func (i *ImageSpecParser) Parse(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()
	sc := bufio.NewScanner(fi)

	lineno := 0

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		lineno++

		if line == "" {
			continue
		}

		// Check for single line comments
		if strings.HasPrefix(line, i.CommentCharacter) {
			continue
		}

		// Check if this is a repo
		if strings.Contains(line, i.RepoSplitCharacter) {
			fields := strings.Split(line, "=")
			value := strings.TrimSpace(strings.Join(fields[1:], "="))
			if value == "" {
				return fmt.Errorf("Missing value for repo declaration '%v' on line '%v'\n", fields[0], lineno)
			}
			// TODO: Add an OpRepo to the stack
			continue
		}

		fmt.Fprintf(os.Stderr, "Line: %s\n", line)
	}

	return errors.New("Not yet implemented!")
}
