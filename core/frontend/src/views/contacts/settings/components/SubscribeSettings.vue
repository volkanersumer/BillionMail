<template>
	<bt-form class="subscribe-settings">
		<n-form-item label="列表类型">
			<div class="w-200px">
				<n-select
					v-model:value="form.double_optin"
					:options="listTypeOptions"
					placeholder="请选择列表类型">
				</n-select>
			</div>
		</n-form-item>
		<n-form-item label="感谢邮件订阅主题">
			<div class="w-420px">
				<n-input v-model:value="form.welcome_subject"></n-input>
			</div>
		</n-form-item>
		<n-form-item>
			<template #label>
				<span class="mr-16px">感谢邮件订阅邮件</span>
				<n-switch v-model:value="form.send_welcome_email" :checked-value="1" :unchecked-value="0">
				</n-switch>
			</template>
			<div class="flex-1">
				<div class="setting-description mb-8px">
					When enabled, a thank you email will be sent to the user when they subscribe via the
					subscription form or API. This will not be triggered for manually added users.
				</div>
				<editor v-model:value="form.welcome_mail_html"></editor>
			</div>
		</n-form-item>
		<n-form-item label="发送感谢订阅的跳转页面链接">
			<div class="w-420px">
				<n-input v-model:value="form.success_url"></n-input>
			</div>
		</n-form-item>
		<n-form-item label="二次确认邮件主题">
			<div class="w-420px">
				<n-input v-model:value="form.confirm_subject"></n-input>
			</div>
		</n-form-item>
		<n-form-item label="二次确认邮件">
			<editor v-model:value="form.confirm_mail_html"></editor>
		</n-form-item>
		<n-form-item label="二次确认的跳转页面链接">
			<div class="w-420px">
				<n-input v-model:value="form.confirm_url"></n-input>
			</div>
		</n-form-item>
		<n-form-item label="订阅已存在的跳转页面链接">
			<div class="w-420px">
				<n-input v-model:value="form.already_url"></n-input>
			</div>
		</n-form-item>
		<n-form-item :show-label="false">
			<n-button type="primary" @click="handleSave">
				{{ $t('common.actions.save') }}
			</n-button>
		</n-form-item>
	</bt-form>
</template>

<script setup lang="ts">
import { SelectOption } from 'naive-ui'
import { useContext } from '../hooks/useContext'

import Editor from './Editor.vue'
import { saveSubscribeSetting } from '@/api/modules/contacts/group'

const form = reactive({
	group_id: 0,
	double_optin: 1,
	send_welcome_email: 1,
	welcome_subject: '',
	thank_you_subject: '',
	welcome_mail_html: '',
	welcome_mail_drag: '',
	success_url: '',
	confirm_subject: '',
	confirm_mail_html: '',
	confirm_mail_drag: '',
	confirm_url: '',
	already_url: '',
})

// 选项数据
const listTypeOptions: SelectOption[] = [
	{ label: '两次确认', value: 1 },
	{ label: '单次确认', value: 0 },
]

const { groupInfo } = useContext()

const handleSave = async () => {
	await saveSubscribeSetting(toRaw(form))
}

watchEffect(() => {
	if (groupInfo.value) {
		form.group_id = groupInfo.value.id
		form.double_optin = groupInfo.value.double_optin
		form.welcome_subject = groupInfo.value.welcome_subject
		form.send_welcome_email = groupInfo.value.send_welcome_email
		form.welcome_mail_html = groupInfo.value.welcome_mail_html
		form.welcome_mail_drag = groupInfo.value.welcome_mail_drag
		form.success_url = groupInfo.value.success_url
		form.confirm_subject = groupInfo.value.confirm_subject
		form.confirm_mail_html = groupInfo.value.confirm_mail_html
		form.confirm_mail_drag = groupInfo.value.confirm_mail_drag
		form.confirm_url = groupInfo.value.confirm_url
		form.already_url = groupInfo.value.already_url
	}
})
</script>

<style lang="scss" scoped>
.subscribe-settings {
	width: 80%;
	min-width: 1000px;
}

.setting-description {
	line-height: 1.4;
	color: #666;
}

.action-buttons {
	width: 100%;
	padding: 23px 24px;
	background-color: var(--color-bg-1);
	border-top: 1px solid var(--color-border-1);
}
</style>
