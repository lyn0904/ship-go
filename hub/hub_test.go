package hub

//nolint:gosec
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lyn0904/ship-go/api"
	"github.com/lyn0904/ship-go/cert"
	"github.com/lyn0904/ship-go/mocks"
	"github.com/lyn0904/ship-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
) // #nosec G505

func TestHubSuite(t *testing.T) {
	suite.Run(t, new(HubSuite))
}

type testStruct struct {
	counter   int
	timeRange connectionInitiationDelayTimeRange
}

type HubSuite struct {
	suite.Suite

	hubReader   *mocks.MockHubReaderInterface
	mdnsService *mocks.MockMdnsInterface

	// serviceProvider  *mocks.ServiceProvider
	// mdnsService      *mocks.MdnsService
	shipConnection *mocks.ShipConnectionInterface
	wsDataWriter   *mocks.WebsocketDataWriterInterface

	remoteSki string

	tests []testStruct

	sut *Hub
}

func (s *HubSuite) BeforeTest(suiteName, testName string) {
	s.remoteSki = "remotetestski"

	s.tests = []testStruct{
		{0, connectionInitiationDelayTimeRanges[0]},
		{1, connectionInitiationDelayTimeRanges[1]},
		{2, connectionInitiationDelayTimeRanges[2]},
		{3, connectionInitiationDelayTimeRanges[2]},
		{4, connectionInitiationDelayTimeRanges[2]},
		{5, connectionInitiationDelayTimeRanges[2]},
		{6, connectionInitiationDelayTimeRanges[2]},
		{7, connectionInitiationDelayTimeRanges[2]},
		{8, connectionInitiationDelayTimeRanges[2]},
		{9, connectionInitiationDelayTimeRanges[2]},
		{10, connectionInitiationDelayTimeRanges[2]},
	}

	ctrl := gomock.NewController(s.T())
	// use gomock mocks instead of mockery, as those will panic with a data race error in these tests

	s.hubReader = mocks.NewMockHubReaderInterface(ctrl)
	// s.serviceProvider = mocks.NewServiceProvider(s.T())
	s.hubReader.EXPECT().RemoteSKIConnected(gomock.Any()).Return().AnyTimes()
	s.hubReader.EXPECT().RemoteSKIDisconnected(gomock.Any()).Return().AnyTimes()
	s.hubReader.EXPECT().ServiceShipIDUpdate(gomock.Any(), gomock.Any()).Return().AnyTimes()
	s.hubReader.EXPECT().ServicePairingDetailUpdate(gomock.Any(), gomock.Any()).Return().AnyTimes()
	s.hubReader.EXPECT().AllowWaitingForTrust(gomock.Any()).Return(false).AnyTimes()

	s.mdnsService = mocks.NewMockMdnsInterface(ctrl)
	s.mdnsService.EXPECT().AnnounceMdnsEntry().Return(nil).AnyTimes()
	s.mdnsService.EXPECT().UnannounceMdnsEntry().Return().AnyTimes()
	s.mdnsService.EXPECT().RequestMdnsEntries().Return().AnyTimes()

	s.wsDataWriter = mocks.NewWebsocketDataWriterInterface(s.T())

	s.shipConnection = mocks.NewShipConnectionInterface(s.T())
	s.shipConnection.EXPECT().CloseConnection(mock.Anything, mock.Anything, mock.Anything).Return().Maybe()
	s.shipConnection.EXPECT().RemoteSKI().Return(s.remoteSki).Maybe()
	s.shipConnection.EXPECT().ApprovePendingHandshake().Return().Maybe()
	s.shipConnection.EXPECT().AbortPendingHandshake().Return().Maybe()
	s.shipConnection.EXPECT().DataHandler().Return(s.wsDataWriter).Maybe()
	s.shipConnection.EXPECT().ShipHandshakeState().Return(model.SmeStateComplete, nil).Maybe()

	localService := api.NewServiceDetails("localSKI")

	certificate, _ := cert.CreateCertificate("unit", "org", "DE", "CN")
	s.sut = NewHub(s.hubReader, s.mdnsService, 4567, certificate, localService)
}

func (s *HubSuite) AfterTest(suiteName, testName string) {
	s.mdnsService.EXPECT().Shutdown().AnyTimes()

	s.sut.Shutdown()
}

