<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
	<header class="wrapper">
		<div class="content">
			<RouterLink :to="{ name: 'Home' }" id="logo">CreatorSiSo.xyz</RouterLink>
			<nav class="navigation">
				<RouterLink :to="{ name: 'Home' }">Home</RouterLink>
				<RouterLink :to="{ name: 'Projects' }">Projects</RouterLink>
				<RouterLink :to="{ name: 'About' }">About</RouterLink>
				<RouterLink :to="{ name: 'Blog' }">Blog</RouterLink>
			</nav>
		</div>
	</header>
</template>

<style scoped lang="scss">
@use '@/assets/styles/variables.scss';

.wrapper {
	padding: 2.4rem 3rem;
	position: sticky;
	top: 0;

	.content {
		display: flex;

		align-items: center;
		justify-content: space-between;
		margin-inline: auto;
		max-width: variables.$content-max-width;

		#logo {
			font-size: 2.4rem;
			font-weight: 800;
		}

		.navigation {
			display: flex;
			gap: 3rem;
			font-size: 1.8rem;
			font-weight: 700;
		}
	}
}

.page-navigation {
	// TODO: Add hamburger menu for mobile devices
	display: none;

	@media (min-width: 500px) {
		display: flex;
		gap: 3rem;
	}
}
</style>
