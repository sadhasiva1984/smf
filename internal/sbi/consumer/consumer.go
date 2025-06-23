package consumer

import (
	"github.com/free5gc/smf/pkg/app"
	"github.com/sadhasiva1984/openapi/amf/Communication"
	"github.com/sadhasiva1984/openapi/chf/ConvergedCharging"
	"github.com/sadhasiva1984/openapi/nrf/NFDiscovery"
	"github.com/sadhasiva1984/openapi/nrf/NFManagement"
	"github.com/sadhasiva1984/openapi/pcf/SMPolicyControl"
	"github.com/sadhasiva1984/openapi/smf/PDUSession"
	"github.com/sadhasiva1984/openapi/udm/SubscriberDataManagement"
	"github.com/sadhasiva1984/openapi/udm/UEContextManagement"
)

type Consumer struct {
	app.App

	// consumer services
	*nsmfService
	*namfService
	*nchfService
	*npcfService
	*nudmService
	*nnrfService
}

func NewConsumer(smf app.App) (*Consumer, error) {
	c := &Consumer{
		App: smf,
	}

	c.nsmfService = &nsmfService{
		consumer:          c,
		PDUSessionClients: make(map[string]*PDUSession.APIClient),
	}

	c.namfService = &namfService{
		consumer:             c,
		CommunicationClients: make(map[string]*Communication.APIClient),
	}

	c.nchfService = &nchfService{
		consumer:                 c,
		ConvergedChargingClients: make(map[string]*ConvergedCharging.APIClient),
	}

	c.nudmService = &nudmService{
		consumer:                        c,
		SubscriberDataManagementClients: make(map[string]*SubscriberDataManagement.APIClient),
		UEContextManagementClients:      make(map[string]*UEContextManagement.APIClient),
	}

	c.nnrfService = &nnrfService{
		consumer:            c,
		NFManagementClients: make(map[string]*NFManagement.APIClient),
		NFDiscoveryClients:  make(map[string]*NFDiscovery.APIClient),
	}

	c.npcfService = &npcfService{
		consumer:               c,
		SMPolicyControlClients: make(map[string]*SMPolicyControl.APIClient),
	}

	return c, nil
}
