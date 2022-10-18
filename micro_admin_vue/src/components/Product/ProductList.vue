<template>
  <div>
    <el-row>
      <el-col :span="5">
        <el-button type="primary" @click="$router.push('/addProduct')">新增</el-button>
      </el-col>
    </el-row>
  
    <el-table
    :data="productlist"
    border
    style="width: 100%">
        <el-table-column
      prop="name"
      label="奖品名称"
      align="center" 
      width="400">
    </el-table-column>
      <el-table-column
      prop="num"
      label="剩余数量"
      align="center" 
      width="200">
    </el-table-column>
    <el-table-column
      label="添加日期"
      align="center" 
      width="400">
      <template slot-scope="scope">
        <span>{{scope.row.createTime|DateForm}}</span>
      </template>
    </el-table-column>
    <el-table-column
      prop="pic"
      label="封面"
      align="center" 
      width="450">
      <template slot-scope="scope">
      <el-image :src="scope.row.pic" lazy></el-image>  
      </template>
    </el-table-column>
    <el-table-column
      fixed="right"
      label="操作"
      align="center" 
      width="200">
      <template slot-scope="scope">
        <el-button @click="$router.push(`/putProduct/${scope.row.name}`)" type="primary" plain size="small">编辑</el-button>
        <el-button @click="deleteClick(scope.row.name)" type="danger"  plain size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  </div>
  
</template>

<script>
import day from 'dayjs'
import router from '@/router'

export default {
    name:'ProductList',
    data(){
      return {
        productlist:[],
         key:''
      }
    },
    filters:{
      DateForm:function(date){
        return date?day(date).format('YYYY年MM月DD日 HH:mm'):'暂无'
      }
    },
    methods:{
      async getList(){
        const {data:res}=await this.$http.get('prize/list')
        console.log(res);
        if(res.code!==200){
            window.sessionStorage.clear()
            this.$router.push('/login')
          this.$message.error(res.$message)
        }
        console.log(res);
        this.productlist=res.data
        this.total=res.count
      },
      handleClick(row){
        router.push(`/addProduct/${row.ID}`)
      },
      deleteClick(name){
       this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const {data:res}=await this.$http.delete(`prize/del`,{
            params:{
              name
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