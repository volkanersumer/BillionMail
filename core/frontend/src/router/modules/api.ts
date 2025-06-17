import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/send',
	component: Layout,
	meta: { sort: 3, key: 'api', title: 'Send API', titleKey: '' },
	children: [
		{
			path: '/send',
			name: 'Api',
			component: () => import('@/views/api/index.vue'),
		},
	],
}

export default route
