<template>
<div>
<div class="center" v-if="List.length>0">
      <vs-table
        v-model="selected"
        >
        <template #thead>
          <vs-tr>
            <vs-th>
              <vs-checkbox
                :indeterminate="selected.length == List.length"
                @change="selected = $vs.checkAll(selected, List)"
              />
            </vs-th>
            <vs-th>
              
            </vs-th>
            <vs-th style="width:45px">
              名称
            </vs-th>
            <vs-th>
              获奖时间
            </vs-th>
          </vs-tr>
        </template>
        <template #tbody>
          <vs-tr
            :key="i"
            v-for="(tr, i) in List"
            :data="tr"
            :is-selected="!!selected.includes(tr)"
          >
            <vs-td checkbox>
              <vs-checkbox :val="tr" v-model="selected" />
            </vs-td>
            <vs-td class="td-head">
               <img :src="tr.pic" alt="" class="pic">
            </vs-td>
            <vs-td>
            {{ tr.name }}
            </vs-td>
            <vs-td>
            {{ tr.CreatTime|formatDate }}
            </vs-td>
          </vs-tr>
        </template>
      </vs-table>
         <vs-button size="xl" class="btn-margin" @click="del()">
          删除订单
         </vs-button>
    </div>
      <Nothing v-else />
      <div class="center">
        <vs-pagination v-model="page" :length="total" />
       </div>
    </div>
</template>

<script>
import Nothing from '@/components/Nothing.vue'
export default {
    name:'Payment',
    components: {
    Nothing
    },
    data(){
        return{
            List:[],
            page:1,
            total:0,
            selected:[],
            totalPrice:0,
        }
    },
      components: {
    Nothing
  },
    methods:{
        async getList(){
            const { data: res } = await this.$http.get('order/list' ,{
                params:{
                    pageNum:this.page,
                }
            })
            console.log(res);
            if (res.code!=200) {
                return alert(res.msg)
            }
            this.List=res.data
            this.total=(parseInt(res.count)+5)/6
        },
        async del(){
          for (let i = 0; i < this.selected.length; i++) {
            const { data: res } = await this.$http.delete(`order/del`,{
                params:{
                    id:this.selected[i].id,
                }
            })
            if (res.code!=200) {
              return alert(res.msg)
            }
            this.getList()
          }
        }
    },
      filters:{
			formatDate:function (value){
				var  dt=new Date(value);//获取日期value值
				let year = dt.getFullYear();
				let month = dt.getMonth()+1;
				let date = dt.getDate();
				return `${year}年${month}月${date}日`;
			},
		},
    mounted(){
        this.getList()
    },
    watch:{
    page(){
       this.getList()
    },
  }
}
</script>

<style>
  .data{
    display: block;
  }
  .pic{
    width:60px !important;
    height:40px !important;
  }
  .td-head{
    width: 60px;
  }
  .btn-margin{
    margin-left:40px !important ;
  }
  element.style {
    width: 30px !important;
}
</style>