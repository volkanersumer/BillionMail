import { RouteRecordRaw } from 'vue-router'
import { isDev } from '@/utils'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/send',
	component: Layout,
	meta: { sort: 3, hidden: !isDev, key: 'api', title: 'Send API', titleKey: '' },
	children: [
		{
			path: '/send',
			name: 'Api',
			component: () => import('@/views/api/index.vue'),
		},
	],
}

export default route
