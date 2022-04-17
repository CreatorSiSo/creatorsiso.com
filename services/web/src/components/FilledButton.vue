<template>
	<button class="button rounded text-white px-5 py-[9px] highlight">
		<slot>Button</slot>
	</button>
</template>

<script setup lang="ts"></script>

<style scoped>
.button {
	background: linear-gradient(to top left, #0f6799 0%, #4bafe7 100%);
}
</style>
