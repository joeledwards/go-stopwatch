package stopwatch

import (
  "testing"
  "fmt"
)

func TestAccumulation(t *testing.T) {
  watch := MakeStopwatch()

  t.Run("Elapsed should increase on each start/stop.", func(t *testing.T) {
    var lastNanos int64 = 0

    for i := 0; i < 10; i++ {
      newNanos := watch.Start().Stop().Elapsed().Nanoseconds()

      if newNanos <= lastNanos {
        t.Fail()
      }

      lastNanos = newNanos
    }
  })

}

func TestReset(t *testing.T) {
  watch := NewStopwatch()

  t.Run("Time should elapse on started watch.", func(t *testing.T) {
    watch.Start()
    watch.Stop()

    if (watch.Elapsed().Nanoseconds() == 0) {
      t.Fail()
    }
  })

  t.Run("No time should elapse on reset watch.", func(t *testing.T) {
    watch.Reset()

    if (watch.Elapsed().Nanoseconds() > 0) {
      t.Fail()
    }
  })

  fmt.Println(watch.String())
}

func TestAllocation(t *testing.T) {
  t.Run("Should work when allocated on the stack.", func(t *testing.T) {
    watch := MakeStopwatch()

    if (watch.Start().Stop().Elapsed().Nanoseconds() < 1) {
      t.Fail()
    }
  })

  t.Run("Should work when allocated on the heap.", func(t *testing.T) {
    watch := NewStopwatch()

    if (watch.Start().Stop().Elapsed().Nanoseconds() < 1) {
      t.Fail()
    }
  })
}

