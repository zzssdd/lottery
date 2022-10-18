<template>
  <div>
    <el-row>
      <el-col :span="30">
        <el-button type="primary" @click="$router.push('/addActive')">新增</el-button>
      </el-col>
    </el-row>
  
    <el-table
    :data="activelist"
    border
    style="width: 100%">
    <el-table-column label="活动名称" prop="name" align="center"  width="280"></el-table-column>
      <el-table-column
      prop="cost"
      label="抽奖积分"
      align="center" 
      width="100">
    </el-table-column>
      <el-table-column
      prop="user"
      label="创建者"
      align="center" 
      width="150">
    </el-table-column>
    <el-table-column
      label="活动开始时间"
      align="center" 
      width="220">
      <template slot-scope="scope">
        <span>{{scope.row.startTime|DateForm}}</span>
      </template>
    </el-table-column>
        <el-table-column
      label="活动结束时间"
      align="center" 
      width="220">
      <template slot-scope="scope">
        <span>{{scope.row.endTime|DateForm}}</span>
      </template>
    </el-table-column>
    <el-table-column
      prop="img"
      label="活动封面"
      align="center" 
      width="350">
      <template slot-scope="scope">
      <el-image :src="scope.row.pic" lazy></el-image>  
      </template>
    </el-table-column>
    <el-table-column
      fixed="right"
      label="编辑"
      align="center" 
      width="200">
      <template slot-scope="scope">
        <el-button @click="$router.push(`/putActive/${scope.row.id}`)" type="primary" plain size="small">编辑</el-button>
        <el-button @click="deleteClick(scope.row.id)" type="danger"  plain size="small">下架活动</el-button>
      </template>
    </el-table-column>
  </el-table>
  <div class="block">
      <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 40]"
      :page-size=pageSize
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
    </div>
  </div>
  
</template>

<script>
import day from 'dayjs'
import router from '@/router'

export default {
    name:'ActiveList',
    data(){
      return {
        activelist:[],
        currentPage: 1, // 当前页码
        total: 0, // 总条数
        pageSize: 10 // 每页的数据条数
      }
    },
    filters:{
      DateForm:function(date){
        return date?day(date).format('YYYY年MM月DD日 HH:mm'):'暂无'
      }
    },
    methods:{
      async getList(){
        const {data:res}=await this.$http.get('activity/list',{
          params:{
            pageSize:this.pageSize,
            currentPage:this.currentPage
          }
        })
        console.log(res);
        if(res.code!==200){
            window.sessionStorage.clear()
            this.$router.push('/login')
          this.$message.error(res.$message)
        }
        this.activelist=res.data
        this.total=res.count
      },
      handleClick(row){
        router.push(`/putActive/${row.id}`)
      },
      deleteClick(id){
       this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`activity/del`,{
            params:{
              id
            }
          })
          console.log(res);
          if (res.code != 200) return this.$message.error(res.message)
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          this.getList()
          }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });          
        });
      },
        handleSizeChange(val){
        this.pageSize=val
        this.getList()
      },
      handleCurrentChange(val){
        this.currentPage=val
        this.getList()
      }
    },
    created(){
      this.getList()
    },
}
</script>

<style scoped>
  .el-image{
    width: 200px;
    height: 100px;
  }
  .el-table{
    margin-top: 20px;
  }
</style>