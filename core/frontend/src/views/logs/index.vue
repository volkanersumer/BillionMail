<template>
	<div class="p-24px">
		<bt-table-layout>
			<template #toolsLeft>
				<div class="w-240px">
					<n-select :filterable="true" :options="typeOptions"></n-select>
				</div>
				<n-button type="primary">{{ $t('common.actions.refresh') }}</n-button>
			</template>
			<template #table>
				<n-data-table v-bind="tableProps" :columns="columns"></n-data-table>
			</template>
			<template #pageRight>
				<bt-table-page v-bind="pageProps" @refresh="fetchTable"> </bt-table-page>
			</template>
		</bt-table-layout>
	</div>
</template>

<script lang="ts" setup>
import { DataTableColumns, SelectOption } from 'naive-ui'
import { formatTime, isObject } from '@/utils'
import { useDataTable } from '@/hooks/useDataTable'
import { getLogsList, getLogsType } from './controller'
import type { Logs } from './types'

const columns = ref<DataTableColumns<Logs>>([
	{
		key: 'username',
		title: 'Username',
	},
	{
		key: 'type',
		title: 'Type',
	},
	{
		key: 'log',
		title: 'Details',
	},
	{
		key: 'addtime',
		title: 'Operating time',
		render: row => {
			return formatTime(row.addtime)
		},
	},
])

const { tableProps, pageProps, fetchTable } = useDataTable<Logs>({
	params: {
		page: 1,
		page_size: 10,
	},
	immediate: true,
	fetchFn: getLogsList,
})

const typeOptions = ref<SelectOption[]>([])

const getTypeOptions = async () => {
	const res = await getLogsType()
	if (isObject<Record<string, string>>(res)) {
		typeOptions.value = Object.entries(res).map(([, value]) => ({
			label: value,
			value: value,
		}))
	}
}

getTypeOptions()
</script>

<style lang="scss" scoped></style>