func (s *HubSuite) Test_NewConnectionsHub() {
	ski := "12af9e"
	localService := api.NewServiceDetails(ski)

	hub := NewHub(s.hubReader, s.mdnsService, 4567, tls.Certificate{}, localService)
	assert.NotNil(s.T(), hub)

	s.mdnsService.EXPECT().Start(gomock.Any()).Return(nil).Times(1)

	hub.Start()

	s.mdnsService.EXPECT().Shutdown().Times(1)

	hub.Shutdown()
}

func (s *HubSuite) Test_AutoAccept() {
	s.mdnsService.EXPECT().SetAutoAccept(gomock.Any()).Return().AnyTimes()

	s.sut.SetAutoAccept(true)
	value := s.sut.IsAutoAcceptEnabled()
	assert.True(s.T(), value)

	s.sut.SetAutoAccept(false)
	value = s.sut.IsAutoAcceptEnabled()
	assert.False(s.T(), value)
}

func (s *HubSuite) Test_SetupRemoteDevice() {
	ski := "12af9e"
	localService := api.NewServiceDetails(ski)

	hub := NewHub(s.hubReader, s.mdnsService, 4567, tls.Certificate{}, localService)
	assert.NotNil(s.T(), hub)

	readerI := mocks.NewShipConnectionDataReaderInterface(s.T())
	s.hubReader.EXPECT().SetupRemoteDevice(gomock.Any(), gomock.Any()).Return(readerI)

	reader := hub.SetupRemoteDevice(ski, nil)

	assert.NotNil(s.T(), reader)
}

func (s *HubSuite) Test_SendWSCloseMessage() {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	s.sut.ServeHTTP(w, req)

	server := httptest.NewServer(s.sut)
	wsURL := strings.Replace(server.URL, "http://", "ws://", -1)

	// Connect to the server
	con, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.Nil(s.T(), err)

	ski := "12af9e"
	localService := api.NewServiceDetails(ski)

	hub := NewHub(s.hubReader, s.mdnsService, 4567, tls.Certificate{}, localService)
	assert.NotNil(s.T(), hub)

	hub.sendWSCloseMessage(con)

	resp.Body.Close()
	_ = con.Close()
	server.CloseClientConnections()
	server.Close()

	time.Sleep(time.Second)
}

func (s *HubSuite) Test_IsRemoteSKIPaired() {
	paired := s.sut.IsRemoteServiceForSKIPaired(s.remoteSki)
	assert.Equal(s.T(), false, paired)

	s.sut.registerConnection(s.shipConnection)
	s.sut.RegisterRemoteSKI(s.remoteSki)

	paired = s.sut.IsRemoteServiceForSKIPaired(s.remoteSki)
	assert.Equal(s.T(), true, paired)

	// remove the connection, so the test doesn't try to close it
	delete(s.sut.connections, s.remoteSki)
	s.sut.UnregisterRemoteSKI(s.remoteSki)
	paired = s.sut.IsRemoteServiceForSKIPaired(s.remoteSki)
	assert.Equal(s.T(), false, paired)

	ski := "12af9e"
	localService := api.NewServiceDetails(ski)

	hub := NewHub(s.hubReader, s.mdnsService, 4567, tls.Certificate{}, localService)
	assert.NotNil(s.T(), hub)

	s.mdnsService.EXPECT().Start(gomock.Any()).Return(nil).Times(1)
	hub.Start()

	hub.UnregisterRemoteSKI(s.remoteSki)
	paired = s.sut.IsRemoteServiceForSKIPaired(s.remoteSki)
	assert.Equal(s.T(), false, paired)

	s.mdnsService.EXPECT().Shutdown().Times(1)
	hub.Shutdown()
}

func (s *HubSuite) Test_RegisterRemoteSKI_AfterStart() {
	s.sut.hasStarted = true

	s.sut.RegisterRemoteSKI(s.remoteSki)
	assert.Equal(s.T(), 0, len(s.sut.connections))

	s.sut.registerConnection(s.shipConnection)
	s.sut.RegisterRemoteSKI(s.remoteSki)
	assert.Equal(s.T(), 1, len(s.sut.connections))
}

