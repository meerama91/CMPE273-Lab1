package sleeps

import (
        "time"
        "testing"
)

func testsleeps(t *testing.T) {
    t1 := time.Now().Unix()
    sleeps(3)
    t2 := time.Now().Unix()

    if (t2 - t1) != 3 {
        t.Error("Incorrect sleep!")
    }
}
