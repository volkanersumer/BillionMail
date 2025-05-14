"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["676"],{4944:function(e,t,a){a.d(t,{L:()=>l});function l(e,t){if(!e)return;let a=document.createElement("a");a.href=e,void 0!==t&&(a.download=t),document.body.appendChild(a),a.click(),document.body.removeChild(a)}},8439:function(e,t,a){a.d(t,{W3:()=>d,ZP:()=>p});var l=a(209),n=a(1321),r=a(4124),o=a(6169),i=a(1579),c=a(3723),s=a(2249);let u=(0,s.cB)("breadcrumb",`
 white-space: nowrap;
 cursor: default;
 line-height: var(--n-item-line-height);
`,[(0,s.c)("ul",`
 list-style: none;
 padding: 0;
 margin: 0;
 `),(0,s.c)("a",`
 color: inherit;
 text-decoration: inherit;
 `),(0,s.cB)("breadcrumb-item",`
 font-size: var(--n-font-size);
 transition: color .3s var(--n-bezier);
 display: inline-flex;
 align-items: center;
 `,[(0,s.cB)("icon",`
 font-size: 18px;
 vertical-align: -.2em;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `),(0,s.c)("&:not(:last-child)",[(0,s.cM)("clickable",[(0,s.cE)("link",`
 cursor: pointer;
 `,[(0,s.c)("&:hover",`
 background-color: var(--n-item-color-hover);
 `),(0,s.c)("&:active",`
 background-color: var(--n-item-color-pressed); 
 `)])])]),(0,s.cE)("link",`
 padding: 4px;
 border-radius: var(--n-item-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 position: relative;
 `,[(0,s.c)("&:hover",`
 color: var(--n-item-text-color-hover);
 `,[(0,s.cB)("icon",`
 color: var(--n-item-text-color-hover);
 `)]),(0,s.c)("&:active",`
 color: var(--n-item-text-color-pressed);
 `,[(0,s.cB)("icon",`
 color: var(--n-item-text-color-pressed);
 `)])]),(0,s.cE)("separator",`
 margin: 0 8px;
 color: var(--n-separator-color);
 transition: color .3s var(--n-bezier);
 user-select: none;
 -webkit-user-select: none;
 `),(0,s.c)("&:last-child",[(0,s.cE)("link",`
 font-weight: var(--n-font-weight-active);
 cursor: unset;
 color: var(--n-item-text-color-active);
 `,[(0,s.cB)("icon",`
 color: var(--n-item-text-color-active);
 `)]),(0,s.cE)("separator",`
 display: none;
 `)])])]),d=(0,i.U)("n-breadcrumb"),m=Object.assign(Object.assign({},n.Z.props),{separator:{type:String,default:"/"}}),p=(0,l.aZ)({name:"Breadcrumb",props:m,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:a}=(0,r.ZP)(e),i=(0,n.Z)("Breadcrumb","-breadcrumb",u,c.Z,e,t);(0,l.JJ)(d,{separatorRef:(0,l.Vh)(e,"separator"),mergedClsPrefixRef:t});let s=(0,l.Fl)(()=>{let{common:{cubicBezierEaseInOut:e},self:{separatorColor:t,itemTextColor:a,itemTextColorHover:l,itemTextColorPressed:n,itemTextColorActive:r,fontSize:o,fontWeightActive:c,itemBorderRadius:s,itemColorHover:u,itemColorPressed:d,itemLineHeight:m}}=i.value;return{"--n-font-size":o,"--n-bezier":e,"--n-item-text-color":a,"--n-item-text-color-hover":l,"--n-item-text-color-pressed":n,"--n-item-text-color-active":r,"--n-separator-color":t,"--n-item-color-hover":u,"--n-item-color-pressed":d,"--n-item-border-radius":s,"--n-font-weight-active":c,"--n-item-line-height":m}}),m=a?(0,o.F)("breadcrumb",void 0,s,e):void 0;return{mergedClsPrefix:t,cssVars:a?void 0:s,themeClass:null==m?void 0:m.themeClass,onRender:null==m?void 0:m.onRender}},render(){var e;return null===(e=this.onRender)||void 0===e||e.call(this),(0,l.h)("nav",{class:[`${this.mergedClsPrefix}-breadcrumb`,this.themeClass],style:this.cssVars,"aria-label":"Breadcrumb"},(0,l.h)("ul",null,this.$slots))}})},6666:function(e,t,a){a.d(t,{Z:()=>i});var l=a(209),n=a(8282),r=a(4934),o=a(8439);let i=(0,l.aZ)({name:"BreadcrumbItem",props:{separator:String,href:String,clickable:{type:Boolean,default:!0},onClick:Function},slots:Object,setup(e,{slots:t}){let a=(0,l.f3)(o.W3,null);if(!a)return()=>null;let{separatorRef:i,mergedClsPrefixRef:c}=a,s=function(e=r.j?window:null){let t=()=>{let{hash:t,host:a,hostname:l,href:n,origin:r,pathname:o,port:i,protocol:c,search:s}=(null==e?void 0:e.location)||{};return{hash:t,host:a,hostname:l,href:n,origin:r,pathname:o,port:i,protocol:c,search:s}},a=(0,l.iH)(t()),n=()=>{a.value=t()};return(0,l.bv)(()=>{e&&(e.addEventListener("popstate",n),e.addEventListener("hashchange",n))}),(0,l.SK)(()=>{e&&(e.removeEventListener("popstate",n),e.removeEventListener("hashchange",n))}),a}(),u=(0,l.Fl)(()=>e.href?"a":"span"),d=(0,l.Fl)(()=>s.value.href===e.href?"location":null);return()=>{let{value:a}=c;return(0,l.h)("li",{class:[`${a}-breadcrumb-item`,e.clickable&&`${a}-breadcrumb-item--clickable`]},(0,l.h)(u.value,{class:`${a}-breadcrumb-item__link`,"aria-current":d.value,href:e.href,onClick:e.onClick},t),(0,l.h)("span",{class:`${a}-breadcrumb-item__separator`,"aria-hidden":"true"},(0,n.gI)(t.separator,()=>{var t;return[null!==(t=e.separator)&&void 0!==t?t:i.value]})))}}})},1555:function(e,t,a){a.d(t,{A8:()=>n,NI:()=>r,QB:()=>c,WF:()=>i,Xc:()=>u,ef:()=>s,ij:()=>o});var l=a(4203);function n(e){return l.e.get("/email_template/list",{params:e})}function r(){return l.e.get("/email_template/all")}function o(e){return l.e.get("/email_template/get",{params:e})}function i(e){return l.e.post("/email_template/create",e,{fetchOptions:{loading:"Creating template, please wait...",successMessage:!0}})}function c(e){return l.e.post("/email_template/update",e,{fetchOptions:{loading:"Updating template, please wait...",successMessage:!0}})}function s(e){return l.e.post("/email_template/copy",e,{fetchOptions:{loading:"Copying template, please wait...",successMessage:!0}})}function u(e){return l.e.post("/email_template/delete",e,{fetchOptions:{loading:"Deleting template, please wait...",successMessage:!0}})}},2550:function(e,t,a){a.d(t,{Z:()=>i});var l=a(8574),n=a(209),r=a(6768);let o=(0,n.aZ)({__name:"index",props:(0,n.Vf)({language:{type:String,default:"html"},theme:{type:String,default:"vs-dark"},options:{type:Object,default:()=>({})}},{value:{default:""},valueModifiers:{}}),emits:(0,n.Vf)(["save"],["update:value"]),setup(e,t){let{expose:a,emit:o}=t,i=(0,n.tT)(e,"value"),c=(0,n.iH)(null),s=null,u=(0,n.iH)(!1);r.Z.config({paths:{vs:`${location.origin}/static/plugin/monaco/min/vs`}});let d=async()=>{if(c.value&&!s)try{u.value=!0;let t=await r.Z.init();(s=t.editor.create(c.value,{value:i.value,language:e.language,theme:e.theme,minimap:{enabled:!1},automaticLayout:!0,scrollBeyondLastLine:!1,fontSize:14,tabSize:2,wordWrap:"on",...e.options})).onDidChangeModelContent(()=>{let e=(null==s?void 0:s.getValue())||"";i.value=e}),s.addCommand(t.KeyMod.CtrlCmd|t.KeyCode.KeyS,function(){o("save")})}finally{u.value=!1}},m=()=>{s&&(s.dispose(),s=null)};return(0,n.YP)(()=>i.value,e=>{s&&e!==s.getValue()&&s.setValue(e)}),(0,n.bv)(()=>{d()}),(0,n.Jd)(()=>{m()}),a({getEditor:()=>s}),(e,t)=>{let a=l.Z;return(0,n.wg)(),(0,n.j4)(a,{class:"h-full",show:(0,n.SU)(u)},{default:(0,n.w5)(()=>[(0,n._)("div",{ref_key:"containerRef",ref:c,class:"monaco-editor-container"},null,512)]),_:1},8,["show"])}}}),i=(0,a(5115).default)(o,[["__scopeId","data-v-5b65e670"]])},9217:function(e,t,a){a.d(t,{Z:()=>r});var l=a(209);let n=["height","srcdoc"],r=(0,l.aZ)({__name:"index",props:{value:{type:String,default:""},height:{type:[String,Number],default:"100%"}},setup:e=>(t,a)=>((0,l.wg)(),(0,l.iD)("iframe",{width:"100%",height:e.height,srcdoc:e.value,style:{border:"none",borderRadius:"4px"}},null,8,n))})},7283:function(e,t,a){a.r(t),a.d(t,{default:()=>z});var l=a(3447),n=a(3951),r=a(6528),o=a(9217),i=a(2550),c=a(6714),s=a(6883),u=a(8454),d=a(8439),m=a(6666),p=a(209),v=a(2965),f=a(1555),h=a(5e3);let b={class:"template-edit"},g={class:"w-400px"},w={class:"editor-area"},_={class:"mb-16px"},k={class:"flex-1 overflow-hidden"},y={class:"preview-title"},x={class:"flex-1 overflow-hidden"},Z={class:"action-buttons"},C=(0,p.aZ)({__name:"edit",setup(e){let t=(0,h.yj)(),a=(0,p.iH)(`${t.params.id}`),C=(0,p.qj)({id:0,temp_name:"",content:""}),z=e=>{let t=new FileReader;t.onload=e=>{var t;C.content=null===(t=e.target)||void 0===t?void 0:t.result},e.file&&t.readAsText(e.file)},S=async()=>{await (0,f.QB)({id:C.id,temp_name:C.temp_name,content:C.content})},W=()=>{window.history.back()};return(async()=>{let e=await (0,f.ij)({id:a.value});(0,v.Kn)(e)&&(C.id=e.id,C.temp_name=e.temp_name,C.content=e.content)})(),(e,t)=>{let a=(0,p.up)("router-link"),v=m.Z,f=d.ZP,h=u.Z,U=s.ZP,$=c.Z,B=i.Z,E=o.Z,j=r.ZP,P=n.Z,L=l.ZP;return(0,p.wg)(),(0,p.iD)("div",b,[(0,p.Wm)(f,{class:"mb-20px"},{default:(0,p.w5)(()=>[(0,p.Wm)(v,null,{default:(0,p.w5)(()=>[(0,p.Wm)(a,{to:"/market/template"},{default:(0,p.w5)(()=>[(0,p.Uk)((0,p.zw)(e.$t("market.template.title")),1)]),_:1})]),_:1}),(0,p.Wm)(v,null,{default:(0,p.w5)(()=>[(0,p.Uk)((0,p.zw)((0,p.SU)(C).temp_name),1)]),_:1})]),_:1}),(0,p.Wm)(P,{class:"template-form","label-placement":"top"},{default:(0,p.w5)(()=>[(0,p.Wm)(U,{label:e.$t("market.template.name")},{default:(0,p.w5)(()=>[(0,p._)("div",g,[(0,p.Wm)(h,{value:(0,p.SU)(C).temp_name,"onUpdate:value":t[0]||(t[0]=e=>(0,p.SU)(C).temp_name=e)},null,8,["value"])])]),_:1},8,["label"]),(0,p.Wm)(U,{class:"editor-form-item",label:e.$t("market.template.content"),"show-feedback":!1},{default:(0,p.w5)(()=>[(0,p._)("div",w,[(0,p._)("div",_,[(0,p.Wm)($,{mode:"button","is-upload":!1,accept:["html"],onChange:z},{default:(0,p.w5)(()=>[(0,p.Uk)((0,p.zw)(e.$t("market.template.uploadHtml")),1)]),_:1})]),(0,p._)("div",k,[(0,p.Wm)(B,{value:(0,p.SU)(C).content,"onUpdate:value":t[1]||(t[1]=e=>(0,p.SU)(C).content=e),language:"html"},null,8,["value"])])]),(0,p.Wm)(j,{class:"preview-area","content-style":{display:"flex",flexDirection:"column"}},{default:(0,p.w5)(()=>[(0,p._)("div",y,(0,p.zw)(e.$t("common.actions.preview")),1),(0,p._)("div",x,[(0,p.Wm)(E,{value:(0,p.SU)(C).content},null,8,["value"])])]),_:1})]),_:1},8,["label"])]),_:1}),(0,p._)("div",Z,[(0,p.Wm)(L,{type:"primary",onClick:S},{default:(0,p.w5)(()=>[(0,p.Uk)((0,p.zw)(e.$t("common.actions.save")),1)]),_:1}),(0,p.Wm)(L,{class:"ml-16px",secondary:"",onClick:W},{default:(0,p.w5)(()=>[(0,p.Uk)((0,p.zw)(e.$t("common.actions.back")),1)]),_:1})])])}}}),z=(0,a(5115).default)(C,[["__scopeId","data-v-74f4ab26"]])}}]);