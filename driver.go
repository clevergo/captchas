// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package captchas

// Driver defines how to generate captchas.
type Driver interface {
	// Generate generates a new captcha, returns an error if failed.
	Generate() (Captcha, error)
}
