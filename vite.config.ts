import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
	resolve: {
		alias: {
			'@/': new URL('./src/', import.meta.url).pathname,
		},
	},

	plugins: [vue()],

	build: {
		target: 'es2015',
		minify: false,
		outDir: 'dist',
	},
})
