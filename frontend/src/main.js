
/**
 * Global styles should be imported before any other Vue files
 */
import 'normalize.css'
import 'vuetify/dist/vuetify.css'
import '@/assets/scss/style.scss'
import '@mdi/font/css/materialdesignicons.css'

/**
 * Entry vue
 */
import Vue from 'vue'
import App from './App.vue'
import { sync } from 'vuex-router-sync'
import router from '@/routers' // Must be a name as `router`
import store from '@/store' // Must be a name as `store`

/**
 * Plugin
 */
import plugins from '@/shared/plugins'
import vueHead from 'vue-head'
import vuetify from './plugins/vuetify'

/**
 * Install plugins
 */
Vue.use(plugins)
Vue.use(vueHead)
Vue.config.productionTip = false
sync(store, router)

new Vue({
  router,
  store,
  vuetify,
  icons: {
    iconfont: 'mdi',
  },
  render: h => h(App)
}).$mount('#app')
