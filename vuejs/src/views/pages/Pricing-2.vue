<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0 mt-n3">
			<div class="pricing-wrapper">
				<v-row>
					<v-col cols="12" sm="6" md="6" lg="3">
						<div class="app-card text-center">
							<div class="app-card-title primary d-custom-flex justify-space-between">
								<h3 class="mb-0 white--text mr-2">{{$t("message.basic")}}</h3>
								<p class="mb-0 fs-12 white--text">{{$t("message.forMostOfTheUsers")}}</p>
							</div>
							<div class="app-full-content">
								<h2 class="mb-5 font-3x">
									<span class="font-xl">$40</span>
									<sub>/Mo</sub>
								</h2>
								<ul class="list-unstyled list-group-flush">
									<li class="list-group-item">10GB of Bandwidth</li>
									<li class="list-group-item">200MB Max File Size</li>
									<li class="list-group-item">2GHZ CPU</li>
									<li class="list-group-item">256MB Memory</li>
									<li class="list-group-item">1 GB Storage</li>
								</ul>
							</div>
							<div class="app-footer">
								<v-btn block color="primary" large>{{$t("message.choosePlan")}}</v-btn>
							</div>
						</div>
					</v-col>
					<v-col cols="12" sm="6" md="6" lg="3">
						<div class="app-card text-center">
							<div class="app-card-title success white--text d-custom-flex justify-space-between">
								<h3 class="mb-0 white--text mr-2">{{$t("message.standard")}}</h3>
								<p class="mb-0 fs-12 fw-normal white--text">{{$t("message.forMostOfTheUsers")}}</p>
							</div>
							<div class="app-full-content">
								<h2 class="mb-5 font-3x">
									<span class="font-xl">$50</span>
									<sub>/Mo</sub>
								</h2>
								<ul class="list-unstyled list-group-flush">
									<li class="list-group-item">10GB of Bandwidth</li>
									<li class="list-group-item">200MB Max File Size</li>
									<li class="list-group-item">2GHZ CPU</li>
									<li class="list-group-item">256MB Memory</li>
									<li class="list-group-item">1 GB Storage</li>
								</ul>
							</div>
							<div class="app-footer">
								<v-btn block color="success" large>{{$t("message.choosePlan")}}</v-btn>
							</div>
						</div>
					</v-col>
					<v-col cols="12" sm="6" md="6" lg="3">
						<div class="app-card text-center">
							<div class="app-card-title warning white--text d-custom-flex justify-space-between">
								<h3 class="mb-0 white--text mr-2">{{$t("message.mega")}}</h3>
								<p class="mb-0 fs-12 fw-normal white--text">{{$t("message.forDeveloper")}}</p>
							</div>
							<div class="app-full-content">
								<h2 class="mb-5 font-3x">
									<span class="font-xl">$70</span>
									<sub>/Mo</sub>
								</h2>
								<ul class="list-unstyled list-group-flush">
									<li class="list-group-item">10GB of Bandwidth</li>
									<li class="list-group-item">200MB Max File Size</li>
									<li class="list-group-item">2GHZ CPU</li>
									<li class="list-group-item">256MB Memory</li>
									<li class="list-group-item">1 GB Storage</li>
								</ul>
							</div>
							<div class="app-footer">
								<v-btn block large color="warning">{{$t("message.choosePlan")}}</v-btn>
							</div>
						</div>
					</v-col>
					<v-col cols="12" sm="6" md="6" lg="3">
						<div class="app-card text-center">
							<div class="app-card-title error white--text d-custom-flex justify-space-between">
								<h3 class="mb-0 white--text mr-2">{{$t("message.master")}}</h3>
								<p class="mb-0 fs-12 fw-normal white--text">{{$t("message.forLargeEnterprises")}}</p>
							</div>
							<div class="app-full-content">
								<h2 class="mb-5 font-3x">
									<span class="font-xl">$100</span>
									<sub>/Mo</sub>
								</h2>
								<ul class="list-unstyled list-group-flush">
									<li class="list-group-item">10GB of Bandwidth</li>
									<li class="list-group-item">200MB Max File Size</li>
									<li class="list-group-item">2GHZ CPU</li>
									<li class="list-group-item">256MB Memory</li>
									<li class="list-group-item">1 GB Storage</li>
								</ul>
							</div>
							<div class="app-footer">
								<v-btn block color="error" large>{{$t("message.choosePlan")}}</v-btn>
							</div>
						</div>
					</v-col>
				</v-row>
			</div>
		</v-container>
	</div>
</template>