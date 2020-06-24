<template>
   <div class="courses-signin">
      <page-title-bar></page-title-bar>
      <v-container fluid class="grid-list-xl">
         <course-banner></course-banner>
         <v-row class="align-center justify-center signIn-wrap">
            <v-col cols="12" sm="12" md="10" lg="10" xl="9">
               <div class="shop-by-category section-gap pad-y-lg">
                  <v-row class="justify-center">
                     <v-col cols="12" sm="8" md="6" lg="6" xl="6">
                        <app-card
                           contentCustomClass="pt-0"
                           :heading="$t('message.userSignIn')"
                        >
                           <v-form>
                              <div class="mrgn-b-md">
                                 <v-text-field
                                 placeholder="Email*"
                                    type="email"
                                    required
                                 >
                                 </v-text-field>
                              </div>
                              <div class="mrgn-b-md">
                                 <v-text-field
                                    placeholder="Password*"
                                    type="password"
                                    required
                                 >
                                 </v-text-field>
                              </div>
                              <v-row class="align-item-center justify-space-between mx-0">
                                 <v-checkbox :label="`Remember Me`"></v-checkbox>
                                 <router-link to="/session/forgot-password" class="mt-3">
                                    Forgot Password ?
                                 </router-link>
                              </v-row>
                              <v-btn color="error" class="danger mx-0 mb-5"  to="/session/login">{{$t('message.signIn')}}</v-btn>
                              <p>Don't have an account? <router-link to="/session/sign-up" class="primary-text">Click here to create one</router-link></p>
                           </v-form>
                        </app-card>
                     </v-col>
                     <v-col cols="12" sm="8" md="6" lg="6" xl="6">
                        <app-card :heading="$t('message.guestCheckout')"  contentCustomClass="pt-0">
                           <p>Proceed to checkout and create an account later.</p>
                           <v-btn 
                              class="error mx-0"
                              :to="`/${getCurrentAppLayoutHandler() + '/courses/payment'}`"
                           >{{$t('message.continueAsGuest')}}</v-btn>
                        </app-card>
                     </v-col>
                  </v-row>
               </div>    
            </v-col>
         </v-row>
      </v-container>
   </div>
</template>

<script>
import CourseBanner from "./CourseWidgets/CourseBanner";
import { getCurrentAppLayout } from "Helpers/helpers";
export default {
   components:{
		CourseBanner
   },
   methods: {
		getCurrentAppLayoutHandler() {
			return getCurrentAppLayout(this.$router);
      }
   }
}
</script>
