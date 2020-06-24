<!-- Vue Chart Page -->
<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container class="grid-list-xl pt-0 mt-n3">
			<v-row>
				<app-card
					:heading="$t('message.barChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<bar-chart :width="370" :height="246"></bar-chart>
				</app-card>
				<app-card
					:heading="$t('message.lineChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<line-chart :width="370" :height="246"></line-chart>
				</app-card>
				<app-card
					:heading="$t('message.doughnutChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<doughnut-chart :width="300" :height="110"></doughnut-chart>
				</app-card>
				<app-card
					:heading="$t('message.gradientLineChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<gradient-line-chart :width="370" :height="246"></gradient-line-chart>
				</app-card>
				<app-card
					:heading="$t('message.polarAreaChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<polar-area-chart :width="370" :height="246"></polar-area-chart>
				</app-card>
				<app-card
					:heading="$t('message.radarChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<radar-chart :width="370" :height="246"></radar-chart>
				</app-card>
				<app-card
					:heading="$t('message.bubbleChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<bubble-chart :width="370" :height="246"></bubble-chart>
				</app-card>
				<app-card
					:heading="$t('message.pieChart')"
					colClasses="col-12 col-sm-6 col-md-6"
					>
					<pie-chart :width="370" :height="110"></pie-chart>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import BarChart from "./vue-chartjs/BarChart";
import BubbleChart from "./vue-chartjs/BubbleChart";
import DoughnutChart from "./vue-chartjs/DoughnutChart";
import LineChart from "./vue-chartjs/LineChart";
import PieChart from "./vue-chartjs/PieChart";
import PolarAreaChart from "./vue-chartjs/PolarAreaChart";
import RadarChart from "./vue-chartjs/RadarChart";
import GradientLineChart from "./vue-chartjs/GradientLineChart";

export default {
  components: {
    BarChart,
    BubbleChart,
    DoughnutChart,
    LineChart,
    PieChart,
    PolarAreaChart,
    RadarChart,
    GradientLineChart
  }
};
</script>
