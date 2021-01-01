package config

import (
  "golang.org/x/sys/unix"
  "os"
  "fmt"
  "bytes"
  "strings"
)

func detectKernelConfig() {
    // Check Kernel Config is available or not.
    // We are replicating BPFTools logic here to check if kernel config is available
    // https://elixir.bootlin.com/linux/v5.7/source/tools/bpf/bpftool/feature.c#L390
    info := unix.Utsname{}
    err := unix.Uname(&info)
    if err != nil {
       fmt.Println("unix.Uname error")
    }
	  
    release := strings.TrimSpace(string(bytes.Trim(info.Release[:], "\x00")))
   
   fmt.Println("release:"+release)

    // Any error checking these files will return Kernel config not found error
    if _, err := os.Stat(fmt.Sprintf("/boot/config-%s", release)); err != nil {
        if _, err = os.Stat("/proc/config.gz"); err != nil {
       	     fmt.Println("kernel config is not found")
        }
    }

    fmt.Println("kernel config has been found")
}

