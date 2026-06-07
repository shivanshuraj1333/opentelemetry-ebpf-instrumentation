// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package timing // import "go.opentelemetry.io/obi/pkg/ebpf/timing"

import (
	"time"
)

// MonoTimeNow mirrors the darwin variant: BPF ktime is not available on
// Windows, so this is a best-effort monotonic-ish timestamp used only where the
// span model expects a nsec value. The Windows agent (ETW-based) does not
// compare against ktime_get_ns().
func MonoTimeNow() time.Duration {
	return time.Duration(time.Now().Nanosecond())
}
