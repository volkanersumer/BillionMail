<template>
	<div class="content-wrapper">
		<div>
			<n-card class="mb-5">
				<!-- title -->
				<div class="page-tit">
					<div class="left-tit">
						<div class="back-tool">
							<i class="i-cuida:mail-outline text-7"></i>
						</div>
						<span class="tit-content"> Mail Domain Configuration </span>
					</div>
				</div>
			</n-card>
			<n-card>
				<!-- form data -->
				<n-form>
					<n-form-item>
						<template #label><span class="form-label">Domain</span></template>
						<n-input v-model:value="domainTit"></n-input>
					</n-form-item>
					<n-form-item label="">
						<template #label><span class="form-label">Domain Quota</span></template>
						<div class="flex justify-between gap-5 items-center w-100%">
							<n-input v-model:value="quota"></n-input>
							<n-select v-model:value="unit" :options="uinitOptions" class="w-20"></n-select>
						</div>
					</n-form-item>
					<n-form-item>
						<template #label><span class="form-label">MailBox Count</span></template>
						<n-input-number v-model:value="mailboxes" :show-button="false"></n-input-number>
					</n-form-item>
					<n-form-item>
						<template #label><span class="form-label">Catch all</span></template>
						<n-input v-model:value="catch_email"></n-input>
					</n-form-item>
				</n-form>
			</n-card>

			<n-card class="mt-5">
				<n-form-item label="Automatically create brand information">
					<n-switch v-model:value="initAi"></n-switch>
				</n-form-item>
				<n-alert v-if="!configurationStatus" type="warning" class="mb-15px" :show-icon="false">
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
						<n-input
							v-for="(_, index) in urls"
							:key="index"
							v-model:value="urls[index]"
							placeholder="">
						</n-input>
						<n-button type="primary" ghost @click="addUrl">
							<template #icon>
								<i class="i-material-symbols:add-circle-outline"></i>
							</template>
							Add more URLs
						</n-button>
					</div>
				</n-form-item>
			</n-card>
		</div>

		<!-- <n-card class="my-5">
            <div class="switch-settings">
                <div class="switch-item">
                    <div class="label">Track Email Opens</div>
                    <n-switch></n-switch>
                </div>
                <div class="switch-item" style="margin-bottom: 0;">
                    <div class="label">Track Link Clicks</div>
                    <n-switch></n-switch>
                </div>
            </div>
        </n-card> -->

		<!-- <n-card>
            <div class="page-tit mb-5">
                <div class="left-tit">
                    <div class="back-tool">
                        <i class="i-ri:rss-fill text-6"></i>
                    </div>
                    <span class="tit-content">
                        Subscription Management
                    </span>
                </div>
            </div>

            <div class="switch-settings">
                <div class="switch-item">
                    <div class="label">Include Unsubscribe Link</div>
                    <n-switch></n-switch>
                </div>
                <div class="switch-item" style="margin-bottom: 0;">
                    <div class="label">Inlcude Preferences Link</div>
                    <n-switch></n-switch>
                </div>
            </div>
        </n-card> -->
	</div>
</template>

<script setup lang="ts">
import { getDomainDetail } from '../controller/domainConfiguration.controller'
import { getEditDomainStoreData } from '../store'
const { domainTit, quota, unit, mailboxes, catch_email, initAi, urls, configurationStatus } =
	getEditDomainStoreData()
const route = useRoute()
const domain = route.params.domain as any
const uinitOptions = ref([
	{
		label: 'B',
		value: 'B',
	},
	{
		label: 'KB',
		value: 'KB',
	},
	{
		label: 'MB',
		value: 'MB',
	},
	{
		label: 'GB',
		value: 'GB',
	},
	{
		label: 'TB',
		value: 'TB',
	},
])

getDomainDetail(domain)

/**
 * @description Add url to urls
 */
function addUrl() {
	urls.value.push('')
}
</script>

<style scoped lang="scss">
@use '@/styles/index' as base;
@use './mixin.scss' as mixin;

.content-wrapper {
	@include mixin.content-wrapper;

	// title
	.page-tit {
		@include mixin.page-tit;
	}

	.form-label {
		@include mixin.form-label;
	}

	// Switch settings
	.switch-settings {
		@mixin settings-label {
			font-size: 14px;
			font-weight: bold;
		}

		.switch-item {
			@include base.row-flex;
			justify-content: space-between;
			margin-bottom: 15px;

			.label {
				color: var(--color-text-2);
				@include settings-label();
			}
		}
	}
}
</style>
