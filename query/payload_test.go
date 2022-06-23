package query

import (
    "testing"
)

func TestPayload(t *testing.T) {
    functions := []func() []string{
        GetNikeProductIDs,
        GetAdidasProductIDs,
        GetPumaProductIDs,
    }

    payloadPool, err := NewPayloadPool(
        '1',
        &PayloadPoolOpt{
            DataFetchMaxConcurrent: len(functions),
        },
    )
    title := "NewPayloadPool returned error value"
    if actual, expected := err, error(nil); actual != nil {
        t.Fatalf("%s, actual != expected, %v != %v", title, actual, expected)
    }

    payloads := []*Payload{}
    for _, f := range functions {
        p := &Payload{Function: PayloadFunctionSignature(f)}
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

    title = "count of ids"
    if actual, expected := len(actualResults), len(expectedIds); actual != expected {
        t.Fatalf("%s, actual != expected, %v != %v", title, actual, expected)
    }

    for _, actualItem := range actualResults {
        if _, actual := expectedIds[actualItem]; actual != true {
            t.Fatalf("%#v not found in list of expected ids", actualItem)
        }
    }
}
