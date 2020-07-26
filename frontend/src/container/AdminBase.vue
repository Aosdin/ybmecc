<template>
  <v-app>
    <header-area />
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
    <footer-area />
  </v-app>
</template>

<script>
import HeaderArea from '../components/common/Header.vue'
import FooterArea from '../components/common/Footer.vue'

export default {
  name: 'FullBase',
  components: {
    HeaderArea,
    FooterArea
  }
}
</script>