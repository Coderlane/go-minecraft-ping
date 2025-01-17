package mcclient

import (
	"bytes"
	"testing"

	"github.com/Coderlane/go-minecraft-ping/client"
	"github.com/golang/mock/gomock"
)

const testStatusResponse = `{
    "version": {
        "name": "1.8.7",
        "protocol": 47
    },
    "players": {
        "max": 10,
        "online": 1,
        "sample": [
            {
                "name": "test",
                "id": "test_id"
            }
        ]
    },
    "description": {
        "text": "Hello world"
    },
    "favicon": "data:image/png;base64,<data>"
}`

type testContext struct {
	ctrl   *gomock.Controller
	client *client.MockClient
	mc     MinecraftClient
}

func newTestContext(t *testing.T) *testContext {
	ctrl := gomock.NewController(t)
	client := client.NewMockClient(ctrl)
	client.EXPECT().Addr().Return("localhost:1234")
	mc, err := NewMinecraftClient(client)
	if err != nil {
		t.Fatal(err)
	}
	return &testContext{
		ctrl:   ctrl,
		client: client,
		mc:     mc,
	}
}

func (tc *testContext) Finish() {
	tc.ctrl.Finish()
}

func TestHandshakeSuccess(t *testing.T) {
	tc := newTestContext(t)
	defer tc.Finish()

	hnd := Handshake{
		Version: -1,
		Address: "localhost",
		Port:    1234,
		State:   ClientStateStatus,
	}
	var buf bytes.Buffer
	if err := hnd.EncodeBinary(&buf); err != nil {
		t.Fatal(err)
	}
	pkt := client.Packet{
		ID:   0,
		Data: buf.Bytes(),
	}
	tc.client.EXPECT().Send(pkt).Return(nil)
	if err := tc.mc.Handshake(ClientStateStatus); err != nil {
		t.Fatal(err)
	}
}

func TestStatusSuccess(t *testing.T) {
	tc := newTestContext(t)
	defer tc.Finish()

	send := client.Packet{
		ID: 0,
	}

	testStr := client.VarString(testStatusResponse)
	var buf bytes.Buffer
	if err := testStr.EncodeBinary(&buf); err != nil {
		t.Fatal(err)
	}

	recv := &client.Packet{
		ID:   0,
		Data: buf.Bytes(),
	}
	tc.client.EXPECT().Send(send).Return(nil)
	tc.client.EXPECT().Recv().Return(recv, nil)
	status, err := tc.mc.Status()
	if err != nil {
		t.Fatal(err)
	}
	if status.Description.Text != "Hello world" {
		t.Errorf("Expected \"Hello world\": Got: \"%s\"\n",
			status.Description.Text)
	}
	t.Logf("%+v\n", status)
}
