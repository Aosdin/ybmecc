<template>
   <div class="courses-detail-wrap" >
      <page-title-bar></page-title-bar>
      <v-container class="grid-list-xl" fluid>
			<course-detail-banner></course-detail-banner>
			<div class="course-detail">
            <v-row class="justify-center align-center">
               <v-col cols="12" sm="12" md="12" lg="12" xl="12">
                  <v-row>
                     <v-col cols="12" sm="12" md="8" lg="9" xl="9">
                        <v-row>
                           <course-detail-learn></course-detail-learn>
                           <course-detail-desciption></course-detail-desciption>  
                           <course-detail-overview></course-detail-overview>
                           <course-detail-instructor></course-detail-instructor>
                           <app-card 
                              colClasses="col-12 col-sm-12 col-md-12 col-lg-12 col-xl-12"
                              customClasses="more-courses-grid"
                              contentCustomClass="layout pt-0"
                              :heading="$t('message.moreCoursesFromJamesColt')"
                           >                     
                              <v-row class="align-start justify-start mx-0 mb-0">
                                 <div class="course-item-wrap">
                                    <course-card 
                                       :height="200"
                                       :width="335"
                                       :data="CourseData.courseDetail.moreCourses"
                                       :colxl="4" :collg="4" :colmd="4" :colsm="6" :cols="12"
                                    ></course-card>
                                 </div>
                              </v-row>
                           </app-card>
                        </v-row>
                     </v-col>
                     <v-col cols="12" sm="12" md="4" lg="3" xl="3" class="course-sidebar">
                        <div class="custom-height-auto">
                           <course-detail-billing></course-detail-billing>
                           <app-card 
                              contentCustomClass="pt-0"
                              heading="Contrary to popular belief ?"
                           >
                              <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
                              <a href="#" class="font-weight-bold primary-text">Nulla eu augue !</a>
                           </app-card>
                        </div>
                     </v-col>
                  </v-row>
					</v-col>
            </v-row>
			</div>
      </v-container>
   </div>
</template>

<script>
import CourseData from "./data";
import CourseDetailLearn from "./CourseWidgets/CourseDetailLearn";
import CourseDetailBanner from "./CourseWidgets/CourseDetailBanner";
import CourseDetailDesciption from "./CourseWidgets/CourseDetailDescription";
import CourseDetailOverview from "./CourseWidgets/CourseDetailOverview";
import CourseDetailInstructor from "./CourseWidgets/CourseDetailInstructor";
import CourseCard from "./CourseWidgets/CourseCard";
import CourseDetailBilling from "./CourseWidgets/CourseDetailBilling";
export default {
	data(){
		return{
			CourseData
		}
	},
   components:{
		CourseDetailBanner,
		CourseDetailLearn,
		CourseDetailDesciption,
		CourseDetailOverview,
		CourseDetailInstructor,
		CourseCard,
		CourseDetailBilling
	}
}
</script>
