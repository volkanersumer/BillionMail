<template>
	<n-button text @click="handleGoHelp">
		<i class="i-custom:help"></i>
		<span class="ml-4px">{{ $t('common.actions.help') }}</span>
	</n-button>
</template>

<script lang="ts" setup>
const { href } = defineProps({
	href: {
		type: String,
		default: '',
	},
})

const handleGoHelp = () => {
	window.open(href)
}
</script>

<style lang="scss" scoped></style>
