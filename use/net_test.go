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
	"os"
	"path/filepath"
	"testing"

	"github.com/shirou/gopsutil/net"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNetUsePlugin(t *testing.T) {

	Convey("Should return 1 when interface dosn't exist", t, func() {

		linkspeed := getLinkSpeedforIface("mock0")
		So(linkspeed, ShouldResemble, 1.0)

	})
	Convey("Should return return proper linkspeed when interface exist", t, func() {

		pwd, err := os.Getwd()
		So(err, ShouldBeNil)
		sysFsNetPath = filepath.Join(pwd, "sys", "class", "net")
		linkspeed := getLinkSpeedforIface("eth0")
		So(linkspeed, ShouldResemble, 125000.0)

	})
	Convey("Should return an error when interface doesn't exist", t, func() {

		stats, err := getNicStatistic("mock0")
		So(stats, ShouldResemble, net.IOCountersStat{Name: "", BytesSent: 0x0, BytesRecv: 0x0, PacketsSent: 0x0, PacketsRecv: 0x0, Errin: 0x0, Errout: 0x0, Dropin: 0x0, Dropout: 0x0, Fifoin: 0x0, Fifoout: 0x0})
		So(err, ShouldNotBeNil)

	})

}
