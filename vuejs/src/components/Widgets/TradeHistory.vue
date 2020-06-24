<template>
   <div>
      <v-card text>
         <div class="table-responsive trade-history-table">
            <app-section-loader :status="loader"></app-section-loader>
            <v-data-table 
               :headers="headers" 
               :items="tradeHistory"
               :items-per-page="5"
            >
            <template v-slot:item="{ item }">
               <tr>
                  <td>
                     <i :class="item.currencyIcon"></i>
                  </td>
                  <td>{{ item.price }}</td>
                  <td :class="item.statusClass">{{ item.status }}</td>
                  <td>{{ item.price }}</td>
                  <td>{{ item.total }}</td>
                  <td>{{ item.date }}</td>
						<td>{{ item.from }}</td>
						<td>
							<div class="d-flex align-items-center">
								<v-dialog max-width="500px">
									<template v-slot:activator="{ on }">
										<v-btn text icon color="grey" small v-on="on" class="my-3 mx-2">
											<i class="material-icons">info</i>
										</v-btn>
									</template>
									<v-card>
										<v-card-text>
											<v-container class="grid-list-md">
												<v-row>
													<v-col cols="12">
														<h5>Transaction Details</h5>
														<p>Transaction with Txn No 76237523 was placed on 21/3/102, 11:10 with Payment done being </p>
														<p>Successful.</p>
														<table>
															<tr>
																<td>
																	<p>Currency :</p>
																</td>
																<td>
																	<div><i :class="item.currencyIcon"></i></div>
																</td>
															</tr>

															<tr>
																<td>
																	<p>Payment Currency :</p>
																</td>
																<td>
																	<p>{{ item.payment_currency }}</p>
																</td>
															</tr>

															<tr>
																<td>
																	<p>Price :</p>
																</td>
																<td>
																<p>{{ item.price }}</p>
																</td>
															</tr>

															<tr>
																<td>
																	<p>Total :</p>
																</td>
																<td>
																<p>{{ item.total }}</p>
																</td>
															</tr>

															<tr>
																<td>
																	<p>To/From :</p>
																</td>
																<td>
																<p>{{ item.from }}</p>
																</td>
															</tr>

															<tr>
																<p>Wallet Address :</p>
																<v-text-field
																	outline
																	solo
																	label="Wallet"
																	value="AXB35H24ISDJHCISDT"
																	:prepend-inner-icon="item.currencyIcon"
																></v-text-field>
															</tr>
														</table>                                      
													</v-col>
												</v-row>
											</v-container>
										</v-card-text>
									</v-card>
								</v-dialog>
							</div>
						</td>
               </tr>
            </template>
            </v-data-table>
         </div>
      </v-card>
   </div>
</template>

<script>
import { tradeHistory } from 'Views/crypto/data.js'

export default {
	data () {
		return {
         tradeHistory,
			loader: false,
         invoice: [],
			headers: [
				{
					text: "Currency",
					sortable: false,
					value: "currency"
				},
				{
					text: "Txn No",
					sortable: false,
					value: "txn No"
				},
				{
					text: "Status",
					sortable: false,
					value: "status"
				},
				{
					text: "Price",
					sortable: false,
					value: "price"
				},
				{
					text: "Total($)",
					sortable: false,
					value: "total($)"
				},
				{
					text: "Date",
					sortable: false,
					value: "date"
				},
				{
					text: "To / From",
					sortable: false,
					value: "to / from"
				},
				{
					text: "More",
					sortable: false,
					value: "more"
				}
        ],
			active: null,
		}
   },
	methods: {
		next () {
			const active = parseInt(this.active)
			this.active = (active < 2 ? active + 1 : 0)
      },
	}
}
</script>