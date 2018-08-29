// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import { Button, Container, Row, Col, Form, Menu, MenuItem, Input, Header, Footer, Main } from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import 'element-ui/lib/theme-chalk/display.css'

Vue.config.productionTip = false

Vue.use(Button)
Vue.use(Container)
Vue.use(Col)
Vue.use(Row)
Vue.use(Form)
Vue.use(Menu)
Vue.use(MenuItem)
Vue.use(Input)
Vue.use(Header)
Vue.use(Footer)
Vue.use(Main)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
