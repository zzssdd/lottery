<template>
  <div>
  
    <el-table
    :data="productlist"
    border
    style="width: 100%">
    <el-table-column label="ID" prop="id" align="center"  width="100"></el-table-column>
        <el-table-column
      prop="name"
      label="奖品名称"
      align="center" 
      width="500">
    </el-table-column>
     <el-table-column
      prop="email"
      label="中奖用户"
      align="center" 
      width="400">
    </el-table-column>
    <el-table-column
      label="中奖日期"
      align="center" 
      width="400">
      <template slot-scope="scope">
        <span>{{scope.row.CreatTime|DateForm}}</span>
      </template>
    </el-table-column>
      <el-table-column
      fixed="right"
      label="操作"
      align="center" 
      width="200">
      <template slot-scope="scope">
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
    name:'OrderList',
    data(){
      return {
        productlist:[],
         key:'',
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
        const {data:res}=await this.$http.get('order/list',{
          params:{
            pageSize:this.pageSize,
            pageNum:this.currentPage
          }
        })
        this.total=res.count
        console.log(res);
        if(res.code!==200){
            window.sessionStorage.clear()
            this.$router.push('/login')
          this.$message.error(res.$message)
        }
        this.productlist=res.data
        this.total=res.count
      },
      handleClick(row){
        router.push(`/addlife/${row.ID}`)
      },
      deleteClick(id){
       this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`order/del`,{
            params:{
              id:id
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
      },
      async deleteClick(id){
        this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`order/del`,{
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