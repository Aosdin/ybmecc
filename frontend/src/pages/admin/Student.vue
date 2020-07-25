<template>
  <div style="margin-top: 100px;">
    <h1>{{ title }}</h1>
    <p v-for="(s, idx) in studentList">{{ s.idx }} = {{ s.koNm }}</p>
  </div>
</template>

<script>
import { getStudentAPI } from '@/store/modules/common/api'
export default {
  name: 'Student',
  head: {
    title: {
      inner: '학생관리'
    },
    // Meta tags
    meta: [
      { name: 'keywords', content: '학생관리' }
    ]
  },
  data () {
    return {
      title: '학생관리',
      studentList: []
    }
  },
  methods: {
    getStudent () {
      getStudentAPI().then(({ data }) => {
        console.log('학생관리:', data.data)
        if (data) {
          this.studentList = data.data
        }
      }).catch((err) => {
        console.error(err)
      })
    }
  },
  mounted() {
    this.getStudent();
  }
}
</script>
