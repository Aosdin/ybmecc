<template>
	<div>
		<div class="top-selling-widget">
			<slick :options="slickOptions">
				<div v-for="(item,index) in exchangeStatistic" :key="index">
					<v-row>
						<v-col xl="7" lg="7" md="8" sm="8" cols="12" class="mb-4">
							<div class="crypto-icon-wrap mb-sm-0">
								<i class="warn--text font-3x" :class="item.icon1"></i>
								<i class="primary--text font-3x" :class="item.icon2"></i>
							</div>
						</v-col>
						<v-col xl="5" lg="5" md="4" sm="4" cols="12" class="mb-4">
							<div>
								<div class="crypto-text-wrap">
									<span class="d-inline-block">{{ item.icon1_name }}</span>
									<span class="d-inline-block ml-2 mr-2">/</span>
									<span class="d-inline-block">{{ item.icon2_name }}</span>
								</div>
								<div class="crypto-result-price">
									<div class="d-custom-flex justify-space-between"><span class="fw-bold">Price:</span>
										<span>{{ item.price }}</span></div>
									<div class="d-custom-flex justify-space-between"><span
											class="fw-bold">Volume:</span><span>{{ item.volume }}</span></div>
									<span class="d-block success--text">{{ item.percentage }}%</span>
								</div>
							</div>
						</v-col>
					</v-row>
					<line-chart-v5 label="Statistics" :chartdata="item.data" :labels="item.chartLabel"
						:pointBackgroundColor="ChartConfig.color.primary" :height="80"
						:pointBorderColor="ChartConfig.color.white" :borderWidth="4" :showGridLines=false
						:enableGradient=false :enableShadow="true" shadowColor="#1565c0"
						:borderColor="ChartConfig.color.primary">
					</line-chart-v5>
				</div>
			</slick>
		</div>
	</div>
</template>
<script>
	import Slick from "vue-slick";
	import { exchangeStatistic } from 'Views/crypto/data.js'
	import LineChartV5 from "Components/Charts/LineChartV5";
	import { ChartConfig } from "Constants/chart-config";


	export default {
		components: {
			Slick,
			LineChartV5
		},
		data() {
			return {
				ChartConfig,
				exchangeStatistic,
				slickOptions: {
					infinite: true,
					speed: 500,
					slidesToShow: 1,
					slidesToScroll: 1,
					autoplay: true,
					dots: false,
					fade: true,
					cssEase: 'linear',
					arrows: false,
					rtl: this.$store.getters.rtlLayouts
				},
			};
		}
	};
</script>