<template class="custom-flex">
   <div class="col-12 col-sm-12 col-md-12 col-lg-12 col-xl-12 courses-overview">
		<v-expansion-panels>
			<v-expansion-panel 
				v-for="(list,key) in  CourseData.courseDetail.topics" 
				:key="key"
			>
				<v-expansion-panel-header>{{list.name}}</v-expansion-panel-header>
				<v-expansion-panel-content>
					<v-list v-for="(details,key) in list.courseDetail" :key="key">
					<v-list-item  :to="`/${getCurrentAppLayoutHandler() + '/courses/courses-detail'}`">
						<v-list-item-action>
							<v-icon>play_circle_filled</v-icon>	
						</v-list-item-action>
						<v-list-item-content>
							<v-list-item-title  class="">
								{{details.name}}							
							</v-list-item-title>
						</v-list-item-content>
						<v-list-item-action>
							<div class="ma-0 d-flex justify-space-between align-center">
								<div class="pr-5"><a href="#" class="fs-12 fw-normal">Preview</a></div>
								<div class="pl-5"><span class="fs-12 fw-normal">{{details.time}}</span></div>
							</div>
						</v-list-item-action>
					</v-list-item>
					<v-divider></v-divider>
				</v-list>
				</v-expansion-panel-content>
				</v-expansion-panel>
			</v-expansion-panels>      
   </div>
</template>

<script>
import CourseData from "../data";
import { getCurrentAppLayout } from "Helpers/helpers";

export default {
   data() {
		return {
         CourseData
      }
   },
	methods: {
		getCurrentAppLayoutHandler() {
			return getCurrentAppLayout(this.$router);
		}
	}
}
</script>