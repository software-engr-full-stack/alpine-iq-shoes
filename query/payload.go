package query

import (
    "sync"
    "fmt"
)

// "Payload..." satisfies the requirement "calling the Get functions ... concurrently"

type PayloadFunctionSignature func() []string

type Payload struct {
    Results []string
    Function PayloadFunctionSignature
}

func (p *Payload) StoreResults(results []string) {
    p.Results = results
}

type PayloadPool interface {
    Run(*Payload)
    Wait()
}

type PayloadPoolOpt struct {
    // The max number of goroutines to be used in executing each of the
    // functions that fetch the ids.
    DataFetchMaxConcurrent int

    // The max number of goroutines to be used in processing each of the ids.
    QueryMaxConcurrent int
}

// TODO: improve this interface. Probably something like
// passing an option struct containing the type of filter
// to use and the parameters for the filter.
func NewPayloadPool(startsWith rune, opt *PayloadPoolOpt) (PayloadPool, error) {
    var empty PayloadPool
    if opt.DataFetchMaxConcurrent < 1 {
        return empty, fmt.Errorf("the number of goroutines must be greater than zero")
    }
    var wg sync.WaitGroup

    pq := make(chan *Payload)

    for i := 0; i < opt.DataFetchMaxConcurrent; i++ {
        wg.Add(1)
        go func() {
            for p := range pq {
                results := p.Function()
                filtered := StartsWith(
                    startsWith,
                    results,
                    &StartsWithOpt{maxConcurrent: opt.QueryMaxConcurrent},
                )
                p.StoreResults(filtered)
            }
            defer wg.Done()
        }()
    }

    return &PayloadRunner{queue: pq, wg: &wg}, nil
}

type PayloadRunner struct {
    queue chan *Payload
    wg *sync.WaitGroup
}

func (cr *PayloadRunner) Run(p *Payload) {
    cr.queue <- p
}

func (cr *PayloadRunner) Wait() {
    close(cr.queue)
    cr.wg.Wait()
}
