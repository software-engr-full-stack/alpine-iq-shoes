package query

import (
    "testing"
)

func BenchmarkPayloadConcurrent(b *testing.B) {
    for n := 0; n < b.N; n++ {
        functions := []func() []string{
            GetNikeProductIDs,
            GetAdidasProductIDs,
            GetPumaProductIDs,
        }

        payloads := []*Payload{}
        for _, f := range functions {
            p := &Payload{Function: PayloadFunctionSignature(f)}
            payloads = append(payloads, p)
        }

        payloadPool, err := NewPayloadPool(
            '1',
            &PayloadPoolOpt{
                DataFetchMaxConcurrent: len(functions),
            },
        )
        if err != nil {
            panic(err)
        }


        for _, p := range payloads {
            payloadPool.Run(p)
        }
        payloadPool.Wait()

        actualResults := []string{}
        for _, p := range payloads {
            actualResults = append(actualResults, p.Results...)
        }
    }
}

func BenchmarkPayloadNonConcurrent(b *testing.B) {
    for n := 0; n < b.N; n++ {
        StartsWith(startsWith, allIDs(), nil)
    }
}
