import actions from './actions'
import getters from './getters'
import mutations from './mutations'

const state = {
  isAPILoading: false,
  drawer: false,
  teacherList: [],
  ui: {
    drawer: false
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
