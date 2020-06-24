<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0">
			<app-card
				:heading="$t('message.defaultInput')"
				contentCustomClass="input-label"
			>
				<v-row>
					<v-col cols="12" sm="4">
						<span class="small pt-4 d-block">Normal with hint text/label</span>
					</v-col>
					<v-col cols="12" sm="8">
					<v-text-field name="input-1" label="Label Text" id="testing"></v-text-field>
					</v-col>
				</v-row>
				<v-row>
					<v-col cols="12" sm="4">
						<span class="small pt-4 d-block">Focus</span>
					</v-col>
					<v-col cols="12" sm="8">
						<v-text-field name="input-2" label="Label Text" value="Input text" class="input-group--focused"></v-text-field>
					</v-col>
				</v-row>
				<v-row>
					<v-col cols="12" sm="4">
						<span class="small pt-4 d-block">Normal with input text + label</span>
					</v-col>
					<v-col cols="12" sm="8">
						<v-text-field name="input-3" label="Label Text" value="Input text"></v-text-field>
					</v-col>
				</v-row>
				<v-row>
					<v-col cols="12" sm="4">
						<span class="small pt-4 d-block">Disabled</span>
					</v-col>
					<v-col cols="12" sm="8">
						<v-text-field name="input-3" label="Label Text" value="Input text" disabled></v-text-field>
					</v-col>
				</v-row>
			</app-card>
		</v-container>
	</div>
</template>
