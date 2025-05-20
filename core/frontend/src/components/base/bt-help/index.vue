<template>
	<n-button text @click="handleGoHelp">
		<i class="i-custom:help"></i>
		<span class="ml-4px">{{ $t('common.actions.help') }}</span>
	</n-button>
</template>

<script lang="ts" setup>
const handleGoHelp = () => {
	window.open('https://github.com/aaPanel/BillionMail')
}
</script>

<style lang="scss" scoped></style>
