// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package helper

import (
	"os"
	"strconv"
	"syscall"
)

type StateOS struct {
	Inode      uint64
	Device     uint64
	Size       int64
	ModifyTime uint64
}

// GetOSState returns the FileStateOS for non windows systems
func GetOSState(info os.FileInfo) StateOS {
	stat := info.Sys().(*syscall.Stat_t)

	// Convert inode and dev to uint64 to be cross platform compatible
	fileState := StateOS{
		Inode:      stat.Ino,
		Device:     stat.Dev,
		Size:       stat.Size,
		ModifyTime: uint64(stat.Mtim.Nano()),
	}
	return fileState
}

// IsSame file checks if the files are identical
func (fs StateOS) IsSame(state StateOS) bool {
	return fs.Inode == state.Inode && fs.Device == state.Device
}

// IsChange file checks if the files are changed
func (fs StateOS) IsChange(state StateOS) bool {
	if !fs.IsSame(state) {
		return true
	}
	return fs.Size != state.Size || fs.ModifyTime != state.ModifyTime
}

func (fs StateOS) IsEmpty() bool {
	return fs.Device == 0 && fs.Inode == 0
}

func (fs StateOS) IsFileChange(state StateOS) bool {
	return fs.Device != state.Device || fs.Inode != state.Inode
}

func (fs StateOS) String() string {
	var buf [64]byte
	current := strconv.AppendUint(buf[:0], fs.Inode, 10)
	current = append(current, '-')
	current = strconv.AppendUint(current, fs.Device, 10)
	current = append(current, '-')
	current = strconv.AppendInt(current, fs.Size, 10)
	return string(current)
}

// ReadOpen opens a file for reading only
func ReadOpen(path string) (*os.File, error) {
	flag := os.O_RDONLY
	perm := os.FileMode(0)
	return os.OpenFile(path, flag, perm) //nolint:gosec
}
