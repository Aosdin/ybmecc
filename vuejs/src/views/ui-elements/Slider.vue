<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0 mt-n3">
			<v-row>
				<app-card
					:heading="$t('message.continuous')"
					colClasses="col-12 col-md-6"
				>
					<v-slider hide-details v-model="value1" step="0"></v-slider>
					<v-slider hide-details v-model="value2" step="0" disabled></v-slider>
				</app-card>
				<app-card
					:heading="$t('message.discrete')"
					colClasses="col-12 col-md-6"
				>
					<v-slider hide-details v-model="value3" thumb-label step="10" ticks></v-slider>
				</app-card>
			</v-row>
			<v-row>
				<app-card 
					:heading="$t('message.customColors')" 
					colClasses="col-12 col-md-6"
				>
					<v-slider :label="ex1.label" v-model="ex1.val" :color="ex1.color"></v-slider>
					<v-slider :label="ex2.label" v-model="ex2.val" :track-color="ex2.color"></v-slider>
					<v-slider :label="ex3.label" v-model="ex3.val" :thumb-color="ex3.color" thumb-label></v-slider>
				</app-card>
				<app-card
					:heading="$t('message.editableNumericValue')"
					colClasses="col-12 col-md-6"
				>
					<v-responsive height="100px" v-bind:style="{ background: `rgb(${red}, ${green}, ${blue})` }"></v-responsive>
					<v-row>
						<v-col cols="9">
							<v-slider hide-details label="R" v-bind:max="255" v-model="red"></v-slider>
						</v-col>
						<v-col cols="3">
							<v-text-field hide-details v-model="red" type="number"></v-text-field>
						</v-col>
						<v-col cols="9">
							<v-slider hide-details label="G" v-bind:max="255" v-model="green"></v-slider>
						</v-col>
						<v-col cols="3">
							<v-text-field hide-details v-model="green" type="number"></v-text-field>
						</v-col>
						<v-col cols="9">
							<v-slider hide-details label="B" v-bind:max="255" v-model="blue"></v-slider>
						</v-col>
						<v-col cols="3">
							<v-text-field hide-details v-model="blue" type="number"></v-text-field>
						</v-col>
					</v-row>
				</app-card>
			</v-row>
			<v-row>
				<app-card 
					:heading="$t('message.icons')" 
					colClasses="col-12 col-md-6"
				>
					<v-row>
						<v-col cols="4">
							<v-subheader>Media volume</v-subheader>
						</v-col>
						<v-col cols="8">
							<v-slider hide-details prepend-icon="volume_up" v-model="media"></v-slider>
						</v-col>
					</v-row>
					<v-row>
						<v-col cols="4">
							<v-subheader>Alarm volume</v-subheader>
						</v-col>
						<v-col cols="8">
							<v-slider hide-details append-icon="alarm" v-model="alarm"></v-slider>
						</v-col>
					</v-row>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
export default {
  data() {
    return {
      value1: 0,
      value2: 0,
      value3: 0,
      media: 0,
      alarm: 0,
      red: 64,
      green: 128,
      blue: 0,
      ex1: { label: "color", val: 25, color: "orange darken-3" },
      ex2: { label: "track-color", val: 75, color: "green lighten-1" },
      ex3: { label: "thumb-color", val: 50, color: "red" }
    };
  }
};
</script>
