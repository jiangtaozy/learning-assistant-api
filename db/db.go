/*
 * Maintained by jemo from 2020.2.15 to now
 * Created by jemo on 2020.2.15 11:01:55
 * Db
 */

package db

import (
  "time"
  "go.etcd.io/etcd/clientv3"
)

var Etcd clientv3.Client

func Init() {
  var err error
  Etcd, err = clientv3.New(clientv3.Config{
    DialTimeout: 2 * time.Second,
    Endpoints: []string{"127.0.0.1:2379"},
  })
  if err != nil {
    log.Fatalln(err.Error())
  }
  defer Etcd.Close()
}
