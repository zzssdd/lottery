package control

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"micro-choose/model"
	"micro-choose/proto"
)

type Choose struct {
}

func (c *Choose) DoChoose(ctx context.Context, in *proto.Request, out *proto.Response) error {
	prize, err := model.Choose()
	if err != nil {
		out.Code = 500
		out.Msg = "奖品已抽完"
		return err
	}
	data := &model.Order{
		Email: in.Email,
		Name:  prize,
	}
	err = model.OrderAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "创建订单失败"
		return err
	}
	out.Code = 200
	out.Msg = "抽奖成功"
	out.Name = prize
	return nil
}

func init() {
	orders, err := model.Ch.Consume("order_queue", "order_consumer", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for order := range orders {
		go OrderApply(order)
	}
}

func OrderApply(order amqp.Delivery) {
	body := string(order.Body)
	byte_data := map[string]interface{}{}
	err := json.Unmarshal([]byte(body), &byte_data)
	if err != nil {
		panic(err)
	}
	email := byte_data["email"].(string)
	prize, err := model.Choose()
	if err != nil {
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "奖品已抽完",
		}
		rep_data, _ := json.Marshal(rep_map)
		model.Set(email, string(rep_data))
		return
	}
	data := &model.Order{
		Email: email,
		Name:  prize,
	}
	err = model.OrderAdd(data)
	if err != nil {
		rep_map := map[string]interface{}{
			"code": 500,
			"msg":  "创建订单失败",
		}
		rep_data, _ := json.Marshal(rep_map)
		model.Set(email, string(rep_data))
		return
	}
	rep_map := map[string]interface{}{
		"code":  200,
		"msg":   "成功",
		"prize": prize,
	}
	rep_data, _ := json.Marshal(rep_map)
	model.Set(email, string(rep_data))
	return
}
