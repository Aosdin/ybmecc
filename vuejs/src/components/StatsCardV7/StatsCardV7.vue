<template>
   <v-col :class="colClasses">
      <div class="app-card vue-card-space" :class="customClasses">
         <v-card-text class="py-0 mb-4">
            <div class="layout justify-space-between">
               <div class="align-items-center d-custom-flex">
                  <div class="crypto-icon-wrap inline-block mr-4 font-2x">
                     <i :class="icon" title="BCN"></i>
                  </div>
                  <div class="inline-block">
                     <h3 class="mb-0">{{heading}}</h3>
                  </div>   
               </div>
            </div>
         </v-card-text>
         <slot></slot>
         <div class="chart-info d-custom-flex align-items-center justify-space-between"> 
            <v-badge :value=false class="primary pa-2">last 4 days</v-badge>
            <span class="primary--text font-md"><i class="zmdi zmdi-long-arrow-up primary--text mr-2"></i>2.3%</span>
         </div>   
      </div>
   </v-col>
</template>
<script>
export default {
  props: ["heading", "amount", "colClasses", "customClasses", "icon"]
};
</script>
