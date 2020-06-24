<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid grid-list-xl pt-0>
			<v-row>
				<!-- Support Request -->
				<app-card
					colClasses="col-xl-4 col-lg-4 col-md-4 col-sm-6 col-12"
					:heading="$t('message.supportRequest')"
					:fullScreen="true"
					:reloadable="true"
					:closeable="true"
					:footer="true"
					:fullBlock="true"
					customClasses="support-widget-wrap"
				>
					<support-request></support-request>
					<div slot="footer" class="justify-space-between footer-flex align-items-center">
						<span class="grey--text d-custom-flex align-items-center">
							<i class="zmdi zmdi-replay mr-2"></i> 
							<span class="fs-12 fw-normal">Updated 10 min ago</span>
						</span>
						<v-btn class="ma-0" color="primary" small>{{$t('message.assignNow')}}</v-btn>
					</div>    
				</app-card>
				<app-card 
					:fullBlock="true"
					colClasses="col-xl-4 col-lg-4 col-md-4 col-sm-6 col-12"
				>
					<blog-layout-two></blog-layout-two>
				</app-card>
				<app-card 
					:fullBlock="true"
					colClasses="col-xl-4 col-lg-4 col-md-4 col-sm-12 col-12"
				>
					<blog-layout-three></blog-layout-three>
				</app-card>
			</v-row>
			<v-row>
				<app-card 
					colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12"
					heading="New Emails"
					:fullBlock="true"
					:closeable="true"
					:reloadable="true"
					:fullScreen="true"
				>
					<new-emails></new-emails>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import SupportRequest from "Components/Widgets/SupportRequest";
import BlogLayoutTwo from "Components/Widgets/BlogLayoutTwo";
import BlogLayoutThree from "Components/Widgets/BlogLayoutThree";
import NewEmails from "Components/Widgets/NewEmails";

import { activeUser } from "../data";

export default {
  components: {
    SupportRequest,
    BlogLayoutTwo,
    BlogLayoutThree,
    NewEmails
  },
  data() {
    return {
      activeUser
    };
  }
};
</script>
