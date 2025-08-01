<template>
	<div>
		<div class="mb-16px">
			<n-button type="primary" @click="getLogs">{{ $t('common.actions.refresh') }}</n-button>
		</div>
		<n-spin :show="loading">
			<bt-logs :code="content" height="640px" :scroll-to-bottom="true"></bt-logs>
		</n-spin>
	</div>
</template>

<script lang="ts" setup>
import { isString } from '@/utils'
import { getErrorLogs } from './controller'

const content = ref('')

const loading = ref(false)

const getLogs = async () => {
	try {
		loading.value = true
		const res = await getErrorLogs()
		content.value = isString(res) ? res : ''
	} finally {
		loading.value = false
	}
}

getLogs()
</script>

<style lang="scss" scoped></style>
