import * as t from './types'
import { getWeatherAPI, getTeacherAPI } from './api'

export default {
  getWeather ({ commit }, key) {
    commit(t.GET_WEATHER_PENDING)
    return Promise.resolve()
      .then(() => {
        return getWeatherAPI(key)
      })
      .then((response) => {
        const result = response.data || {}
        // const { data } = result || {}
        // const status = response.status || []
        commit(t.GET_WEATHER_SUCCESS, { result })
        return response
      })
      .catch((err) => {
        commit(t.GET_WEATHER_FAIL)
        throw err
      })
  },
  getTeacher ({ commit }) {
    commit(t.GET_TEACHER_PENDING)
    return Promise.resolve()
      .then(() => {
        return getTeacherAPI()
      })
      .then((response) => {
        const result = response.data || {}
        // const { data } = result || {}
        // const status = response.status || []
        commit(t.GET_TEACHER_SUCCESS, { result })
        return response
      })
      .catch((err) => {
        commit(t.GET_TEACHER_FAIL)
        throw err
      })
  },
  setDrawer ({ commit }, value) {
    return commit(t.SET_DRAWER, value)
  },
  setUI ({ commit }, key, value) {
    return commit(t.SET_UI, key, value)
  }
}
