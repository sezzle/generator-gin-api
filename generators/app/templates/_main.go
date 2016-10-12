package main

import (
  "flag"
  _ "net/http/pprof"

  "<%= myrepoUrl %>/<%= myappName %>/gin"
  "<%= myrepoUrl %>/<%= myappName %>/gorm"

  "github.com/golang/glog"
)

func main() {
  //Snag all flags that our application is run on.
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")

  //Initalize our db.
	glog.Info("Initalizing db...")
	db, err := gorm.InitDB()
	if err != nil {
		glog.Fatal("Could not initalize db", err.Error())
	}


  //Defer this so that if our application exits, we close the db.
	//Double check this.

	defer db.Close()

	glog.Info("Initalizing Models...")

	err = gorm.Migrate()
	if err != nil {
		glog.Fatal("Could not run object migrations.")
	}

  	gin.Run()

}
