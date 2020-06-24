<template>
   <div class="hover-wrapper">
      <page-title-bar></page-title-bar>
		<v-container class="grid-list-xl pt-0 mt-n3">
			<v-row>				
            <app-card
					colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12"
				>
               <div class="mb-5">
						<p>Easily alternate styles to provide a unique design.</p>
					</div>   
               <v-timeline>
                  <v-timeline-item
                     color="purple lighten-2"
                     fill-dot
                     right
                  >
                     <v-card>
                        <v-card-title class="purple lighten-2">
                           <v-icon
                              size="42"
                              class="mr-3 white--text"
                           >
                              mdi-magnify
                           </v-icon>
                           <h2 class="display-1 white--text font-weight-light">Title 1</h2>
                        </v-card-title>
                        <v-container>
                           <v-row>
                              <v-col xs="10">
                                 Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit.
                              </v-col>
                              <v-col xs="2">
                                 <v-icon size="64">mdi-calendar-text</v-icon>
                              </v-col>
                           </v-row>
                        </v-container>
                     </v-card>
                  </v-timeline-item>

                  <v-timeline-item
                     color="amber lighten-1"
                     fill-dot
                     left
                     small
                  >
                     <v-card>
                        <v-card-title class="amber lighten-1 justify-end">
                           <h2 class="display-1 mr-3 white--text font-weight-light">Title 2</h2>
                           <v-icon
                              size="42"
                              class="white--text"
                           >
                              mdi-home-outline
                           </v-icon>
                        </v-card-title>
                        <v-container>
                           <v-row>
                              <v-col xs="8">
                                 Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit.
                              </v-col>
                              <v-col xs="4">
                                 Lorem ipsum dolor sit amet, no nam oblique veritus.
                              </v-col>
                           </v-row>
                        </v-container>
                     </v-card>
                  </v-timeline-item>

                  <v-timeline-item
                     color="cyan lighten-1"
                     fill-dot
                     right
                  >
                     <v-card>
                        <v-card-title class="cyan lighten-1">
                           <v-icon
                              class="mr-3 white--text"
                              size="42"
                           >
                              mdi-email-outline
                           </v-icon>
                           <h2 class="display-1 white--text font-weight-light">Title 3</h2>
                        </v-card-title>
                        <v-container>
                           <v-row>
                              <v-col
                                 v-for="n in 3"
                                 :key="n"
                                 xs4
                              >
                                 Lorem ipsum dolor sit amet, no nam oblique veritus no nam oblique.
                              </v-col>
                           </v-row>
                        </v-container>
                     </v-card>
                  </v-timeline-item>

                  <v-timeline-item
                     color="red lighten-1"
                     fill-dot
                     left
                     small
                  >
                     <v-card>
                        <v-card-title class="red lighten-1 justify-end">
                           <h2 class="display-1 mr-3 white--text font-weight-light">Title 4</h2>
                           <v-icon
                              size="42"
                              class="white--text"
                           >
                              mdi-account-multiple-outline
                           </v-icon>
                        </v-card-title>
                        <v-container>
                           <v-row>
                              <v-col xs="2">
                                 <v-icon size="64">mdi-server-network</v-icon>
                              </v-col>
                              <v-col xs="10">
                                 Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus.
                              </v-col>
                           </v-row>
                        </v-container>
                     </v-card>
                  </v-timeline-item>

                  <v-timeline-item
                     color="green lighten-1"
                     fill-dot
                     right
                  >
                     <v-card>
                        <v-card-title class="green lighten-1">
                           <v-icon
                              class="mr-3 white--text"
                              size="42"
                           >
                              mdi-phone-in-talk
                           </v-icon>
                           <h2 class="display-1 white--text font-weight-light">Title 5</h2>
                        </v-card-title>
                        <v-container>
                           <v-row>
                              <v-col>
                                 Lorem ipsum dolor sit amet, no nam oblique veritus. Commune scaevola imperdiet nec ut, sed euismod convenire principes at. Est et nobis iisque percipit, an vim zril disputando voluptatibus, vix an salutandi sententiae.
                              </v-col>
                           </v-row>
                        </v-container>
                     </v-card>
                  </v-timeline-item>
               </v-timeline>
            </app-card>
         </v-row>
      </v-container>
   </div>
</template>