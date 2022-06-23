package query

import (
    "time"
    "fmt"
)

const startsWith = '1'

var expectedIds = map[string]bool{
    "10_nike": true,
    "10_adidas": true,
    "15_adidas": true,
    "10_puma": true,
    "12_puma": true,
    "14_puma": true,
    "16_puma": true,
    "18_puma": true,
}

func allIDs() []string {
    all := GetNikeProductIDs()
    all = append(all, GetAdidasProductIDs()...)
    all = append(all, GetPumaProductIDs()...)

    return all
}

func GetNikeProductIDs() (out []string) {
    for i := 0; i < 100; i += 10 {
        // Sleeping to simulate real world delays
        time.Sleep(time.Duration(1) * time.Second)
        out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
    }

    return out
}

func GetAdidasProductIDs() (out []string) {
    for i := 0; i < 100; i += 5 {
        out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
    }

    // Sleeping to simulate real world delays
    time.Sleep(time.Duration(2) * time.Second)

    return out
}

func GetPumaProductIDs() (out []string) {
    for i := 0; i < 100; i += 2 {
        out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
    }

    // Sleeping to simulate real world delays
    time.Sleep(time.Duration(3) * time.Second)

    return out
}
