import '@/assets/styles/base.scss'

import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'

import App from './App.vue'
import Home from './views/Home.vue'

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{
			path: '/',
			name: 'Home',
			component: Home,
		},
		{
			path: '/about',
			name: 'About',
			component: () => import('@/views/About.vue'),
		},
		{
			path: '/projects',
			name: 'Projects',
			component: () => import('@/views/Projects.vue'),
		},
		{
			path: '/blog',
			name: 'Blog',
			component: () => import('@/views/Blog.vue'),
		},
		{
			path: '/:pathMatch(.*)*',
			name: 'NotFound',
			component: () => import('@/views/NotFound.vue'),
		},
		{
			path: '/s',
			name: 'DesignSystem',
			component: () => import('@/views/DesignSystem.vue'),
		},
	],
})

createApp(App).use(router).mount('#app')
