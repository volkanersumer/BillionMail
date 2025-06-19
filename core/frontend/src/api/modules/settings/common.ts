import { instance } from '@/api'

export const getSystemConfig = () => {
	return instance.get('/settings/get_system_config')
}

export const getTimezoneList = () => {
	return instance.get('/settings/get_timezone_list')
}

export const setSystemConfigKey = (params: { key: string; value: string }) => {
	return instance.post('/settings/set_system_config_key', params, {
		fetchOptions: {
			loading: 'Setting the system configuration, please wait...',
			successMessage: true,
		},
	})
}

export const setSslConfig = (params: { certPem: string; privateKey: string }) => {
	return instance.post('/settings/set_ssl_config', params, {
		fetchOptions: {
			loading: 'Setting SSL, please wait...',
			successMessage: true,
		},
	})
}

export const applyCert = () => {
	return instance.post(
		'/ssl/console_apply_cert',
		{},
		{
			fetchOptions: {
				loading: 'Applying certificate, please wait...',
				successMessage: true,
			},
		}
	)
}
