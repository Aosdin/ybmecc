<template>
  <v-app-bar
          :clipped-left="$vuetify.breakpoint.lgAndUp"
          app
          color="blue darken-3"
          dark
  >
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title
            style="width: 300px"
            class="ml-0 pl-4"
    >
      <span class="hidden-sm-and-down"></span>
    </v-toolbar-title>
  </v-app-bar>
</template>

<script>
export default {
  name: 'Header'
}
</script>
