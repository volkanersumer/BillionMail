<template>
	<div class="p-24px">
		<div class="bt-title">Send API</div>
		<n-card class="mb-24px" :content-style="{ padding: '16px 24px' }">
			<div class="flex flex-wrap items-center gap-16px">
				<div class="flex items-center">
					<span class="text-default">时间范围：</span>
					<bt-time
						v-model:value="tableParams.time_range"
						default-type="last7days"
						@change="handleRefreshData">
					</bt-time>
				</div>
				<div class="flex items-center">
					<span class="text-default">状态：</span>
					<div class="w-120px">
						<n-select
							v-model:value="tableParams.active"
							:options="activeOptions"
							@update:value="refreshTable">
						</n-select>
					</div>
				</div>
				<div class="flex items-center">
					<span class="text-default">搜索：</span>
					<div class="w-200px">
						<n-input
							v-model:value="tableParams.keyword"
							placeholder="请输入API名称"
							@keyup.enter="refreshTable">
						</n-input>
					</div>
				</div>
				<div class="flex gap-8px">
					<n-button type="primary" ghost @click="refreshTable">
						{{ $t('common.actions.refresh') }}
					</n-button>
					<n-button type="primary" ghost @click="handleRefreshData">刷新数据</n-button>
				</div>
			</div>
		</n-card>
		<api-overview ref="overviewRef" :time="tableParams.time_range"></api-overview>
		<bt-table-layout>
			<template #toolsLeft>
				<n-button type="primary" @click="handleAdd">新建API</n-button>
			</template>
			<template #table>
				<n-data-table v-bind="tableConfig" :columns="columns"></n-data-table>
			</template>
			<template #pageRight>
				<bt-table-page v-bind="pageConfig" @refresh="refreshTable"> </bt-table-page>
			</template>
			<template #modal>
				<form-modal></form-modal>
				<test-modal></test-modal>
			</template>
		</bt-table-layout>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NButton, NFlex, NTag } from 'naive-ui'
import { confirm } from '@/utils'
import { useDataTable } from '@/hooks/useDataTable'
import { useModal } from '@/hooks/modal/useModal'
import { deleteApi, getApiList } from '@/api/modules/api'
import type { Api } from './types/base'

import ApiOverview from './components/Overview.vue'
import ApiKey from './components/ApiKey.vue'
import ApiForm from './components/ApiForm.vue'
import ApiTest from './components/ApiTest.vue'

const { t } = useI18n()

const overviewRef = useTemplateRef('overviewRef')

const activeOptions = [
	{ label: '全部', value: -1 },
	{ label: '可用', value: 1 },
	{ label: '停用', value: 0 },
]

const columns = ref<DataTableColumns<Api>>([
	{
		key: 'api_name',
		title: 'API名称',
		maxWidth: 120,
		render: row => (
			<NButton
				type="primary"
				text
				onClick={() => {
					handleEdit(row)
				}}>
				{row.api_name}
			</NButton>
		),
	},
	{
		key: 'api_key',
		title: 'API密钥',
		width: '28%',
		maxWidth: 160,
		render: row => <ApiKey value={row.api_key}></ApiKey>,
	},
	{
		key: 'send_count',
		title: '发送量',
		width: '9%',
		maxWidth: 100,
	},
	{
		key: 'open_rate',
		title: '打开率',
		width: '9%',
		maxWidth: 100,
		render: row => {
			return `${row.open_rate}%`
		},
	},
	{
		key: 'click_rate',
		title: '点击率',
		width: '9%',
		maxWidth: 100,
		render: row => {
			return `${row.click_rate}%`
		},
	},
	{
		key: 'bounce_rate',
		title: '退信率',
		width: '9%',
		maxWidth: 100,
		render: row => {
			return `${row.bounce_rate}%`
		},
	},
	{
		key: 'active',
		title: '状态',
		width: '8%',
		maxWidth: 80,
		render: row => (
			<NTag size="small" type={row.active ? 'primary' : 'warning'} bordered={false}>
				{row.active ? '可用' : '停用'}
			</NTag>
		),
	},
	{
		title: t('common.columns.actions'),
		key: 'actions',
		align: 'right',
		width: 160,
		render: row => (
			<NFlex inline={true}>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleTest(row)
					}}>
					{t('common.actions.test')}
				</NButton>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleEdit(row)
					}}>
					{t('common.actions.edit')}
				</NButton>
				<NButton
					type="error"
					text={true}
					onClick={() => {
						handleDelete(row)
					}}>
					{t('common.actions.delete')}
				</NButton>
			</NFlex>
		),
	},
])

const { tableParams, tableConfig, pageConfig, refreshTable } = useDataTable<Api>({
	params: {
		page: 1,
		page_size: 10,
		keyword: '',
		active: -1,
		time_range: [0, 0],
	},
	immediate: false,
	fetchFn: getApiList,
	useParams: params => {
		return {
			page: params.page,
			page_size: params.page_size,
			keyword: params.keyword,
			active: params.active,
			start_time: Math.floor(params.time_range[0] / 1000),
			end_time: Math.floor(params.time_range[1] / 1000),
		}
	},
})

const handleRefreshData = async () => {
	overviewRef.value?.getStats()
	refreshTable()
}

const [FormModal, formModalApi] = useModal({
	component: ApiForm,
	state: {
		isEdit: false,
		refresh: refreshTable,
	},
})

const handleAdd = () => {
	formModalApi.setState({
		row: null,
		isEdit: false,
	})
	formModalApi.open()
}

const handleEdit = (row: Api) => {
	formModalApi.setState({
		row,
		isEdit: true,
	})
	formModalApi.open()
}

const [TestModal, testModalApi] = useModal({
	component: ApiTest,
	state: {
		row: null,
	},
})

const handleTest = (row: Api) => {
	testModalApi.setState({ row })
	testModalApi.open()
}

const handleDelete = (row: Api) => {
	confirm({
		title: `删除API【${row.api_name}】`,
		content: '确认删除该API吗？',
		onConfirm: async () => {
			await deleteApi({ id: row.id })
			refreshTable()
		},
	})
}

onMounted(() => {
	refreshTable()
})
</script>

<style lang="scss" scoped></style>
