// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package captchas

import "context"

// Store defines how to save and load captcha information.
type Store interface {
	// Get returns the answer of the given captcha ID, returns
	// an error if failed. Clear indicates whether delete the
	// captcha after fetching.
	Get(ctx context.Context, id string, clear bool) (string, error)

	// Set saves the captcha ID and answer, returns error
	// if failed.
	Set(ctx context.Context, id, answer string) error
}
