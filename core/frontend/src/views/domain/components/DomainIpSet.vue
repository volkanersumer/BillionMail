<template>
	<modal :title="title" :width="500" :footer="false">
		<div class="pt-16px px-16px text-basic">
			<div class="flex justify-between items-center mb-16px">
				<div>专用IP：{{ ip }}</div>
				<div class="flex items-center">
					配置文件：
					<ip-status :status="configStatus"></ip-status>
				</div>
				<div class="flex items-center">
					连接状态：
					<ip-status :status="connectStatus"></ip-status>
				</div>
			</div>
			<n-alert class="mb-16px" type="warning" :show-icon="false" :bordered="false">
				你可以通过手工在SSH终端执行以下命令使专用IP连接状态激活
				如果你有多个域名需要配置专用IP，我们建议你先把所有域名配
				置完成后在执行以下命令，避免邮件服务器多次中断
			</n-alert>

			<n-spin size="small" :show="loading">
				<n-alert class="command-alert mb-16px" :show-icon="false" :bordered="false">
					{{ command || '--' }}
				</n-alert>
			</n-spin>

			<div class="flex justify-center gap-16px">
				<n-button type="primary" :loading="loading" :disabled="loading" @click="getCommand">
					获取命令
				</n-button>
				<n-button @click="onCopyCommand">复制命令</n-button>
			</div>
		</div>

		<bt-modal v-model:show="showModal" :width="420" title="操作提醒" :footer="false">
			<n-alert class="mb-16px" type="warning" :show-icon="false" :bordered="false">
				如果执行命令失败，您可以执行该命令进行修复：<br />"{{ fixCommand }}"
			</n-alert>
			<div class="flex justify-center">
				<n-button type="primary" @click="showModal = false">我知道了</n-button>
			</div>
		</bt-modal>
	</modal>
</template>

<script lang="ts" setup>
import { isObject, Message } from '@/utils'
import { useCopy } from '@/hooks/useCopy'
import { useModal } from '@/hooks/modal/useModal'
import { getDomainIpCommand } from '@/api/modules/domain'
import { MailDomain } from '../interface'
import IpStatus from './IpStatus.vue'

const domain = ref('')

const ip = ref('')

const status = ref('')

const configStatus = computed(() => {
	if (status.value === 'failed') return 'failed'
	return 'applied'
})

const connectStatus = computed(() => {
	if (status.value === 'applied') return 'applied'
	return 'failed'
})

const title = computed(() => {
	return domain.value ? `专用IP设置【${domain.value}】` : '专用IP设置'
})

const loading = ref(false)

const isFetchCommand = ref(false)

const command = ref('')

const fixCommand = ref('')

const getCommand = async () => {
	try {
		loading.value = true
		const res = await getDomainIpCommand()
		if (isObject<{ command: string; fix_command: string }>(res)) {
			command.value = res.command
			fixCommand.value = res.fix_command
			isFetchCommand.value = true
		}
	} finally {
		loading.value = false
	}
}

const showModal = ref(false)

const { copyText } = useCopy()

const onCopyCommand = () => {
	if (!isFetchCommand.value) {
		Message.error('请先获取命令')
		return
	}

	copyText(command.value, false)
	showModal.value = true
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const { row } = modalApi.getState<{ row: MailDomain }>()
			domain.value = row.domain
			ip.value = row.multi_ip_domains?.outbound_ip || ''
			status.value = row.multi_ip_domains?.status || ''
		}
	},
})
</script>

<style lang="scss" scoped>
.command-alert {
	--n-color: var(--color-command);
	--n-content-text-color: #fff;
}
</style>
