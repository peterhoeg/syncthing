// Copyright (C) 2014 Jakob Borg and other contributors. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"os"
	"strings"
)

var (
	debugNet = strings.Contains(os.Getenv("STTRACE"), "net") || os.Getenv("STTRACE") == "all"
)
