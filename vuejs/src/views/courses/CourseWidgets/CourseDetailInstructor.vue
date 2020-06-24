<template>
    <app-card  
        class="about-instructor"
        v-if="CourseData.courseDetail"
        :heading="$t('message.aboutInstructor')"
			colClasses="col-12 col-sm-12 col-md-12 col-lg-12 col-xl-12"
			contentCustomClass="pt-0"
    >
		<div class="d-flex align-center justify-start mx-0 mt-0 mb-4">
			<div class="image-wrap mb-2">
				<img :src="CourseData.courseDetail.instructorInformation.image">
			</div>
			<div class="instructor-meta">
				<h4>{{CourseData.courseDetail.instructorInformation.name}}</h4>  
				<v-list 
					class="incentive-list"
					v-for="(feature,key) of CourseData.courseDetail.instructorInformation.features"
					:key="key"
				>
					<v-list-item class="d-custom-flex">
						<v-list-item-action class="my-0 mr-1">
							<v-icon>{{feature.icon}}</v-icon>
						</v-list-item-action>
						<v-list-item-content class="py-0">
							<v-list-item-title><p class="mb-0">{{feature.feature}}</p></v-list-item-title>
						</v-list-item-content>
					</v-list-item>
				</v-list>
			</div>
		</div>
		<p>{{ CourseData.courseDetail.instructorInformation.content}}</p>
   </app-card>
</template>

<script>
import CourseData from "../data";
export default {
   data(){
		return{
         CourseData
      }
   }
}
</script>
<style scoped>
	.image-wrap{
		width:130px;
	}
	.instructor-meta{
		width:calc(100% - 130px);
	}
	@media (max-width: 600px){
		.instructor-meta{
			width:100%;
		}
	}
</style>