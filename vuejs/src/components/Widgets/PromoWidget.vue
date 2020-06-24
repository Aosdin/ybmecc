<template>
   <div class="promo-widget">
		<app-card customClasses="border-info">
			<v-row>
				<v-col xl="3" lg="3" md="2" sm="3" cols="4" b-50>
					<div class="d-flex justify-center align-center h-100">
						<img class="img-responsive" alt="transform" src="/static/img/Freetransform.png" />
					</div>
				</v-col>
				<v-col xl="9" lg="9" md="10" sm="9" cols="8" b-50>
					<div>
						<h4>Pixel Perfect Design</h4>
						<p class="fs-12 grey--text">Far far away, behind the word mountains, far from the countries Vokalia and Consonantia. Even the all-powerful Pointing has no control about the blind texts.Far far away, behind the word mountains, far from the countries Vokalia and Consonantia.</p>
						<v-btn color="info" class="fs-12 fw-normal">{{$t('message.purchaseVuely')}}</v-btn>
					</div>
				</v-col>
			</v-row>
		</app-card>
	</div>
</template>
