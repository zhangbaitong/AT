package oauth

import (
	"common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type AtStorage struct {
	clients   map[string]osin.Client
	authorize map[string]*osin.AuthorizeData
	access    map[string]*osin.AccessData
	refresh   map[string]string
}

func NewATStorage() *AtStorage {
	r := &AtStorage{}

	conn := common.GetRedisPool().Get()
	_, err := conn.Do("set", "1234", "1234")
	if err != nil {
		fmt.Println(err)
	}
	conn.Do("set", "1234:secret", "aabbccdd")
	conn.Do("set", "1234:redirecturl", "http://localhost:14000/appauth")
	defer conn.Close()

	return r
}

func (s *AtStorage) Clone() osin.Storage {
	return s
}

func (s *AtStorage) Close() {
}

// client
// key
// 	- id
// 	- id:secret
// 	- id:redirecturl

// - authorize:code

func (s *AtStorage) GetClient(id string) (osin.Client, error) {
	fmt.Printf("GetClient: %s\n", id)
	if id == "" {
		return nil, nil
	}
	conn := common.GetRedisPool().Get()
	cid, err := redis.String(conn.Do("get", id))
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	if cid == "" {
		return nil, errors.New("Client not found")
	}
	csecret, _ := redis.String(conn.Do("get", id+":secret"))
	credirecturl, _ := redis.String(conn.Do("get", id+":redirecturl"))
	defer conn.Close()
	return &osin.DefaultClient{
		Id:          cid,
		Secret:      csecret,
		RedirectUri: credirecturl,
	}, nil
}

func (s *AtStorage) SetClient(id string, client osin.Client) error {
	fmt.Printf("SetClient: %s\n", id)

	conn := common.GetRedisPool().Get()
	_, err := conn.Do("set", id, id)
	if err != nil {
		fmt.Println(err)
		return errors.New("Client set failed")
	}
	conn.Do("set", id+":secret", client.GetSecret())
	conn.Do("set", id+":redirecturl", client.GetRedirectUri())
	defer conn.Close()
	return nil
}

func (s *AtStorage) SaveAuthorize(data *osin.AuthorizeData) error {
	fmt.Printf("SaveAuthorize: %s\n", data.Code)

	conn := common.GetRedisPool().Get()
	ret, _ := json.Marshal(data)
	conn.Do("set", "authorize:"+data.Code, ret)
	defer conn.Close()
	return nil
}

func (s *AtStorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	fmt.Printf("LoadAuthorize: %s\n", code)

	conn := common.GetRedisPool().Get()
	codeData, err := redis.String(conn.Do("get", "authorize:"+code))
	if err != nil {
		return nil, errors.New("Authorize not found")
	}
	defer conn.Close()
	var authorizeData *osin.AuthorizeData
	json.Unmarshal([]byte(codeData), &authorizeData)

	return authorizeData, nil
}

func (s *AtStorage) RemoveAuthorize(code string) error {
	fmt.Printf("RemoveAuthorize: %s\n", code)

	conn := common.GetRedisPool().Get()
	_, err := conn.Do("del", "authorize:"+code)
	if err != nil {
		return errors.New("Del Authorize Faild")
	}
	defer conn.Close()
	return nil
}

func (s *AtStorage) SaveAccess(data *osin.AccessData) error {
	fmt.Printf("SaveAccess: %s\n", data.AccessToken)

	conn := common.GetRedisPool().Get()
	conn.Do("set", "access:"+data.AccessToken, data)
	defer conn.Close()
	if data.RefreshToken != "" {
		conn.Do("set", "refresh:"+data.RefreshToken, data.AccessToken)
	}
	return nil
}

func (s *AtStorage) LoadAccess(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadAccess: %s\n", code)

	conn := common.GetRedisPool().Get()
	codeData, err := redis.String(conn.Do("get", "access:"+code))
	if err != nil {
		return nil, errors.New("access not found")
	}
	defer conn.Close()
	var accessData *osin.AccessData
	json.Unmarshal([]byte(codeData), &accessData)

	return accessData, nil

}

func (s *AtStorage) RemoveAccess(code string) error {
	fmt.Printf("RemoveAccess: %s\n", code)

	conn := common.GetRedisPool().Get()
	_, err := conn.Do("del", "access:"+code)
	if err != nil {
		return errors.New("Del access Faild")
	}
	defer conn.Close()
	return nil

}

func (s *AtStorage) LoadRefresh(code string) (*osin.AccessData, error) {
	fmt.Printf("LoadRefresh: %s\n", code)

	conn := common.GetRedisPool().Get()
	refreshData, err := redis.String(conn.Do("set", "refresh:"+code))
	if err != nil {
		return nil, errors.New("Refresh not found")
	}
	defer conn.Close()
	return s.LoadAccess(refreshData)
}

func (s *AtStorage) RemoveRefresh(code string) error {
	fmt.Printf("RemoveRefresh: %s\n", code)

	conn := common.GetRedisPool().Get()
	_, err := conn.Do("del", "refresh:"+code)
	if err != nil {
		return errors.New("Del access Faild")
	}
	defer conn.Close()
	return nil
}
