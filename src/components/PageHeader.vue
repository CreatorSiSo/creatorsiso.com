<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
	<header class="page-header">
		<div class="page-header__content">
			<RouterLink :to="{ name: 'Home' }">CreatorSiSo.xyz</RouterLink>
			<nav class="page-navigation">
				<RouterLink :to="{ name: 'Home' }">Home</RouterLink>
				<RouterLink :to="{ name: 'Projects' }">Projects</RouterLink>
				<RouterLink :to="{ name: 'About' }">About</RouterLink>
				<RouterLink :to="{ name: 'Blog' }">Blog</RouterLink>
			</nav>
		</div>
	</header>
</template>

<style scoped lang="scss">
.page-header {
	padding: 2.4rem 3rem;
	position: sticky;
	top: 0;

	.page-header__content {
		display: flex;
		justify-content: space-between;
		max-width: 1500px;
		margin-inline: auto;
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
