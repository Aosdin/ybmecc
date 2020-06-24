<!-- Vue EChart Page -->
<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container class="grid-list-xl pt-0 mt-n3">
			<v-row>
				<app-card
					:heading="$t('message.barChart')"
					contentCustomClass="overflow-hidden"
					colClasses="col-12 col-sm-6 col-md-6"
					:fullBlock="true"
					>
					<bar-chart></bar-chart>
				</app-card>
				<app-card
					:heading="$t('message.doughnutChart')"
					contentCustomClass="overflow-hidden"
					colClasses="col-12 col-sm-6 col-md-6"
					:fullBlock="true"
					>
					<doughnut-chart></doughnut-chart>
				</app-card>
				<app-card
					:heading="$t('message.lineChart')"
					contentCustomClass="overflow-hidden"
					colClasses="col-12 col-sm-6 col-md-6"
					:fullBlock="true"
					>
					<line-chart></line-chart>
				</app-card>
				<app-card
					:heading="$t('message.pieChart')"
					contentCustomClass="overflow-hidden"
					colClasses="col-12 col-sm-6 col-md-6"
					:fullBlock="true"
					>
					<pie-chart></pie-chart>
				</app-card>
			</v-row>
			<v-row>
				<app-card
					:heading="$t('message.funnelChart')"
					contentCustomClass="overflow-hidden"
					colClasses="col-12"
					:fullBlock="true"
					>
					<funnel-chart></funnel-chart>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import BarChart from "./vue-echarts/BarChart";
import DoughnutChart from "./vue-echarts/DoughnutChart";
import LineChart from "./vue-echarts/LineChart";
import PieChart from "./vue-echarts/PieChart";
import FunnelChart from "./vue-echarts/FunnelChart";

export default {
  components: {
    BarChart,
    PieChart,
    DoughnutChart,
    LineChart,
    FunnelChart
  }
};
</script>