func (s *HubSuite) Test_HandleConnectionClosed() {
	s.sut.HandleConnectionClosed(s.shipConnection, false)

	s.sut.registerConnection(s.shipConnection)

	s.sut.HandleConnectionClosed(s.shipConnection, true)

	assert.Equal(s.T(), 0, len(s.sut.connections))
}

func (s *HubSuite) Test_Mdns() {
	s.sut.checkAutoReannounce()

	pairedServices := s.sut.numberPairedServices()
	assert.Equal(s.T(), 0, len(s.sut.connections))
	assert.Equal(s.T(), 0, pairedServices)

	s.sut.RegisterRemoteSKI(s.remoteSki)
	pairedServices = s.sut.numberPairedServices()
	assert.Equal(s.T(), 0, len(s.sut.connections))
	assert.Equal(s.T(), 1, pairedServices)
}

func (s *HubSuite) Test_Ship() {
	s.sut.HandleShipHandshakeStateUpdate(s.remoteSki, model.ShipState{
		State: model.SmeStateError,
		Error: errors.New("test"),
	})

	s.sut.HandleShipHandshakeStateUpdate(s.remoteSki, model.ShipState{
		State: model.SmeHelloStateOk,
	})

	s.sut.ReportServiceShipID(s.remoteSki, "test")

	accept := s.sut.IsAutoAcceptEnabled()
	assert.Equal(s.T(), false, accept)

	trust := s.sut.AllowWaitingForTrust(s.remoteSki)
	assert.Equal(s.T(), true, trust)

	trust = s.sut.AllowWaitingForTrust("test")
	assert.Equal(s.T(), false, trust)

	detail := s.sut.PairingDetailForSki(s.remoteSki)
	assert.NotNil(s.T(), detail)

	s.sut.registerConnection(s.shipConnection)

	detail = s.sut.PairingDetailForSki(s.remoteSki)
	assert.NotNil(s.T(), detail)
}

func (s *HubSuite) Test_MapShipMessageExchangeState() {
	state := s.sut.mapShipMessageExchangeState(model.CmiStateInitStart, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateQueued, state)

	state = s.sut.mapShipMessageExchangeState(model.CmiStateClientSend, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateInitiated, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeHelloStateReadyInit, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateInProgress, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeHelloStatePendingListen, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateReceivedPairingRequest, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeHelloStateOk, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateTrusted, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeHelloStateAbort, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateNone, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeHelloStateRemoteAbortDone, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateRemoteDeniedTrust, state)

	state = s.sut.mapShipMessageExchangeState(model.SmePinStateCheckInit, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStatePin, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeAccessMethodsRequest, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateInProgress, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeStateComplete, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateCompleted, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeStateError, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateError, state)

	state = s.sut.mapShipMessageExchangeState(model.SmeProtHStateTimeout, s.remoteSki)
	assert.Equal(s.T(), api.ConnectionStateInProgress, state)
}

func (s *HubSuite) Test_DisconnectSKI() {
	s.sut.DisconnectSKI(s.remoteSki, "none")
}

func (s *HubSuite) Test_RegisterConnection() {
	s.sut.registerConnection(s.shipConnection)
	assert.Equal(s.T(), 1, len(s.sut.connections))
	con := s.sut.connectionForSKI(s.remoteSki)
	assert.NotNil(s.T(), con)
}

func (s *HubSuite) Test_VerifyPeerCertificate() {
	testCert, _ := cert.CreateCertificate("unit", "org", "DE", "CN")
	var rawCerts [][]byte
	rawCerts = append(rawCerts, testCert.Certificate...)
	err := s.sut.verifyPeerCertificate(rawCerts, nil)
	assert.Nil(s.T(), err)

	rawCerts = nil
	rawCerts = append(rawCerts, []byte{100})
	err = s.sut.verifyPeerCertificate(rawCerts, nil)
	assert.NotNil(s.T(), err)

	rawCerts = nil
	invalidCert, _ := createInvalidCertificate("unit", "org", "DE", "CN")
	rawCerts = append(rawCerts, invalidCert.Certificate...)

	err = s.sut.verifyPeerCertificate(rawCerts, nil)
	assert.NotNil(s.T(), err)
}

