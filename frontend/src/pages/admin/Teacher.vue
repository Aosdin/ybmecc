<template>
  <v-row align="center" class="pa-2">
    <!--선생님목록-->
    <v-row class="col-12">
      <h1 class="col-12">선생님 리스트</h1>
      <v-data-table
              :headers="headers"
              :items="teacherList"
              :items-per-page="10"
              class="elevation-1"
      ></v-data-table>
    </v-row>
    <!--//선생님목록-->
    <v-row class="col-12">
      <v-form
              ref="form"
              v-model="valid"
      >
        <h1>선생님 추가</h1>

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
        <v-text-field
                v-model="iptKoName"
                :counter="10"
                label="Korea name"
                required
        ></v-text-field>
        <v-text-field
                v-model="iptEnName"
                :counter="10"
                label="English name"
                required
        ></v-text-field>
        <v-select
                v-model="SeniorSelect"
                :items="SeniorItems"
                item-text="text"
                item-value="value"
                label="Senior"
                required
        ></v-select>
        <v-select
                v-model="LevelSelect"
                :items="LevelItems"
                item-text="text"
                item-value="value"
                label="Level"
                required
        ></v-select>
        <v-btn
                :disabled="!valid"
                color="success"
                class="mr-4"
                @click="addTeacher"
        >
          등록
        </v-btn>
      </v-form>
    </v-row>
  </v-row>
</template>

<script>
import { getTeacherAPI, addTeacherAPI, getSeniorTeacherAPI } from '@/store/modules/common/api'
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
  computed: {
    SeniorItems () {
      let a = [{value: 0, text: '선택하세요'}]
      this.seniorTeacherList.map(s => {
        a.push({value: s.idx, text: s.enNm},)
      })
      return a
    }
  },
  data () {
    return {
      teacherList: [],
      seniorTeacherList: [],
      headers: [
        { text: 'ID', value: 'tcId' },
        { text: 'Ko name', value: 'koNm'},
        { text: 'En name', value: 'enNm' },
        { text: 'Level', value: 'tcLv' },
        { text: 'Senior', value: 'seniorIdx' }
      ],
      valid: false,
      iptId: '',
      iptPwd: '',
      iptKoName: '',
      iptEnName: '',
      LevelSelect: {value: 0, text: '선택하세요'},
      LevelItems: [{value: 0, text: '선택하세요'}, {value: 1, text: '1'}, {value: 2, text: '2'}, {value: 3, text: '3'}],
      SeniorSelect: {value: 0, text: '선택하세요'},
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
    },
    addTeacher () {
      addTeacherAPI({
        //POST전송데이터
        tcId: this.iptId,
        tcPwd: this.iptPwd,
        koNm: this.iptKoName,
        enNm: this.iptEnName,
        seniorIdx: this.SeniorSelect,
        tcLv: this.LevelSelect,
      }).then(({ data }) => {
        console.log('선생님등록 여부:', data)
        if (data.success) {
          alert('선생님등록에 성공하였습니다.')
          this.getTeacher();
        //   this.teacherList = data.data
        } else {
          alert(data.msg)
        }
      }).catch((err) => {
        console.error(err)
      })
    }
  },
  mounted() {
    this.getTeacher();
    this.getSeniorTeacher();
  }
}
</script>
