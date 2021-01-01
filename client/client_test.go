package client

import (
	"net"
	"reflect"
	"testing"
)

type testServer struct {
	listener net.Listener
}

func newTestServer(t *testing.T) *testServer {
	t.Helper()
	ln, err := net.Listen("tcp", ":")
	if err != nil {
		t.Fatal(err)
	}
	return &testServer{
		listener: ln,
	}
}

func (ts *testServer) Start() {
	go func() {
		for {
			conn, err := ts.listener.Accept()
			if err != nil {
				return
			}
			go func() {
				for {
					var pkt Packet
					if err := pkt.DecodeBinary(conn); err != nil {
						return
					}
					if err := pkt.EncodeBinary(conn); err != nil {
						return
					}
				}
			}()
		}
	}()
}

func (ts *testServer) Stop() error {
	return ts.listener.Close()
}

func TestClientEcho(t *testing.T) {
	srv := newTestServer(t)
	srv.Start()
	defer srv.Stop()

	cnt, err := NewClient(srv.listener.Addr().String())
	defer cnt.Close()
	if err != nil {
		t.Fatal(err)
	}

	input := Packet{
		ID:   1,
		Data: []byte("test"),
	}
	if err = cnt.Send(input); err != nil {
		t.Fatal(err)
	}

	output, err := cnt.Recv()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(input, *output) {
		t.Errorf("Expected: %v Got: %v\n", input, *output)
	}
}

func TestClientFailure(t *testing.T) {
	cnt, err := NewClient("*invalid$")
	if cnt != nil {
		t.Errorf("Expected null client")
	}
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestClientSendRecvFailure(t *testing.T) {
	srv := newTestServer(t)
	srv.Start()
	defer srv.Stop()

	cnt, err := NewClient(srv.listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	cnt.Close()
	var pkt Packet
	if err = cnt.Send(pkt); err == nil {
		t.Errorf("Expected to fail to send packet.")
	}
	if _, err := cnt.Recv(); err == nil {
		t.Errorf("Expected to fail to recv packet.")
	}
}
