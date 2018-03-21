package main

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "github.com/Tomasen/realip"
    "net"
    "errors"
)

func externalIP() (string, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }
        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue // not an ipv4 address
            }
            return ip.String(), nil
        }
    }
    return "", errors.New("are you connected to the network?")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    ip, err := externalIP()
    if err != nil {
        fmt.Fprintf(os.Stdout, "error: %s\n", err)
    }
    hostname, _ := os.Hostname()
    fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            reqOrigin := realip.FromRequest(r)
            fmt.Fprintf(w, "I'm %s\nMy IP is %s\n", hostname, ip)
            fmt.Fprintf(os.Stdout, "- Request from %s\n", reqOrigin)
        }
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}