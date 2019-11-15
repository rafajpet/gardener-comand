## gardener-command

Simple util for control iotivity-lite device [gardener](https://github.com/rafajpet/gardener)

## Device type and resource type

 - device: oic.d.switch.device
 - resource: core.switch.1

## Build
    
    ```bash
    dep ensure -v --vendor-only
    go build
    ```
    
## Get info about switch
    
    ```bash
    ./gardener-comand 
    ```
    
## Turn on/off switch

    ```bash
    ./gardener-comand --on
    ./gardener-comand --off
    ```
    
