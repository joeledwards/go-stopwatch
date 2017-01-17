package stopwatch

import (
  "fmt"
  "time"
)

type Stopwatch struct {
  running bool
  accumulated time.Duration
  whence time.Time
}

func (watch *Stopwatch) IsRunning() bool {
  return watch.running
}

func (watch *Stopwatch) Start() *Stopwatch {
  if !watch.running {
    watch.running = true
    watch.whence = time.Now()
  }

  return watch
}

func (watch *Stopwatch) Stop() *Stopwatch {
  if watch.running {
    watch.accumulated += time.Now().Sub(watch.whence)
    watch.running = false
  }

  return watch
}

func (watch *Stopwatch) Reset() *Stopwatch {
  if watch.running {
    watch.Stop()
  }

  watch.accumulated = 0
  watch.running = false

  return watch
}

func (watch *Stopwatch) Elapsed() time.Duration {
  if watch.running {
    return watch.accumulated + time.Now().Sub(watch.whence)
  } else {
    return watch.accumulated
  }
}

func (watch *Stopwatch) Format() string {
  elapsed := watch.Elapsed()
  hours := int64(elapsed.Hours())
  minutes := int64(elapsed.Minutes())
  seconds := int64(elapsed.Seconds())
  nanos := elapsed.Nanoseconds()
  micros := nanos / 1000
  millis := micros / 1000

  switch {
    case hours > 0:
      return fmt.Sprintf("%d h, %d min", hours, minutes)
    case minutes > 0:
      return fmt.Sprintf("%d min, %d s", minutes, seconds)
    case seconds > 0:
      return fmt.Sprintf("%d.%03d s", seconds, millis % 1000)
    case millis > 0:
      return fmt.Sprintf("%d.%03d ms", millis, micros % 1000)
    default:
      return fmt.Sprintf("%d.%03d us", micros, nanos % 1000)
  }
}

func (watch *Stopwatch) String() string {
  var state string = "stopped"

  if watch.running {
    state = "running"
  }

  return fmt.Sprintf("Stopwatch[%s : %s]", state, watch.Format())
}

func NewStopwatch() Stopwatch {
  return new(Stopwatch{false, 0, time.Now()})
}

func MakeStopwatch() Stopwatch {
  return Stopwatch{false, 0, time.Now()}
}

