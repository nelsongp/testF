package process

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/updateRegisters"
	"time"
)

type infoBot struct {
	log        kitlog.Logger
	ctx        context.Context
	updateRegs updateRegisters.UpdateRegs
}

func NewInfoBot(log kitlog.Logger, ctx context.Context, updateRegs updateRegisters.UpdateRegs) *infoBot {
	return &infoBot{
		log:        log,
		ctx:        ctx,
		updateRegs: updateRegs,
	}
}

func (i *infoBot) BotInfoInit() {
	i.log.Log("Initializing bot", "Start", "infoBot", "hola")
	loc, _ := time.LoadLocation("America/El_Salvador")
	s := gocron.NewScheduler(loc)
	s.Every("30s").Do(i.updateRegs.UpdateRegisters)
	s.StartAsync()
}

func printTest() {
	fmt.Println("Testeando esto")
}
