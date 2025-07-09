import { FormInst } from "naive-ui";
import { AddProviderFormData, Model, ModelStore, Provider } from "../dto";

export function useModelManagerStore(): ModelStore {
    const providerList = ref<Provider[]>([])
    const modelList = ref<Model[]>([])
    const modelName = ref("")
    const addProviderRef = ref<any>()
    const currentProvider = ref<Provider>({
        supplierTitle: "",
        supplierName: "",
        baseUrl: "",
        baseUrlExample: "",
        isUseUrlExample: false,
        apiKey: "",
        home: "",
        help: "",
        status: false,
    })
    const configurationLoading = ref(false)
    const addProviderFormData = ref<AddProviderFormData>({
        supplierTitle: "",
        supplierName: "",
        baseUrl: "",
        apiKey: ""
    })
    const addProviderFormDataRef = ref<FormInst | null | undefined>()


    return {
        providerList,
        modelList,
        modelName,
        addProviderRef,
        currentProvider,
        configurationLoading,
        addProviderFormData,
        addProviderFormDataRef
    }
}