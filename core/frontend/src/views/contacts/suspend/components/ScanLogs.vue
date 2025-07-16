<template>
	<modal title="Scan Logs" width="720" :footer="false">
		<div class="pt-12px mb-8px">
			<n-button class="mb-8px" type="primary">
				{{ $t('common.actions.refresh') }}
			</n-button>
		</div>
		<div class="h-400px">
			<bt-logs :loading="loading" :log="content"></bt-logs>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { useModal } from '@/hooks/modal/useModal'

const loading = ref(false)

const content = ref(`None`)

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {},
})
</script>

<style lang="scss" scoped></style>
