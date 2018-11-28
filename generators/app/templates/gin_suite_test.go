package gin_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

  "github.com/golang/glog"
	"<%= myAppPath %>/config"
  testapi "<%= myAppPath %>/gin"
	testgorm "<%= myAppPath %>/gorm"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "10")
}

var (
	s               *gin.Engine
	response        *httptest.ResponseRecorder
	request         *http.Request
	endpointURL     string = ""
	endpointMethod  string
	endpointHeaders http.Header
	form            interface{}
)

type errorReply map[string][]string
type keyValueResp map[string]string

func TestGin(t *testing.T) {
	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		// Initialize an in memory database (WIP)
		config.SetSettingsFromViper()
		var testDBName string = config.DbName
		// db.LogMode(true)
		db, err := testgorm.InitDB()

		//Drop old test db
		dropDBStmt := fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, testDBName)
		db.Exec(dropDBStmt)

		//TODO: lock tables

		//Create Database
		createStmt := fmt.Sprintf(`CREATE DATABASE %s;`, testDBName)
		result := db.Exec(createStmt)
		if result.Error != nil {
			glog.Error(result.Error)
		}

		//Select new test database
		useStmnt := fmt.Sprintf(`USE %s;`, testDBName)
		result = db.Exec(useStmnt)
		if result.Error != nil {
			glog.Error(result.Error)
		}

		err = testgorm.Migrate()
		if err != nil {
			glog.Fatal("Could not run object migrations.")
		}

		gin.SetMode(gin.ReleaseMode)
		s = testapi.InitRoutes()

	})

	AfterSuite(func() {
		//Drop old test db
		//dropDBStmt := fmt.Sprintf(`DROP DATABASE IF EXISTS %s`, testDBName)
	})

	RunSpecs(t, "Gin Suite")
}

//DecodeTestJson Decodes response body to interface provided
func DecodeTestJson(response *httptest.ResponseRecorder, decodeStruct interface{}) error {
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		glog.Error(err)
		return err
	}

	decoder := json.NewDecoder(bytes.NewBuffer(htmlData))
	if err = decoder.Decode(decodeStruct); err != nil {
		return err
	}

	return nil
}
