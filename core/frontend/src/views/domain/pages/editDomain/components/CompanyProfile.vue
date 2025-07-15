<template>
	<div class="content-wrapper">
		<n-card class="mb-5">
			<div class="page-tit">
				<div class="left-tit">
					<div class="left-tit">
						<div class="back-tool">
							<i class="i-ri:apps-fill text-6"></i>
						</div>
						<span class="tit-content"> Company Detail </span>
					</div>
				</div>
			</div>
		</n-card>

		<n-card class="mb-5">
			<n-form>
				<div class="flex justify-between gap-5">
					<n-form-item class="flex-1">
						<template #label><span class="form-label">Legal Company Name</span></template>
						<n-input v-model:value="legal_company_name"></n-input>
					</n-form-item>
					<n-form-item class="flex-1">
						<template #label><span class="form-label">web_site</span></template>
						<n-input v-model:value="web_site"></n-input>
					</n-form-item>
				</div>
				<n-form-item>
					<template #label><span class="form-label">Organization Number</span></template>
					<n-input></n-input>
				</n-form-item>
				<n-form-item>
					<template #label><span class="form-label">Additional Company Information</span></template>
					<n-input v-model:value="company_profile" type="textarea" :rows="8"></n-input>
				</n-form-item>
			</n-form>
		</n-card>

		<n-card>
			<div class="page-tit mb-20px">
				<div class="left-tit">
					<div class="back-tool">
						<i class="i-ri:apps-fill text-6"></i>
					</div>
					<span class="tit-content"> Contact Information </span>
				</div>
			</div>

			<n-form>
				<n-form-item>
					<template #label><span class="form-label">Contact Email</span></template>
					<n-input v-model:value="email"></n-input>
				</n-form-item>
				<n-form-item>
					<template #label><span class="form-label">Phone Number</span></template>
					<n-input v-model:value="phone"></n-input>
				</n-form-item>
				<n-form-item>
					<template #label><span class="form-label">Support URL</span></template>
					<n-input v-model:value="support_url"></n-input>
				</n-form-item>
			</n-form>
		</n-card>
	</div>
</template>

<script setup lang="ts">
import { getCompanuyProfile } from '../controller/companyProfile.controller'
import { getEditDomainStoreData } from '../store'
const { legal_company_name, web_site, phone, company_profile, email, support_url } =
	getEditDomainStoreData()
const route = useRoute()
const domain = route.params.domain as string
getCompanuyProfile(domain)
</script>

<style scoped lang="scss">
@use '@/styles/index' as base;
@use './mixin.scss' as mixin;

.content-wrapper {
	@include mixin.content-wrapper;

	.page-tit {
		@include mixin.page-tit;
	}

	.form-label {
		@include mixin.form-label;
	}
}
</style>
