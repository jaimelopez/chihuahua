![](https://image.ibb.co/i2vvmo/chihuahua.png)

Chihuahua is a command line tool which executes golang microbenchmarks and compares them with previous results to determine whether performance is good enough to deploy new changes in your application.

Results of previous executions can be stored with different drivers, for now `elasticsearch` or directly to `filesystem`.

# Installation
```go get -u -f github.com/jaimelopez/chihuahua```

# Usage
- `group`       group name of metrics to store
- `storage`     driver to store results (elastic/file)
- `destination` storage destination
- `fromfile`    take results to analyze from file instead running benchmarks
- `duration`    time to execute benchmarks (default 1s)
- `threshold`   threshold percent to determine whether performance is good enough (default 15)
- `save`        results will be saved if are higher than previous ones
- `force`       forces to save results even if they are worse
- `results`     prints results as table directly thru standard output
- `debug`       shows traces for debugging

## From file
There is a way in which you can specify the output (printed directly to a file) of executed benchmarks so you can take that data instead of executing them via `Chihuahua`.

Example:
```
cd application/benchmarks
go test -bench . -run NONE -benchmem > /var/tmp/current-results
chihuahua -group myapp -storage file -destination /var/tmp/benchs -fromfile /var/tmp/current-results -results -save 
```

## Filesystem storage
Destination should be a folder where results will be stored.  
A file per group will be generated inside folder.

`Notice that folder must exist.`

Example:
```
chihuahua -group myapp -storage file -destination /var/tmp/benchs -results -save 
```

## Elasticsearch storage
Elastic structure is a bit hardcoded, having 3 different indexes one per metric with the following names:
- `mygroup-ns`            with nano-seconds per operation metric
- `mygroup-mallocs`       with number of allocations per operation metric
- `mygroup-mallocbytes`   with total bytes allocated per operation

The structure of each document is:
```json
{
    "name-of-benchmark-1": "value",
    "name-of-benchmark-2": "value",
    "name-of-benchmark-3": "value",
    "@timestamp": "2018-07-12 07:07:22.439511876 +0000 UTC"
}
```

Example:
```
chihuahua -group servicebus -storage elastic -destination https://user:password@locahost:9243 -save
```

# Exit codes
Chihuahua was designed to integrate it within a continous integration tool so it returns different exit codes depending whether performance is good enough (or force is specified) or not.

`exit code -1` when error during executing  
`exit code 0` when everything is ok and performance is good  
`exit code 1` when performance is not good enough