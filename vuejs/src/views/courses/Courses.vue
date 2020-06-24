<template>
	<div class="courses-wrap">
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl">
			<course-banner></course-banner>
			<v-row class="align-center justify-center fill-height courses-inner">
				<v-col sm="12" cols="12" md="12" lg="12" xl="12">
					<div class="courses-grid-sec">
						<v-row class="custom-align-stretch fill-heigh course-grid-layout">
							<app-card :fullScreen="true" :closeable="true"
								colClasses="col-12 col-sm-12 col-md-12 col-lg-4 col-xl-4"
								customClasses="custom-grey collection-group px-0 pb-0">
								<h2>Best collection from best tutors from around the globe</h2>
								<p>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
							</app-card>
							<app-card :fullScreen="true" :closeable="true"
								colClasses="col-12 col-sm-12 col-md-12 col-lg-8 col-xl-8" customClasses="tab-wrap pa-0" content>
								<v-tabs v-if="CourseData.courses" slider-color="primary">
									<v-tab v-if="CourseData.courses.type == CourseData.courses.management" ripple>
										{{$t('message.management')}}
									</v-tab>
									<v-tab v-if="CourseData.courses.type == CourseData.courses.design" ripple>
										{{$t('message.design')}}
									</v-tab>
									<v-tab v-if="CourseData.courses.type == CourseData.courses.development" ripple>
										{{$t('message.development')}}
									</v-tab>
									<v-tab-item v-if="CourseData.courses.type == CourseData.courses.management">
										<course-card :data="isManagement.slice(0, 3)" :colxl="4" :collg="4" :colmd="4" :colsm="6"
											:cols="12"></course-card>
									</v-tab-item>
									<v-tab-item v-if="CourseData.courses.type == CourseData.courses.design">
										<course-card :data="isDesign" :colxl="4" :collg="4" :colmd="4" :colsm="6" :cols="12">
										</course-card>
									</v-tab-item>
									<v-tab-item v-if="CourseData.courses.type == CourseData.courses.development">
										<course-card :data="isDevelop" :colxl="4" :collg="4" :colmd="4" :colsm="6" :cols="12">
										</course-card>
									</v-tab-item>
								</v-tabs>
								<div class="view-all-link">
									<router-link :to="`/${getCurrentAppLayoutHandler() + '/courses/courses-list'}`">View All
									</router-link>
								</div>
							</app-card>
						</v-row>
					</div>
					<div class="courses-grid">
						<v-row class="align-start justify-start fill-height mb-0">
							<v-col sm="12" s="12" md="12" lg="12" xl="12">
								<div class="popularity tab-wrap">
									<v-tabs v-if="CourseData.courses" slider-color="primary">
										<v-tab v-if="CourseData.courses.popular == CourseData.courses.top" ripple>
											{{$t('message.top')}}
										</v-tab>
										<v-tab v-if="CourseData.courses.popular == CourseData.courses.new" ripple>
											{{$t('message.new')}}
										</v-tab>
										<v-tab v-if="CourseData.courses.popular == CourseData.courses.trending" ripple>
											{{$t('message.trending')}}
										</v-tab>
										<v-tab-item v-if="CourseData.courses.popular == CourseData.courses.top">
											<course-card :data="isTop" :colxl="3" :collg="3" :colmd="4" :colsm="6"
												:cols="12"></course-card>
										</v-tab-item>
										<v-tab-item v-if="CourseData.courses.popular == CourseData.courses.new">
											<course-card :data="isNew" :colxl="3" :collg="3" :colmd="4" :colsm="6"
												:cols="12"></course-card>
										</v-tab-item>
										<v-tab-item v-if="CourseData.courses.popular == CourseData.courses.trending">
											<course-card :data="isTrending" :colxl="3" :collg="3" :colmd="4" :colsm="6"
												:cols="12"></course-card>
										</v-tab-item>
									</v-tabs>
								</div>
							</v-col>
						</v-row>
					</div>
					<div class="instructor-card-wrap">
						<h3>{{$t('message.popularInstructors')}}</h3>
						<instructor-card></instructor-card>
					</div>
				</v-col>
			</v-row>
		</v-container>
	</div>
</template>

<script>
	import CourseBanner from "./CourseWidgets/CourseBanner";
	import InstructorCard from "./CourseWidgets/InstructorCard";
	import CourseData from "./data";
	import CourseCard from "./CourseWidgets/CourseCard";
	import { getCurrentAppLayout } from "Helpers/helpers";
	export default {
		data() {
			return {
				CourseData,
			}
		},
		computed: {
			isManagement() {
				return this.CourseData.courses.filter(item => {
					return item.type == 'management'
				})
			},
			isDesign() {
				return this.CourseData.courses.filter(item => {
					return item.type == 'design'
				})
			},
			isDevelop() {
				return this.CourseData.courses.filter(item => {
					return item.type == 'develop'
				})
			},
			isTop() {
				return this.CourseData.courses.filter(item => {
					return item.popular == 'top'
				})
			},
			isNew() {
				return this.CourseData.courses.filter(item => {
					return item.popular == 'new'
				})
			},
			isTrending() {
				return this.CourseData.courses.filter(item => {
					return item.popular == 'trending'
				})
			}

		},
		components: {
			CourseBanner,
			InstructorCard,
			CourseCard
		},
		methods: {
			getCurrentAppLayoutHandler() {
				return getCurrentAppLayout(this.$router);
			}
		}
	}
</script>