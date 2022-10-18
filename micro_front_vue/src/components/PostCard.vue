<template>
  <router-link :to="`/choose/${post.id}?cost=${post.cost}`">
    <vs-card class="post-card">
      <template #title>
        <h3>{{ post.name }}</h3>
      </template>
      <template #img>
        <img :src="post.pic" alt />
      </template>
      <template #text>
        <p class="post-txt">所需积分：{{ post.cost }}</p>
        <p class="post-txt">开始时间：{{ post.startTime }}</p>
        <p class="post-txt">结束时间：{{ post.endTime }}</p>
        <p class="post-txt">活动描述：{{ post.desc }}</p>
        <br />
        <small class="post-txt right">活动发起者：{{ post.user }}</small>
      </template>
      <template #interactions>
      </template>
    </vs-card>
  </router-link>
</template>

<script>

export default {
  name: 'PostCard',
  props: [
    'post',
  ],
  filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			}
		}
}
</script>

<style>
.post-txt {
  overflow: hidden;
  text-overflow: ellipsis;
  display: auto;
}

.post-card .vs-card {
  height: 20%!important;
  width: 100% !important;
  max-width: 100% !important;
  max-height: 40% !important;
  overflow: hidden;
}

.post-card-tag {
  position: absolute;
  bottom: 20px;
}
.right{
  float: right;
}
</style>
