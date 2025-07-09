import { FormInst } from "naive-ui"

export type Provider = {
    supplierTitle: string
    supplierName: string
    baseUrl: string
    baseUrlExample: string
    isUseUrlExample: boolean
    apiKey: string
    home: string
    help: string
    status: boolean
}

export type Model = {
    supplierName: string
    modelName: string
    capability: string[]
    title: string
}

export type AddProviderFormData = {
    supplierTitle: string
    supplierName: string
    baseUrl: string
    apiKey: string
}

export type ModelStore = {
    providerList: Ref<Provider[]>
    modelList: Ref<Model[]>
    modelName: Ref<string>
    currentProvider: Ref<Provider>
    configurationLoading: Ref<boolean>
    addProviderFormData: Ref<AddProviderFormData>,
    addProviderFormDataRef: Ref<FormInst | null | undefined>
    addProviderRef: Ref<any>
}