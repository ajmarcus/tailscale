// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !linux

package pidowner

func ownerOfPID(pid int) (userID string, err error) { return "", ErrNotImplemented }
