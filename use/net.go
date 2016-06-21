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
	"fmt"
	"regexp"
	"time"

	"github.com/pkg/errors"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/shirou/gopsutil/net"
)

// NetStat contains values of Network previous measurments
type NetStat struct {
	last      net.IOCountersStat
	current   net.IOCountersStat
	ifaceName string
}

// Utilization returns utilization of Memory
func (n *NetStat) Utilization() (float64, error) {
	var err error
	linkSpeed := getLinkSpeedforIface(n.ifaceName)

	n.last, err = getNicStatistic(n.ifaceName)
	if err != nil {
		return 0.0, err
	}
	time.Sleep(waitTime)
	n.current, err = getNicStatistic(n.ifaceName)
	if err != nil {
		return 0.0, err
	}

	value := (n.current.BytesSent - n.last.BytesSent) + (n.current.BytesRecv - n.last.BytesRecv)
	return float64(value) / float64(linkSpeed), nil

}

// Saturation returns saturation of Memory
func (n *NetStat) Saturation() (float64, error) {
	var err error
	linkSpeed := getLinkSpeedforIface(n.ifaceName)

	n.last, err = getNicStatistic(n.ifaceName)
	if err != nil {
		return 0.0, err
	}
	time.Sleep(waitTime)
	n.current, err = getNicStatistic(n.ifaceName)
	if err != nil {
		return 0.0, err
	}

	value := (n.current.Fifoout - n.last.Fifoout) + (n.current.Fifoin - n.last.Fifoin)
	return float64(value) / float64(linkSpeed), nil

}

func (u *Use) networkStat(ns core.Namespace) (*plugin.MetricType, error) {
	ifaceName := ns.Strings()[3]

	switch {
	case regexp.MustCompile(`^/intel/use/network/.*/utilization$`).MatchString(ns.String()):

		netStat := NetStat{ifaceName: ifaceName}
		metric, err := netStat.Utilization()
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      metric,
		}, nil

	case regexp.MustCompile(`^/intel/use/network/.*/saturation$`).MatchString(ns.String()):
		netStat := NetStat{ifaceName: ifaceName}
		metric, err := netStat.Saturation()
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      metric,
		}, nil
	}

	return nil, errors.Errorf("Unknown error processing %v", ns)
}

func getNicStatistic(iface string) (net.IOCountersStat, error) {
	nets, err := net.IOCounters(true)
	if err != nil {
		return net.IOCountersStat{}, err

	}

	for _, net := range nets {
		if net.Name == iface {
			return net, nil
		}
	}
	return net.IOCountersStat{}, errors.New("Can't find interface")
}

func getNetIOCounterMetricTypes() ([]plugin.MetricType, error) {
	var mts []plugin.MetricType
	//per nic
	nets, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	for _, net := range nets {
		for _, name := range metricLabels {
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("intel", "use", "network", net.Name, name)})
		}
	}
	return mts, nil
}
func getLinkSpeedforIface(iface string) float64 {
	path := fmt.Sprintf("%s/%s/speed", sysFsNetPath, iface)

	speed, err := readInt(path)
	if err != nil {
		return 1
	}
	//Fixed value for 10ms
	return float64(speed) * 12.5
}
