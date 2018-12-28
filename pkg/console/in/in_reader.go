/*
 * Copyright 2018 Christian Bargmann <chris@cbrgm.net>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package in

import (
	"bytes"
	"io"
)

// LineReader is an unbuffered line reader
type LineReader struct {
	r io.Reader
}

// NewReader creates a new line reader
func NewReader(r io.Reader) *LineReader {
	return &LineReader{r: r}
}

// Read implements io.Reader
func (lr LineReader) Read(p []byte) (int, error) {
	return lr.r.Read(p)
}

// ReadLine reads one line without buffering
func (lr LineReader) ReadLine() (string, error) {
	out := &bytes.Buffer{}
	buf := make([]byte, 1)
	for {
		n, err := lr.r.Read(buf)
		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				return out.String(), nil
			}
			_ = out.WriteByte(buf[i])
		}

		if err != nil {
			if err == io.EOF {
				return out.String(), nil
			}
			return out.String(), err
		}
	}
}
