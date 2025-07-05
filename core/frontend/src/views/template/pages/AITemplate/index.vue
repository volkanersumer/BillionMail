<template>
    <div class="wrapper">
        <div class="page-tit">
            <div class="back-tool">
                <i class="i-ci:arrow-left-lg text-4"></i>
                <div class="back-btn">back</div>
            </div>
            <span class="tit-content">
                AI Template
            </span>
        </div>

        <div class="ai-container">
            <div class="chat-area">
                <div ref="answerRef" class="answer">
                    <n-card style="height: 100%;">
                        <n-scrollbar :style="{ height: `${answerHeight! - 40}px` }">
                            <div class="answer-container">
                                <div class="ask-content">
                                    <div class="pic">
                                        <i
                                            class="i-material-symbols:account-circle text-6 text-[var(--color-primary-hover-1)]"></i>
                                    </div>
                                    <div class="text">
                                        我有几个问题，帮我实现以下效果：<br />
                                        1. 定制一个618宝塔面板的优惠活动<br />
                                        2. 主题色为绿色<br />
                                        3. 价格区间可以参考官网设计的价格，但需要注意的是，如果企业版本和活动版本的优惠冲突，那就以企业版本的优先<br />
                                    </div>
                                </div>
                                <div class="answer-content">
                                    <div class="pic">
                                        <span class="text-4 text-[var(--color-primary-hover-1)]">BillionMail</span>
                                    </div>
                                    <div class="text">
                                        <div class="mb-2.5">好的，根据您的要求，我为您生成了以下代码结构，您可以直接复制运行: </div>
                                        <div class="code-area">
                                            <div class="code-tool">
                                                <span>{{ languageTit }}</span>
                                                <div class="tools">
                                                    <div class="tool-item">
                                                        <i class="i-material-symbols:chrome-restore-rounded text-5"></i>
                                                        <span>Copy</span>
                                                    </div>
                                                    <div class="tool-item">
                                                        <i class="i-mingcute:eye-2-fill text-5"></i>
                                                        <span>View</span>
                                                    </div>
                                                </div>
                                            </div>
                                            <div v-html="codeHtml"></div>
                                        </div>
                                    </div>
                                    <div class="answer-tool">
                                        <div class="tool-item">
                                            <i class="i-ri:refresh-line text-5"></i>
                                            <span>Regenerate</span>
                                        </div>
                                        <div class="tool-item">
                                            <i class="i-material-symbols:chrome-restore-rounded text-5"></i>
                                            <span>Copy</span>
                                        </div>
                                        <div class="tool-item">
                                            <i class="i-ri:information-2-fill text-5"></i>
                                            <span>Info</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </n-scrollbar>
                    </n-card>
                </div>
                <div class="question">
                    <n-card style="height: 100%;">
                        <n-input type="textarea" class="question-input"></n-input>
                        <div class="send-btn">
                            <i
                                class="i-icon-park-twotone:new-picture text-8 text-[var(--color-primary-1)] hover-text-[var(--color-primary-hover-1)]"></i>
                            <i
                                class="i-icon-park-twotone:arrow-circle-up text-8 text-[var(--color-primary-1)] hover-text-[var(--color-primary-hover-1)]"></i>
                        </div>
                    </n-card>
                </div>
            </div>
            <div class="view-area">
                <n-card style="height: 100%;"></n-card>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import markdownit from "markdown-it"
    import hljs from 'highlight.js';
    import { useTemplateRef } from 'vue';
    import "./highlight.theme.css"
    const answerRef = useTemplateRef("answerRef")
    const answerHeight = computed(() => {
        return answerRef.value?.offsetHeight
    })
    const languageTit  = ref("")
    const md = markdownit({
        html: true,
        linkify: true,
        typographer: true,
        langPrefix: 'language-',
        highlight: function (str, lang): any {
            languageTit.value = lang
            if (lang && hljs.getLanguage(lang)) {
                try {
                    return hljs.highlight(lang, str).value;
                } catch (__) { console.warn(__) }
            }

            return md.utils.escapeHtml(str);// 使用额外的默认转义
        }
    })


const codeStr = `
\`\`\`vue
let a = 20;
function calc(){
    for(let i=0;i<a;i++){
        console.log(\`这是用于测试的代码-\${i}\`)
    }
}

class TestClass {
    private name:string;
    public age:nuumber

    constructor(name:string,age:number){
        this.name = name;
        this.age = age
    }
}
\`\`\`
`;
    const codeHtml = ref("")
    codeHtml.value = md.render(codeStr)
</script>

<style scoped lang="scss">
    @use "@/styles/index" as base;
    @use "../../mixin.scss";

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
            grid-template-columns: 500px 1fr;
            gap: 15px;

            .chat-area {
                display: grid;
                grid-template-rows: 1fr 230px;
                gap: 15px;

                .answer {
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
                                transition:.1s all ease-in-out;

                                &:hover {
                                    background: var(--color-primary-1);
                                    color: #fff;
                                }

                                &:active{
                                    background: var(--color-primary-3);
                                }
                            }

                            .text {
                                font-size: 14px;

                                .code-area {
                                    .code-tool {
                                        @include base.row-flex;
                                        justify-content: space-between;
                                        align-items: center;
                                        height: 34px;
                                        background: var(--color-card-text-1);
                                        color: var(--color-bg-1);
                                        padding: 0 0 0 10px;

                                        .tools {
                                            @include base.row-flex-start;
                                            gap: 10px;

                                            .tool-item {
                                                height: 34px;
                                                @include item-tool();
                                                border-radius: 0;
                                            }
                                        }
                                    }

                                    :deep(pre) {
                                        background: var(--color-menu-1);
                                        margin: 0;
                                        padding: 10px;
                                        font-size: 14px;
                                    }
                                }
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

            .view-area {
                /* background: linear-gradient(-45deg, rgba(238, 152, 82, 0.3),rgba(35, 166, 213,.3), rgba(231, 60, 60, 0.3),  rgba(35, 213, 171,.3));
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
                } */
            }
        }
    }




</style>