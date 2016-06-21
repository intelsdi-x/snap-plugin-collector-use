//
// +build unit

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package use

import (
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemoryUsePlugin(t *testing.T) {
	procPath = "proc"
	memInfoPath = filepath.Join(procPath, "meminfo")
	Convey("Read memory data should return proper MemInfo", t, func() {
		file, err := readStatForMemInfo()
		So(file, ShouldResemble, map[string]int64{"MemTotal": 16310784, "MemFree": 15344340})
		So(err, ShouldBeNil)

	})
	procPath = "/some/proc"
	memInfoPath = filepath.Join(procPath, "meminfo")
	Convey("Read memory data when file not available should return error", t, func() {
		file, err := readStatForMemInfo()
		So(file, ShouldResemble, map[string]int64{})
		So(err.Error(), ShouldResemble, "open /some/proc/meminfo: no such file or directory")

	})
	procPath = "proc"
	vmStatPath = filepath.Join(procPath, "vmstat")
	Convey("Read vm memory data should return proper MemInfo", t, func() {
		file, err := readStatForVMStat()
		So(file, ShouldResemble, map[string]int64{"SwapIn": 0, "SwapOut": 10.0})
		So(err, ShouldBeNil)

	})
	procPath = "/some/proc"
	vmStatPath = filepath.Join(procPath, "vmstat")
	Convey("Read vm memory data when file not available should return error", t, func() {
		file, err := readStatForVMStat()
		So(file, ShouldResemble, map[string]int64{})
		So(err.Error(), ShouldResemble, "open /some/proc/vmstat: no such file or directory")

	})
	procPath = "proc"
	vmStatPath = filepath.Join(procPath, "vmstat")
	memInfoPath = filepath.Join(procPath, "meminfo")
	Convey("get Utilization should return proper value", t, func() {
		m := MemInfo{}
		utilization, err := m.Utilization()
		So(utilization, ShouldResemble, 5.925184221678123)
		So(err, ShouldBeNil)
	})

	Convey("get Saturation should return proper value", t, func() {
		m := MemInfo{}
		saturation, err := m.Saturation()
		So(saturation, ShouldResemble, 0.0)
		So(err, ShouldBeNil)
	})
}
