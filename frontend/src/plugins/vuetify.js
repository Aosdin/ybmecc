import Vue from 'vue'
import Vuetify from 'vuetify'
Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#222',
        secondary: '#c30e2e',
        accent: '#f03c3c',
        error: '#c30e2e',
        info: '#b656c8',
        success: '#76ba50',
        warning: '#e6e6e6',
        anchor: 'inherit'
      }
    }
  }
})
