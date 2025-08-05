package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jdetok/bball-etl-go/etl"
	"github.com/jdetok/golib/errd"
	"github.com/jdetok/golib/logd"
	"github.com/jdetok/golib/pgresd"
)

func main() {
	// start time variable for logging
	var sTime time.Time = time.Now()

	// Conf variable, hold logger, db, etc
	var cnf etl.Conf

	e := errd.InitErr() // start error handler

	// initialize logger
	l, err := logd.InitLogger("z_log", "nightly_etl")
	if err != nil {
		e.Msg = "error initializing logger"
		log.Fatal(e.BuildErr(err))
	}
	cnf.L = l // assign to cnf

	// postgres connection
	pg := pgresd.GetEnvPG()
	pg.MakeConnStr()
	db, err := pg.Conn()
	if err != nil {
		e.Msg = "error connecting to postgres"
		cnf.L.WriteLog(e.Msg)
		log.Fatal(e.BuildErr(err))
	}

	cnf.DB = db // asign to cnf
	cnf.DB.SetMaxOpenConns(40)
	cnf.DB.SetMaxIdleConns(40)
	cnf.RowCnt = 0 // START ROW COUNTER AT 0 BEFORE ETL STARTS

	if err = etl.RunNightlyETL(cnf); err != nil {
		e.Msg = fmt.Sprintf(
			"error with %v nightly etl", etl.Yesterday(time.Now()))
		cnf.L.WriteLog(e.Msg)
		log.Fatal(e.BuildErr(err))
	}

	// write errors to the log
	if len(cnf.Errs) > 0 {
		cnf.L.WriteLog(fmt.Sprintln("ERRORS:"))
		for _, e := range cnf.Errs {
			cnf.L.WriteLog(fmt.Sprintln(e))
		}
	}

	// email log file to myself
	EmailLog(cnf.L)
	if err != nil {
		e.Msg = "error emailing log"
		cnf.L.WriteLog(e.Msg)
		log.Fatal(e.BuildErr(err))
	}

	// log NIGHTLY process complete
	cnf.L.WriteLog(
		fmt.Sprint(
			"process complete",
			fmt.Sprintf(
				"\n ---- start time: %v", sTime),
			fmt.Sprintf(
				"\n ---- cmplt time: %v", time.Now()),
			fmt.Sprintf(
				"\n ---- duration: %v", time.Since(sTime)),
			fmt.Sprintf(
				"\n---- nightly etl for %v complete | total rows affected: %d",
				etl.Yesterday(time.Now()), cnf.RowCnt,
			),
		),
	)
}
