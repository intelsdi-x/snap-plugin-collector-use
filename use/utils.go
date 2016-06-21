package use

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/host"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string
	var offset uint

	r := bufio.NewReader(f)
	n := -1
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		ret = append(ret, strings.Trim(line, "\n"))
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
	}

	return ret, nil

}

func run(cmd string, args []string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	var b bytes.Buffer
	command.Stdout = &b
	if err := command.Run(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func readInt(filename string) (int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// The int files that this is concerned with should only be one liners.
	line, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func hostTags() (map[string]string, error) {

	tags := make(map[string]string)

	hostInfo, err := host.Info()
	if err != nil {
		return tags, err
	}

	tags["hostname"] = hostInfo.Hostname
	tags["os"] = hostInfo.OS
	tags["platform"] = hostInfo.Platform
	tags["platform_family"] = hostInfo.PlatformFamily
	tags["platform_version"] = hostInfo.PlatformVersion
	tags["virtualization_role"] = hostInfo.VirtualizationRole
	tags["virtualization_system"] = hostInfo.VirtualizationSystem

	return tags, nil

}
