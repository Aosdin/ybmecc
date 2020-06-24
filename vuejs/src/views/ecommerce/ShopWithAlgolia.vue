<template>
	<div class="shop-wrapper">
		<page-title-bar></page-title-bar>
		<v-container fluid class="px-4 py-0 grid-list-xl mt-n3">
			<ais-instant-search :search-client="searchClient" index-name="ikea">
				<v-row>
					<v-col xl="9" lg="9" md="9" sm="12" cols="12">
						<product-items></product-items>
					</v-col>
					<v-col xl="3" lg="3" md="3" sm="0" xs="0" class="shop-filter">
						<sidebar-filters></sidebar-filters>
					</v-col>
				</v-row>
			</ais-instant-search>
		</v-container>
	</div>
</template>
<script>
	import ProductItems from "Components/Shop/ProductItems";
	import SidebarFilters from "Components/Shop/SidebarFilters";

	import algoliasearch from 'algoliasearch/lite';
	import 'instantsearch.css/themes/algolia-min.css';

	const searchClient = algoliasearch('latency', '6be0576ff61c053d5f9a3225e2a90f76');

	export default {
		components: {
			ProductItems,
			SidebarFilters
		},
		data() {
			return {
				searchClient,
			};
		},
		computed: {
			ecommerceSidebarFilter: {
				get() {
					return this.$store.getters.ecommerceSidebarFilter;
				},
				set(val) {
					this.$store.dispatch("toggleEcommerceSidebarFilter", val);
				}
			}
		}
	};
</script>