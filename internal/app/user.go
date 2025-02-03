package app

import (
	"context"
	"fmt"
	"strconv"

	pb "cache/pkg/proto"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

// TODO: can use the embed proto message and map to the struct.
func (s *server) GetUsers(_ context.Context, req *pb.GetUserReq) (*pb.GetUsersResp, error) {
	logrus.Infof("GetUser-Request: %+v \n", req)
	var id *int32
	if req.GetId() != nil {
		id = &req.GetId().Value
	} else {
		id = nil
	}

	v := app.Cache.GetUserCache(int(*id))
	logrus.Infoln(v)

	ch := app.Cache.UserCache.Range()
	for k := range ch {
		logrus.Debugf("k=%v, v=%v", k, v)
	}

	logrus.Infof("GetUser-Response: %+v \n", v)
	var d = pb.GetUsersResp{}
	var tmp map[string]interface{}
	var t pb.User

	mapstructure.Decode(v, &tmp)
	logrus.Debug(tmp)
	mapstructure.Decode(tmp, &t)
	logrus.Debug(t)
	d.Users = append(d.Users, &t)
	return &d, nil

	// config := &mapstructure.DecoderConfig{
	// 	TagName: "mapstructure",
	// 	Result:  new(pb.User), // 指定结果为TargetStruct指针
	// }
	// decoder, _ := mapstructure.NewDecoder(config)
	// var d pb.GetUsersResp

	// for _, s := range res {
	// 	var t pb.User
	// 	// 重用Decoder，更新Result为目标元素的地址
	// 	config.Result = &t
	// 	err := decoder.Decode(s)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	logrus.Infoln(t)
	// 	d.Users = append(d.Users, &t)
	// }
	// return &d, nil
}

// RegisterFunc: GetUserCache
func GetUserCache(args ...string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("GetCache need one arg")
	}
	key := args[0]
	vI, err := strconv.Atoi(key)
	if err != nil {
		return "", err
	}
	v := app.Cache.GetUserCache(vI)
	return fmt.Sprint(v), nil
}
