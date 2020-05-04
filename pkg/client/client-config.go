/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package client

import "fmt"

// Toggle HTTP output
var verbose bool

func SetVerbosity(verbose bool) {
	verbose = verbose
}

func Verbose(msg string) {
	if verbose {
		fmt.Println(fmt.Sprintf("[HTTP]: %s", msg))
	}
}

var GlobalRetriesPolicy Retries

func SetGlobalRetries(retries Retries) {
	GlobalRetriesPolicy = retries
}

type Retries struct {
	Timeout       int
	CustomMessage map[string]int
}
