// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package retryutil

import (
	"fmt"
	"time"

	"openpitrix.io/logger"
)

func Retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		if sleep > 0 {
			time.Sleep(sleep)
		}

		logger.Warnf(nil, "Will retry %d because of error: %+v", i, err)
	}
	return fmt.Errorf("failed after %d attempts, error: %+v", attempts, err)
}
