package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nelsongp/testF/internal/botinfo"
	"github.com/nelsongp/testF/internal/botinfo/process"
	"github.com/nelsongp/testF/internal/getInfoNodeRed/getInfoNode"
	handlerShowNode "github.com/nelsongp/testF/internal/getInfoNodeRed/platform/handler"
	"github.com/nelsongp/testF/internal/testEndpoint/getTestEndpoint"
	"github.com/nelsongp/testF/internal/testEndpoint/platform/handler"
	mysql2 "github.com/nelsongp/testF/internal/updateRegisters/platform/storage/mysql"
	updateRegs2 "github.com/nelsongp/testF/internal/updateRegisters/updateRegs"
	"github.com/nelsongp/testF/kit/platform/server"
	"os"
)

func Run() {
	port := "9292"
	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	strConnection := getStrConnection()
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		kitlogger.Log("unable to open database connection %s", err.Error())
	}
	defer db.Close()

	///get test endpoint
	getTestSvc := getTestEndpoint.NewTestEndpointService(kitlogger)
	getTestSub := handler.MakeTestEndpoint(getTestSvc)
	getTestSub = handler.GetTestResponseMiddleware(kitlogger)(getTestSub)
	getTestHandler := handler.NewHttpGetTestResponseHandler("/test", getTestSub)
	///end get test endpoint

	//showNodeRedInfo endpoint
	showNodeRedSvc := getInfoNode.NewShowInfoNodeRedService(kitlogger)
	showNodeRedSub := handlerShowNode.MakeRetrieveNodeRedInfo(showNodeRedSvc)
	showNodeRedSub = handlerShowNode.ShowInfoNodeRedTransportMiddleware(kitlogger)(showNodeRedSub)
	showNodeRedHandler := handlerShowNode.NewHttpShowNodeRedInfoHandler("/showNodeRed", showNodeRedSub)
	//End showNodeRedInfo endpoint

	//bot info init
	selectNotRep := mysql2.NewNotProcessed(db, kitlogger)
	updateRepo := mysql2.NewUpdateRepo(db, kitlogger)
	updateRegs := updateRegs2.NewUpdateRegsService(kitlogger, updateRepo, selectNotRep)
	var botinit botinfo.IbotProcess
	botinit = process.NewInfoBot(kitlogger, context.Background(), updateRegs)
	go botinit.BotInfoInit()
	//end bot info
	kitlogger.Log("initiate", "test", "log")
	svc := server.NewServer(kitlogger)
	svc.RegisterRoutes("/test", getTestHandler)
	svc.RegisterRoutes("/showNodeRed", showNodeRedHandler)
	svc.Run(port)
}

func getStrConnection() string {
	host := "127.0.0.1:3306"
	user := "root"
	pass := "root"
	dbname := "test"
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
