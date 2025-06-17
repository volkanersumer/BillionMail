<template>
	<modal :title="title" width="520">
		<bt-form ref="formRef" :model="form" :rules="rules" class="pt-12px">
			<n-form-item label="API名称" path="api_name">
				<n-input v-model:value="form.api_name" placeholder="请输入API名称"></n-input>
			</n-form-item>
			<n-form-item :label="$t('market.task.edit.from')" path="addresser">
				<from-select v-model:value="form.addresser" v-model:name="form.full_name"> </from-select>
			</n-form-item>
			<n-form-item :label="$t('market.task.edit.displayName')" path="full_name">
				<n-input
					v-model:value="form.full_name"
					:placeholder="$t('market.task.edit.displayNamePlaceholder')">
				</n-input>
			</n-form-item>
			<n-form-item :label="$t('market.task.edit.subject')" path="subject">
				<n-input
					v-model:value="form.subject"
					:placeholder="$t('market.task.edit.subjectPlaceholder')">
				</n-input>
			</n-form-item>
			<n-form-item :label="$t('market.task.edit.template')" path="template_id">
				<div class="flex-1">
					<template-select v-model:value="form.template_id" v-model:content="form.template_content">
					</template-select>
				</div>
				<n-button text type="primary" class="ml-12px" @click="handlePreviewTemplate">
					{{ $t('common.actions.preview') }}
				</n-button>
			</n-form-item>
			<n-form-item :label="$t('market.task.edit.unsubscribeLink')">
				<n-switch v-model:value="form.unsubscribe" :checked-value="1" :unchecked-value="0">
				</n-switch>
			</n-form-item>
			<n-form-item label="状态">
				<n-switch v-model:value="form.active" :checked-value="1" :unchecked-value="0"></n-switch>
			</n-form-item>
			<preview-modal />
		</bt-form>
	</modal>
</template>

<script lang="ts" setup>
import { FormRules } from 'naive-ui'
import { useModal } from '@/hooks/modal/useModal'
import { createApi, updateApi } from '@/api/modules/api'
import type { Api } from '../types/base'

import FromSelect from '@/views/market/task/components/FromSelect.vue'
import TemplateSelect from '@/views/market/task/components/TemplateSelect.vue'
import TemplatePreview from '@/views/market/template/components/TemplatePreview.vue'

const isEdit = ref(false)

const title = computed(() => {
	return isEdit.value ? 'Edit API' : 'Add API'
})

const formRef = useTemplateRef('formRef')

const form = reactive({
	id: 0,
	api_name: '',
	template_id: null as number | null,
	template_content: '',
	subject: '',
	addresser: null as string | null,
	full_name: '',
	unsubscribe: 1,
	active: 1,
})

const rules: FormRules = {
	api_name: {
		required: true,
		message: '请输入API名称',
	},
	addresser: {
		required: true,
		message: '请选择发送者',
	},
	full_name: {
		required: true,
		message: '请输入发送者名称',
	},
	subject: {
		required: true,
		message: '请输入主题',
	},
	template_id: {
		required: true,
		message: '请选择模板',
	},
}

const [PreviewModal, previewModalApi] = useModal({
	component: TemplatePreview,
})

const handlePreviewTemplate = () => {
	previewModalApi.setState({ html: form.template_content })
	previewModalApi.open()
}

const resetForm = () => {
	form.id = 0
	form.api_name = ''
	form.template_id = null
	form.template_content = ''
	form.subject = ''
	form.addresser = null
	form.full_name = ''
	form.unsubscribe = 1
	form.active = 1
}

const getParams = () => {
	return {
		api_name: form.api_name,
		template_id: form.template_id || 0,
		subject: form.subject,
		addresser: form.addresser || '',
		full_name: form.full_name,
		unsubscribe: form.unsubscribe,
		active: form.active,
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ isEdit: boolean; row: Api }>()
			const { row } = state
			isEdit.value = state.isEdit
			if (row) {
				form.id = row.id
				form.api_name = row.api_name
				form.template_id = row.template_id
				form.subject = row.subject
				form.addresser = row.addresser
				form.full_name = row.full_name
				form.unsubscribe = row.unsubscribe
				form.active = row.active
			}
		} else {
			resetForm()
		}
	},
	onConfirm: async () => {
		await formRef.value?.validate()
		const params = getParams()
		if (isEdit.value) {
			await updateApi({ ...params, id: form.id })
		} else {
			await createApi(params)
		}
		const state = modalApi.getState<{ refresh: Function }>()
		state.refresh()
	},
})
</script>
