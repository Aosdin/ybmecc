<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container class="grid-list-xl pt-0 mt-n3">
			<v-row>
				<app-card :heading="$t('message.flatButton')" customClasses="mb-30" colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12">
					<div class="mb-5">
						<p>The <code>v-overlay</code> component is used to provide emphasis on a particular element or parts
							of it.</p>
					</div>
					<div class="text-center">
						<v-btn color="error" @click="overlay1 = !overlay1">
							Show Overlay
						</v-btn>
						<v-overlay :value="overlay1">
							<v-btn icon @click="overlay1 = false">
								<v-icon>mdi-close</v-icon>
							</v-btn>
						</v-overlay>
					</div>
				</app-card>

				<app-card :heading="$t('message.playground')" customClasses="mb-30" colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12">
					<v-row justify="center">
						<v-col cols="12">
							<v-btn class="mt-12" color="primary" @click="overlay = !overlay">
								Show Overlay
							</v-btn>
							<v-overlay :absolute="absolute" :opacity="opacity" :value="overlay" :z-index="zIndex">
								<v-btn color="primary" @click="overlay = false">
									Hide Overlay
								</v-btn>
							</v-overlay>
						</v-col>

						<div justify="center" class="overlay-input">
							<v-radio-group row >
								<v-checkbox v-model="absolute" label="Absolute"></v-checkbox>
								<v-checkbox v-model="overlay" class="mx-6" label="value"></v-checkbox>
								<v-text-field v-model="opacity" label="Opacity" max="1" min="0" step=".01" style="width: 125px"
									type="number" @keydown="false"></v-text-field>
								<v-text-field v-model="zIndex" label="z-index" max="9999" min="-9999" step="1"
									style="width: 125px" type="number" @keydown="false"></v-text-field>
							</v-radio-group>
						</div>
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
				overlay1: false,
				absolute: false,
				opacity: 0.46,
				overlay: false,
				zIndex: 5,
			}
		}
	}
</script>