func (s *HubSuite) Test_ServeHTTP_01() {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	s.sut.ServeHTTP(w, req)

	server := httptest.NewServer(s.sut)
	wsURL := strings.Replace(server.URL, "http://", "ws://", -1)

	// Connect to the server
	con, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.Nil(s.T(), err)
	resp.Body.Close()
	_ = con.Close()

	dialer := &websocket.Dialer{
		Subprotocols: []string{api.ShipWebsocketSubProtocol},
	}
	con, resp, err = dialer.Dial(wsURL, nil)
	assert.Nil(s.T(), err)

	resp.Body.Close()
	_ = con.Close()
	server.CloseClientConnections()
	server.Close()

	time.Sleep(time.Second)
}

func (s *HubSuite) Test_ServeHTTP_02() {
	server := httptest.NewUnstartedServer(s.sut)
	server.TLS = &tls.Config{
		Certificates:       []tls.Certificate{s.sut.certifciate},
		ClientAuth:         tls.RequireAnyClientCert,
		CipherSuites:       cert.CipherSuites, // #nosec G402
		InsecureSkipVerify: true,              // #nosec G402
	}
	server.StartTLS()
	wsURL := strings.Replace(server.URL, "https://", "wss://", -1)

	invalidCert, _ := createInvalidCertificate("unit", "org", "DE", "CN")
	dialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{invalidCert},
			InsecureSkipVerify: true,              // #nosec G402
			CipherSuites:       cert.CipherSuites, // #nosec G402
		},
		Subprotocols: []string{api.ShipWebsocketSubProtocol},
	}
	con, resp, err := dialer.Dial(wsURL, nil)
	assert.Nil(s.T(), err)

	resp.Body.Close()
	_ = con.Close()

	validCert, _ := cert.CreateCertificate("unit", "org", "DE", "CN")
	dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{validCert},
			InsecureSkipVerify: true,              // #nosec G402
			CipherSuites:       cert.CipherSuites, // #nosec G402
		},
		Subprotocols: []string{api.ShipWebsocketSubProtocol},
	}
	con, resp, err = dialer.Dial(wsURL, nil)
	assert.Nil(s.T(), err)

	resp.Body.Close()
	_ = con.Close()
	server.CloseClientConnections()
	server.Close()

	time.Sleep(time.Second)
}

func (s *HubSuite) Test_ConnectFoundService_01() {
	service := s.sut.ServiceForSKI(s.remoteSki)

	err := s.sut.connectFoundService(service, "localhost", "80", "/ship")
	assert.NotNil(s.T(), err)

	server := httptest.NewServer(s.sut)
	url, err := url.Parse(server.URL)
	assert.Nil(s.T(), err)

	err = s.sut.connectFoundService(service, url.Hostname(), url.Port(), url.Path)
	assert.NotNil(s.T(), err)

	server.CloseClientConnections()
	server.Close()

	time.Sleep(time.Second)
}

func (s *HubSuite) Test_ConnectFoundService_02() {
	service := s.sut.ServiceForSKI(s.remoteSki)

	server := httptest.NewUnstartedServer(s.sut)
	invalidCert, _ := createInvalidCertificate("unit", "org", "DE", "CN")
	server.TLS = &tls.Config{
		Certificates:       []tls.Certificate{invalidCert},
		ClientAuth:         tls.RequireAnyClientCert,
		CipherSuites:       cert.CipherSuites, // #nosec G402
		InsecureSkipVerify: true,              // #nosec G402
	}
	server.StartTLS()

	url, err := url.Parse(server.URL)
	assert.Nil(s.T(), err)

	err = s.sut.connectFoundService(service, url.Hostname(), url.Port(), url.Path)
	assert.NotNil(s.T(), err)

	server.CloseClientConnections()
	server.Close()

	time.Sleep(time.Second)
}

