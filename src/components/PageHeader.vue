<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
	<header class="px-12 py-[2.4rem] sticky top-0">
		<div class="max-w-screen-2xl mx-auto flex items-center justify-between">
			<RouterLink :to="{ name: 'Home' }" id="logo">CreatorSiSo.xyz</RouterLink>

			<!-- TODO: Add hamburger menu for mobile devices -->
			<nav id="navigation" class="flex gap-12">
				<RouterLink :to="{ name: 'Home' }">Home</RouterLink>
				<RouterLink :to="{ name: 'Projects' }">Projects</RouterLink>
				<RouterLink :to="{ name: 'About' }">About</RouterLink>
				<RouterLink :to="{ name: 'Blog' }">Blog</RouterLink>
			</nav>
		</div>
	</header>
</template>

<style scoped lang="scss">
#logo {
	font-size: 2.4rem;
	font-weight: 800;
}

#navigation {
	font-size: 1.8rem;
	font-weight: 700;
}
</style>
