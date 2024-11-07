package ship

import (
	"fmt"

	"github.com/lyn0904/ship-go/model"
)

// Handshake initialization covers the states cmiState...

// CMI_STATE_INIT_START
func (c *ShipConnection) handshakeInit_cmiStateInitStart() {
	switch c.role {
	case ShipRoleClient:
		// CMI_STATE_CLIENT_SEND
		c.setState(model.CmiStateClientSend, nil)
		if err := c.dataWriter.WriteMessageToWebsocketConnection(model.ShipInit); err != nil {
			c.endHandshakeWithError(err)
			return
		}
		c.setState(model.CmiStateClientWait, nil)
	case ShipRoleServer:
		c.setState(model.CmiStateServerWait, nil)
	}

	c.setHandshakeTimer(timeoutTimerTypeWaitForReady, cmiTimeout)
}

// CMI_STATE_SERVER_WAIT
func (c *ShipConnection) handshakeInit_cmiStateServerWait(message []byte) {
	c.setState(model.CmiStateServerEvaluate, nil)

	if !c.handshakeInit_cmiStateEvaluate(message) {
		return
	}

	if err := c.dataWriter.WriteMessageToWebsocketConnection(model.ShipInit); err != nil {
		c.endHandshakeWithError(err)
		return
	}

	c.setAndHandleState(model.SmeHelloState)
}

// CMI_STATE_CLIENT_WAIT
func (c *ShipConnection) handshakeInit_cmiStateClientWait(message []byte) {
	c.setState(model.CmiStateClientEvaluate, nil)

	if !c.handshakeInit_cmiStateEvaluate(message) {
		return
	}

	c.setAndHandleState(model.SmeHelloState)
}

// CMI_STATE_SERVER_EVALUATE
// CMI_STATE_CLIENT_EVALUATE
// returns false in case of an error
func (c *ShipConnection) handshakeInit_cmiStateEvaluate(message []byte) bool {
	msgType, data := c.parseMessage(message, false)

	if msgType != model.MsgTypeInit {
		c.endHandshakeWithError(fmt.Errorf("Invalid SHIP MessageType, expected 0 and got %s", string(msgType)))
		return false
	}
	if len(data) > 0 && data[0] != byte(0) {
		c.endHandshakeWithError(fmt.Errorf("Invalid SHIP MessageValue, expected 0 and got %s", string(data)))
		return false
	}

	return true
}
