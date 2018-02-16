package fabric

// Basic imports
import (
	"context"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// ProtocolTransportTestSuite -
type ProtocolTransportTestSuite struct {
	suite.Suite
}

func (suite *ProtocolTransportTestSuite) SetupTest() {}

func (suite *ProtocolTransportTestSuite) TestName() {
	tls := &transportWrapper{
		protocolNames: []string{},
	}

	name := tls.Name()
	suite.Assert().Equal("", name)
}

func (suite *ProtocolTransportTestSuite) TestHandleSuccess() {
	wrp := &transportWrapper{
		protocolNames: []string{},
	}

	protocol := &MockProtocol{}
	protocol.On("Name").Return("test")
	var handler HandlerFunc = func(ctx context.Context, c Conn) error {
		return nil
	}
	var negotiator NegotiatorFunc = func(ctx context.Context, c Conn) error {
		return nil
	}
	protocol.On("Handle", mock.Anything).Return(handler)
	protocol.On("Negotiate", mock.Anything).Return(negotiator)

	addr := NewAddress("test")
	mockConn := &MockConn{}
	mockConn.On("GetAddress").Return(addr)
	suite.Assert().Equal("test", addr.CurrentProtocol())

	ctx := context.Background()
	err := wrp.Handle(protocol.Handle(nil))(ctx, mockConn)
	suite.Assert().Nil(err)
	protocol.AssertCalled(suite.T(), "Handle", mock.Anything)
}

func (suite *ProtocolTransportTestSuite) TestNegotiateSuccess() {
	wrp := &transportWrapper{
		protocolNames: []string{},
	}

	protocol := &MockProtocol{}
	protocol.On("Name").Return("test")
	var handler HandlerFunc = func(ctx context.Context, c Conn) error {
		return nil
	}
	var negotiator NegotiatorFunc = func(ctx context.Context, c Conn) error {
		return nil
	}
	protocol.On("Handle", mock.Anything).Return(handler)
	protocol.On("Negotiate", mock.Anything).Return(negotiator)

	addr := NewAddress("test")
	mockConn := &MockConn{}
	mockConn.On("GetAddress").Return(addr)
	suite.Assert().Equal("test", addr.CurrentProtocol())

	ctx := context.Background()
	err := wrp.Negotiate(protocol.Negotiate(nil))(ctx, mockConn)
	suite.Assert().Nil(err)
	protocol.AssertCalled(suite.T(), "Negotiate", mock.Anything)
}

func TestProtocolTransportTestSuite(t *testing.T) {
	suite.Run(t, new(ProtocolTransportTestSuite))
}
