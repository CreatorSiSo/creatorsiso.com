<template>
	<HeroSection id="section1"></HeroSection>
</template>

<script setup lang="ts">
import HeroSection from '@/components/HeroSection.vue'
</script>

<style lang="scss">
#section1 {
	height: calc(100vh - 70px);
	box-shadow: inset 0 0 0 1px black;
}
</style>
