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

func TestDiskUsePlugin(t *testing.T) {
	procPath = "proc"
	diskStatPath = filepath.Join(procPath, "diskstats")
	Convey("Read disk data should return proper timeio", t, func() {
		file, err := readStatForDisk("sda", "timeio")
		So(err, ShouldBeNil)
		var expectedValue int64
		expectedValue = 208788
		So(file, ShouldResemble, expectedValue)
		So(err, ShouldBeNil)

	})
	Convey("Read disk data should return proper weightedtimeio", t, func() {
		file, err := readStatForDisk("sda", "weightedtimeio")
		So(err, ShouldBeNil)
		var expectedValue int64
		expectedValue = 10474225
		So(file, ShouldResemble, expectedValue)
		So(err, ShouldBeNil)

	})
	procPath = "/some/proc"
	diskStatPath = filepath.Join(procPath, "diskstats")
	Convey("Read memory data when file not available should return error", t, func() {
		file, err := readStatForDisk("sda", "timeio")
		var expectedValue int64
		So(file, ShouldResemble, expectedValue)
		So(err.Error(), ShouldResemble, "open /some/proc/diskstats: no such file or directory")

	})

	procPath = "proc"
	diskStatPath = filepath.Join(procPath, "diskstats")
	Convey("get Utilization should return proper value", t, func() {
		d := DiskStat{diskName: "sda"}
		utilization, err := d.Utilization()
		So(utilization, ShouldResemble, 0.0)
		So(err, ShouldBeNil)
	})

	Convey("get Saturation should return proper value", t, func() {
		d := DiskStat{diskName: "sda"}
		saturation, err := d.Saturation()
		So(saturation, ShouldResemble, 0.0)
		So(err, ShouldBeNil)
	})
}
