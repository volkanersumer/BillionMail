import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/rspamd_admin',
	component: Layout,
	meta: { sort: 3, key: 'rspamd', title: 'Rspamd', titleKey: '' },
	children: [
		{
			path: '/rspamd_admin',
			name: 'Rspamd',
			component: () => import('@/views/rspamd/index.vue'),
		},
	],
}

export default route
