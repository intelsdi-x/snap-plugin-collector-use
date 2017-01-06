# snap collector plugin - use

## Collected Metrics
This plugin has the ability to gather the following metrics:


Namespace | Data Type | Formula | Threshold |Description |
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