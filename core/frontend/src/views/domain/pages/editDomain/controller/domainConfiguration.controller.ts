import { instance } from "@/api"
import { instanceOptions } from "./companyProfile.controller"
import { getEditDomainStoreData } from "../store"
import { getByteUnit } from '@/utils'


/**
 * @description Get domain detail
 */
export async function getDomainDetail(domain: string) {
    const {
        domainTit,
        quota,
        unit,
        mailboxes,
        catch_email,
        hasGotDomainConfiguration,
        urls
    } = getEditDomainStoreData()
    if (hasGotDomainConfiguration.value) return
    try {
        const { list } = await instance.get("/domains/list", { params: { keyword: domain, page: 1, page_size: 1 }, ...instanceOptions }) as { total: number, list: Record<string, any>[] }
        const quotaAndUnit = getByteUnit(list[0].quota, true, 2).split(" ")
        domainTit.value = list[0].domain
        quota.value = quotaAndUnit[0]
        unit.value = quotaAndUnit[1]
        mailboxes.value = list[0].mailboxes
        catch_email.value = list[0].email
        urls.value = list[0].urls
        await configurationStatus()
        hasGotDomainConfiguration.value = true
    } catch (error) {
        console.warn(error)
    }
}

/**
 * @description Configuration status
 */
export async function configurationStatus() {
    const { configurationStatus } = getEditDomainStoreData()
    try {
        const res = await instance.post("/askai/supplier/status") as Record<string, boolean>
        configurationStatus.value = res.is_configured
    } catch (error) {
        console.warn(error)
    }
}

/**
 * @description Update aomain 
 */
export async function updateDomain() {
    const {
        domainTit,
        quota,
        unit,
        urls,
        mailboxes,
        catch_email
    } = getEditDomainStoreData()
    const quotaParam = calcBaseUnits(Number(quota.value), unit.value)
    try {
        await instance.post("/domains/update", {
            domain: domainTit.value,
            mailboxes: mailboxes.value,
            mailboxQuota: quotaParam,
            quota: quotaParam,
            rateLimit: 1,
            active: 1,
            email: catch_email.value,
            urls: urls.value
        }, instanceOptions)
    } catch (error) {
        console.warn(error)
    }

}

/**
 * @description Calculate based on units
 */
function calcBaseUnits(quota: number, unit: string) {
    switch (unit) {
        case "KB":
            return quota * 1024
        case "MB":
            return quota * 1024 * 1024
        case "GB":
            return quota * 1024 * 1024 * 1024
        case "TB":
            return quota * 1024 * 1024 * 1024 * 1024
        default:
            return quota
    }
}

/**
 * @description Sync domain to url
 */
export function syncToUrl(domainVal: string) {
    const { urls } = getEditDomainStoreData()
    const httpStr = urls.value[0].match(/http:\/\/?/g)
    const httpsStr = urls.value[0].match(/https:\/\/?/g)
    if(httpStr){
        urls.value[0] = httpStr[0] + domainVal
    }

    if(httpsStr){
        urls.value[0] = httpsStr[0] + domainVal
    }

    if(!httpStr && !httpsStr){
        urls.value[0] = 'https://' + domainVal
    }
}

/**
 * @description Remove url
 */
export function removeUrl(index: number) {
    const { urls } = getEditDomainStoreData()
    
    urls.value.splice(index, 1)
}

/**
 * @description Reset all api status
 */
export function resetAllApiStatus() {
    const { hasGotAisettings, hasGotCompanyProfile, hasGotDomainConfiguration, hasGotFootersettings, hasGotProjectDetail, hasGotSitemap, hasGotStyling } = getEditDomainStoreData()
    hasGotAisettings.value = false
    hasGotCompanyProfile.value = false
    hasGotDomainConfiguration.value = false
    hasGotFootersettings.value = false
    hasGotProjectDetail.value = false
    hasGotSitemap.value = false
    hasGotStyling.value = false
}