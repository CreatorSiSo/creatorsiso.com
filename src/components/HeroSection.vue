<template>
	<section class="wrapper text_sans">
		<div class="content">
			<p class="text_upper overline">
				I create
				{{ randomOverlineAdjective }} {{ randomOverlineCreation }} with
			</p>
			<h1 class="headline">
				<span class="text_mono">Tech</span>
				<span> & </span>
				<span class="text_serif">Art</span>
			</h1>
		</div>
	</section>
</template>

<script setup lang="ts">
import { getRandomElement } from '@/functions/array.js'

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
@use '@/assets/styles/variables.scss';

.wrapper {
	height: 100%;

	.content {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		margin-inline: auto;
		gap: 1rem;
		max-width: variables.$content-max-width;

		// TODO: Use typography from the design system
		.overline {
			font-weight: 600;
			font-size: 2.7rem;
		}

		.headline {
			font-weight: 600;
			font-size: 14.5rem;
			line-height: 15rem;

			.text_serif {
				font-weight: 700;
			}
		}
	}
}
</style>
