package msg

import (
	"context"
	"encoding/binary"
	"kafka-messager/internal/domain"
	"kafka-messager/internal/infra/config"
	"kafka-messager/internal/infra/repo"
	"log"
	"strings"

	"github.com/lovoo/goka"
	"go.uber.org/zap"
)

type Processor struct {
	gp       *goka.Processor
	conf     config.Config
	log      *zap.SugaredLogger
	banWords []string
	bwr      repo.BanWord
}

func maskBanWords(msg string, banWords []string) string {
	res := msg
	for _, word := range banWords {
		res = strings.Replace(res, word, "***", -1)
	}
	return res
}

func (p *Processor) getProcess() func(ctx goka.Context, msg any) {
	return func(ctx goka.Context, msg any) {
		var (
			m  *domain.Msg
			ok bool
		)

		if m, ok = msg.(*domain.Msg); !ok || m == nil {
			return
		}

		m.RecipientId = int64(binary.BigEndian.Uint32([]byte(ctx.Key())))
		m.Message = maskBanWords(m.Message, p.banWords)

		ctx.SetValue(m)
	}
}

func (p *Processor) UpdateBanWords(ctx context.Context) {
	bw, err := p.bwr.GetList(ctx)
	if err != nil {
		log.Fatal(err)
	}

	p.banWords = make([]string, 0)
	for _, b := range bw {
		p.banWords = append(p.banWords, b.Word)
	}
}

func (p *Processor) Run(ctx context.Context) {
	g := goka.DefineGroup(goka.Group(p.conf.MsgFiltered),
		goka.Input(goka.Stream(p.conf.MsgFilteredBlockUsersTopic), new(Codec), p.getProcess()),
		goka.Persist(new(Codec)),
	)

	gp, err := goka.NewProcessor(p.conf.Brokers, g)
	if err != nil {
		log.Fatal(err)
	}
	p.gp = gp

	err = p.gp.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}

	p.gp = gp
}

func (p *Processor) Stop() {
	if p.gp != nil {
		p.gp.Stop()
	}
}

func NewProcessor(conf config.Config, log *zap.SugaredLogger, bwr repo.BanWord) *Processor {
	var gp *goka.Processor
	banWords := make([]string, 0)

	return &Processor{gp, conf, log, banWords, bwr}
}
