<template>
	<div class="promo-widget">
		<app-card customClasses="border-primary">
			<v-row>
				<v-col xl="3" lg="3" md="2" sm="3" cols="4" b-50>
					<div class="d-flex justify-center align-center h-100">
						<img class="img-responsive" alt="comment" src="/static/img/Comments.png" />
					</div>
				</v-col>
				<v-col xl="9" lg="9" md="10" sm="9" cols="8" b-50>
					<div>
						<h4>24x7 Customer Support</h4>
						<p class="fs-12 fw-normal grey--text">Far far away, behind the word mountains, far from the countries
							Vokalia and Consonantia. Even the all-powerful Pointing has no control about the blind texts.Far
							far away, behind the word mountains, far from the countries Vokalia and Consonantia.</p>
						<v-btn color="primary">{{$t('message.letsGetInTouch')}}</v-btn>
					</div>
				</v-col>
			</v-row>
		</app-card>
	</div>
</template>