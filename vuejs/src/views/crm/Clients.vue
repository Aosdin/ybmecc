<template>
	<div class="contact-wrapper">
      <page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0">
         <v-row class="align-items-center search-wrap">
            <v-col cols="12" md="12" lg="12"  class="align-items-center pt-0">
               <app-card customClasses="mb-0 py-12">
                  <v-row >
                     <v-col cols="12" md="12" lg="3"  class="align-center">
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
                              <v-btn color="primary m-0" medium @click="openDialog()" >Add<i class="material-icons">add</i></v-btn>
                           </div>
                        </div>
                     </v-col>
                   </v-row>
               </app-card>
            </v-col>   
         </v-row>   
         <v-row>
            <app-card
               colClasses="col-xl-12 col-lg-12 col-md-12 col-12 "
               customClasses="mb-0 client-wrapper"
            >
               <clients-details :clientData="clientData"></clients-details>
            </app-card>
         </v-row>
		</v-container>
      <template>
         <v-dialog v-model="open" max-width="600">
           <v-card class="pa-6">
               <v-text-field label="Name" v-model="editName" required></v-text-field>
               <v-text-field label="Email" v-model="editEmail" required></v-text-field>
               <v-text-field label="Phone Number"  v-model="editPhoneNo" required></v-text-field>
               <v-text-field label="Location" v-model="editLocation" required></v-text-field>

               <v-card-actions class="pa-0">
                  <v-spacer></v-spacer>
                  <v-btn color="primary" @click="open = false">Cancel</v-btn>
                  <v-btn color="error" @click="addClient()">Add</v-btn>
               </v-card-actions>
           </v-card>
         </v-dialog>
      </template>
	</div>
</template>

<script>
import ClientsDetails from 'Components/Widgets/ClientsDetails'
import { clientData } from 'Views/crm/data.js'

export default {
  components: {
     ClientsDetails
  },
  data() {
    return {
      clientData,
      open: false,
      editName:'',
      editEmail: '',
      editPhoneNo:'',
      editLocation:''          
    };
  },
  methods:{
    openDialog(){
       this.open = true;
    },
    close(){
       this.open = false;
    },
    addClient(){
      this.openDialog();
      var clientsArr = clientData.Clients;
      
      var newClient = {
        image : "/static/avatars/user-1.jpg",
        name: this.editName,
        e_mail: this.editEmail,
        phone_number: this.editPhoneNo,
        country: this.editLocation,
        tag: "recently_added"
      }
      clientsArr.push(newClient);
      this.editName='';
      this.editEmail='';
      this.editPhoneNo='';
      this.editLocation='';

      this.close();
    }
  }
};
</script>
