<template>
   <app-card 
      class="course-info-wrap"
      :heading="$t('message.whatYoWillLearn')"
      colClasses="col-12 col-sm-12 col-md-12 col-lg-12 col-xl-12"
      contentCustomClass="pt-0"
   >
      <div class="course-info-box">
         <ul class="info-box-list">
            <li v-for= "(info, key) in CourseData.courseDetail.learn" :key="key"><i class="material-icons">check</i><span>{{info}}</span></li>
         </ul>
      </div>
   </app-card>
</template>

<script>
import CourseData from "../data";
export default {
   data() {
		return{
         CourseData
      }
   }
}
</script>