/*
 * Maintained by jemo from 2020.2.22 to now
 * Created by jemo on 2020.2.22 23:30:37
 * Db
 */

package config

import (
  "strings"
  "io/ioutil"
)

var AlismsAccessKeyId string
var AlismsAccessKeySecret string

func Init() {
  if alismsAccessKey, err := ioutil.ReadFile("pem/alismsAccessKey.csv"); err == nil {
    alismsAccessKeyArray := strings.Split(strings.Split(string(alismsAccessKey), "\n")[1], ",")
    AlismsAccessKeyId = alismsAccessKeyArray[0]
    AlismsAccessKeySecret = strings.Trim(alismsAccessKeyArray[1], "\r")
  } else {
    panic(err)
  }
}
