<template>
  <div style="margin-top: 100px;">
    <h1>{{ title }}</h1>
    <p v-for="(t, idx) in teacherList">{{ t.idx }} = {{ t.koNm }}</p>
  </div>
</template>

<script>
import { getTeacherAPI } from '@/store/modules/common/api'
export default {
  name: 'Teacher',
  head: {
    title: {
      inner: '선생님관리'
    },
    // Meta tags
    meta: [
      { name: 'keywords', content: '선생님관리' }
    ]
  },
  data () {
    return {
      title: '선생님관리',
      teacherList: []
    }
  },
  methods: {
    getTeacher () {
      getTeacherAPI().then(({ data }) => {
        console.log('선생님목록:', data)
        if (data) {
          this.teacherList = data
        }
      }).catch((err) => {
        console.error(err)
      })
    }
  },
  mounted() {
    this.getTeacher();
  }
}
</script>
