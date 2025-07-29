<template>
	<bt-modal
		v-model:show="show"
		:width="520"
		:title="$t('overview.sendQueue.buttons.paramConfig')"
		@confirm="onConfirm">
		<bt-form class="pt-8px">
			<n-form-item label="Bounce Queue Lifetime">
				<n-input v-model:value="form.bounce_queue_lifetime" />
			</n-form-item>
			<n-form-item label="Maximal Queue Lifetime">
				<n-input v-model:value="form.maximal_queue_lifetime" />
			</n-form-item>
			<n-form-item label="Minimal Backoff Time">
				<n-input v-model:value="form.minimal_backoff_time" />
			</n-form-item>
			<n-form-item label="Maximal Backoff Time">
				<n-input v-model:value="form.maximal_backoff_time" />
			</n-form-item>
			<n-form-item label="Queue Run Delay">
				<n-input v-model:value="form.queue_run_delay" />
			</n-form-item>
			<n-form-item label="Trigger Timeout">
				<n-input v-model:value="form.trigger_timeout" />
			</n-form-item>
		</bt-form>
	</bt-modal>
</template>

<script lang="ts" setup>
import { isObject } from '@/utils'
import { getSendQueueConfig, setSendQueueConfig } from '@/api/modules/overview'

const show = ref(false)

const form = reactive({
	bounce_queue_lifetime: '',
	maximal_queue_lifetime: '',
	minimal_backoff_time: '',
	maximal_backoff_time: '',
	queue_run_delay: '',
	trigger_timeout: '',
})

const getConfig = async () => {
	const res = await getSendQueueConfig()
	if (isObject(res)) {
		Object.assign(form, res)
	}
}

const onConfirm = async () => {
	await setSendQueueConfig(toRaw(form))
}

const open = () => {
	show.value = true
	getConfig()
}

defineExpose({
	open,
})
</script>

<style lang="scss" scoped></style>
