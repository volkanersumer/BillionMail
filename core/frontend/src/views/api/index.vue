<template>
	<div class="p-24px">
		<div class="flex items-center justify-between mb-24px">
			<div class="bt-title mb-0px!">Send API</div>
			<div class="flex gap-12px">
				<n-button type="primary" ghost>刷新数据</n-button>
				<n-button type="primary">新建API</n-button>
			</div>
		</div>
		<n-card class="mb-24px" :content-style="{ padding: '16px 24px' }">
			<div class="flex flex-wrap items-center gap-16px">
				<div class="flex items-center">
					<span class="text-default">时间范围：</span>
					<bt-time></bt-time>
				</div>
				<div class="flex items-center">
					<span class="text-default">状态：</span>
					<div class="w-120px">
						<n-select v-model:value="tableParams.active" :options="activeOptions"></n-select>
					</div>
				</div>
				<div class="flex items-center">
					<span class="text-default">搜索：</span>
					<div class="w-200px">
						<n-input v-model:value="tableParams.keyword" placeholder="请输入API名称"></n-input>
					</div>
				</div>
				<div>
					<n-button type="primary" ghost>搜索</n-button>
				</div>
			</div>
		</n-card>
		<api-overview></api-overview>
		<bt-table-layout>
			<template #table>
				<n-data-table v-bind="tableConfig" :columns="columns"></n-data-table>
			</template>
			<template #pageRight>
				<bt-table-page v-bind="pageConfig" @refresh="refreshTable"> </bt-table-page>
			</template>
		</bt-table-layout>
	</div>
</template>

<script lang="ts" setup>
import { DataTableColumns } from 'naive-ui'
import { useDataTable } from '@/hooks/useDataTable'
import { getApiList } from '@/api/modules/api'
import type { Api } from './types/base'

import ApiOverview from './components/Overview.vue'

const activeOptions = [
	{ label: '全部', value: -1 },
	{ label: '可用', value: 1 },
	{ label: '停用', value: 0 },
]

const columns = ref<DataTableColumns<Api>>([
	{
		key: 'api_name',
		title: 'API名称',
	},
	{
		key: 'send_count',
		title: '发送量',
	},
	{
		key: 'open_rate',
		title: '打开率',
		render: row => {
			return `${row.open_rate}%`
		},
	},
	{
		key: 'click_rate',
		title: '点击率',
		render: row => {
			return `${row.click_rate}%`
		},
	},
	{
		key: 'bounce_rate',
		title: '退信率',
		render: row => {
			return `${row.bounce_rate}%`
		},
	},
])

const params = {
	page: 1,
	page_size: 10,
	keyword: '',
	active: -1,
}

const { tableParams, tableConfig, pageConfig, refreshTable } = useDataTable<Api, typeof params>({
	params,
	fetchFn: getApiList,
})
</script>

<style lang="scss" scoped></style>
