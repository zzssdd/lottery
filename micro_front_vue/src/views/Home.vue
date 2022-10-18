<template>
  <div class="main_page">
    <div class="container">
      <hr/>
      <div class="grid-3_xs-1_sm-2_md-2" v-if="posts.length">
        <div
          :key="index"
          v-for="(post, index) in posts.slice(Math.max(posts.length - 6, 0))"
          class="col"
        >
          <PostCard :post="post"/>
        </div>
      </div>
      <Nothing v-else />
        <div class="center">
        <vs-pagination v-model="page" :length="total" />
       </div>
    </div>
  </div>
</template>

<script>
import PostCard from '@/components/PostCard.vue'
import Nothing from '@/components/Nothing.vue'

export default {
  name: 'Home',
  components: {
    PostCard,
    Nothing
  },
  methods:{
    async getPost(){
      const {data:res}=await this.$http.get(`activity/list`,{
        params:{
          pageNum:this.page
        }
      })
      console.log(res);
      this.posts=res.data
      this.total=(parseInt(res.count)+5)/6
      console.log(this.total);
    },
  },
  data: function () {
    return {
      posts:[],
      page:1,
      total:0,
    }
  },
  created(){
    this.getPost()
  },
  mounted: function () {
    this.changeTitle('主页')
  },
  watch:{
    page(){
       this.getPost()
    }
  }
}
</script>
