[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-use.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-use )
[![Go Report Card](https://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-use)](https://goreportcard.com/report/intelsdi-x/snap-plugin-collector-use)

# Snap collector plugin - use

This plugin supports collecting utilization, saturation and error metrics from GNU/Linux OS. 

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
* [golang 1.7+](https://golang.org/dl/)  - needed only for building
* This Plugin compatible with kernel > 2.6
* Linux/x86_64

### Operating systems
* Linux/amd64


### Installation

#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-collector-use/releases) page. For Snap, check [here](https://github.com/intelsdi-x/snap/releases).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-use
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-use.git
```

Build the snap use plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage

Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started).

Load the plugin and create a task, see example in [Examples](#examples).

## Documentation

The Utilization Saturation and Errors (USE) Method is a methodology for analyzing the performance of any system. It directs the construction of a checklist, which for server analysis can be used for quickly identifying resource bottlenecks or errors. It begins by posing questions, and then seeks answers, instead of beginning with given metrics (partial answers) and trying to work backwards (1). Brendan D. Gregg is an author of USE methodology.

### Collected Metrics

List of collected metrics is described in [METRICS.md](METRICS.md).

### Examples 

Example of running snap use collector and writing data to file.

Ensure [snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapteld -l 1 -t 0 &`

Download and load snap plugins:

```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-use/latest/linux/x86_64/snap-plugin-collector-use
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-use
$ snaptel plugin load snap-plugin-publisher-file
```

Create a task manifest file  (exemplary files in [examples/tasks/] (examples/tasks/)):
```yaml
--- 
schedule: 
  interval: 1s
  type: simple
version: 1
workflow: 
  collect: 
    metrics: 
      /intel/use/compute/saturation: {}
      /intel/use/compute/utilization: {}
      /intel/use/memory/saturation: {}
      /intel/use/memory/utilization: {}
      /intel/use/network/eth0/saturation: {}
      /intel/use/network/eth0/utilization: {}
      /intel/use/storage/sda/saturation: {}
      /intel/use/storage/sda/utilization: {}
    publish: 
      - 
        plugin_name: file
        config: 
          file: /tmp/use_metrics

```
Download an [example task file](https://github.com/intelsdi-x/snap-plugin-collector-use/blob/master/examples/tasks/) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-use/master/examples/tasks/use-file.yml
$ snaptel task create -t use-file.yml
Using task manifest to create task
Task created
ID: 480323af-15b0-4af8-a526-eb2ca6d8ae67
Name: Task-480323af-15b0-4af8-a526-eb2ca6d8ae67
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
$ snaptel task watch 480323af-15b0-4af8-a526-eb2ca6d8ae67
```

This data is published to a file `/tmp/use_metrics` per task specification

Stop task:
```
$ snaptel task stop 480323af-15b0-4af8-a526-eb2ca6d8ae67
Task stopped:
ID: 480323af-15b0-4af8-a526-eb2ca6d8ae67
```

## Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-use/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-use/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack](http://slack.snap-telemetry.io).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.

(1) http://www.brendangregg.com/usemethod.html

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
