<template>
	<div class="wrapper">
		<div class="page-tit">
			<div class="back-tool">
				<i class="i-ci:arrow-left-lg text-4"></i>
				<div class="back-btn">back</div>
			</div>
			<span class="tit-content"> AI Template </span>
		</div>

		<div class="ai-container">
			<div class="chat-area">
				<div class="model-select">
					<n-card>
						<div class="flex justify-start items-center gap-15px">
							<!-- <span>Choose Model</span> -->
							<n-select
								v-model:value="currentModelTitle"
								class="flex-1"
								:options="modelList"
								label-field="title"
								value-field="title"
								@update:value="changeModel">
							</n-select>
						</div>
					</n-card>
				</div>

				<div ref="answerRef" class="answer" @wheel="handleScroll">
					<n-card class="h-100%">
						<n-scrollbar ref="chatScrollRef" style="height: calc(100vh - 522px)">
							<div ref="scrollWrapperRef">
								<template v-for="item in chatRecord" :key="item[0]">
									<div class="answer-container">
										<div class="ask-content">
											<div class="pic">
												<i
													class="i-material-symbols:account-circle text-6 text-[var(--color-primary-hover-1)]"></i>
											</div>
											<div class="text">
												{{ item[0].split('_+_')[0] }}
											</div>
										</div>
										<!-- <n-spin description="Loading please wait..." :show="generateShow"> -->
										<div class="answer-content">
											<div class="pic">
												<span class="text-4 text-[var(--color-primary-hover-1)]">BillionMail</span>
											</div>
											<div class="text">
												<MarkdownRender
													v-for="(_item, _index) in item[1]"
													:key="_index"
													:content="_item"
													:chat-record-key="item[0]"
													class="mb-4"
													@code-render="handleCodeRender">
												</MarkdownRender>
											</div>
											<div class="answer-tool" @click="handleRegenerate(item[0].toString())">
												<div class="tool-item">
													<i class="i-ri:refresh-line text-5"></i>
													<span>Regenerate</span>
												</div>
												<div class="tool-item">
													<i class="i-ri:information-2-fill text-5"></i>
													<span>Info</span>
												</div>
											</div>
										</div>
										<!-- </n-spin> -->
									</div>
								</template>
							</div>
						</n-scrollbar>
					</n-card>
				</div>
				<div class="question">
					<n-card style="height: 100%">
						<n-input
							v-model:value="questionContent"
							type="textarea"
							class="question-input"
							@keydown.enter="sendChat(store)"></n-input>
						<div class="send-btn">
							<!-- <i
                                class="i-icon-park-twotone:new-picture text-8 text-[var(--color-primary-1)] hover-text-[var(--color-primary-hover-1)]"></i> -->
							<i
								v-if="!isChat"
								class="i-icon-park-twotone:arrow-circle-up text-8 text-[var(--color-primary-1)] hover-text-[var(--color-primary-hover-1)]"
								@click="sendChat(store)"></i>
							<i
								v-else
								class="i-svg-spinners:pulse-2 text-8 text-[var(--color-primary-1)] hover-text-[var(--color-primary-hover-1)]"
								@click="stopChat(store)"></i>
						</div>
					</n-card>
				</div>
			</div>
			<div class="view-area">
				<n-card style="height: 100%">
					<div style="height: calc(100vh - 177px)">
						<div class="h-10 flex justify-start items-center gap-2.5">
							<n-button-group size="small">
								<n-button
									:type="previewStatus == 'view' ? 'primary' : 'default'"
									@click="previewStatus = 'view'">
									<template #icon>
										<i class="i-mingcute:eye-2-fill text-5"></i>
									</template>
									View
								</n-button>
								<n-button
									:type="previewStatus == 'edit' ? 'primary' : 'default'"
									@click="previewStatus = 'edit'">
									<template #icon>
										<i class="i-tdesign:code text-5"></i>
									</template>
									Edit
								</n-button>
							</n-button-group>
							<span v-if="previewStatus == 'edit'">[Ctrl + s] to save your change</span>
						</div>
						<n-scrollbar style="height: calc(100% - 40px)">
							<div v-if="previewStatus == 'view'" v-html="previewCode"></div>
							<BtEditor
								v-else
								:value="previewCode"
								style="height: calc(100vh - 217px)"
								@save="saveCodeChange(store)">
							</BtEditor>
						</n-scrollbar>
					</div>
				</n-card>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {
	initialTemplateInfo,
	sendChat,
	removeHtmlCodeBlockMarkers,
	getHtmlTemplateContent,
	saveCodeChange,
	stopChat,
} from './controller'
import MarkdownRender from './components/MarkdownRender.vue'
import { useTemplateStore } from './store'
import { TemplateStore } from './dto'
import BtEditor from '@/components/base/bt-editor/index.vue'
const store = useTemplateStore()
provide<TemplateStore>('modelStore', store)
const {
	chatId,
	modelList,
	currentModelTitle,
	previewCode,
	questionContent,
	chatRecord,
	currentChatRecordKey,
	currentModel,
	isChat,
	chatScrollRef,
	scrollWrapperRef,
	scrollable,
} = store
const route = useRoute()

