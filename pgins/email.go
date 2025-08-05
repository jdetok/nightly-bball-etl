package main

import (
	"fmt"

	"github.com/jdetok/golib/logd"
	"github.com/jdetok/golib/maild"
)

func EmailLog(l logd.Logger) error {
	m := maild.MakeMail(
		[]string{"jdekock17@gmail.com"},
		"Go bball ETL log attached",
		"the Go bball ETL process ran. The log is attached.",
	)
	l.WriteLog(fmt.Sprintf("attempting to email %s to %s", l.LogF, m.MlTo[0]))
	return m.SendMIMEEmail(l.LogF)
}
