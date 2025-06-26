<template>
	<iframe class="w-full" :src="src"></iframe>
</template>

<script lang="ts" setup>
import { isDev, apiUrl } from '@/utils'

const src = ref(`${isDev ? apiUrl : window.location.origin}/rspamd/`)
</script>

<style lang="scss" scoped>
iframe {
	height: calc(100% - 2px);
}
</style>
