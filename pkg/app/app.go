package app

import (
	nssf_context "github.com/nycu-ucr/nssf/internal/context"
	"github.com/nycu-ucr/nssf/pkg/factory"
)

type NssfApp interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *nssf_context.NSSFContext
	Config() *factory.Config
}
