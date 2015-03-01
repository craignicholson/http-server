// This application will accept data from the go-weatherunderground application
// and post the data to mongodb.
// TODO:  Should be clean the data here or change the schema for our own
//        database?

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
)

// http://mholt.github.io/json-to-go/
type WeatherHistory struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			History int `json:"history"`
		} `json:"features"`
	} `json:"response"`
	History struct {
		Date struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"date"`
		Utcdate struct {
			Pretty string `json:"pretty"`
			Year   string `json:"year"`
			Mon    string `json:"mon"`
			Mday   string `json:"mday"`
			Hour   string `json:"hour"`
			Min    string `json:"min"`
			Tzname string `json:"tzname"`
		} `json:"utcdate"`
		Observations []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"date"`
			Utcdate struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"utcdate"`
			Tempm      string `json:"tempm"`
			Tempi      string `json:"tempi"`
			Dewptm     string `json:"dewptm"`
			Dewpti     string `json:"dewpti"`
			Hum        string `json:"hum"`
			Wspdm      string `json:"wspdm"`
			Wspdi      string `json:"wspdi"`
			Wgustm     string `json:"wgustm"`
			Wgusti     string `json:"wgusti"`
			Wdird      string `json:"wdird"`
			Wdire      string `json:"wdire"`
			Vism       string `json:"vism"`
			Visi       string `json:"visi"`
			Pressurem  string `json:"pressurem"`
			Pressurei  string `json:"pressurei"`
			Windchillm string `json:"windchillm"`
			Windchilli string `json:"windchilli"`
			Heatindexm string `json:"heatindexm"`
			Heatindexi string `json:"heatindexi"`
			Precipm    string `json:"precipm"`
			Precipi    string `json:"precipi"`
			Conds      string `json:"conds"`
			Icon       string `json:"icon"`
			Fog        string `json:"fog"`
			Rain       string `json:"rain"`
			Snow       string `json:"snow"`
			Hail       string `json:"hail"`
			Thunder    string `json:"thunder"`
			Tornado    string `json:"tornado"`
			Metar      string `json:"metar"`
		} `json:"observations"`
		Dailysummary []struct {
			Date struct {
				Pretty string `json:"pretty"`
				Year   string `json:"year"`
				Mon    string `json:"mon"`
				Mday   string `json:"mday"`
				Hour   string `json:"hour"`
				Min    string `json:"min"`
				Tzname string `json:"tzname"`
			} `json:"date"`
			Fog                                string `json:"fog"`
			Rain                               string `json:"rain"`
			Snow                               string `json:"snow"`
			Snowfallm                          string `json:"snowfallm"`
			Snowfalli                          string `json:"snowfalli"`
			Monthtodatesnowfallm               string `json:"monthtodatesnowfallm"`
			Monthtodatesnowfalli               string `json:"monthtodatesnowfalli"`
			Since1julsnowfallm                 string `json:"since1julsnowfallm"`
			Since1julsnowfalli                 string `json:"since1julsnowfalli"`
			Snowdepthm                         string `json:"snowdepthm"`
			Snowdepthi                         string `json:"snowdepthi"`
			Hail                               string `json:"hail"`
			Thunder                            string `json:"thunder"`
			Tornado                            string `json:"tornado"`
			Meantempm                          string `json:"meantempm"`
			Meantempi                          string `json:"meantempi"`
			Meandewptm                         string `json:"meandewptm"`
			Meandewpti                         string `json:"meandewpti"`
			Meanpressurem                      string `json:"meanpressurem"`
			Meanpressurei                      string `json:"meanpressurei"`
			Meanwindspdm                       string `json:"meanwindspdm"`
			Meanwindspdi                       string `json:"meanwindspdi"`
			Meanwdire                          string `json:"meanwdire"`
			Meanwdird                          string `json:"meanwdird"`
			Meanvism                           string `json:"meanvism"`
			Meanvisi                           string `json:"meanvisi"`
			Humidity                           string `json:"humidity"`
			Maxtempm                           string `json:"maxtempm"`
			Maxtempi                           string `json:"maxtempi"`
			Mintempm                           string `json:"mintempm"`
			Mintempi                           string `json:"mintempi"`
			Maxhumidity                        string `json:"maxhumidity"`
			Minhumidity                        string `json:"minhumidity"`
			Maxdewptm                          string `json:"maxdewptm"`
			Maxdewpti                          string `json:"maxdewpti"`
			Mindewptm                          string `json:"mindewptm"`
			Mindewpti                          string `json:"mindewpti"`
			Maxpressurem                       string `json:"maxpressurem"`
			Maxpressurei                       string `json:"maxpressurei"`
			Minpressurem                       string `json:"minpressurem"`
			Minpressurei                       string `json:"minpressurei"`
			Maxwspdm                           string `json:"maxwspdm"`
			Maxwspdi                           string `json:"maxwspdi"`
			Minwspdm                           string `json:"minwspdm"`
			Minwspdi                           string `json:"minwspdi"`
			Maxvism                            string `json:"maxvism"`
			Maxvisi                            string `json:"maxvisi"`
			Minvism                            string `json:"minvism"`
			Minvisi                            string `json:"minvisi"`
			Gdegreedays                        string `json:"gdegreedays"`
			Heatingdegreedays                  string `json:"heatingdegreedays"`
			Coolingdegreedays                  string `json:"coolingdegreedays"`
			Precipm                            string `json:"precipm"`
			Precipi                            string `json:"precipi"`
			Precipsource                       string `json:"precipsource"`
			Heatingdegreedaysnormal            string `json:"heatingdegreedaysnormal"`
			Monthtodateheatingdegreedays       string `json:"monthtodateheatingdegreedays"`
			Monthtodateheatingdegreedaysnormal string `json:"monthtodateheatingdegreedaysnormal"`
			Since1sepheatingdegreedays         string `json:"since1sepheatingdegreedays"`
			Since1sepheatingdegreedaysnormal   string `json:"since1sepheatingdegreedaysnormal"`
			Since1julheatingdegreedays         string `json:"since1julheatingdegreedays"`
			Since1julheatingdegreedaysnormal   string `json:"since1julheatingdegreedaysnormal"`
			Coolingdegreedaysnormal            string `json:"coolingdegreedaysnormal"`
			Monthtodatecoolingdegreedays       string `json:"monthtodatecoolingdegreedays"`
			Monthtodatecoolingdegreedaysnormal string `json:"monthtodatecoolingdegreedaysnormal"`
			Since1sepcoolingdegreedays         string `json:"since1sepcoolingdegreedays"`
			Since1sepcoolingdegreedaysnormal   string `json:"since1sepcoolingdegreedaysnormal"`
			Since1jancoolingdegreedays         string `json:"since1jancoolingdegreedays"`
			Since1jancoolingdegreedaysnormal   string `json:"since1jancoolingdegreedaysnormal"`
		} `json:"dailysummary"`
	} `json:"history"`
}

// This only writes out to the web browser...
// We can do a handler for each diff. data set
func handler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if len(data) < -1 {
		fmt.Printf("%s", data)
	}
	var h WeatherHistory
	err = json.Unmarshal(data, &h)
	if err != nil {
		panic(err)
	}
	InsertToDatabase(h)

}

// TODO Add Basic Auth to check for user and password
// TODO Add Rate Limiting and Throttling
func main() {
	http.HandleFunc("/", handler)

	// TODO set port in a config
	http.ListenAndServe(":8080", nil)
}

// TODO Set Mongodb in config
func InsertToDatabase(data WeatherHistory) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

  // Database name and collection, these need to be
  // hardcoded because we need them to always be the
  // same to have consitency across the applications
	c := session.DB("weatherunderground").C("history")

	err = c.Insert(&data)
	if err != nil {
		log.Fatal(err)
	}

}
