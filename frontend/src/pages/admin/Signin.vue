<template>
  <v-row align="center" class="pa-2">
    <v-row class="col-12">
      <v-form
              ref="form"
              v-model="valid"
      >
        <h1>로그인</h1>

        <v-text-field
                v-model.trim="iptId"
                label="ID"
                required
        ></v-text-field>

        <v-text-field
                v-model.trim="iptPwd"
                label="Password"
                required
        ></v-text-field>
        <v-btn
                :disabled="!valid"
                color="success"
                class="mr-4"
                @click="addTeacher"
        >
          로그인
        </v-btn>
      </v-form>
    </v-row>
  </v-row>
</template>

<script>
import { getTeacherAPI, addTeacherAPI, getSeniorTeacherAPI } from '@/store/modules/common/api'
export default {
  name: 'Signin',
  head: {
    title: {
      inner: 'Signin'
    },
    // Meta tags
    meta: [
      { name: 'keywords', content: 'Signin' }
    ]
  },
  computed: {
  },
  data () {
    return {
      valid: false,
      iptId: '',
      iptPwd: '',
    }
  },
  methods: {
    getSeniorTeacher () {
      getSeniorTeacherAPI().then(({ data }) => {
        console.log('시니어선생님목록:', data)
        if (data.success) {
          this.seniorTeacherList = data.data
        } else {
          alert('시니어선생님목록 가져오기에 실패했습니다.')
        }
      }).catch((err) => {
        console.error(err)
      })
    },
    getTeacher () {
      getTeacherAPI().then(({ data }) => {
        console.log('선생님목록:', data)
        if (data.success) {
          this.teacherList = data.data
        } else {
          alert('선생님목록 가져오기에 실패했습니다.')
        }
      }).catch((err) => {
        console.error(err)
      })
    }
  },
  mounted() {
  }
}
</script>
