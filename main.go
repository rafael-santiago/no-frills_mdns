// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under BSD-style license found in the
// LICENSE file in the root directory of this source tree.
package main

// INFO(Rafael): Any doubt take a look into 'internal/mdns/mdns.go'
//               It is BSD 3-Clause License, if you want use it feel free.
//               If you want to give me some credit give me it. If you want
//               to hide my work, to look like as you make this entire shit,
//               I don't give you any shit, too. I will continue being able to
//               do my stuff without you ha-ha-ha! But what about you without
//               me? Huh?! Opensource must be a two-way road (just saying)...

import (
    "internal/mdns"
    "os"
    "time"
    "fmt"
)

func main() {
    MDNSHosts := make([]mdns.MDNSHost, 0)
    MDNSHosts = append(MDNSHosts,
                       mdns.MDNSHost {
                            "deepthrought.local",
                            []byte { 42, 42, 42, 42 },
                            600,
                       },
                )
    MDNSHosts = append(MDNSHosts,
                       mdns.MDNSHost {
                            "hal9000.local",
                            []byte { 9, 0, 0, 0 },
                            9000,
                       },
                )
    goinHome := make(chan bool)
    err := mdns.MDNSServerStart(MDNSHosts, goinHome)
    if err != nil {
        fmt.Fprint(os.Stderr, "%s", err.Error())
        os.Exit(1)
    }
    fmt.Print("Delivering resolutions for:")
    str := []string{ ".\n", "," }
    golangSometimesEhMeioVirjona := func(eval bool) int {
        if eval {
            return 1
        }
        return 0
    }
    for h, hosts := range MDNSHosts {
        fmt.Printf(" '%s'%s", hosts.Name,
                   str[golangSometimesEhMeioVirjona((h+1) != len(MDNSHosts))])
    }
    time.Sleep(2 * time.Minute)
    goinHome <- true
    fmt.Println("That's all folks!")
    os.Exit(0)
}
