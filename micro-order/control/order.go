package control

import (
	"context"
	"micro-order/model"
	"micro-order/proto"
)

type Order struct {
}

func (o *Order) OrderList(ctx context.Context, in *proto.ListRequest, out *proto.ListResponse) error {
	orders, count, err := model.OrderList(int(in.PageNum), int(in.PageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "获取订单列表失败！"
		return err
	}
	rep_datas := []*proto.OrderInfo{}
	for _, v := range orders {
		prize, _ := model.PrizeInfo(v.Name)
		rep_data := &proto.OrderInfo{
			Id:        int32(v.ID),
			Email:     v.Email,
			Name:      prize.Name,
			Pic:       prize.Pic,
			CreatTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Code = 200
	out.Msg = "获取列表成功"
	out.Orders = rep_datas
	out.Count = int32(count)
	return nil
}
func (o *Order) OrderUser(ctx context.Context, in *proto.UserRequest, out *proto.UserResponse) error {
	orders, count, err := model.OrderByEmail(in.Email, int(in.PageNum), int(in.PageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "获取用户订单失败"
		return err
	}
	rep_datas := []*proto.UserOrder{}
	for _, v := range orders {
		prize, _ := model.PrizeInfo(v.Name)
		rep_data := &proto.UserOrder{
			Name:      prize.Name,
			Pic:       prize.Pic,
			CreatTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Code = 200
	out.Msg = "获取用户订单成功"
	out.Orders = rep_datas
	out.Count = int32(count)
	return nil
}
func (o *Order) OrderAdd(ctx context.Context, in *proto.AddRequest, out *proto.Response) error {
	data := &model.Order{
		Email: in.Email,
		Name:  in.Name,
	}
	err := model.OrderAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "订单增加成功"
		return err
	}
	out.Code = 200
	out.Msg = "订单增加成功"
	return nil
}
func (o *Order) OrderDel(ctx context.Context, in *proto.IdRequest, out *proto.Response) error {
	err := model.OrderDel(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "订单删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "订单删除成功"
	return nil
}
