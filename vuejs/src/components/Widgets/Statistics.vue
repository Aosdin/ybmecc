<template>
   <div class="statistics-chart-wrap">
      <v-row>
         <line-chart-v6
            :style="{height: '400px',width:'100%', position: 'relative'}"
            :dataSet= "data"
            :labels= "labels"
            :labelX="labelX"
            :labelY="labelY"
            :borderColor="ChartConfig.color.primary"
            :label="label"
            >
         </line-chart-v6>
      </v-row>
   </div>
</template>

<script>
import LineChartV6 from "Components/Charts/LineChartV6";
import { ChartConfig } from 'Constants/chart-config'

export default {
  props: ['data', 'labels', 'label', 'labelX', 'labelY'],
   components:{
      LineChartV6
   },
  data() {
    return {
       ChartConfig
    };
  }
};
</script>
