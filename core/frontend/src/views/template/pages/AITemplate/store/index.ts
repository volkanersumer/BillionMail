import { ChatInfo, Model, TemplateStore } from "../dto"

export function useTemplateStore(): TemplateStore {
    const sourceDomain = ref("")
    const chatId = ref("")
    const questionContent = ref("")
    const chatInfo = ref<ChatInfo>({
        chatId: "",
        create_time: 0,
        files: [],
        messages: [],
        modelId: "",
        prompt: "",
        supplierName: "",
        title: "",
        update_time: 0,
    })
    const currentChatRecordKey = ref<string>("")
    const modelList = ref<Model[]>([])
    const currentModel = ref<Model>({
        modelId: "",
        supplierName: "",
        capability: [],
        title: "",
        status: false,
        max_tokens: 0
    })
    const currentModelTitle = ref<string>("")
    const answerContent = ref<string[]>([])
    const previewCode = ref<string>("")
    const generateShow = ref<boolean>(false)
    const chatRecord = ref<Map<string, string[]>>(new Map())
    const isChat = ref(false)
    const chatScrollRef = ref()
    const scrollWrapperRef = ref()
    const scrollable = ref(true)

    return {
        sourceDomain,
        chatId,
        questionContent,
        chatInfo,
        currentChatRecordKey,
        modelList,
        currentModel,
        currentModelTitle,
        answerContent,
        previewCode,
        generateShow,
        chatRecord,
        isChat,
        chatScrollRef,
        scrollWrapperRef,
        scrollable
    }
}