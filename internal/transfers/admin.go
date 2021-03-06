// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package transfers

import (
	"github.com/moov-io/base/admin"

	"github.com/go-kit/kit/log"
)

// RegisterAdminRoutes will add HTTP handlers for paygate's admin HTTP server
func RegisterAdminRoutes(logger log.Logger, svc *admin.Server, repo Repository) {
	svc.AddHandler("/transfers/{transferId}/status", updateTransferStatus(logger, repo))
}
