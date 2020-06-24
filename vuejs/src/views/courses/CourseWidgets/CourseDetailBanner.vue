<template>
   <div class="banner-image-wrap banner-detail grey darken-4 ma-0">
      <v-container class="grid-list-xl" fluid>
         <div class="banner-content-wrap">
            <v-row class="align-center justify-center fill-height">
               <v-col cols="12" md="12" lg="12" xl="12">
                  <v-row class="align-center justify-center fill-height">
                     <v-col cols="12" sm="6" md="7" lg="7" xl="8">
                        <h2 class="white--text">{{CourseData.courseDetail.heading}}</h2>
                        <h4 class="white--text">{{CourseData.courseDetail.content}}</h4>
                        <div class="meta-info">
                           <p class="white--text">
                              <span class="bestseller-tag d-inline-block mr-4 mb-1">{{$t('message.bestseller')}}</span>
                              <span class="rating-wrap d-inline-block"></span>
                              <span class="student-count d-inline-block"><span class="count font-weight-bold">{{CourseData.courseDetail.bestSeller}}</span> students enrolled</span>
                           </p>
                           <p class="white--text">
                              <span class="white--text">Created By </span>
                              <span class="instructor-name d-inline-block"><a href="#" class="white--text"> {{CourseData.courseDetail.createdBy}} </a></span>
                              <span class="updated-time d-inline-block"> Last Updated {{CourseData.courseDetail.lastUpdates}}</span>
                           </p>
                           <p class="white--text">
                              <span class="language layout align-start ma-0"><v-icon class="cmr-8">chat_bubble_outline</v-icon>{{CourseData.courseDetail.language}}</span>
                           </p>
                        </div>
                     </v-col>
                     <v-col cols="12" sm="6" md="5" lg="4" xl="3">
                        <app-card contentCustomClass="pa-4">                        
                           <div class="video-wrap overlay-wrap d-inline-flex">
                              <img class="banner-video w-100" src="/static/img/about2.png" width="230" height="235" alt="video">
                              <div class="overlay-content layout align-center justify-center">
                                 <v-dialog
                                    v-model="dialog"
                                    width="500"
                                    height="300"
                                    >
                                    <template v-slot:activator="{ on }">
                                       <v-btn  v-on="on" icon>
                                          <v-icon>play_circle_filled</v-icon>
                                       </v-btn>
                                    </template>
                                    <iframe  :src="CourseData.courseDetail.demoVideoUrl" frameborder="0" allowfullscreen></iframe>
                                 </v-dialog>
                              </div>
                           </div>
                        </app-card>
                     </v-col>
                  </v-row>
               </v-col>
            </v-row>
         </div>
      </v-container>
   </div>
</template>

<script>
import CourseData from "../data";
export default {
   data(){
		return {
         CourseData,
         dialog: false
      }
   }
}
</script>