func (s *HubSuite) Test_ConnectFoundService_03() {
	service := s.sut.ServiceForSKI(s.remoteSki)

	server := httptest.NewUnstartedServer(s.sut)
	server.TLS = &tls.Config{
		Certificates:       []tls.Certificate{s.sut.certifciate},
		ClientAuth:         tls.RequireAnyClientCert,
		CipherSuites:       cert.CipherSuites, // #nosec G402
		InsecureSkipVerify: true,              // #nosec G402
	}
	server.StartTLS()

	url, err := url.Parse(server.URL)
	assert.Nil(s.T(), err)

	err = s.sut.connectFoundService(service, url.Hostname(), url.Port(), url.Path)
	assert.NotNil(s.T(), err)

	time.Sleep(time.Second)

	server.CloseClientConnections()
	server.Close()
}

func (s *HubSuite) Test_KeepThisConnection() {
	service := s.sut.ServiceForSKI(s.remoteSki)

	result := s.sut.keepThisConnection(nil, false, service)
	assert.Equal(s.T(), true, result)

	s.sut.registerConnection(s.shipConnection)

	result = s.sut.keepThisConnection(nil, false, service)
	assert.Equal(s.T(), false, result)

	result = s.sut.keepThisConnection(nil, true, service)
	assert.Equal(s.T(), true, result)
}

func (s *HubSuite) Test_prepareConnectionInitiation() {
	entry := &api.MdnsEntry{
		Ski:  s.remoteSki,
		Host: "somehost",
	}
	service := s.sut.ServiceForSKI(s.remoteSki)

	s.sut.prepareConnectionInitation(s.remoteSki, 0, entry)

	s.sut.setConnectionAttemptRunning(s.remoteSki, true)

	counter := s.sut.increaseConnectionAttemptCounter(s.remoteSki)
	assert.Equal(s.T(), 0, counter)
	s.sut.prepareConnectionInitation(s.remoteSki, 0, entry)

	s.sut.UnregisterRemoteSKI(s.remoteSki)
	service.ConnectionStateDetail().SetState(api.ConnectionStateQueued)

	counter = s.sut.increaseConnectionAttemptCounter(s.remoteSki)
	assert.Equal(s.T(), 0, counter)

	s.sut.prepareConnectionInitation(s.remoteSki, 0, entry)
}

func (s *HubSuite) Test_InitiateConnection() {
	entry := &api.MdnsEntry{
		Ski:  s.remoteSki,
		Host: "somehost",
	}
	service := s.sut.ServiceForSKI(s.remoteSki)

	result := s.sut.initateConnection(service, entry)
	assert.Equal(s.T(), false, result)

	entry.Addresses = []net.IP{[]byte("127.0.0.1")}

	result = s.sut.initateConnection(service, entry)
	assert.Equal(s.T(), false, result)

	s.sut.RegisterRemoteSKI(s.remoteSki)
	service.ConnectionStateDetail().SetState(api.ConnectionStateQueued)

	result = s.sut.initateConnection(service, entry)
	assert.Equal(s.T(), false, result)
}

func (s *HubSuite) Test_checkHasStarted() {
	checked := s.sut.checkHasStarted()
	assert.Equal(s.T(), s.sut.hasStarted, checked)
}

func (s *HubSuite) Test_IncreaseConnectionAttemptCounter() {
	for _, test := range s.tests {
		s.sut.increaseConnectionAttemptCounter(s.remoteSki)

		s.sut.muxConAttempt.Lock()
		counter, exists := s.sut.connectionAttemptCounter[s.remoteSki]
		timeRange := connectionInitiationDelayTimeRanges[counter]
		s.sut.muxConAttempt.Unlock()

		assert.Equal(s.T(), true, exists)
		assert.Equal(s.T(), test.timeRange.min, timeRange.min)
		assert.Equal(s.T(), test.timeRange.max, timeRange.max)
	}
}

func (s *HubSuite) Test_RemoveConnectionAttemptCounter() {
	s.sut.increaseConnectionAttemptCounter(s.remoteSki)
	_, exists := s.sut.connectionAttemptCounter[s.remoteSki]
	assert.Equal(s.T(), true, exists)

	s.sut.removeConnectionAttemptCounter(s.remoteSki)
	_, exists = s.sut.connectionAttemptCounter[s.remoteSki]
	assert.Equal(s.T(), false, exists)
}

