<template>
	<section class="h-full text_sans">
		<div class="max-w-screen-2xl flex flex-col items-start mx-auto gap-4">
			<p class="uppercase hero-section__overline">
				I create
				{{ randomOverlineAdjective }} {{ randomOverlineCreation }} with
			</p>
			<h1 class="hero-section__headline">
				<span class="font-mono">Tech</span>
				<span> & </span>
				<span class="font-serif">Art</span>
			</h1>
		</div>
	</section>
</template>

<script setup lang="ts">
import { getRandomElement } from '@/functions/array'

const randomOverlineAdjective = getRandomElement([
	'immersive',
	'interactive',
	// 'impressive',
	'fun',
])

const randomOverlineCreation = getRandomElement([
	'Experiences',
	'Tools',
	// 'Stuff',
	'Things',
])
</script>

<style scoped lang="scss">
// TODO: Use typography from the design system
.hero-section__overline {
	font-weight: 600;
	font-size: 2.7rem;
}

.hero-section__headline {
	font-weight: 600;
	font-size: 14.5rem;
	line-height: 15rem;

	.font-serif {
		font-weight: 700;
	}
}
</style>
