<template>
  <div>
    <el-table :data="UserList" border style="width: 100%">
      <el-table-column label="ID" align="center"  prop="id" width="250"></el-table-column>
      <el-table-column label="注册时间" align="center"  width="400">
        <template slot-scope="scope" align="center" >
          <span>{{ scope.row.createTime | DateForm }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="email" align="center" label="邮箱" width="400">
      </el-table-column>
      <el-table-column   align="center" label="剩余积分" width="300">
        <template slot-scope="scope">
        <span v-show="scope.row.id!==editId">{{scope.row.score}}</span>
        <input type="text" v-show="scope.row.id===editId" v-model="scope.row.score" @blur="FinEdit(scope.row)">
        </template>
      </el-table-column>
      <el-table-column
      fixed="right"
      label="操作"
      align="center" 
      width="300">
        <template slot-scope="scope">
        <el-button @click="editClick(scope.row.id)" type="primary"  plain size="small">编辑</el-button>
        <el-button @click="deleteClick(scope.row.id)" type="danger"  plain size="small">删除</el-button>
      </template>
    </el-table-column>
    </el-table>
      <div class="block">
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 40]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
  </div>
  </div>
  
</template>

<script>
import day from "dayjs";

export default {
  name: "Users",
  data() {
    return {
      UserList: [],
      total:0,
      currentPage:1,
      pageSize:10,
      editId:-1,
    };
  },
  filters: {
    DateForm: function (date) {
      return date ? day(date).format("YYYY年MM月DD日 HH:mm") : "暂无";
    },
  },
  methods: {
    async getUsersList() {
      const { data: res } = await this.$http.get("user/list",{
        params:{
          pageNum:this.currentPage,
          pageSize:this.pageSize,
        }
      });
      console.log(res);
      if (res.code !== 200) {
          // window.sessionStorage.clear();
          // this.$router.push("/login");
        this.$message.error(res.$message);
      }
      this.UserList = res.data;
      this.total=res.count
    },
    editClick(id){
      this.editId=id
    },
    handleSizeChange(val){
      this.pageSize=val
      this.getUsersList()
    },
    handleCurrentChange(val){
      this.currentPage=val
      this.getUsersList()
    },
    async FinEdit(row){
      this.editId=-1
      var form_data=new FormData
      form_data.append("id",row.id)
      form_data.append("score",row.score)
      console.log(row.id);
      console.log(row.score);
      const {data:res}=await this.$http.put(`user/edit`,form_data)
      console.log(res);
        if (res.code !== 200) {
          // window.sessionStorage.clear();
          // this.$router.push("/login");
        this.$message.error(res.msg);
      }
      this.$message({
            type: 'success',
            message: res.msg
          });
    },
    async deleteClick(id){
        const { data: res } = await this.$http.delete("user/del",{
        params:{
          id
        }
      });
      console.log(res);
      if (res.code !== 200) {
          // window.sessionStorage.clear();
          // this.$router.push("/login");
        this.$message.error(res.$message);
      }
      this.getUsersList()
    },
  },
  created() {
    this.getUsersList();
  },
};
</script>

<style scoped>
.el-image {
  width: 200px;
  height: 100px;
}
</style>
