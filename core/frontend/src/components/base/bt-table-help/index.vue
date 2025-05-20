<template>
	<n-empty>
		<template #default>
			<div class="text-14px text-desc text-center">
				<span>Submit </span>
				<n-button class="text-14px" type="primary" text @click="handleGoIssues">
					<span>issues</span>
					<i class="i-mdi:arrow-right ml-1px"></i>
				</n-button>
				<span> for questions or suggestions, and light up a </span>
				<n-button class="text-14px" type="primary" text @click="handleGoStar">
					Star
					<i class="i-mdi:arrow-right ml-1px"></i>
				</n-button>
				<span> on Github to support the project.</span>
				<br />
				<span>We highly value your active participation. Thank you.</span>
			</div>
		</template>
	</n-empty>
</template>

<script lang="ts" setup>
const handleGoIssues = () => {
	window.open('https://github.com/aaPanel/BillionMail/issues')
}

const handleGoStar = () => {
	window.open('https://github.com/aaPanel/BillionMail')
}
</script>

<style lang="scss" scoped></style>
