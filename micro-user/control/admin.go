package control

import (
	"context"
	"micro-user/model"
	"micro-user/proto/admin"
	"micro-user/utils"
)

type Admin struct {
}

func (a *Admin) AdminLogin(ctx context.Context, in *admin.LoginRequest, out *admin.LoginResponse) error {
	data := &model.Admin{
		Username: in.Username,
		Password: utils.Md5Pw(in.Password),
	}
	is_ok := model.AdminLogin(data)
	if !is_ok {
		out.Code = 500
		out.Msg = "账号密码输入错误"
		return nil
	}
	out.Code = 200
	out.Msg = "登陆成功"
	return nil
}
func (a *Admin) UserList(ctx context.Context, in *admin.UsersRequest, out *admin.UsersResponse) error {
	count, users, err := model.UserList(int(in.PageNum), int(in.PageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "用户列表获取失败"
		return err
	}
	rep_datas := []*admin.UserInfo{}
	for _, v := range users {
		rep_data := &admin.UserInfo{
			Id:         int32(v.ID),
			Email:      v.Email,
			Score:      int32(v.Score),
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_datas = append(rep_datas, rep_data)
	}
	out.Users = rep_datas
	out.Code = 200
	out.Msg = "用户列表获取成功"
	out.Count = int32(count)
	return nil
}
func (a *Admin) UserDel(ctx context.Context, in *admin.DelRequest, out *admin.DelResponse) error {
	err := model.UserDel(int(in.Id))
	if err != nil {
		out.Code = 500
		out.Msg = "用户删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "用户删除成功"
	return nil
}
