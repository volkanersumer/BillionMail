import { RouteRecordRaw } from 'vue-router'
import { isDev } from '@/utils'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/logs',
	component: Layout,
	meta: { key: 'logs', title: 'Logs', hidden: !isDev, titleKey: '' },
	children: [
		{
			path: '/logs',
			name: 'Logs',
			component: () => import('@/views/logs/index.vue'),
		},
	],
}

export default route
