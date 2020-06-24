<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container class="grid-list-xl fluid pt-0 mt-n3">
			<app-card
				:heading="$t('message.grid')"
				contentCustomClass="text-center"
			>
				<v-row>
					<v-col sm="12" cols="12">
						<v-card color="primary" class="theme--dark">
						<v-card-text>12</v-card-text>
						</v-card>
					</v-col>
					<v-col cols="6" sm="6" v-for="{n, index} in 2" :key="index">
						<v-card color="primary" class="theme--dark">
						<v-card-text>6</v-card-text>
						</v-card>
					</v-col>
					<v-col sm="4" cols="4" v-for="{n, index} in 3" :key="index">
						<v-card color="primary" class="theme--dark">
						<v-card-text>4</v-card-text>
						</v-card>
					</v-col>
					<v-col sm="3" cols="3" v-for="{n, index} in 4" :key="index">
						<v-card color="primary" class="theme--dark">
						<v-card-text>3</v-card-text>
						</v-card>
					</v-col>
					<v-col sm="2" cols="2" v-for="{n, index} in 6" :key="index">
						<v-card color="primary" class="theme--dark">
						<v-card-text>2</v-card-text>
						</v-card>
					</v-col>
					<v-col sm="1" cols="1" v-for="{n,index} in 12" :key="index">
						<v-card color="primary" class="theme--dark">
						<v-card-text>1</v-card-text>
						</v-card>
					</v-col>
				</v-row>
			</app-card>
		</v-container>
	</div>
</template>
