package control

import (
	"context"
	"fmt"
	"micro-prize/model"
	"micro-prize/proto"
)

type Prize struct {
}

func (p *Prize) PrizeList(ctx context.Context, in *proto.ListRequest, out *proto.ListResponse) error {
	prizes, err := model.PrizeList()
	if err != nil {
		out.Code = 500
		out.Msg = "获取奖品列表失败"
		return err
	}
	rep_datas := []*proto.Prize{}
	for _, v := range prizes {
		rep_data := &proto.Prize{
			Name:       v.Name,
			Num:        int32(v.Num),
			Pic:        v.Pic,
			CreateTime: v.CreateTime,
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Code = 200
	out.Msg = "获取奖品列表成功"
	out.Prizes = rep_datas
	return nil
}

func (p *Prize) PrizeAdd(ctx context.Context, in *proto.Prize, out *proto.Response) error {
	data := &model.Prize{
		Name:       in.Name,
		Num:        int(in.Num),
		Pic:        in.Pic,
		CreateTime: in.CreateTime,
	}
	err := model.PrizeAdd(data)
	if err != nil {
		fmt.Println(err)
		out.Code = 500
		out.Msg = "增加奖品失败"
		return err
	}
	out.Code = 200
	out.Msg = "增加商品成功"
	return nil

}
func (p *Prize) PrizeDel(ctx context.Context, in *proto.NameRequest, out *proto.Response) error {
	err := model.PrizeDel(in.Name)
	if err != nil {
		fmt.Println(err)
		out.Code = 500
		out.Msg = "商品删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "商品删除成功"
	return nil
}
func (p *Prize) PrizeEdit(ctx context.Context, in *proto.Prize, out *proto.Response) error {
	data := &model.Prize{
		Name: in.Name,
		Num:  int(in.Num),
		Pic:  in.Pic,
	}
	err := model.PrizeEdit(in.Name, data)
	if err != nil {
		out.Code = 500
		out.Msg = "修改奖品失败"
		return err
	}
	out.Code = 200
	out.Msg = "修改商品成功"
	return nil
}
func (p *Prize) PrizeInfo(ctx context.Context, in *proto.NameRequest, out *proto.InfoResponse) error {
	prize, err := model.PrizeInfo(in.Name)
	if err != nil {
		out.Code = 500
		out.Msg = "获取奖品信息失败"
		return err
	}
	rep_data := &proto.Prize{
		Name:       prize.Name,
		Num:        int32(prize.Num),
		Pic:        prize.Pic,
		CreateTime: prize.CreateTime,
	}
	out.Code = 200
	out.Msg = "获取商品信息成功"
	out.Prize = rep_data
	return nil
}
