<template>
	<modal :title="title" :width="600">
		<div class="pt-16px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item path="domain">
					<template #label>
						<div class="flex items-center justify-start gap-1.25">
							<span>{{ t('domain.form.domain') }}</span>
							<n-popover trigger="hover">
								<template #trigger>
									<i class="i-custom:help text-primary cursor-pointer"></i>
								</template>
								<bt-tips>
									<li class="text-error">{{ t('domain.form.tips.aRecordFailed') }}</li>
									<li>{{ t('domain.form.tips.dnsSetup') }}</li>
									<li>{{ t('domain.form.tips.cloudflare') }}</li>
								</bt-tips>
							</n-popover>
						</div>
					</template>
					<n-input v-model:value="form.domain" :disabled="isEdit"
						:placeholder="t('domain.form.domainPlaceholder')" @update:value="syncToUrls">
					</n-input>
				</n-form-item>
				<n-form-item v-if="false" label="A记录">
					<n-input v-model:value="form.a_record" :disabled="isEdit" placeholder="请输入A记录，例如：mail.aapanel.com">
					</n-input>
				</n-form-item>
				<n-form-item :label="t('domain.form.quota')">
					<n-input-number v-model:value="form.quota" class="flex-1" :min="0" :show-button="false">
					</n-input-number>
					<div class="w-100px ml-20px">
						<n-select v-model:value="form.quota_unit" :options="unitOptions"></n-select>
					</div>
				</n-form-item>
				<n-form-item :label="t('domain.form.mailboxCount')">
					<n-input-number v-model:value="form.mailboxes" class="flex-1" :min="0" :show-button="false">
					</n-input-number>
				</n-form-item>
				<n-form-item :label="t('domain.form.globalCatch')">
					<n-input v-model:value="form.email" :placeholder="t('domain.form.globalCatchPlaceholder')">
					</n-input>
				</n-form-item>

				<n-form-item label="Automatically create brand information">
					<n-switch v-model:value="initAi"></n-switch>
				</n-form-item>
				<n-alert v-if="!aiConfigurationStatus" style="margin: 0 0 15px 0" type="warning" :show-icon="false">
					<div class="w-100% flex justify-between items-center">
						<span class="mr-5">To use this feature, you need to integrate an AI model first.</span>
						<n-button type="primary">Integrate immediately</n-button>
					</div>
				</n-alert>
				<div class="text-[#777]">
					AI-driven information import automatically analyzes and creates brand information from
					email domains. After creation, you can modify it at your discretion.
				</div>
				<bt-tips style="margin-bottom: 15px">
					<li>Extract brand and color</li>
					<li>Analyze content structure</li>
					<li>Import images and content</li>
					<li>Use a custom logo (optional)</li>
				</bt-tips>
				<n-form-item label="Specify a domain name">
					<div class="w-100% flex flex-col gap-2.5">
						<!-- <n-input v-for="(item, index) in urls" :key="index" v-model:value="urls[index]" placeholder="">
						</n-input> -->
						<div v-for="(_, index) in urls" :key="index" class="flex justify-start items-center gap-2.5">
							<n-input v-model:value="urls[index]" placeholder="">
							</n-input>
							<div v-if="index != 0" class="close" @click="removeUrl(index)">
								<i class="i-material-symbols:close-rounded text-5"></i>
							</div>
							<!-- <div v-else class="w-34px h-34px"></div> -->
						</div>

						<n-button type="primary" ghost @click="addUrl">
							<template #icon>
								<i class="i-material-symbols:add-circle-outline"></i>
							</template>
							Add more URLs
						</n-button>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
	import { FormRules } from 'naive-ui'
	import { getByteUnit, getNumber } from '@/utils'
	import { useModal } from '@/hooks/modal/useModal'
	import { createDomain, initAiConfiguration, updateDomain, checkAiConfiguration } from '@/api/modules/domain'
	import type { MailDomain } from '../interface'

	const { t } = useI18n()
	const isEdit = ref(false)
	const title = computed(() => {
		return isEdit.value ? t('domain.form.editTitle') : t('domain.form.addTitle')
	})
	const aiConfigurationStatus = ref(false)

	const formRef = useTemplateRef('formRef')
	const initAi = ref(false)
	const urls = ref([''])
	const form = reactive({
		domain: '',
		a_record: '',
		quota: 5,
		quota_unit: 'GB',
		mailboxes: 50,
		email: '',
	})

	const unitOptions = [
		{ label: 'GB', value: 'GB' },
		{ label: 'MB', value: 'MB' },
	]

	const rules: FormRules = {
		domain: {
			trigger: ['blur', 'input'],
			validator: () => {
				if (form.domain.trim() === '') {
					return new Error(t('domain.form.validation.domainRequired'))
				}
				return true
			},
		},
	}

	/**
	 * @description Calculate the byte number based on the domain quota and unit
	 * @param quota
	 * @param quota_unit
	 */
	const getQuotaByte = (quota: number, quota_unit: string) => {
		switch (quota_unit) {
			case 'GB':
				return quota * 1024 * 1024 * 1024
			case 'MB':
				return quota * 1024 * 1024
			default:
				return quota
		}
	}

	/**
	 * @description Sync domain to urls
	 */
	const syncToUrls = (domainVal: string) => {
		urls.value[0] = domainVal
	}

	/**
	 * @description Remove url
	 */
	const removeUrl = (index: number) => {
		urls.value.splice(index, 1)
	}

	const [Modal, modalApi] = useModal({
		onChangeState: async isOpen => {
			if (isOpen) {
				const state = modalApi.getState<{ isEdit: boolean; row: MailDomain | null }>()
				isEdit.value = state.isEdit
				const res = await checkAiConfiguration() as Record<string, any>
				aiConfigurationStatus.value = res.is_configured
				if (state.row) {
					const { row } = state
					form.domain = row.domain
					form.a_record = row.a_record
					const quota = getByteUnit(row.quota)
					form.quota = getNumber(quota.split(' ')[0])
					form.quota_unit = quota.split(' ')[1]
					form.mailboxes = row.mailboxes
					form.email = row.email
				}
			} else {
				form.domain = ''
				form.a_record = ''
				form.quota = 5
				form.quota_unit = 'GB'
				form.mailboxes = 50
				form.email = ''
				urls.value = ['']
			}
		},
		onConfirm: async () => {
			await formRef.value?.validate()
			if (isEdit.value) {
				await updateDomain({
					domain: form.domain,
					quota: getQuotaByte(form.quota, form.quota_unit),
					mailboxes: form.mailboxes,
					email: form.email,
					urls: urls.value
				})
			} else {
				await createDomain({
					domain: form.domain,
					quota: getQuotaByte(form.quota, form.quota_unit),
					mailboxes: form.mailboxes,
					email: form.email,
					urls: urls.value
				})
				// Init Ai configuration (via. src\api\modules\domain.ts)
				if (initAi.value) {
					await initAiConfiguration({
						domain: form.domain,
						urls: urls.value,
					})
				}
			}
			const state = modalApi.getState<{ refresh: Function }>()
			state.refresh()
		},
	})

	/**
	 * @description Add url to urls
	 */
	function addUrl() {
		urls.value.push('')
	}
</script>

<style lang="scss" scoped>
	@use "@/views/domain/pages/editDomain/components/mixin.scss";

	.close {
		@include mixin.operation-close
	}
</style>
