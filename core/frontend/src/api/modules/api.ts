import { instance } from '@/api'

export const getApiList = (params: {
	page: number
	page_size: number
	keyword: string
	active: number
}) => {
	return instance.get('/batch_mail/api/list', { params })
}
