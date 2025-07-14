export type ChatInfo = {
    chatId: string
    create_time: number
    files: string[]
    messages: ChatMessage[]
    modelId: string
    prompt: string
    supplierName: string
    title: string
    update_time: number
}

export type ChatMessage = {
    chatId: string
    content: string
    create_time: number
    end_time: number
    finish_reason: string
    html_content: string
    messageId: string
    model_id: string
    ppid: string
    reasoning: string
    role: string
    supplier_name: string
    time_consuming: number
    usage: {
        completion_tokens: number
        prompt_tokens: number
        total_tokens: number
    }
}


export type Model = {
    modelId: string
    supplierName: string
    capability: string[]
    title: string
    status: boolean
    max_tokens: number
}

export type SSEContent = {
    chatId: string
    content: string
    is_end: boolean
    is_html: boolean
    reasoning: string
    role: string
    update_time: number
}


export type TemplateStore = {
    sourceDomain: Ref<string>
    chatId: Ref<string>
    questionContent: Ref<string>
    chatInfo: Ref<ChatInfo>
    currentChatRecordKey: Ref<string>
    modelList: Ref<Model[]>
    currentModel: Ref<Model>
    currentModelTitle: Ref<string>
    answerContent: Ref<string[]>
    previewCode: Ref<string>
    generateShow: Ref<boolean>
    chatRecord: Ref<Map<string, string[]>>
}