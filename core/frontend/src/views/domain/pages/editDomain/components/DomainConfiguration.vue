<template>
	<div class="content-wrapper">
		<div>
			<n-card>
				<div class="page-tit mb-5">
					<div class="left-tit">
						<div class="back-tool">
							<i class="i-cuida:mail-outline text-28px"></i>
						</div>
						<span class="tit-content">{{ $t('domain.edit.domainConfiguration.title') }}</span>
					</div>
				</div>
				<!-- form data -->
				<n-form>
					<n-form-item :label="$t('domain.edit.domainConfiguration.domain')">
						<n-input v-model:value="domainTit" :disabled="true" @input="syncToUrl"></n-input>
					</n-form-item>
					<n-form-item label="域名专用IP">
						<div class="flex-1">
							<div class="mb-12px text-desc">
								* 域名 aapanel.com
								发送和接受邮件时会在iptables处指定此专属IP，如果你的服务器只有一个IP的话不需要设置此功能
							</div>
							<div class="flex gap-16px">
								<div class="flex-1">
									<n-input v-model:value="domainIp" placeholder="未设置"></n-input>
								</div>
								<n-button type="primary" ghost @click="testConnection">连接测试</n-button>
							</div>
						</div>
					</n-form-item>
					<n-form-item :label="$t('domain.edit.domainConfiguration.catchAll')">
						<div class="flex-1">
							<div class="mb-12px text-desc">
								* 在您这里填写邮箱后，这个域名下所有未知的邮箱收到邮件将会转发到这个邮箱里面
							</div>
							<n-input v-model:value="catch_email"></n-input>
						</div>
					</n-form-item>
					<n-form-item :label="$t('domain.edit.domainConfiguration.domainQuota')">
						<div class="flex-1">
							<div class="mb-12px text-desc">* 设置此域名的可用空间</div>
							<div class="flex gap-16px">
								<div class="flex-1">
									<n-input v-model:value="quota"></n-input>
								</div>
								<div class="w-80px">
									<n-select v-model:value="unit" :options="uintOptions"></n-select>
								</div>
							</div>
						</div>
					</n-form-item>
					<n-form-item :label="$t('domain.edit.domainConfiguration.mailboxCount')">
						<div class="flex-1">
							<div class="mb-12px text-desc">* 设置此域名可以创建多少个邮箱</div>
							<n-input-number v-model:value="mailboxes" class="w-full" :show-button="false">
							</n-input-number>
						</div>
					</n-form-item>
				</n-form>
			</n-card>

			<n-card v-if="createdBrandInfo" class="mt-5">
				<div class="page-tit mb-5">
					<div class="left-tit">
						<div class="back-tool">
							<i class="i-cuida:mail-outline text-7"></i>
						</div>
						<span class="tit-content">{{ $t('domain.edit.domainConfiguration.brandInfo') }}</span>
					</div>
					<div class="right-tit">
						<n-switch v-model:value="brandInfo" @update:value="switchBrandInfo"></n-switch>
					</div>
				</div>

				<div class="w-100% flex flex-col gap-2.5">
					<div
						v-for="(_, index) in urls"
						:key="index"
						class="flex justify-start items-center gap-2.5">
						<n-input
							v-model:value="urls[index]"
							:placeholder="$t('domain.form.urlsPlacement')"
							:disabled="!supplierStatus"
							readonly>
						</n-input>
						<div v-if="index != 0" class="close" @click="removeUrl(index)">
							<i class="i-material-symbols:close-rounded text-5"></i>
						</div>
					</div>
				</div>
			</n-card>

			<n-card v-else class="mt-5">
				<div class="fw-bold text-4 mb-5 text-[#777] flex justify-between items-center">
					<span>{{ $t('domain.edit.domainConfiguration.noBrandInfo') }}</span>
				</div>
				<n-alert v-if="!supplierStatus" type="warning" class="mb-15px" :show-icon="false">
					<div class="w-100% flex justify-between items-center">
						<span class="mr-5">
							{{ $t('domain.edit.domainConfiguration.aiIntegrationWarning') }}
						</span>
						<n-button ghost type="primary" @click="jumpToAiSettings">
							{{ $t('domain.edit.domainConfiguration.integrateImmediately') }}
						</n-button>
					</div>
				</n-alert>
				<div class="mb-8px text-default">
					{{ $t('domain.edit.domainConfiguration.aiDescription') }}
				</div>
				<bt-tips style="margin-bottom: 15px">
					<li>{{ $t('domain.edit.domainConfiguration.features.extractBrand') }}</li>
					<li>{{ $t('domain.edit.domainConfiguration.features.analyzeStructure') }}</li>
					<li>{{ $t('domain.edit.domainConfiguration.features.importContent') }}</li>
					<li>{{ $t('domain.edit.domainConfiguration.features.customLogo') }}</li>
				</bt-tips>
				<n-form-item :label="$t('domain.edit.domainConfiguration.specifyDomain')">
					<div class="w-100% flex flex-col gap-2.5">
						<div
							v-for="(_, index) in urls"
							:key="index"
							class="flex justify-start items-center gap-2.5">
							<n-input
								v-model:value="urls[index]"
								:placeholder="$t('domain.form.urlsPlacement')"
								:disabled="!supplierStatus">
							</n-input>
							<div v-if="index != 0" class="close" @click="removeUrl(index)">
								<i class="i-material-symbols:close-rounded text-5"></i>
							</div>
						</div>
						<!-- <div class="text-[#777]">{{ $t("domain.form.maxUrl") }}</div> -->
						<!-- <n-button type="primary" :disabled="!supplierStatus" ghost @click="addUrl">
							<template #icon>
								<i class="i-material-symbols:add-circle-outline"></i>
							</template>
							{{ $t('domain.edit.domainConfiguration.addMoreUrls') }}
						</n-button> -->
					</div>
				</n-form-item>
				<n-button type="primary" ghost @click="createBrandInfo">
					{{ $t('domain.edit.domainConfiguration.createNow') }}
				</n-button>
			</n-card>
			<n-button type="primary" class="w-100% mt-5" @click="updateDomain">
				<template #icon>
					<i class="i-mingcute:save-2-line text-5"></i>
				</template>
				{{ $t('domain.edit.save') }}
			</n-button>
		</div>
	</div>

	<!-- Wait for init brand info -->
	<WaitAndCheckDomainStatus ref="waitAndCheckDomainStatusRef" />
