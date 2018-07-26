// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'

global.jQuery = require('jquery');
var $ = global.jQuery;
window.$ = $;

var Bootstrap = require('bootstrap3/dist/js/bootstrap')
Vue.use(Bootstrap)

var BootstrapVue = require('bootstrap-vue')
Vue.use(BootstrapVue)

import { library } from '@fortawesome/fontawesome-svg-core'
import { faEye } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
library.add(faEye)
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.config.productionTip = false



import App from './App.vue'
new Vue({ // eslint-disable-line no-new
  el: '#app',
  render: (h) => h(App)
})
