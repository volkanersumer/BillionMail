<template>
	<div class="flex gap-16px mb-24px">
		<n-card class="flex-1">
			<div class="title">
				<i class="i-fa:paper-plane"></i>
				<span>发送总量</span>
			</div>
			<div class="value">84,200</div>
		</n-card>
		<n-card class="flex-1">
			<div class="title">
				<i class="i-fa:envelope-open"></i>
				<span>平均打开率</span>
			</div>
			<div class="value">32.4%</div>
		</n-card>
		<n-card class="flex-1">
			<div class="title">
				<i class="i-fa:mouse-pointer"></i>
				<span>平均点击率</span>
			</div>
			<div class="value">15.7%</div>
		</n-card>
		<n-card class="flex-1">
			<div class="title">
				<i class="i-fa:ban"></i>
				<span>平均退信率</span>
			</div>
			<div class="value">0.5%</div>
		</n-card>
	</div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped>
.n-card {
	--n-padding-top: 20px;
	--n-padding-left: 20px;
	--n-padding-right: 20px;
	--n-padding-bottom: 20px;
}

.title {
	display: flex;
	align-items: center;
	margin-bottom: 8px;
	gap: 8px;
	font-size: 14px;
	color: var(--color-text-3);
}

.value {
	font-size: 28px;
	font-weight: 600;
	margin-bottom: 4px;
	font-family: 'Montserrat', sans-serif;
	color: var(--color-text-1);
}
</style>
