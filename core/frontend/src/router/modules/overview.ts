import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/overview',
	component: Layout,
	meta: { sort: 0, key: 'overview', title: 'Overview', titleKey: 'layout.menu.overview' },
	children: [
		{
			path: '/overview',
			name: 'Overview',
			component: () => import('@/views/overview/index.vue'),
		},
		{
			path: '/overview/send-queue',
			name: 'SendQueue',
			component: () => import('@/views/overview/pages/send-queue/index.vue'),
		},
	],
}

export default route
