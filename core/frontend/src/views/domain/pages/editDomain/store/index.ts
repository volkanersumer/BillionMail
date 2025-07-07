import { defineStore, storeToRefs } from "pinia";
import { KnowledgeBase, SiteInfo } from "../dto";

const editDomainStore = defineStore("editDomainStoroe", () => {
    // domain configuration
    const domainTit = ref("")
    const quota = ref("0")
    const unit = ref("B")
    const mailboxes = ref(0)
    const catch_email = ref("")

    // Project detail
    const hasGotProjectDetail = ref(false)
    const domain = ref("")
    const description = ref("")
    const favicon = ref("")
    const industry = ref("")
    const primary_logo = ref("")
    const project_name = ref("")
    const secondary_logo = ref("")
    const knowledge_base = ref<KnowledgeBase[]>([])
    const activeKnowledge = ref<KnowledgeBase>({
        title: "",
        content: "",
        kid: "",
        update_timme: 0
    })
    const knowledgeTitle = ref("")
    const knowledgeContent = ref("")
    const knowledgeModalShow = ref(false)
    const isEditKnowledge = ref(false)

    // Company profile
    const hasGotCompanyProfile = ref(false)
    const legal_company_name = ref("")
    const web_site = ref("")
    const company_profile = ref("")
    const email = ref("")
    const phone = ref("")
    const support_url = ref("")

    // Styleing
    const hasGotStyleing = ref(false)
    const accent_color = ref("")
    const text_color = ref("")
    const page_background = ref("")
    const container_background = ref("")
    const link_social_color = ref("")
    const link_footer_color = ref("")
    const heading_font = ref("")
    const body_font = ref("")

    // Sitemap
    const hasGotSitemap = ref(false)
    const sitemap = ref<SiteInfo[]>([
        {
            title: "",
            uri_path: ""
        }
    ])

    // Footer settings
    const hasGotFootersettings = ref(false)
    const copyright_text = ref("")
    const disclaimer = ref("")

    // Ai settings
    const hasGotAisettings = ref(false)
    const prompt = ref("")

    return {
        domainTit,
        quota,
        unit,
        mailboxes,
        catch_email,
        
        hasGotProjectDetail,
        domain,
        description,
        favicon,
        industry,
        primary_logo,
        project_name,
        secondary_logo,
        knowledge_base,
        activeKnowledge,
        knowledgeTitle,
        knowledgeContent,
        knowledgeModalShow,
        isEditKnowledge,

        hasGotCompanyProfile,
        legal_company_name,
        web_site,
        company_profile,
        email,
        phone,
        support_url,

        hasGotStyleing,
        accent_color,
        text_color,
        page_background,
        container_background,
        link_social_color,
        link_footer_color,
        heading_font,
        body_font,

        hasGotSitemap,
        sitemap,

        hasGotFootersettings,
        copyright_text,
        disclaimer,

        hasGotAisettings,
        prompt
    }
})

export function getEditDomainStoreData() {
    const store = editDomainStore()
    return {
        ...storeToRefs(store)
    }
}