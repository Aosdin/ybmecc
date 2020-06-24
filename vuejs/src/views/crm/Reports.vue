<template>
	<div>
      <page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0">
         <v-row>
            <app-card
               :heading="$t('message.invoices')"
					colClasses="col-xl-6 col-lg-6 col-md-12 col-sm-12 col-12 pt-0"
					customClasses="mb-0"
				>
               <invoice></invoice>
            </app-card>
            
            <app-card
               :heading="$t('message.payments')"
					colClasses="col-xl-6 col-lg-6 col-md-12 col-sm-12 col-12 pt-0"
					customClasses="mb-0"
				>
               <payments></payments>
            </app-card>
			</v-row>

         <v-row>
            <app-card
					colClasses="col-xl-12 col-lg-12 col-md-12 col-sm-12 col-12"
					customClasses="mb-0"
				>
            <TabsAndTable></TabsAndTable>
            </app-card>
			</v-row>

         <v-row>
            <app-card
               :heading="$t('message.taxRates')"
					colClasses="col-xl-6 col-lg-6 col-md-12 col-sm-12 col-12"
					customClasses="mb-0"
				>
               <TaxRates></TaxRates>
            </app-card>

            <app-card
               :heading="$t('message.addTickets')"
					colClasses="col-xl-6 col-lg-6 col-md-12 col-sm-12 col-12"
					customClasses="mb-0"
				>
               <add-tickets></add-tickets>
            </app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
import TabsAndTable from "Components/Widgets/TabsAndTable";
import Invoice from "Components/Widgets/Invoice";
import Payments from "Components/Widgets/Payments";
import TaxRates from "Components/Widgets/TaxRates";
import AddTickets from "Components/Widgets/AddTickets";

export default {
  components: {
     TabsAndTable,
     Invoice,
     Payments,
     TaxRates,
     AddTickets
  }
};
</script>
