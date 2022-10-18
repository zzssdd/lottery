package control

import (
	"context"
	"fmt"
	"micro-acticity/model"
	"micro-acticity/proto"
	"time"
)

type Activity struct {
}

func (a *Activity) ActiveAdd(ctx context.Context, in *proto.Request, out *proto.AddResponse) error {
	start, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	end, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	data := &model.Activity{
		Name:  in.Name,
		Desc:  in.Desc,
		Pic:   in.Pic,
		User:  in.User,
		Cost:  int(in.Cost),
		Type:  int(in.Type),
		Start: start,
		End:   end,
	}
	err, id := model.ActiveAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "添加活动失败"
		return err
	}
	out.Code = 200
	out.Msg = "添加活动成功"
	out.Id = int32(id)
	return nil
}
func (a *Activity) ActiveDel(ctx context.Context, in *proto.IdRequest, out *proto.Response) error {
	err := model.ActiveDel(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "删除活动失败"
		return err
	}
	out.Code = 200
	out.Msg = "删除活动成功"
	return nil
}
func (a *Activity) ActiveEdit(ctx context.Context, in *proto.EditRequest, out *proto.Response) error {
	start, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	end, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	data := &model.Activity{
		Name:  in.Name,
		Desc:  in.Desc,
		Pic:   in.Pic,
		User:  in.User,
		Cost:  int(in.Cost),
		Type:  int(in.Type),
		Start: start,
		End:   end,
	}
	err := model.ActiveEdit(data, int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "修改活动失败"
		return err
	}
	out.Code = 200
	out.Msg = "修改活动成功"
	return nil
}
func (a *Activity) ActiveList(ctx context.Context, in *proto.ListRequest, out *proto.ListResponse) error {
	actives, count, err := model.ActiveList(int(in.PageNum), int(in.PageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "获取活动列表失败"
		return err
	}
	rep_datas := []*proto.ActiveInfo{}
	for _, v := range actives {
		rep_data := &proto.ActiveInfo{
			Id:        int32(v.ID),
			Name:      v.Name,
			Desc:      v.Desc,
			Pic:       v.Pic,
			User:      v.User,
			Cost:      int32(v.Cost),
			Type:      int32(v.Type),
			StartTime: v.Start.Format("2006-01-02 15:04:05"),
			EndTime:   v.End.Format("2006-01-02 15:04:05"),
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Code = 200
	out.Msg = "获取活动列表成功"
	out.Acts = rep_datas
	out.Count = int32(count)
	return nil
}
func (a *Activity) ActiveInfo(ctx context.Context, in *proto.IdRequest, out *proto.InfoResponse) error {
	active, err := model.ActiveInfo(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "获取活动信息失败"
		return err
	}
	rep_data := &proto.ActiveInfo{
		Id:        int32(active.ID),
		Name:      active.Name,
		Desc:      active.Desc,
		Pic:       active.Pic,
		User:      active.User,
		Cost:      int32(active.Cost),
		Type:      int32(active.Type),
		StartTime: active.Start.Format("2006-01-02 15:04:05"),
		EndTime:   active.End.Format("2006-01-02 15:04:05"),
	}
	prizes, err := model.ActivePrizes(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "获取活动奖品失败"
		return err
	}
	rep_prizes := []*proto.PrizeInfo{}
	for _, v := range prizes {
		rep_prize := &proto.PrizeInfo{
			Name:       v.Name,
			Num:        int32(v.Num),
			Pic:        v.Pic,
			CreateTime: v.CreateTime,
		}
		rep_prizes = append(rep_prizes, rep_prize)
	}
	out.Code = 200
	out.Msg = "获取活动信息成功"
	out.Prizes = rep_prizes
	out.Act = rep_data
	return nil
}

func (a *Activity) ActivePrizes(ctx context.Context, in *proto.IdRequest, out *proto.PrizesResponse) error {
	prizes, err := model.ActivePrizes(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "获取活动奖品失败"
		return err
	}
	rep_prizes := []*proto.PrizeInfo{}
	for _, v := range prizes {
		rep_prize := &proto.PrizeInfo{
			Name:       v.Name,
			Num:        int32(v.Num),
			Pic:        v.Pic,
			CreateTime: v.CreateTime,
		}
		rep_prizes = append(rep_prizes, rep_prize)
	}
	out.Code = 200
	out.Msg = "获取活动奖品成功"
	out.Prizes = rep_prizes
	return nil
}
func (a *Activity) ActiveAddPrize(ctx context.Context, in *proto.PrizeAddRequest, out *proto.Response) error {
	err := model.ActiveAddPrize(int(in.Id), in.Name)
	if err != nil {
		fmt.Println(err)
		out.Code = 500
		out.Msg = "增加活动奖品失败"
		return err
	}
	out.Code = 200
	out.Msg = "增加活动奖品成功"
	return nil
}
