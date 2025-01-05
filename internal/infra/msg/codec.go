package msg

import (
	"encoding/json"
	"errors"
	"kafka-messager/internal/domain"
)

type Codec struct {
	topic string
	sch   *Schema
}

func (mc Codec) Encode(value any) ([]byte, error) {
	if _, isMsg := value.(*domain.Msg); !isMsg {
		return nil, errors.New("value is not Msg")
	}
	return json.Marshal(value)
}

func (mc Codec) Decode(data []byte) (any, error) {
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

func NewMsgCodec(topic string, sch *Schema) Codec {
	return Codec{topic, sch}
}
