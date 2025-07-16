import { instance } from '@/api'
import type { SuspendParams } from '@/views/contacts/suspend/types/base'

export const getSuspendList = (params: SuspendParams) => {
	return instance.get('/abnormal_recipient/list', { params })
}

export const deleteSuspend = (params: { id: number }) => {
	return instance.post('/abnormal_recipient/delete', params, {
		fetchOptions: {
			loading: 'Deleting, please wait...',
			successMessage: true,
		},
	})
}
