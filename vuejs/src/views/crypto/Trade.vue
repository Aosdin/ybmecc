<template>
   <div>
		<page-title-bar></page-title-bar>
      <v-container fluid class="grid-list-xl pt-0">
         <crypto-slider></crypto-slider>
         <v-row>
				<!-- Wallet Address -->
            <app-card
               :heading="$t('message.walletAddress')"
               colClasses="col-xl-4 col-lg-4 col-md-6 col-sm-6 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <wallet-address></wallet-address>
            </app-card>
				<!-- Buy Crypto Currency -->
            <app-card
                  :heading="$t('message.buyCryptocurrency')"
                  colClasses="col-xl-4 col-lg-4 col-md-6 col-sm-6 col-12"
                  customClasses="mb-0"
                  :fullScreen="true"
                  :reloadable="true"
                  :closeable="true"
               >
               <buy-cryptocurrency></buy-cryptocurrency>
            </app-card>
				<!-- Expenditure Stats -->
            <app-card
                  :heading="$t('message.expenditureStats')"
                  colClasses="col-xl-4 col-lg-4 col-md-12 col-sm-12 col-12"
                  customClasses="mb-0"
                  :fullScreen="true"
                  :reloadable="true"
                  :closeable="true"
               >
               <expenditure-stats></expenditure-stats>
            </app-card>
         </v-row>
         <v-row>
				<!-- Sell Crypto Currency -->
            <app-card
               :heading="$t('message.sellCryptocurrency')"
               colClasses="col-xl-4 col-lg-4 col-md-6 col-sm-6 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <sell-cryptocurrency></sell-cryptocurrency>
            </app-card>
				<!-- Transfer Crypto Currency -->
            <app-card
               :heading="$t('message.transferCryptocurrency')"
               colClasses="col-xl-4 col-lg-4 col-md-6 col-sm-6 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <transfer-cryptocurrency></transfer-cryptocurrency>
            </app-card>
				<!-- Exchange Rate -->
            <app-card
               :heading="$t('message.exchangeRate')"
               colClasses="col-xl-4 col-lg-4 col-md-12 col-sm-12 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <exchange-rate></exchange-rate>
            </app-card>
         </v-row>
         <v-row>
				<!-- Buy Or Sell -->
            <app-card
               :heading="$t('message.buyOrSell')"
               colClasses="col-xl-8 col-lg-8 col-md-12 col-sm-12 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <buy-or-sell></buy-or-sell>
            </app-card>
				<!-- Recent Trades -->
            <app-card
               :heading="$t('message.recentTrades')"
               colClasses="col-xl-4 col-lg-4 col-md-12 col-sm-12 col-12"
               customClasses="mb-0" 
               :fullScreen="true"
               :reloadable="true"
               :closeable="true"
            >
               <recent-trades></recent-trades>
            </app-card>
         </v-row>
      </v-container>
   </div>
</template>

<script>
import WalletAddress from "Components/Widgets/WalletAddress";
import BuyCryptocurrency from "Components/Widgets/BuyCryptocurrency";
import SellCryptocurrency from "Components/Widgets/SellCryptocurrency";
import TransferCryptocurrency from "Components/Widgets/TransferCryptocurrency";
import ExchangeRate from "Components/Widgets/ExchangeRate";
import RecentTrades from "Components/Widgets/RecentTrades";
import BuyOrSell from "Components/Widgets/BuyOrSell";
import ExpenditureStats from "Components/Charts/PieChartWithLegend";

export default {
   components:{
      WalletAddress,
      BuyCryptocurrency,
      SellCryptocurrency,
      TransferCryptocurrency,
      ExchangeRate,
      RecentTrades,
      BuyOrSell,
      ExpenditureStats
   }
}
</script>