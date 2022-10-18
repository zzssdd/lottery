<template>
  <div>
    <vs-navbar shadow-scroll fixed center-collapsed not-line>
      <template #left>
        <vs-button flat icon style="margin-right: 20px" class="menu-btn" @click="activeSidebar = !activeSidebar">
          <i class='bx bx-menu'></i>
        </vs-button>
          <h5>Yogen的抽奖系统</h5>
      </template>
      <template #right>
         <vs-button
        :active="active"
        @click="active=true"
        v-if="!is_login"
      >
        登录
      </vs-button>
        <vs-navbar-item>
          <vs-switch
            @click="changeTheme()"
          >
            <template #on>
                <i class='bx bxs-sun'></i>
            </template>
            <template #off>
              <i class='bx bxs-moon'></i>
            </template>
          </vs-switch>
        </vs-navbar-item>
      </template>
     </vs-navbar>
      <vs-sidebar
        v-model="activeItem"
        :open.sync="activeSidebar"
      >
        <template #logo>
          Yogen的抽奖系统
        </template>

        <vs-sidebar-item
          v-for="(link, index) in links"
          :key="index"
          :to="link.url"
          :id="link.name"
        >
          {{ link.name }}
        </vs-sidebar-item>

        <vs-sidebar-item :active="active" @click.native="Logout">
        退出
      </vs-sidebar-item>
      </vs-sidebar>
    <div style="margin-bottom: 100px"></div>

      <vs-dialog v-model="active">
        <template #header>
          <h4 class="not-margin">
            请登录
          </h4>
        </template>


        <div class="con-content">
          <vs-input v-model="email" placeholder="邮箱">
            <template #icon>
              @
            </template>
          </vs-input>
          <vs-input type="password" v-model="password" placeholder="密码">
            <template #icon>
              <i class='bx bxs-lock'></i>
            </template>
          </vs-input>
        </div>

        <template #footer>
          <div class="footer-dialog">
            <vs-button block @click="Login">
              登录
            </vs-button>
             <vs-button block @click="active2=true">
              注册
            </vs-button>
          </div>
        </template>
      </vs-dialog>
      <vs-dialog v-model="active2">
        <template #header>
          <h4 class="not-margin">
           请注册
          </h4>
        </template>


        <div class="con-form">
          <vs-input v-model="register.email" placeholder="邮箱">
            <template #icon>
              @
            </template>
          </vs-input>
          <vs-input type="password" v-model="register.password" placeholder="密码">
            <template #icon>
              <i class='bx bxs-lock'></i>
            </template>
          </vs-input>
           <vs-input type="password" v-model="register.repassword" placeholder="确认密码">
            <template #icon>
              <i class='bx bxs-lock'></i>
            </template>
          </vs-input>
          <div class="code">
            <vs-input id="codeinput" v-model="register.code" placeholder="验证码">
          </vs-input>
            <vs-button id="codebutton" @click="sendEmail">
                获取验证码
            </vs-button>
          </div>
        </div>

        <template #footer>
          <div class="footer-dialog">
            <vs-button block @click="Postregister">
            注册
            </vs-button>
          </div>
        </template>
      </vs-dialog>
    </div>

</template>

<script>

export default {
  name: 'Navbar',
  data: () => {
    return {
      activeItem: 'Home',
      activeSidebar: false,
      dark: false,
      active:false,
      active2:false,
      email:undefined,
      is_login:false,
      email:undefined,
      password:undefined,
      register:{
        email:undefined,
        password:undefined,
        repassword:undefined,
        code:undefined,
      },
      links:[
        {
          "name":"活动主页",
          "url":"/"
        },
        {
          "name":"我的奖品",
          "url":"/payment"
        },
      ]
    }
  },
  methods: {
    changeTheme: function () {
      if (this.dark) {
        document.getElementsByTagName('body')[0].setAttribute('vs-theme', 'light')
        this.dark = false
      } else {
        document.getElementsByTagName('body')[0].setAttribute('vs-theme', 'dark')
        this.dark = true
      }
    },
    async Login(){
      const form_data=new FormData()
      form_data.append("email",this.email)
      form_data.append("password",this.password)
      const { data: res } = await this.$http.post('user/login',form_data)
      if (res.code!=200) {
      return alert(res.msg)
      }
      window.sessionStorage.setItem('email',res.user)
      window.sessionStorage.setItem('token',res.token)
      this.is_login=true
      this.active=false
    },
    Logout(){
      window.sessionStorage.removeItem('email')
      window.sessionStorage.removeItem('token')
      this.is_login=false
    },
    getUser(){
    this.email=window.sessionStorage.getItem('email')
    if (!this.email) {
      this.email=undefined
      this.is_login=false
    }else{
      this.is_login=true
    }
  },
    async sendEmail(){
      const form_data=new FormData()
      form_data.append("email",this.register.email)
      const { data: res } = await this.$http.post('user/email',form_data)
      return alert(res.msg)
    },
    async Postregister(){
      console.log(this.register);
      if (this.register.password.length<6) {
        return alert("密码长度不小于6")
      }
      if (this.register.password!=this.register.repassword) {
        return alert("两次密码输入不一致")
      }
      if (this.register.code.length!=6) {
        return alert("验证码格式错误")
      }
      const form_data=new FormData()
      form_data.append("email",this.register.email)
      form_data.append("password",this.register.password)
      form_data.append("code",this.register.code)
       const { data: res } = await this.$http.post('user/registry',form_data)
       if (res.code==200) {
        this.active2=false
       }
       return alert(res.msg)
    }
    },
  mounted(){
    this.getUser()
  },
}
</script>

<style>
.vs-input {
    width: 100% !important;
    margin-top: 20px;
}
.linkNav{
  margin: 0 auto;
}
#codeinput{
  width: 60%;
  float: left;
  margin-bottom:20px ;
}
#codebutton{
  float:right;
  margin-top:20px ;
  margin-bottom: 20px;
  width: 30%;
  height: 100%;
}
</style>