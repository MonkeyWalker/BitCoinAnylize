package myutils

import (
	//"net/http"
	//"io/ioutil"
)
const name = "https://query.yahooapis.com/v1/public/yql" +
	"?q=select * from yahoo.finance.xchange where pair in (\"CNYJPY\")&format=json&diagnostics=true&env=store://datatables.org/alltableswithkeys&callback="

const PREFIX  = "http://developer.yahoo.com/yql/console/?q=select Bid from yahoo.finance.xchange where pair in (";

const SUFFIX  = ")&env=store://datatables.org/alltableswithkeys";
func queryExchange(from string, to string)  {
/*	url := PREFIX + from + to + SUFFIX;
	resp,err := http.Get(url);
	if (err != nil){
		println(err)
	}
	json,err := ioutil.ReadAll(resp.Body);
	if (err != nil){
		println(err)
	}
	json.Unmarshal(json)
	return .*/
}