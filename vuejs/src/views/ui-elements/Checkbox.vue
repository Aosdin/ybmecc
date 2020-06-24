<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid grid-list-xl pt-0>
			<v-row>
				<app-card 
					:heading="$t('message.checkboxesBoolean')"
					colClasses="col-12 col-md-6"
					>
					<v-checkbox color="primary" :label="`Checkbox : ${checkbox}`" v-model="checkbox"></v-checkbox>
					<v-checkbox color="primary" :label="`Intermediate: ${intermediate.toString()}`" v-model="intermediate"></v-checkbox>
				</app-card>
				<app-card 
					:heading="$t('message.checkboxesArray')"
					colClasses="col-12 col-md-6"
					>
					<p>{{ name1 }}</p>
					<v-checkbox color="primary" label="John" v-model="name1" value="John"></v-checkbox>
					<v-checkbox color="primary" label="Jacob" v-model="name1" value="Jacob"></v-checkbox>
				</app-card>
			</v-row>
			<v-row>
				<app-card 
					:heading="$t('message.checkboxesStates')"
					colClasses="col-12 col-md-6"
					>
					<v-row>
						<v-col cols="4">on</v-col>
						<v-col cols="4">off</v-col>
						<v-col cols="4">indeterminate</v-col>
					</v-row>
					<v-row>
						<v-col cols="4" class="col-height-auto">
							<v-checkbox color="primary" input-value="true" value></v-checkbox>
						</v-col>
						<v-col cols="4" class="col-height-auto">
							<v-checkbox color="primary" value></v-checkbox>
						</v-col>
						<v-col cols="4" class="col-height-auto">
							<v-checkbox color="primary" value indeterminate></v-checkbox>
						</v-col>
					</v-row>
					<v-row>
						<v-col cols="4" class="col-height-auto">on disabled</v-col>
						<v-col cols="4" class="col-height-auto">off disabled</v-col>
					</v-row>
					<v-row>
						<v-col cols="4" class="col-height-auto">
							<v-checkbox color="primary" input-value="true" value disabled></v-checkbox>
						</v-col>
						<v-col cols="4" class="col-height-auto">
							<v-checkbox color="primary" value disabled></v-checkbox>
						</v-col>
					</v-row>
				</app-card>
				<app-card 
					:heading="$t('message.checkboxesColors')"
					colClasses="col-12 col-md-6"
					>
					<v-row>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="red" v-model="ex4" color="red" value="red" hide-details>
							</v-checkbox>
							<v-checkbox label="red darken-3" v-model="ex4" color="red darken-3" value="red darken-3" hide-details>
							</v-checkbox>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="indigo" v-model="ex4" color="indigo" value="indigo" hide-details>
							</v-checkbox>
							<v-checkbox label="indigo darken-3" v-model="ex4" color="indigo darken-3" value="indigo darken-3" hide-details>
							</v-checkbox>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="orange" v-model="ex4" color="orange" value="orange" hide-details>
							</v-checkbox>
							<v-checkbox label="orange darken-3" v-model="ex4" color="orange darken-3" value="orange darken-3" hide-details>
							</v-checkbox>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="primary" v-model="ex4" color="primary" value="primary" hide-details>
							</v-checkbox>
							<v-checkbox label="secondary" v-model="ex4" color="secondary" value="secondary" hide-details>
							</v-checkbox>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="success" v-model="ex4" color="success" value="success" hide-details>
							</v-checkbox>
							<v-checkbox label="info" v-model="ex4" color="info" value="info" hide-details>
							</v-checkbox>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-checkbox label="warning" v-model="ex4" color="warning" value="warning" hide-details>
							</v-checkbox>
							<v-checkbox label="error" v-model="ex4" color="error" value="error" hide-details>
							</v-checkbox>
						</v-col>
					</v-row>
				</app-card>
			</v-row>
			<v-row>
				<app-card 
					:heading="$t('message.switches')"
					colClasses="col-lg-12 col-height-auto"
					>
					<v-row>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="red" v-model="ex11" color="red" value="red" hide-details>
							</v-switch>
							<v-switch label="red darken-3" v-model="ex11" color="red darken-3" value="red darken-3" hide-details>
							</v-switch>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="indigo" v-model="ex11" color="indigo" value="indigo" hide-details>
							</v-switch>
							<v-switch label="indigo darken-3" v-model="ex11" color="indigo darken-3" value="indigo darken-3" hide-details>
							</v-switch>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="orange" v-model="ex11" color="orange" value="orange" hide-details>
							</v-switch>
							<v-switch label="orange darken-3" v-model="ex11" color="orange darken-3" value="orange darken-3" hide-details>
							</v-switch>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="primary" v-model="ex11" color="primary" value="primary" hide-details>
							</v-switch>
							<v-switch label="secondary" v-model="ex11" color="secondary" value="secondary" hide-details>
							</v-switch>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="success" v-model="ex11" color="success" value="success" hide-details>
							</v-switch>
							<v-switch label="info" v-model="ex11" color="info" value="info" hide-details>
							</v-switch>
						</v-col>
						<v-col cols="12" sm="4" md="4" class="col-height-auto">
							<v-switch label="warning" v-model="ex11" color="warning" value="warning" hide-details>
							</v-switch>
							<v-switch label="error" v-model="ex11" color="error" value="error" hide-details>
							</v-switch>
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
      checkbox: false,
      intermediate: true,
      name1: ["John"],
      ex4: [
        "red",
        "indigo",
        "orange",
        "primary",
        "secondary",
        "success",
        "info",
        "warning",
        "error",
        "red darken-3",
        "indigo darken-3",
        "orange darken-3"
      ],
      ex11: [
        "red",
        "indigo",
        "orange",
        "primary",
        "secondary",
        "success",
        "info",
        "warning",
        "error",
        "red darken-3",
        "indigo darken-3",
        "orange darken-3"
      ]
    };
  }
};
</script>
