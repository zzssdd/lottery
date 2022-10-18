import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/ui'
import './plugin/http'


Vue.config.productionTip = false


new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
