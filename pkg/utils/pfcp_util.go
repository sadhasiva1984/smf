package utils

import (
	"context"
	"time"

	smf_context "github.com/sadhasiva1984/smf/internal/context"
	"github.com/sadhasiva1984/smf/internal/logger"
	"github.com/sadhasiva1984/smf/internal/pfcp"
	"github.com/sadhasiva1984/smf/internal/pfcp/udp"
	"github.com/sadhasiva1984/smf/pkg/service"
)

func InitPFCPFunc(pCtx context.Context) (func(a *service.SmfApp), func()) {
	smfContext := smf_context.GetSelf()

	pfcpStart := func(a *service.SmfApp) {
		// Initialize PFCP server
		smfContext.PfcpContext, smfContext.PfcpCancelFunc = context.WithCancel(pCtx)

		udp.Run(pfcp.Dispatch)

		// Wait for PFCP start
		time.Sleep(1000 * time.Millisecond)

		for _, upNode := range smf_context.GetSelf().UserPlaneInformation.UPFs {
			go a.Processor().ToBeAssociatedWithUPF(smfContext.PfcpContext, upNode.UPF)
		}
	}

	pfcpStop := func() {
		smfContext.PfcpCancelFunc()
		err := udp.Server.Close()
		if err != nil {
			logger.Log.Errorf("udp server close failed %+v", err)
		}
	}

	return pfcpStart, pfcpStop
}
