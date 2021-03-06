import * as t from './types'

export default {
  [t.GET_WEATHER_PENDING]: (state) => {
    state.isAPILoading = true
  },
  [t.GET_WEATHER_SUCCESS]: (state, { result }) => {
    state.isAPILoading = false
    state.weather = result
  },
  [t.GET_WEATHER_FAIL]: (state) => {
    state.isAPILoading = false
  },
  [t.GET_TEACHER_PENDING]: (state) => {
    state.isAPILoading = true
  },
  [t.GET_TEACHER_SUCCESS]: (state, { result }) => {
    state.isAPILoading = false
    console.log(result)
    state.teacherList = result
  },
  [t.GET_TEACHER_FAIL]: (state) => {
    state.isAPILoading = false
  },
  [t.SET_DRAWER]: (state, { b }) => {
    state.drawer = b
  },
  [t.SET_UI]: (state, { key, value }) => {
    state[key] = value
  }
}
