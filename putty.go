package main

import (
    "fmt"
    "golang.org/x/sys/windows/registry"
)

func main() {
    key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions`, registry.ALL_ACCESS)

    defer key.Close()
    if exists {
        println(`[+] key is vaild`)
        keys, _ := key.ReadSubKeyNames(0)
	    for _, key_subkey := range keys {
	        k, e, _ := registry.CreateKey(registry.CURRENT_USER, `Software\SimonTatham\PuTTY\Sessions\` + key_subkey, registry.ALL_ACCESS)
	        if e {
	        	hostname, _, _ := k.GetStringValue(`HostName`)
	        	fmt.Println("[+] name: " + key_subkey + ", hostname: " + hostname)
	        }
    		
	    }
    } else {
        println(`[-] key is invaild`)
    }
}