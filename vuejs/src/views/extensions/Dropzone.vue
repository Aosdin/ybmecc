<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="mt-n3 pt-0">
			<v-row>
				<app-card colClasses="col-12 col-sm-12 col-md-12 col-lg-12" :fullScreen="true" :reloadable="true"
					:closeable="true" customClasses="blog-widget">
					<dropzone></dropzone>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
	import Dropzone from "Components/Widgets/Dropzone";
	export default {
		components: {
			Dropzone
		}
	}
</script>