</template>

<script setup lang="ts">
import WaitAndCheckDomainStatus from '@/views/domain/components/WaitAndCheckDomainStatus.vue'
import { checkAiConfiguration } from '@/api/modules/domain'
import {
	getDomainDetail,
	syncToUrl,
	removeUrl,
	createBrandInfo,
	switchBrandInfo,
	updateDomain,
	testConnection,
} from '../controller/domainConfiguration.controller'
import { getEditDomainStoreData } from '../store'

const router = useRouter()

const supplierStatus = ref(false)

const {
	domainTit,
	domainIp,
	quota,
	unit,
	mailboxes,
	catch_email,
	urls,
	brandInfo,
	createdBrandInfo,
	waitAndCheckDomainStatusRef,
} = getEditDomainStoreData()

const route = useRoute()

const domain = route.params.domain as any

const uintOptions = ref([
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
 * @description Check AI configuration
 */
async function checkAiConfig() {
	const res = (await checkAiConfiguration()) as Record<string, any>
	if (res) {
		supplierStatus.value = res.is_configured
	}
}

checkAiConfig()

/**
 * @description Add url to urls
 */
// function addUrl() {
// 	urls.value.push('')
// }

/**
 * @description Jump to AI settings page
 */
function jumpToAiSettings() {
	router.push({
		name: 'AiModel',
	})
}
</script>

<style scoped lang="scss">
@use '@/styles/index' as base;
@use './mixin.scss' as mixin;

.content-wrapper {
	@include mixin.content-wrapper;
	padding-bottom: 32px;

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

	.close {
		@include mixin.operation-close;
	}
}
</style>
