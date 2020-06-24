<template>
	<div>
      <page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0">
			<v-row class="align-items-center search-wrap">
            <v-col cols="12" md="12" lg="12"  class="align-items-center pt-0">
               <app-card
                  customClasses="mb-0 py-12"
               >
                  <v-row >
                     <v-col cols="12" md="12" lg="3" class="align-center">
								<h2 class="mb-0">Search</h2> 
                     </v-col>      
                     <v-col cols="12" md="12" lg="9" class="pb-0"> 
                        <div class="d-flex search-field-wrap">   
                           <div class="w-75">
                              <v-text-field 
                                 class=" pt-0 pr-3"
                                 label="Search Projects"
                              >
                              </v-text-field>
                           </div>
                           <div>
                              <v-btn color="primary" class="my-0 ml-0 mr-2" medium>Search</v-btn>
                              <v-btn color="primary m-0" medium>Add<i class="material-icons">add</i></v-btn>
                           </div>
                        </div>
                     </v-col>
                   </v-row>
               </app-card>
            </v-col>   
         </v-row>   
         <div class="d-flex justify-space-between align-items-center pa-6 project-grid-title">
            <div class="title">
               <h3 class="mb-0">{{$t('message'+'.'+viewType)}}</h3>
            </div>
            <div class="text-right project-icon"> 
               <v-icon class="mr-2" :class="{active: isActive == 'grid'}" style="cursor: pointer;" @click="girdView(1)">apps</v-icon>
               <v-icon :class="{active: isActive == 'list'}" style="cursor: pointer;" @click="listView(2)">list</v-icon>
            </div>
         </div>      
			<v-row>
            <app-card
               v-for="(data,index) in projectData" :key="index"
               :heading="$t('message'+'.'+data.name)"
					colClasses="col-xl-4 col-lg-4 col-md-4 col-sm-6 col-12" 
					customClasses="mb-0"
					:fullBlock="true"
					:fullScreen="true"
					:reloadable="true"
					:closeable="true"
               v-show="selectedView == 'grid'"
				>
					<ProjectGridView :managementData="data"></ProjectGridView>
				</app-card>  
			</v-row>
         <v-row v-show="selectedView == 'list'">
            <app-card
               colClasses="col-xl-12 col-lg-12 col-md-12 col-12 col-12"
               customClasses="mb-0"
            >             
            <ProjectListView ></ProjectListView>
            </app-card>
         </v-row>
		</v-container>
	</div>
</template>

<script>

// widgets
import ProjectGridView from "Components/Widgets/ProjectGrid";
import ProjectListView from "Components/Widgets/ProjectList";
import api from "Api";


export default {
   components: {
      ProjectGridView,
      ProjectListView
   },
   data() {
      return {
         projectData: [],
         viewType: "projectGrid",
         selectedView: "grid",
         isActive: 'grid'
      };
   },
   mounted() {
      this.getProjectData();
   },
   methods:{
      getProjectData() {
         this.loader = true;
         api
            .get("vuely/projectDetails.js")
            .then(response => {
               this.loader = false;
               this.projectData = response.data;
            })
            .catch(error => {
               console.log("error" + error);
            });
      },
      listView(){
         this.viewType = "projectList";
         this.selectedView = "list";
         this.isActive = 'list';
      },
      girdView(){
         this.viewType = "projectGrid";
         this.selectedView = "grid";
         this.isActive = "grid";         
      }
  }
};
</script>