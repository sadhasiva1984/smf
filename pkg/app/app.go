package app

import (
	smf_context "github.com/sadhasiva1984/smf/internal/context"
	"github.com/sadhasiva1984/smf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *smf_context.SMFContext
	Config() *factory.Config
}
