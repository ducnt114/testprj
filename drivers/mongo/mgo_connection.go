package mongo

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type MongoConnection struct {
	Session *mgo.Session
}

func NewConnection() (*MongoConnection, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:       []string{fmt.Sprintf("%v:%v", viper.GetString("mongo.host"), viper.GetInt64("mongo.port"))},
		Database:    viper.GetString("mongo.db"),
		FailFast:    viper.GetBool("mongo.fail_fast"),
		Username:    "",
		Password:    "",
		Service:     "",
		ServiceHost: "",
		Mechanism:   "",
		PoolLimit:   viper.GetInt("mongo.pool_limit"),
		Direct:      false,
		Timeout: time.Second * func() time.Duration {
			if viper.GetDuration("mongo.timeout") <= 0 {
				return 5
			}
			return viper.GetDuration("mongo.timeout")
		}(),
	}

	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Printf("Could not create connection to mongodb %v:%d/%v with timeout %v seconds with error %v",
			viper.GetString("mongo.host"), viper.GetString("mongo.port"),
			viper.GetString("mongo.db"), viper.GetDuration("mongo.timeout"), err)
		return nil, err
	}
	err = session.Ping()
	if err != nil {
		log.Printf("Could not ping to mongodb database, details: %v", err)
		return nil, err
	}
	return &MongoConnection{Session: session}, nil
}
