/**
 * @Author: huangw1
 * @Date: 2019/6/14 17:53
 */

package nosql

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func String(replay interface{}, err error) (string, error) {
	return redis.String(replay, err)
}

func Int(replay interface{}, err error) (int, error) {
	return redis.Int(replay, err)
}

func Int64(replay interface{}, err error) (int64, error) {
	return redis.Int64(replay, err)
}

func Bool(replay interface{}, err error) (bool, error) {
	return redis.Bool(replay, err)
}

type Nosql struct {
	pool      *redis.Pool
	prefix    string
	marshal   func(v interface{}) ([]byte, error)
	unmarshal func([]byte, interface{}) error
}

type Options struct {
	Network     string
	Addr        string
	Password    string
	Db          int
	MaxActive   int
	MaxIdle     int
	IdleTimeout int
	Prefix      string
	Marshal     func(v interface{}) ([]byte, error)
	Unmarshal   func([]byte, interface{}) error
}

func New(options Options) *Nosql {
	nosql := new(Nosql)
	nosql.RunAndGC(options)
	return nosql
}

func (n *Nosql) RunAndGC(options Options) {
	if options.Network == "" {
		options.Network = "tcp"
	}
	if options.Addr == "" {
		options.Addr = "127.0.0.1:6379"
	}
	if options.MaxIdle == 0 {
		options.MaxIdle = 3
	}
	if options.IdleTimeout == 0 {
		options.IdleTimeout = 300
	}
	if options.Marshal == nil {
		options.Marshal = json.Marshal
	}
	if options.Unmarshal == nil {
		options.Unmarshal = json.Unmarshal
	}
	pool := &redis.Pool{
		MaxActive:   options.MaxActive,
		MaxIdle:     options.MaxIdle,
		IdleTimeout: time.Duration(options.IdleTimeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(options.Network, options.Addr)
			if err != nil {
				return nil, err
			}
			if options.Password != "" {
				if _, err := conn.Do("AUTH", options.Password); err != nil {
					conn.Close()
					return nil, err
				}
			}
			if _, err := conn.Do("SELECT", options.Db); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},

		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
	n.pool = pool
}

func (n *Nosql) Do(command string, args ...interface{}) (reply interface{}, err error) {
	conn := n.pool.Get()
	defer conn.Close()
	return conn.Do(command, args...)
}

func (n *Nosql) getKey(key string) string {
	return n.prefix + key
}

func (n *Nosql) Get(key string) (replay interface{}, err error) {
	return n.Do("GET", n.getKey(key))
}

func (n *Nosql) GetString(key string) (string, error) {
	return String(n.Get(key))
}

func (n *Nosql) GetInt(key string) (int, error) {
	return Int(n.Get(key))
}

func (n *Nosql) GetInt64(key string) (int64, error) {
	return Int64(n.Get(key))
}

func (n *Nosql) GetBool(key string) (bool, error) {
	return Bool(n.Get(key))
}

func (n *Nosql) GetObject(key string, val interface{}) error {
	replay, err := n.Get(key)
	return n.decode(replay, err, val)
}

func (n *Nosql) Set(key string, val interface{}, expire int) error {
	v, err := n.encode(val)
	if err != nil {
		return err
	}
	if expire > 0 {
		_, err := n.Do("SETEX", n.getKey(key), expire, v)
		return err
	}
	_, err = n.Do("SET", n.getKey(key), v)
	return err
}

func (n *Nosql) Del(key string) error {
	_, err := n.Do("DEL", n.getKey(key))
	return err
}

func (n *Nosql) Exists(key string) (bool, error) {
	return Bool(n.Do("EXISTS", n.getKey(key)))
}

func (n *Nosql) Flush() error {
	_, err := n.Do("FLUSH")
	return err
}

func (n *Nosql) Incr(key string) (val int64, err error) {
	return Int64(n.Do("INCR", n.getKey(key)))
}

func (n *Nosql) IncrBy(key string, amount int64) (val int64, err error) {
	return Int64(n.Do("INCRBY", n.getKey(key), amount))
}

func (n *Nosql) Decr(key string) (val int64, err error) {
	return Int64(n.Do("DECR", n.getKey(key)))
}

func (n *Nosql) DecrBy(key string, amount int64) (val int64, err error) {
	return Int64(n.Do("DECRBY", n.getKey(key), amount))
}

func (n *Nosql) HMset(key string, val interface{}, expire int) error {
	conn := n.pool.Get()
	defer conn.Close()
	err := conn.Send("HMSET", redis.Args{}.Add(n.getKey(key)).AddFlat(val)...)
	if err != nil {
		return err
	}
	if expire > 0 {
		err = conn.Send("EXPIRE", n.getKey(key), expire)
	}
	if err != nil {
		return err
	}
	_, err = conn.Receive()
	return err
}

func (n *Nosql) HSet(key string, val interface{}) (interface{}, error) {
	v, err := n.encode(val)
	if err != nil {
		return nil, err
	}
	return n.Do("HSET", n.getKey(key), v)
}

func (n *Nosql) HGet(key string) (interface{}, error) {
	return n.Do("HGET", n.getKey(key))
}

func (n *Nosql) HGetString(key string) (string, error) {
	return String(n.HGet(key))
}

func (n *Nosql) HGetInt(key string) (int, error) {
	return Int(n.HGet(key))
}

func (n *Nosql) HGetInt64(key string) (int64, error) {
	return Int64(n.HGet(key))
}

func (n *Nosql) HGetBool(key string) (bool, error) {
	return Bool(n.HGet(key))
}

func (n *Nosql) HGetObject(key string, val interface{}) error {
	str, err := n.HGet(key)
	return n.decode(str, err, val)
}

func (n *Nosql) HGetAll(key string, val interface{}) error {
	v, err := redis.Values(n.Do("HGETALL", n.getKey(key)))
	if err != nil {
		return err
	}
	return redis.ScanStruct(v, val)
}

func (n *Nosql) decode(replay interface{}, err error, val interface{}) error {
	str, err := String(replay, err)
	if err != nil {
		return err
	}
	return n.unmarshal([]byte(str), val)
}

func (n *Nosql) encode(replay interface{}) (interface{}, error) {
	switch t := replay.(type) {
	case string, int, uint, int8, int16, int32, int64, float32, float64, bool:
		return t, nil
	default:
		b, err := n.marshal(t)
		if err != nil {
			return nil, err
		}
		return string(b), nil
	}
}

func (n *Nosql) closePool(options Options) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGKILL)
	go func() {
		<-quit
		n.pool.Close()
		os.Exit(1)
	}()

}
