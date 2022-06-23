package query

import (
    "sync"
)

type StartsWithOpt struct {
    maxConcurrent int
}

const defaultMaxConcurrent = 1000
func StartsWith(char rune, input []string, opt *StartsWithOpt) []string {
    if opt == nil {
        opt = &StartsWithOpt{maxConcurrent: defaultMaxConcurrent}
    }
    if opt.maxConcurrent < 1 {
        opt.maxConcurrent = defaultMaxConcurrent
    }

    var result []string

    if opt.maxConcurrent > 0 {
        return qcon(char, input, opt.maxConcurrent)
    }

    for _, item := range input {
        if isFirstRuneAMatch(char, item) {
            result = append(result, item)
        }
    }
    return result
}

func isFirstRuneAMatch(char rune, str string) bool {
    inp := []rune(str)
    return inp[0] == char
}

// This is the function that satisfies the requirement "inserting into "output" concurrently"
func qcon(startsWith rune, ids []string, maxConcurrent int) []string {
    var wg sync.WaitGroup

    in := make(chan string)

    wg.Add(1)
    go func() {
        for _, id := range ids {
            in <- id
        }
        close(in)
        wg.Done()
    }()

    out := make(chan string)

    for i := 0; i < maxConcurrent; i++ {
        wg.Add(1)
        go func() {
            for t := range in {
                if isFirstRuneAMatch(startsWith, t) {
                    out <- t
                }
            }
            wg.Done()
        }()
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    results := []string{}
    for t := range out {
        results = append(results, t)
    }

    return results
}
