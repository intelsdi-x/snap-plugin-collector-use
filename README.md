[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-use.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-use )
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-use)](http://goreportcard.com/report/intelsdi-x/snap-plugin-collector-use)
# snap collector plugin - use

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started


### System Requirements
* Plugin supports only Linux systems

### Installation
#### Download use plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [Github Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-use

Clone repo into `$GOPATH/src/github/intelsdi-x/`:
```
$ git clone https://github.com/<yourGithubID>/intelsdi-x/snap-plugin-collector-use
```
Build the plugin by running make in repo:
```
$ make
```
This builds the plugin in `/build/rootfs`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

### Examples
Example running use plugin, passthru processor, and writing data into an csv file.

Documentation for snap file publisher can be found [here](https://github.com/intelsdi-x/snap)

In one terminal window, open the snap daemon :
```
$ snapd -t 0 -l 1
```
The option "-l 1" it is for setting the debugging log level and "-t 0" is for disabling plugin signing.

In another terminal window:

Load collector and processor plugins
```
$ snapctl plugin load $SNAP_USE_PLUGIN/build/rootfs/snap-plugin-collector-use
$ snapctl plugin load $SNAP_PATH/build/plugin/snap-plugin-publisher-file
$ snapctl plugin load $SNAP_PATH/build/plugin/snap-plugin-processor-passthru
```

See available metrics for your system
```
$ snapctl metric list
```

Create a task file. For example, sample-use-task.json:

```
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/use/compute/utilization": {},
                "/intel/use/compute/saturation": {},
                "/intel/use/memory/utilization": {},
                "/intel/use/memory/saturation": {},
                "/intel/use/network/eth0/utilization": {},
                "/intel/use/network/eth0/saturation": {},
                "/intel/use/storage/sda/utilization": {},
                "/intel/use/storage/sda/saturation": {}
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```

## Documentation

The Utilization Saturation and Errors (USE) Method is a methodology for analyzing the performance of any system. It directs the construction of a checklist, which for server analysis can be used for quickly identifying resource bottlenecks or errors. It begins by posing questions, and then seeks answers, instead of beginning with given metrics (partial answers) and trying to work backwards (1). Brendan D. Gregg is an author of USE methodology.


### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Formula | |Description
----------|-----------|-----------|-----------|-----------|
/intel/use/compute/utilization | float64| 100 -idle | Normalized over cores 0 - 100 % | Compute utilization
/intel/use/compute/saturation | float64| load1/nr of cpus | Not normalized 0 - 100 % | Compute saturation
/intel/use/storage/{device_name}/utilization| float64| iostat % util | 0 - max %| Storage utilization
/intel/use/storage/{device_name}/saturation| float64| iostat avg-queue-size | 0 - max % | Storage utilization
/intel/use/storage/{device_name}/errors| float64| /sys/devices/.../ioerr_cnt | 0 - max %  | Storage errors
/intel/use/memory/utilization | float64| main_memory - memory_used | 0 - 100% | Memory utilization
/intel/use/memory/saturation | float64| memstat si/ memstat so | 0 - max %  | Memory saturation
/intel/use/network/{device_name}/utilization| float64| (tx + rcv bytes)/ bandwith % | 0 - 100% | Network device Utilization
/intel/use/network/{device_name}/saturation| float64| (tx + rcv overrun) - # of pkts % | 0 - max % | Network device Utilization


### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-use/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-use/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-use/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.

(1) http://www.brendangregg.com/usemethod.html