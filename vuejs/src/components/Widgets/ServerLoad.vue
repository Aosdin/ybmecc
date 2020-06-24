<template>
	<div class="pos-relative h-100 overflow-hidden server-load-widget">
		<div class="d-custom-flex justify-space-between align-items-center mb-4">
			<div>
				<span class="fs-12 fw-normal grey--text">Space</span>
				<h5 class="mb-0"><span class="primary--text">98GB </span><span class="info--text">/ 124GB</span></h5>
			</div>
			<div>
				<span class="fs-12 fw-normal grey--text">Bandwidth</span>
        <h5 class="mb-0"><span class="error--text">512MB </span><span class="info--text">/ 1024MB</span></h5>
			</div>
		</div>
    <div>
      <ECharts :options="gauge" autoresize id="barChart" style="width:100%; height:260px"></ECharts>
    </div>
	</div>
</template>

<script>
import ECharts from "vue-echarts";
import "echarts/lib/chart/gauge";
import "echarts/lib/component/title";
import { ChartConfig } from "Constants/chart-config";

export default {
  props: ["styles"],
  name: "buyers-stats",
  components: {
    ECharts
  },
  data() {
    return {
      gauge: {
        tooltip : {
            formatter: "{a} <br/>{b} : {c}%"
        },
        toolbox: {
            feature: {
                restore: {},
                saveAsImage: {}
            }
        },
        series: [
            {
                name: 'Server Load',
                type: 'gauge',
                startAngle:225,
                endAngle:-45,
                detail: {
                  formatter:'{value}%',
                  fontSize:22 
                },
                axisLine: {
                  show: true,
                  lineStyle: {
                    color: [
                      [0.4, ChartConfig.color.primary], 
                      [0.8, ChartConfig.color.info], 
                      [1, ChartConfig.color.warning]
                    ],
                    width: 20
                  }
                },
                data: [{value: 50, name: 'Server Load'}],
                radius : 130
            }
        ]

      }
    };
  },
  mounted(){
    this.gaugeMove();
  },
  methods:{
    gaugeMove(){
      const self = this;
      setInterval(function () {
          self.gauge.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
      },2000);
    }
  }
}
</script>