import axios from 'axios'

const axiosInstance = function () {
  return axios.create({
    baseURL: `http://localhost:1323`,
    timeout: 50000,
    withCredentials: true,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTk1OTI0MzE0LCJuYW1lIjoic2NoZW1hayJ9.9exy8nhC5mIA-eL8gAlGn4mAv9cIQXrE0_ItsQCrkv4'
    }
  })
}
export function getWeatherAPI (key) {
  return new Promise((resolve, reject) => {
    return axiosInstance().get(`/data/2.5/weather?q=Seoul&appid=${key}`)
      .then(response => resolve(response))
      .catch(err => reject(err))
  })
}
export function getSeniorTeacherAPI () {
  return new Promise((resolve, reject) => {
    return axiosInstance().get(`/api/tcSeniorUser`)
      .then(response => resolve(response))
      .catch(err => reject(err))
  })
}
export function getTeacherAPI () {
  return new Promise((resolve, reject) => {
    return axiosInstance().get(`/api/tcUser`)
      .then(response => resolve(response))
      .catch(err => reject(err))
  })
}
export function addTeacherAPI (payload) {
  return new Promise((resolve, reject) => {
    return axiosInstance().post(`/api/tcUser`, payload)
      .then(response => resolve(response))
      .catch(err => reject(err))
  })
}
export function getStudentAPI () {
  return new Promise((resolve, reject) => {
    return axiosInstance().get(`/api/sdUser`)
        .then(response => resolve(response))
        .catch(err => reject(err))
  })
}
