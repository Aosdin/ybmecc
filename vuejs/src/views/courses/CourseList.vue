<template>
   <div class="courses-wrap">
		<page-title-bar></page-title-bar>
      <v-container fluid class="grid-list-xl">
         <course-banner></course-banner>
         <v-row class="align-center justify-center detail-course-list">
             <v-col cols="12" md="12" lg="12" xl="12">
               <v-row class="align-start fill-height">
                  <v-col cols="12" md="12" lg="12" xl="12">
                     <div class="popularity tab-wrap">
                        <v-tabs 
                           v-if="CourseData.courses"
                           color=""
                           slider-color="primary"
                        >
                           <v-tab v-if="CourseData.courses.popular == CourseData.courses.top"	ripple>
                              {{$t('message.top')}}
                           </v-tab>
                           <v-tab v-if="CourseData.courses.popular == CourseData.courses.new" ripple>
                              {{$t('message.new')}}
                           </v-tab>
                           <v-tab v-if="CourseData.courses.popular == CourseData.courses.trending" ripple>
                              {{$t('message.trending')}}
                           </v-tab>
                           <v-tab-item 
                              v-if="CourseData.courses.popular == CourseData.courses.top"
                           >
                              <course-card 
                                 :data="isTop"
                                 :colxl="3" :collg="3" :colmd="4" :colsm="6" :cols="12"
                                 :width="305"
                              ></course-card>
                           </v-tab-item>
                           <v-tab-item 
                              v-if="CourseData.courses.popular == CourseData.courses.new"
                           >
                              <course-card 
                                 :data="isNew"
                                 :colxl="3" :collg="3" :colmd="4" :colsm="6" :cols="12"
                                 :width="305"
                              ></course-card>
                           </v-tab-item>
                           <v-tab-item
                              v-if="CourseData.courses.popular == CourseData.courses.trending"
                           >
                              <course-card 
                                 :data="isTrending"
                                 :colxl="3" :collg="3" :colmd="4" :colsm="6" :cols="12"
                                 :width="305"
                              ></course-card>
                           </v-tab-item>
                        </v-tabs>
                     </div>                   
                  </v-col> 
                  <v-col cols="12" md="10" lg="10" xl="12" class="instructor-card-wrap">
                     <div>
                        <h3>{{$t('message.popularInstructors')}}</h3>
                     </div>
                     <instructor-card></instructor-card>
                  </v-col>
               </v-row>
            </v-col>
         </v-row>
      </v-container>
   </div>
</template>

<script>
import CourseBanner from "./CourseWidgets/CourseBanner";
import InstructorCard from "./CourseWidgets/InstructorCard";
import CourseCard from "./CourseWidgets/CourseCard";
import CourseData from "./data";

export default {
   data() {
		return{
			CourseData
		}
	},
	components: {
		CourseBanner,
      InstructorCard,
      CourseCard
   },
   computed: {
		isTop() {
			return this.CourseData.courses.filter(item=> {
            return item.popular=='top'
			})
		},
		isNew() {
			return this.CourseData.courses.filter(item=> {
            return item.popular=='new'
			})
		},
		isTrending() {
			return this.CourseData.courses.filter(item=> {
            return item.popular=='trending'
			})
		}
   }
}
</script>