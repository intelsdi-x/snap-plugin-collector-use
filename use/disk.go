package use

import (
	"regexp"
	"strings"
	"time"

	"fmt"
	"strconv"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/pkg/errors"
)

// DiskStat struct for storing disk metric Data
type DiskStat struct {
	last     int64
	current  int64
	diskName string
}

// Utilization returns utilization of Disk Device
func (d *DiskStat) Utilization() (float64, error) {
	var err error
	d.last, err = readStatForDisk(d.diskName, "timeio")
	if err != nil {
		return 0.0, err
	}
	time.Sleep(waitTime)
	d.current, err = readStatForDisk(d.diskName, "timeio")
	if err != nil {
		return 0.0, err
	}
	return float64(d.current-d.last) / 10.0, nil
}

// Saturation returns saturation of Disk Device
func (d *DiskStat) Saturation() (float64, error) {
	var err error
	d.last, err = readStatForDisk(d.diskName, "weightedtimeio")
	if err != nil {
		return 0.0, err
	}
	d.current, err = readStatForDisk(d.diskName, "weightedtimeio")
	if err != nil {
		return 0.0, err
	}
	// 10ms * 10 ticks
	return float64(d.current-d.last) / 100.0, nil
}

func getDiskMetricTypes() ([]plugin.MetricType, error) {
	var mts []plugin.MetricType

	for _, diskName := range listDisks() {
		for _, name := range metricLabels {
			mts = append(mts, plugin.MetricType{Namespace_: core.NewNamespace("intel", "use", "storage", diskName, name)})
		}

	}
	return mts, nil
}

func listDisks() []string {
	cmd := "lsblk"
	args := []string{"-d", "--noheadings", "--list", "-o", "NAME"}
	output, err := run(cmd, args)
	if err != nil {
		return []string{}
	}
	ret := strings.Split(string(output), "\n")
	return ret
}

func readStatForDisk(diskName string, statType string) (int64, error) {
	lines, err := readLines(diskStatPath)
	if err != nil {
		return 0, err
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		if diskName == fields[2] {
			switch statType {
			case "timeio":
				return strconv.ParseInt((fields[12]), 10, 64)
			case "weightedtimeio":
				return strconv.ParseInt((fields[13]), 10, 64)
			}
		}
	}

	return 0, fmt.Errorf("Can't find a disk %s.\n", diskName)
}

func (u *Use) diskStat(ns core.Namespace) (*plugin.MetricType, error) {
	diskName := ns.Strings()[3]
	switch {
	case regexp.MustCompile(`^/intel/use/storage/.*/utilization$`).MatchString(ns.String()):
		diskStat := DiskStat{diskName: diskName}
		metric, err := diskStat.Utilization()
		if err != nil {
			return nil, err
		}
		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      metric,
		}, nil
	case regexp.MustCompile(`^/intel/use/storage/.*/saturation$`).MatchString(ns.String()):
		diskStat := DiskStat{diskName: diskName}
		metric, err := diskStat.Saturation()
		if err != nil {
			return nil, err
		}

		return &plugin.MetricType{
			Namespace_: ns,
			Data_:      float64(metric),
		}, nil
	}

	return nil, errors.Errorf("Unknown error processing %v", ns)
}
