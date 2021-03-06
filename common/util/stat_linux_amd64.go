/*
 * Copyright The Dragonfly Authors.
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

package util

import (
	"syscall"
	"time"
)

// Atime returns the last access time in time.Time.
func Atime(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}

// AtimeSec returns the last access time in seconds.
func AtimeSec(stat *syscall.Stat_t) int64 {
	return stat.Atim.Sec
}

// Ctime returns the create time in time.Time.
func Ctime(stat *syscall.Stat_t) time.Time {
	return time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
}

// CtimeSec returns the create time in seconds.
func CtimeSec(stat *syscall.Stat_t) int64 {
	return stat.Ctim.Sec
}
