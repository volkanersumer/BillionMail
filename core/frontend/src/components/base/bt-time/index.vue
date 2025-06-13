<template>
	<div class="date-filters">
		<div class="w-120px">
			<n-select v-model:value="activeFilter" :options="filterOptions" @update:value="setFilter">
			</n-select>
		</div>
		<div class="w-240px">
			<n-date-picker
				:value="dateRange"
				type="daterange"
				:is-date-disabled="dateDisabled"
				@update:value="onDateRangeChange">
			</n-date-picker>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { getDayTimeRange } from '@/utils'
import { addDays, startOfDay, endOfDay } from 'date-fns'

const { t } = useI18n()

const dateRange = defineModel<[number, number]>('dateRange')

const activeFilter = ref('today')

const filterOptions = [
	{ label: t('common.time.today'), value: 'today' },
	{ label: t('common.time.yesterday'), value: 'yesterday' },
	{ label: t('common.time.last7days'), value: 'last7days' },
	{ label: t('common.time.custom'), value: 'custom' },
]

const getLastSevenDaysRange = (): [number, number] => {
	const today = new Date()
	const end = endOfDay(today)
	const start = startOfDay(addDays(today, -6))
	return [start.getTime(), end.getTime()]
}

const setFilter = (filter: string) => {
	switch (filter) {
		case 'today':
			dateRange.value = getDayTimeRange()
			break
		case 'yesterday':
			dateRange.value = getDayTimeRange(addDays(new Date(), -1))
			break
		case 'last7days':
			dateRange.value = getLastSevenDaysRange()
			break
	}
}

const dateDisabled = (ts: number) => {
	const today = new Date()
	const before = addDays(today, -30)
	return startOfDay(before).getTime() > ts || endOfDay(today).getTime() < ts
}

const onDateRangeChange = (val: [number, number]) => {
	activeFilter.value = 'custom'
	dateRange.value = [startOfDay(val[0]).getTime(), endOfDay(val[1]).getTime()]
}
</script>

<style lang="scss" scoped>
.date-filters {
	display: flex;
	gap: 8px;
}
</style>
