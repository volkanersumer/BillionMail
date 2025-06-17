import { instance } from '@/api'

export const getApiList = (params: {
	page: number
	page_size: number
	keyword: string
	active: number
	start_time: number
	end_time: number
}) => {
	return instance.get('/batch_mail/api/list', { params })
}

export const getOverviewStats = (params: { start_time: number; end_time: number }) => {
	return instance.get('/batch_mail/api/overview_stats', { params })
}

type APIParams = {
	api_name: string
	template_id: number
	subject: string
	addresser: string
	full_name: string
	unsubscribe: number
	active: number
}

export const createApi = (params: APIParams) => {
	return instance.post('/batch_mail/api/create', params, {
		fetchOptions: {
			loading: '正在创建，请稍候...',
			successMessage: true,
		},
	})
}

export const updateApi = (params: APIParams & { id: number }) => {
	return instance.post('/batch_mail/api/update', params, {
		fetchOptions: {
			loading: '正在编辑，请稍候...',
			successMessage: true,
		},
	})
}

export const deleteApi = (params: { id: number }) => {
	return instance.post('/batch_mail/api/delete', params, {
		fetchOptions: {
			loading: '正在删除，请稍候...',
			successMessage: true,
		},
	})
}

export const testApi = (key: string, params: { recipient: string }) => {
	return instance.post('/batch_mail/api/send', params, {
		headers: {
			'x-api-key': key,
		},
		fetchOptions: {
			loading: '正在测试，请稍候...',
			successMessage: true,
		},
	})
}
