//
// +build small

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
	"regexp"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsePlugin(t *testing.T) {
	Convey("Meta should return metadata for the plugin", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, name)
		So(meta.Version, ShouldResemble, version)
		So(meta.Type, ShouldResemble, plugin.CollectorPluginType)
	})

	Convey("Create Use Collector", t, func() {
		useCol := NewUseCollector()
		Convey("So NewUseCollector should not be nil", func() {
			So(useCol, ShouldNotBeNil)
		})
		Convey("So NewUseCollector should be of Use type", func() {
			So(useCol, ShouldHaveSameTypeAs, &Use{})
		})
		Convey("useCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := useCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
	})

	Convey("Get Metrics ", t, func() {
		useCol := NewUseCollector()
		var cfg = plugin.ConfigType{}

		Convey("So should return 26 types of metrics", func() {
			metrics, err := useCol.GetMetricTypes(cfg)
			So(len(metrics), ShouldBeGreaterThan, 13)
			So(err, ShouldBeNil)
		})
		Convey("So should check namespace", func() {
			metrics, err := useCol.GetMetricTypes(cfg)
			vcpuNamespace := metrics[0].Namespace().String()
			vcpu := regexp.MustCompile(`^/intel/use/compute/utilization`)
			So(true, ShouldEqual, vcpu.MatchString(vcpuNamespace))
			So(err, ShouldBeNil)

			vcpuNamespace1 := metrics[1].Namespace().String()
			vcpu1 := regexp.MustCompile(`^/intel/use/compute/saturation`)
			So(true, ShouldEqual, vcpu1.MatchString(vcpuNamespace1))
			So(err, ShouldBeNil)
		})

	})
	Convey("Collect Metrics", t, func() {
		useCol := &Use{}

		cfgNode := cdata.NewNode()

		pwd, err := os.Getwd()
		procPath = filepath.Join(pwd, "proc")
		diskStatPath = filepath.Join(procPath, "diskstats")
		cpuStatPath = filepath.Join(procPath, "stat")
		loadAvgPath = filepath.Join(procPath, "loadavg")
		memInfoPath = filepath.Join(procPath, "meminfo")
		vmStatPath = filepath.Join(procPath, "vmstat")
		So(err, ShouldBeNil)
		Convey("So should get memory saturation metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "memory", "saturation"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType float64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get memory utilization metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "memory", "utilization"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType float64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get compute utilization metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "compute", "utilization"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get compute saturation metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "compute", "saturation"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType float64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get disk utilization metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "storage", "sda", "utilization"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			So(len(collect), ShouldResemble, 1)
		})
		Convey("So should get disk saturation metrics", func() {
			metrics := []plugin.MetricType{{
				Namespace_: core.NewNamespace("intel", "use", "storage", "sda", "saturation"),
				Config_:    cfgNode,
			}}
			collect, err := useCol.CollectMetrics(metrics)
			So(err, ShouldBeNil)
			So(collect[0].Data_, ShouldNotBeNil)
			var expectedType float64
			So(collect[0].Data_, ShouldHaveSameTypeAs, expectedType)
			So(len(collect), ShouldResemble, 1)
		})
	})
}
