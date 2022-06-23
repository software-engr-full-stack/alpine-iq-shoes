package query

import (
    "testing"
)

func TestStartsWith(t *testing.T) {
    actualResults := StartsWith(startsWith, allIDs(), nil)

    title := "count of ids"
    if actual, expected := len(actualResults), len(expectedIds); actual != expected {
        t.Fatalf("%s, actual != expected, %v != %v", title, actual, expected)
    }

    for _, actualItem := range actualResults {
        if _, actual := expectedIds[actualItem]; actual != true {
            t.Fatalf("%#v not found in list of expected ids", actualItem)
        }
    }
}
