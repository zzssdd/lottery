<template>
  <el-card class="box-card">
    <el-form
      ref="activeForm"
      :model="activeInfo"
      :inline="true"
      label-position="top"
      :rules="activeInfoRules"
    >
      <el-row>
        <el-col :span="7">
          <el-form-item label="活动名称：">
            <el-input size="big" v-model="activeInfo.name" style="width: 300px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="抽奖积分：">
            <el-input size="big" v-model="activeInfo.cost" style="width: 150px"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="活动开始时间：">
      <el-date-picker
      v-model="activeInfo.startTime"
      type="datetime"
      format="yyyy-MM-dd HH:mm:ss"
      placeholder="选择日期时间"
      value-format="timestamp">
    </el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="5">
          <el-form-item label="活动结束时间：">
         <el-date-picker
      v-model="activeInfo.endTime"
      type="datetime"
       format="yyyy-MM-dd HH:mm:ss"
      placeholder="选择日期时间"
      value-format="timestamp">
    </el-date-picker>
          </el-form-item>
        </el-col>

        <el-col :span="7">
          <el-form-item label="抽奖类型：">
            <el-select v-model="activeInfo.type" placeholder="请选择抽奖类型">
              <el-option
                v-for="item in [0,1,2,3]"
                :key="item"
                :label="type[item]"
                :value="item"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>

        <el-col :span="5">
          <div>
            <p> 请选择活动奖品:</p>
          </div>
          <el-checkbox-group v-model="prizesChecked" v-for="item in prizeList" :key="item.name">
          <el-checkbox :label="item.name"></el-checkbox>
          </el-checkbox-group>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="14">
          <el-form-item label="活动描述：">
            <el-input type="textarea" :rows="8" style="width: 400px" v-model="activeInfo.desc"></el-input>
          </el-form-item>
        </el-col>
           <el-col :span="10">
          <el-form-item label="活动封面"  >
           <el-upload  action="" :on-change="handleelchange"  :auto-upload="false" class="upload-demo" multiple drag>
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
            <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
             <i class="el-icon-plus"></i>
            </el-upload>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <div class="btn">
      <el-button type="primary" @click="submitForm">完成</el-button>
      <el-button type="danger" @click="back">返回</el-button>
    </div>
  </el-card>
</template>

<script>
import day from 'dayjs'
export default {
  name: "AddActive",
  props: ["id"],
  data() {
    return {
      activeInfo: {
        name:"",
        cost: undefined,
        num: undefined,
        startTime: "",
        endTime: "",
        desc: "",
        type:undefined,
        pic:"",
      },
      type:["大转盘","扭蛋机","老虎机","掷筛子"],
      prizesChecked:[],
      prizeList:[],
      activeInfoRules: {
        startTime: [
          { required: true, message: "请输入文章标题", trigger: "change" }
        ],
        endTime: [
          { required: true, message: "请选择文章分类", trigger: "change" }
        ],
        desc: [
          { required: true, message: "请输入文章描述", trigger: "change" },
          { max: 120, message: "描述最多可写120个字符", trigger: "change" }
        ]
      }
    };
  },

  methods: {
    async getPrizeList(){
        const {data:res}=await this.$http.get(`prize/list`)
        console.log(res);
        this.prizeList=res.data
    },
    async getInfo(id){
      const {data:res}=await this.$http.get(`activity/info`,{
        params:{
          id
        }
      })
      console.log(res);
      if(res.code!=200){
        this.$message.error(res.message)
        return
      }
      this.activeInfo=res.activity;
    },
    DateForm(date){
        return date?day(date).format('YYYY-MM-DD HH:mm:ss'):'暂无'
      },
    async submitForm() {
      this.$refs.activeForm.validate(async valid => {
        if (valid) {
          const form_data=new FormData();
          form_data.append("name",this.activeInfo.name)
          form_data.append("cost",this.activeInfo.cost)
          form_data.append("type",this.activeInfo.type)
          form_data.append("startTime",this.DateForm(this.activeInfo.startTime))
          form_data.append("endTime",this.DateForm(this.activeInfo.endTime))
          form_data.append("desc",this.activeInfo.desc)
          form_data.append("pic",this.activeInfo.pic)
          if(this.id){
          form_data.append("id",this.id)
          const { data: res } = await this.$http.put(`activity/edit`,form_data)
          if (res.code !== 200) return this.$message.error(res.message)
          this.$message({
            type:'success',
            message:'更新活动成功！'
          }
          )       
          }else{
            console.log(this.activeInfo.pic);
          const { data: res } = await this.$http.post(
            "activity/add",
            form_data
          );
          console.log(res);
          if (res.code !== 200) return this.$message.error(res.message);
          this.id=res.id
            this.$message({
            type: "success",
            message: "添加活动成功！"
          });
          }
          this.addPrize()
          this.$router.push("/activeList");
        } else {
          return this.$message.error("输入数据非法，请重新输入");
        }
      });
    },
    async addPrize(){
    const form_data=new FormData();
    form_data.append("id",this.id)
    form_data.append("name",JSON.stringify(this.prizesChecked))
    const{data:res}=await this.$http.post(`activity/addPrizes`,form_data)
    console.log(res);
    if (res.code!=200) {
      return this.$message.error(res.message)
    }
    },
    back() {
      this.$router.push("/activeList");
    },
    handleelchange(file,fileList){
      this.activeInfo.pic=file.raw
    }
  },
  created() {
    this.getPrizeList()
      if(this.id){
      this.getInfo(this.id)
    }
  }
};
</script>

<style scoped>
</style>
