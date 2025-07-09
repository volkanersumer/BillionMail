import { instance } from "@/api"
import { Model, ModelStore, Provider } from "../dto"
export const instanceOptions = {
    fetchOptions: {
        loading: 'Loading... Please wait.',
        successMessage: true,
    }
}


/**
 * @description Get provider list
 */
export async function getProviderList(modelStore: ModelStore) {
    const { providerList, currentProvider } = modelStore
    try {
        const res = await instance.post("/askai/supplier/list") as Provider[];
        providerList.value = res
        if (res.length) {
            currentProvider.value = res[0]
        }
    } catch (error) {
        console.warn(error)
    }
}

/**
 * @description Get model list
 */
export async function getModelList(supplierName: string, modelStore: ModelStore) {
    const { modelList, modelName, configurationLoading } = modelStore
    try {
        configurationLoading.value = true
        const res = await instance.post("/askai/supplier/models", { supplierName }) as Model[]
        modelList.value = res
        modelName.value = res[0] ? res[0].title : ""
    } catch (error) {
        console.warn(error)
    } finally {
        configurationLoading.value = false
    }
}

/***
 * @description Confirm add provider
 */
export async function confirmAddProvider(modelStore: ModelStore) {
    const { addProviderFormData, addProviderFormDataRef, addProviderRef } = modelStore
    try {
        await addProviderFormDataRef.value?.validate()
        await instance.post("/askai/supplier/add_supplier", addProviderFormData.value, instanceOptions)
        addProviderRef.value.close()
        await getProviderList(modelStore)
    } catch (error) {
        console.warn(error)
    }
}