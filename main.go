package main

import (
    "fmt"
    "aplineiq.com/codetest/query"
)

const useTestFunctions = false

func main() {
    functions := []func() []string{
        GetNikeProductIDs,
        GetAdidasProductIDs,
        GetPumaProductIDs,
    }

    if useTestFunctions {
        functions = []func() []string{
            query.GetNikeProductIDs,
            query.GetAdidasProductIDs,
            query.GetPumaProductIDs,
        }
    }
    const queryMaxConcurrent = 10
    payloadPool, err := query.NewPayloadPool(
        '1',
        &query.PayloadPoolOpt{
            DataFetchMaxConcurrent: len(functions),

            QueryMaxConcurrent: queryMaxConcurrent,
        },
    )
    if err != nil {
        panic(err)
    }

    payloads := []*query.Payload{}
    for _, f := range functions {
        p := &query.Payload{Function: query.PayloadFunctionSignature(f)}
        payloads = append(payloads, p)
    }

    for _, p := range payloads {
        payloadPool.Run(p)
    }
    payloadPool.Wait()

    actualResults := []string{}
    for _, p := range payloads {
        actualResults = append(actualResults, p.Results...)
    }

    for _, actualItem := range actualResults {
        fmt.Println(actualItem)
    }
}

// func nonconc() {
//     all := query.GetNikeProductIDs()
//     all = append(all, query.GetAdidasProductIDs()...)
//     all = append(all, query.GetPumaProductIDs()...)

//     actualResults := query.StartsWith('1', all, nil)
//     for _, id := range actualResults {
//         fmt.Printf("%#v\n", id)
//     }
// }

func GetNikeProductIDs() (out []string) {
    for i := 0; i < 100; i += 10 {
        out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
    }
    return out
}

func GetAdidasProductIDs() (out []string) {
    for i := 0; i < 100; i += 5 {
        out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
    }
    return out
}

func GetPumaProductIDs() (out []string) {
    for i := 0; i < 100; i += 2 {
        out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
    }
    return out
}