const previewStatus = ref('view')

chatId.value = route.params.chatId as string

/**
 * @description Handle model select
 */
function changeModel(_: string, options: any) {
	currentModel.value = options
}

/**
 * @description Get html template code content
 */
getHtmlTemplateContent(store)

/**
 * @description Initial chatInfo and modelList
 */
initialTemplateInfo(store)

/**
 * @description Render page code
 */
function handleCodeRender(data: { code: string; key: string }) {
	if (data.key == currentChatRecordKey.value) {
		// console.log(data.key, currentChatRecordKey.value)
		previewCode.value = removeHtmlCodeBlockMarkers(data.code)
	}
}

/**
 * @description Regenerate
 */
function handleRegenerate(key: string) {
	questionContent.value = key.split('_+_')[0]
	sendChat(store)
}

/**
 * @description Listen scroll wrapper for height-change
 */
function scrollWrapperHeightChange(timer?: any, index = 0) {
	let timeoutTimer = timer || null
	let startIndex = index
	if (startIndex >= 6) {
		return
	}
	if (timeoutTimer) {
		clearTimeout(timer)
		timeoutTimer = null
	}
	if (scrollWrapperRef.value.offsetHeight) {
		chatScrollRef.value.scrollTo({ left: 0, top: scrollWrapperRef.value.offsetHeight })
		startIndex++
		return
	} else {
		startIndex++
		timeoutTimer = setTimeout(() => {
			scrollWrapperHeightChange(timeoutTimer, startIndex)
		}, 100)
	}
}

function handleScroll(event: WheelEvent) {
	if (event.deltaY < 0) {
		scrollable.value = false
	}
}

onMounted(() => {
	scrollWrapperHeightChange()
})
</script>

<style scoped lang="scss">
@use '@/styles/index' as base;
@use '../../mixin.scss';

.wrapper {
	@mixin content-card {
		border-radius: 5px;
		padding: 20px;
	}

	@include mixin.wrapper;

	.page-tit {
		@include mixin.page-tit;
	}

	.ai-container {
		display: grid;
		grid-template-columns: 900px 1fr;
		gap: 15px;

		.chat-area {
			display: grid;
			grid-template-rows: 76px 1fr 230px;
			gap: 15px;

			.answer {
				width: 100%;
				height: calc(100vh - 472px);

				.answer-container {
					padding: 30px 10px;

					.ask-content,
					.answer-content {
						@include content-card();
						background: var(--color-bg-3);
						display: grid;
						grid-template-rows: 30px 1fr;
						gap: 10px;
						margin-bottom: 20px;

						.pic {
							@include base.row-flex;
							align-items: center;
						}

						@mixin item-tool {
							@include base.row-flex-start;
							cursor: pointer;
							gap: 5px;
							padding: 5px 10px;
							align-items: center;
							font-size: 12px;
							border-radius: 3px;
							transition: 0.1s all ease-in-out;

							&:hover {
								background: var(--color-primary-1);
								color: #fff;
							}

							&:active {
								background: var(--color-primary-3);
							}
						}

						.text {
							font-size: 14px;
						}

						.answer-tool {
							@include base.row-flex-end;
							border-top: 1px solid var(--color-border-1);
							height: 30px;
							padding-top: 20px;

							.tool-item {
								@include item-tool();
							}
						}
					}

					.answer-content {
						background: var(--color-bg-2);
					}
				}
			}

			.question {
				position: relative;

				.question-input {
					height: 100%;
					font-size: 14px;
				}

				.send-btn {
					position: absolute;
					right: 30px;
					bottom: 30px;
					@include base.row-flex-start;
					gap: 10px;

					i {
						cursor: pointer;
					}
				}
			}
		}

		/* .view-area {
			background: linear-gradient(-45deg, rgba(238, 152, 82, 0.3),rgba(35, 166, 213,.3), rgba(231, 60, 60, 0.3),  rgba(35, 213, 171,.3));
                background-size: 400% 400%;
                animation: gradient 15s ease infinite;

                @keyframes gradient {
                    0% {
                        background-position: 0% 50%;
                    }

                    50% {
                        background-position: 100% 50%;
                    }

                    100% {
                        background-position: 0% 50%;
                    }
                } 
		} */
	}
}
</style>
