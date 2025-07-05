import { instance } from "@/api";
import { getEditDomainStoreData } from "../store";
import { instanceOptions } from "./companyProfile.controller"
/**
 * @description Get styleing info
 * 
 * @param { string } domain
 */
export async function getStylingInfo(domain: string) {
    const {
        accent_color,
        text_color,
        page_background,
        container_background,
        link_social_color,
        link_footer_color,
        heading_font,
        body_font,
        hasGotStyleing
    } = getEditDomainStoreData()
    if (hasGotStyleing.value) return
    try {
        const res = await instance.post("/askai/project/get_style_config", { domain }, instanceOptions) as Record<string, any>
        accent_color.value = res.accent_color
        text_color.value = res.text_color
        page_background.value = res.page_background
        container_background.value = res.container_background
        link_social_color.value = res.link_social_color
        link_footer_color.value = res.link_footer_color
        heading_font.value = res.heading_font
        body_font.value = res.body_font

        hasGotStyleing.value = true
    } catch (error) {
        console.warn(error)
    }
}

/**
 * @description Update styleing info
 */
export async function updateStyleingInfo(domain: string) {
    const {
        accent_color,
        text_color,
        page_background,
        container_background,
    } = getEditDomainStoreData()
    try {
        await instance.post("/askai/project/modify_style_config", {
            domain,
            accent_color: accent_color.value,
            text_color: text_color.value,
            page_background: page_background.value,
            container_background: container_background.value,
        }, instanceOptions) as Record<string, any>
    } catch (error) {
        console.warn(error)
    }
}