func (s *HubSuite) Test_GetCurrentConnectionAttemptCounter() {
	s.sut.increaseConnectionAttemptCounter(s.remoteSki)
	_, exists := s.sut.connectionAttemptCounter[s.remoteSki]
	assert.Equal(s.T(), exists, true)
	s.sut.increaseConnectionAttemptCounter(s.remoteSki)

	value, exists := s.sut.getCurrentConnectionAttemptCounter(s.remoteSki)
	assert.Equal(s.T(), 1, value)
	assert.Equal(s.T(), true, exists)
}

func (s *HubSuite) Test_GetConnectionInitiationDelayTime() {
	counter, duration := s.sut.getConnectionInitiationDelayTime(s.remoteSki)
	assert.Equal(s.T(), 0, counter)
	assert.LessOrEqual(s.T(), float64(s.tests[counter].timeRange.min), float64(duration/time.Second))
	assert.GreaterOrEqual(s.T(), float64(s.tests[counter].timeRange.max), float64(duration/time.Second))
}

func (s *HubSuite) Test_ConnectionAttemptRunning() {
	s.sut.setConnectionAttemptRunning(s.remoteSki, true)
	status := s.sut.isConnectionAttemptRunning(s.remoteSki)
	assert.Equal(s.T(), true, status)
	s.sut.setConnectionAttemptRunning(s.remoteSki, false)
	status = s.sut.isConnectionAttemptRunning(s.remoteSki)
	assert.Equal(s.T(), false, status)
}

func (s *HubSuite) Test_CancelPairingWithSKI() {
	s.sut.CancelPairingWithSKI(s.remoteSki)
	assert.Equal(s.T(), 0, len(s.sut.connections))
	assert.Equal(s.T(), 0, len(s.sut.connectionAttemptRunning))

	s.sut.registerConnection(s.shipConnection)
	assert.Equal(s.T(), 1, len(s.sut.connections))

	s.sut.CancelPairingWithSKI(s.remoteSki)
	assert.Equal(s.T(), 0, len(s.sut.connectionAttemptRunning))
}

func (s *HubSuite) Test_ReportMdnsEntries() {
	testski1 := "test1"
	testski2 := "test2"

	entries := make(map[string]*api.MdnsEntry)

	s.hubReader.EXPECT().VisibleRemoteServicesUpdated(gomock.Any()).AnyTimes()
	s.sut.ReportMdnsEntries(entries, true)

	entries[testski1] = &api.MdnsEntry{
		Ski: testski1,
	}
	service1 := s.sut.ServiceForSKI(testski1)
	service1.SetTrusted(true)
	service1.SetIPv4("127.0.0.1")

	entries[testski2] = &api.MdnsEntry{
		Ski: testski2,
	}
	service2 := s.sut.ServiceForSKI(testski2)
	service2.SetTrusted(true)
	service2.SetIPv4("127.0.0.1")

	s.sut.ReportMdnsEntries(entries, true)
}

func createInvalidCertificate(organizationalUnit, organization, country, commonName string) (tls.Certificate, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Create the EEBUS service SKI using the private key
	asn1, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return tls.Certificate{}, err
	}
	// SHIP 12.2: Required to be created according to RFC 3280 4.2.1.2
	// #nosec G401
	ski := sha1.Sum(asn1)

	subject := pkix.Name{
		OrganizationalUnit: []string{organizationalUnit},
		Organization:       []string{organization},
		Country:            []string{country},
		CommonName:         commonName,
	}

	// Create a random serial big int value
	maxValue := new(big.Int)
	maxValue.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(maxValue, big.NewInt(1))
	serialNumber, err := rand.Int(rand.Reader, maxValue)
	if err != nil {
		return tls.Certificate{}, err
	}

	template := x509.Certificate{
		SignatureAlgorithm:    x509.ECDSAWithSHA256,
		SerialNumber:          serialNumber,
		Subject:               subject,
		NotBefore:             time.Now(),                                // Valid starting now
		NotAfter:              time.Now().Add(time.Hour * 24 * 365 * 10), // Valid for 10 years
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		SubjectKeyId:          ski[:19],
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return tls.Certificate{}, err
	}

	tlsCertificate := tls.Certificate{
		Certificate:                  [][]byte{certBytes},
		PrivateKey:                   privateKey,
		SupportedSignatureAlgorithms: []tls.SignatureScheme{tls.ECDSAWithP256AndSHA256},
	}

	return tlsCertificate, nil
}
