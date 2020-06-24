<template>
	<div>
		<page-title-bar></page-title-bar>
		<v-container fluid class="grid-list-xl pt-0 mt-n3">
			<v-row>
				<app-card 
					:heading="$t('message.activator')" 
					colClasses="col-12 col-sm-4 col-md-4"
				>
					<v-menu offset-y>
						<template v-slot:activator="{ on }">
							<v-btn color="primary" dark v-on="on">Dropdown</v-btn>
						</template>
						<v-list>
							<v-list-item v-for="item in list1" :key="item.title" >
								<v-list-item-title>{{ item.title }}</v-list-item-title>
							</v-list-item>
						</v-list>
					</v-menu>
				</app-card>
				<app-card 
					:heading="$t('message.hover')" 
					colClasses="col-12 col-sm-4 col-md-4"
				>
					<v-menu open-on-hover top offset-y>
						<template v-slot:activator="{ on }">
						<v-btn color="error" dark v-on="on">Dropdown</v-btn>
						</template>
						<v-list>
							<v-list-item v-for="item in list1" :key="item.title" >
								<v-list-item-title>{{ item.title }}</v-list-item-title>
							</v-list-item>
						</v-list>
					</v-menu>
				</app-card>
				<app-card 
					:heading="$t('message.menus')" 
					colClasses="col-12 col-sm-4 col-md-4"
				>
					<v-menu bottom right>
						<template v-slot:activator="{ on }">
							<v-btn icon v-on="on">
								<v-icon>more_vert</v-icon>
							</v-btn>
						</template>
						<v-list>
							<v-list-item v-for="item in list1" :key="item.title" >
								<v-list-item-title>{{ item.title }}</v-list-item-title>
							</v-list-item>
						</v-list>
					</v-menu>
				</app-card>
			</v-row>
		</v-container>
	</div>
</template>

<script>
export default {
  data: () => ({
    list1: [
      { title: "Refresh" },
      { title: "Settings" },
      { title: "Help" },
      { title: "Sign Out" }
    ]
  })
};
</script>
