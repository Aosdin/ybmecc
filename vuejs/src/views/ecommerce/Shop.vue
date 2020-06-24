<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid pt-0 grid-list-xl>
			<!-- search box -->
         <v-row class="align-items-center search-wrap">
            <v-col cols="12" sm="12" md="12" lg="12" class="align-items-center pt-0">
               <app-card
                  customClasses="mb-0 py-12"
               >
                  <v-row>
                     <v-col cols="12" sm="12" md="12" lg="3" class="align-center">
								<h2 class="mb-0">Search</h2>
                     </v-col>      
                     <v-col cols="12" sm="12" md="12" lg="9" class="pb-0"> 
                        <div class="d-flex search-field-wrap">   
                           <div class="w-75">
                              <v-text-field 
                                 class=" pt-0 pr-4"
                                 label="Search Projects"
                                 >
                              </v-text-field>
                           </div>
                           <div>
                              <v-btn color="primary" class="my-0 ml-0 mr-2" medium>Search</v-btn>
                           </div>
                        </div>
                     </v-col>
                   </v-row>
               </app-card>
            </v-col>   
         </v-row>
         <!-- Actions -->
         <v-row class="pa-3 align-items-center justify-space-between">
				<div class="title">
					<h3 class="mb-0">{{$t('message'+'.'+viewType)}}</h3>
				</div>
            <div class="text-right project-icon"> 
               <v-icon class="mr-2" :class="{active: isActive == 'grid'}" style="cursor: pointer;" @click="girdView()">apps</v-icon>
               <v-icon :class="{active: isActive == 'list'}" style="cursor: pointer;" @click="listView()">list</v-icon>
            </div>
			</v-row>
			<div class="d-md-inline-flex mb-5">
				<v-select :items="type" label="Type" class="mr-md-5 select-width-150"></v-select>
				<v-select :items="recent" label="Recent" class="mr-md-5 select-width-150 "></v-select>
				<v-select :items="noOfItems" label="No of Items" class="select-width-150"></v-select>
			</div>
         <!-- grid view			 -->
         <div v-show="selectedView == 'grid'">
               <products :productsData="productsData"></products>
			</div> 
         <!-- list view -->
			<v-row class="align-center search-wrap">
				<v-col cols="12" sm="10" md="8" lg="8" class="mx-auto">
					<div v-show="selectedView == 'list'">
						<productslist :productsData="productsData"></productslist>
					</div>
				</v-col>
			</v-row>
		</v-container>
	</div>
</template>
<script>
	import Products from 'Components/Widgets/Products'
   import Productslist from 'Components/Widgets/ProductsList'
	import { productsData } from 'Views/ecommerce/data.js'
	export default {
		components: {
         Products,
         Productslist  
      },
      data() {
      return {
         productsData,
         type: ['Men', 'Women', 'Gadgets', 'Accessories'],
         recent:['This Week', 'This Month', 'Past Month'],
         noOfItems: ['10', '20', '30'],
         viewType: "projectGrid",
         selectedView: "grid",
         isActive: 'grid'
      };
   },
     methods:{
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
	}
</script>
