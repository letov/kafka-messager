package emitter

import (
	"encoding/json"
	"errors"
	"kafka-messager/internal/domain"
)

type MsgCodec struct{}

func (mc MsgCodec) Encode(value any) ([]byte, error) {
	if _, isMsg := value.(*domain.Msg); !isMsg {
		return nil, errors.New("value is not Msg")
	}
	return json.Marshal(value)
}

func (mc MsgCodec) Decode(data []byte) (any, error) {
	var (
		m   domain.Msg
		err error
	)
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, errors.New("unmarshal Msg failed")
	}
	return &m, nil
}
