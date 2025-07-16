<template>
	<n-spin class="h-full" :show="loading">
		<highlightjs ref="logsRef" class="bt-log" language="log" :code="log"> </highlightjs>
	</n-spin>
</template>

<script lang="ts" setup>
interface Props {
	log: string
	loading?: boolean
}

const props = withDefaults(defineProps<Props>(), {
	log: '',
	loading: false,
})

const logsRef = ref()

watch(
	() => props.log,
	() => {
		nextTick(() => {
			const dom = logsRef.value.$el as HTMLElement
			const codeDom = dom.querySelector('code')
			codeDom?.scrollTo({ top: 9999999 })
		})
	},
	{
		immediate: true,
	}
)
</script>

<style lang="scss" scoped>
.bt-log {
	height: 100%;
	font-size: 13px;
	overflow: auto;
	background: none;
	border: none;
	border-radius: 0;

	:deep(.hljs) {
		height: 100%;
		border-radius: 4px;
		white-space: pre-wrap;
	}
}
</style>
