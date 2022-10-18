<template>
  <el-card class="box-card">
    <el-form
      ref="artForm"
      :model="productInfo"
      :inline="true"
      label-position="top"
      :rules="productRules"
    >
      <el-row>
        <el-col :span="20">
          <el-form-item label="奖品名称：">
            <el-input
              size="big"
              v-model="productInfo.name"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
          <el-col :span="20">
          <el-form-item label="奖品数量：">
            <el-input
              size="big"
              v-model="productInfo.num"
              style="width: 250px"
            ></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="8">
          <el-form-item label="奖品封面"  >
           <input type="file"/>
          </el-form-item>
          <el-image :src="productInfo.pic" v-show="productInfo.pic"></el-image>
        </el-col>
      </el-row>
    </el-form>
     <div class="btn">
        <el-button type="primary" @click="submitForm"
          >完成</el-button
        >
        <el-button type="danger" @click="back"
          >返回</el-button
        >
        </div>
  </el-card>
</template>

<script>
export default {
  name: "AddProduct",
  props:['name'],
  data() {
    return {
      productInfo: {
        name: '',
        desc: '',
        price:undefined,
        unit:'',
        num:undefined,
        pic: '',
      },
       productRules: {
        name: [{ required: true, message: '请输入奖品名称', trigger: 'change' }],
      },
    };
  },
  methods: {
    async getInfo(name){
      const {data:res}=await this.$http.get(`prize/info`,{
        params:{
          name
        }
      })
      console.log(res);
      if(res.code!=200){
        this.$message.error(res.message)
        return
      }
      this.productInfo=res.data;
    },
    async submitForm(){
          this.$refs.artForm.validate(async (valid) => {
          if(valid){
          const form_data=new FormData();
          form_data.append("name",this.productInfo.name)
          form_data.append("num",this.productInfo.num)
        form_data.append("pic",document.querySelector('input[type=file]').files[0])
          if(this.name){
          const { data: res } = await this.$http.put(`prize/edit`, form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$router.push('/productList')
          this.$message({
            type:'success',
            message:'更新奖品成功！'
          })
        }else{
          const { data: res } = await this.$http.post('prize/add',form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$router.push('/productList')
          this.$message({
            type:'success',
            message:'添加奖品成功！'
          })
        };
        }else{
          return this.$message.error('输入数据非法，请重新输入')
        }})
    },
    back(){
      this.$router.push('/productList')
    }
  },
  created(){
    if(this.name){
      this.getInfo(this.name)
    }
  }
};
</script>

<style scoped>
</style>
