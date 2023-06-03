# README

make `wget -O <file>.tar.gz https://github.com/<file>.tar.gz` faster or work in your server that don't have VPN or proxy support.


## install

```
curl -sf https://gobinaries.com/dfang/xwget | sh
```


```
alias wget='xwget'
```

## test 

`time xwget -O 'OpenJDK8U-jdk_x64_linux_hotspot_8u372b07.tar.gz\n' https://github.com/adoptium/temurin8-binaries/releases/download/jdk8u372-b07/OpenJDK8U-jdk_x64_linux_hotspot_8u372b07.tar.gz`


## thanks

https://ghproxy.com