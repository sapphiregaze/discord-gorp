package rpc

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gorilla/websocket"

	"github.com/sapphiregaze/discord-gorp/pkg/config"
	"github.com/sapphiregaze/discord-gorp/pkg/logger"
)

type RPCClient struct {
	conn *websocket.Conn
	pid  int
}

func NewClient() (*RPCClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:6463", nil)
	if err != nil {
		return nil, err
	}

	return &RPCClient{
		conn: conn,
		pid:  os.Getpid(),
	}, nil
}

func (r *RPCClient) Close() {
	r.conn.Close()
	logger.Info("Disconnected from Discord RPC server")
}

func (r *RPCClient) SetActivity(activity *config.Activity) error {
	payload := map[string]interface{}{
		"cmd": "SET_ACTIVITY",
		"args": map[string]interface{}{
			"pid":      r.pid,
			"activity": activity,
		},
		"nonce": generateNonce(),
	}

	message, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = r.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}

	logger.Info("Sent activity update to Discord RPC server")
	return nil
}

func generateNonce() string {
	return time.Now().Format("20060102150405")
}
