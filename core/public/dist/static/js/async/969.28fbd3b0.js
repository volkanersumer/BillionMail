"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["969"],{65977:function(e,t,a){a.d(t,{Z:()=>i});var l=a(58786);let i=(0,l.aZ)({name:"Add",render:()=>(0,l.h)("svg",{width:"512",height:"512",viewBox:"0 0 512 512",fill:"none",xmlns:"http://www.w3.org/2000/svg"},(0,l.h)("path",{d:"M256 112V400M400 256H112",stroke:"currentColor","stroke-width":"32","stroke-linecap":"round","stroke-linejoin":"round"}))})},60309:function(e,t,a){a.d(t,{L:()=>l});function l(e,t){if(!e)return;let a=document.createElement("a");a.href=e,void 0!==t&&(a.download=t),document.body.appendChild(a),a.click(),document.body.removeChild(a)}},79537:function(e,t,a){a.d(t,{Z:()=>n});var l=a(58786),i=a(54470),r=a(27046),s=a(71309);let o=(0,s.cB)("input-group",`
 display: inline-flex;
 width: 100%;
 flex-wrap: nowrap;
 vertical-align: bottom;
`,[(0,s.c)(">",[(0,s.cB)("input",[(0,s.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,s.c)("&:not(:first-child)",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 margin-left: -1px!important;
 `)]),(0,s.cB)("button",[(0,s.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `,[(0,s.cE)("state-border, border",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `)]),(0,s.c)("&:not(:first-child)",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `,[(0,s.cE)("state-border, border",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `)])]),(0,s.c)("*",[(0,s.c)("&:not(:last-child)",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `,[(0,s.c)(">",[(0,s.cB)("input",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,s.cB)("base-selection",[(0,s.cB)("base-selection-label",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,s.cB)("base-selection-tags",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `),(0,s.cE)("box-shadow, border, state-border",`
 border-top-right-radius: 0!important;
 border-bottom-right-radius: 0!important;
 `)])])]),(0,s.c)("&:not(:first-child)",`
 margin-left: -1px!important;
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `,[(0,s.c)(">",[(0,s.cB)("input",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,s.cB)("base-selection",[(0,s.cB)("base-selection-label",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,s.cB)("base-selection-tags",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `),(0,s.cE)("box-shadow, border, state-border",`
 border-top-left-radius: 0!important;
 border-bottom-left-radius: 0!important;
 `)])])])])])]),n=(0,l.aZ)({name:"InputGroup",props:{},setup(e){let{mergedClsPrefixRef:t}=(0,i.ZP)(e);return(0,r.Z)("-input-group",o,t),{mergedClsPrefix:t}},render(){let{mergedClsPrefix:e}=this;return(0,l.h)("div",{class:`${e}-input-group`},this.$slots)}})},60579:function(e,t,a){a.d(t,{Z:()=>f});var l=a(65083),i=a(23125),r=a(58786),s=a(62594),o=a(56946),n=a(54470),p=a(53198),u=a(71309),d=a(53573),c=a(66480);let m=(0,u.c)([(0,u.c)("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),(0,u.cB)("spin-container",`
 position: relative;
 `,[(0,u.cB)("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[(0,c.h)()])]),(0,u.cB)("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),(0,u.cB)("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[(0,u.cM)("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),(0,u.cB)("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),(0,u.cB)("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[(0,u.cM)("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),v={small:20,medium:18,large:16},b=Object.assign(Object.assign({},o.Z.props),{contentClass:String,contentStyle:[Object,String],description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number}),f=(0,r.aZ)({name:"Spin",props:b,slots:Object,setup(e){let{mergedClsPrefixRef:t,inlineThemeDisabled:a}=(0,n.ZP)(e),s=(0,o.Z)("Spin","-spin",m,d.Z,e,t),c=(0,r.Fl)(()=>{let{size:t}=e,{common:{cubicBezierEaseInOut:a},self:i}=s.value,{opacitySpinning:r,color:o,textColor:n}=i;return{"--n-bezier":a,"--n-opacity-spinning":r,"--n-size":"number"==typeof t?(0,l.BL)(t):i[(0,u.Tl)("size",t)],"--n-color":o,"--n-text-color":n}}),b=a?(0,p.F)("spin",(0,r.Fl)(()=>{let{size:t}=e;return"number"==typeof t?String(t):t[0]}),c,e):void 0,f=(0,i.Z)(e,["spinning","show"]),h=(0,r.iH)(!1);return(0,r.m0)(t=>{let a;if(f.value){let{delay:l}=e;if(l){a=window.setTimeout(()=>{h.value=!0},l),t(()=>{clearTimeout(a)});return}}h.value=f.value}),{mergedClsPrefix:t,active:h,mergedStrokeWidth:(0,r.Fl)(()=>{let{strokeWidth:t}=e;if(void 0!==t)return t;let{size:a}=e;return v["number"==typeof a?"medium":a]}),cssVars:a?void 0:c,themeClass:null==b?void 0:b.themeClass,onRender:null==b?void 0:b.onRender}},render(){var e,t;let{$slots:a,mergedClsPrefix:l,description:i}=this,o=a.icon&&this.rotate,n=(i||a.description)&&(0,r.h)("div",{class:`${l}-spin-description`},i||(null==(e=a.description)?void 0:e.call(a))),p=a.icon?(0,r.h)("div",{class:[`${l}-spin-body`,this.themeClass]},(0,r.h)("div",{class:[`${l}-spin`,o&&`${l}-spin--rotate`],style:a.default?"":this.cssVars},a.icon()),n):(0,r.h)("div",{class:[`${l}-spin-body`,this.themeClass]},(0,r.h)(s.Z,{clsPrefix:l,style:a.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${l}-spin`}),n);return null==(t=this.onRender)||t.call(this),a.default?(0,r.h)("div",{class:[`${l}-spin-container`,this.themeClass],style:this.cssVars},(0,r.h)("div",{class:[`${l}-spin-content`,this.active&&`${l}-spin-content--spinning`,this.contentClass],style:this.contentStyle},a),(0,r.h)(r.uT,{name:"fade-in-transition"},{default:()=>this.active?p:null})):p}})},50144:function(e,t,a){a.d(t,{Z:()=>p});var l=a(58786),i=a(56946),r=a(54470),s=a(49942),o=a(44108);let n=Object.assign(Object.assign({},s.Kd),i.Z.props),p=(0,l.aZ)({name:"Tooltip",props:n,slots:Object,__popover__:!0,setup(e){let{mergedClsPrefixRef:t}=(0,r.ZP)(e),a=(0,i.Z)("Tooltip","-tooltip",void 0,o.Z,e,t),s=(0,l.iH)(null);return Object.assign(Object.assign({},{syncPosition(){s.value.syncPosition()},setShow(e){s.value.setShow(e)}}),{popoverRef:s,mergedTheme:a,popoverThemeOverrides:(0,l.Fl)(()=>a.value.self)})},render(){let{mergedTheme:e,internalExtraClass:t}=this;return(0,l.h)(s.ZP,Object.assign(Object.assign({},this.$props),{theme:e.peers.Popover,themeOverrides:e.peerOverrides.Popover,builtinThemeOverrides:this.popoverThemeOverrides,internalExtraClass:t.concat("tooltip"),ref:"popoverRef"}),this.$slots)}})},77713:function(e,t,a){a.r(t),a.d(t,{default:()=>et});var l=a(57869),i=a(58786),r=a(60579),s=a(98048),o=a(6455),n=a(79537),p=a(73084),u=a(29875),d=a(15496),c=a(40972),m=a(5057),v=a(34541);let b={fetchOptions:{loading:"Loading... Please wait.",successMessage:!0}};async function f(e){let{providerList:t,currentProvider:a}=e;try{let l=await v.e.post("/askai/supplier/list");t.value=l,l.length&&(a.value=l[0],await h(a.value.supplierName,e))}catch(e){console.warn(e)}}async function h(e,t){let{modelList:a,configurationLoading:l}=t;try{l.value=!0,a.value=await v.e.post("/askai/supplier/models",{supplierName:e})}catch(e){console.warn(e)}finally{l.value=!1}}async function g(e){let{addProviderFormData:t,addProviderFormDataRef:a,addProviderRef:l}=e;try{var i;await (null==(i=a.value)?void 0:i.validate()),await v.e.post("/askai/supplier/add_supplier",t.value,b),l.value.close(),await f(e)}catch(e){console.warn(e)}}async function _(e){let{currentProvider:t,addModelFormData:a,addModelFormRef:l,addModelRef:i}=e;try{var r;console.log(a),await (null==(r=l.value)?void 0:r.validate()),await v.e.post("/askai/supplier/add_model",{...a.value,supplierName:t.value.supplierName},b),await h(t.value.supplierName,e),i.value.close()}catch(e){console.warn(e)}}async function w(e){let{currentProvider:t,providerList:a}=e;try{let e=await v.e.post("/askai/supplier/get_supplier_config",{supplierName:t.value.supplierName},b),l=a.value.find(e=>e.supplierName==t.value.supplierName);l&&(l.apiKey=e.apiKey,l.baseUrl=e.baseUrl)}catch(e){console.warn(e)}}async function y(e,t){let{currentProvider:a}=t;try{await v.e.post("/askai/supplier/set_supplier_status",{supplierName:a.value.supplierName,status:e},b)}catch(e){console.warn(e)}}async function U(e){let{currentProvider:t}=e;try{await v.e.post("/askai/supplier/testing",{supplierName:t.value.supplierName,baseUrl:t.value.baseUrl,apiKey:t.value.apiKey},b)}catch(e){console.warn(e)}}async function k(e){let{currentProvider:t}=e;try{await v.e.post("/askai/supplier/set_supplier_config",{supplierName:t.value.supplierName,baseUrl:t.value.baseUrl,apiKey:t.value.apiKey},b)}catch(e){console.warn(e)}}async function S(e,t,a){let{currentProvider:l}=a;try{await v.e.post("/askai/supplier/set_model_status",{supplierName:l.value.supplierName,modelId:e,status:t},b)}catch(e){console.warn(e)}}async function x(e){let{currentProvider:t}=e;try{await v.e.post("/askai/supplier/remove_supplier",{supplierName:t.value.supplierName},b),await f(e)}catch(e){console.warn(e)}}let Z={class:"flex justify-end gap-5"},W=(0,i.aZ)({__name:"AddProvider",setup(e,t){let{expose:a}=t,l=(0,i.f3)("modelStore"),{addProviderFormData:r,addProviderFormDataRef:n}=l,u=(0,i.iH)(!1),c={supplierTitle:[{required:!0,message:"supplierTitle is required",trigger:"blur"}],supplierName:[{required:!0,message:"supplierName is required",trigger:"blur"}],baseUrl:[{required:!0,message:"baseUrl is required",trigger:"blur"}],apiKey:[{required:!0,message:"apiKey is required",trigger:"blur"}]};return a({open:function(){u.value=!0},close:function(){u.value=!1}}),(e,t)=>{let a=p.Z,v=o.ZP,b=s.Z,f=d.ZP,h=m.Z;return(0,i.wg)(),(0,i.j4)(h,{show:(0,i.SU)(u),"onUpdate:show":t[5]||(t[5]=e=>(0,i.dq)(u)?u.value=e:null),preset:"card","close-on-esc":!1,"mask-closable":!1,draggable:"",class:"w-150",title:"Add Provider"},{footer:(0,i.w5)(()=>[(0,i._)("div",Z,[(0,i.Wm)(f,null,{default:(0,i.w5)(()=>t[6]||(t[6]=[(0,i.Uk)("Cancel")])),_:1,__:[6]}),(0,i.Wm)(f,{type:"primary",onClick:t[4]||(t[4]=e=>(0,i.SU)(g)((0,i.SU)(l)))},{default:(0,i.w5)(()=>t[7]||(t[7]=[(0,i.Uk)("Confirm")])),_:1,__:[7]})])]),default:(0,i.w5)(()=>[(0,i.Wm)(b,{ref_key:"addProviderFormDataRef",ref:n,model:(0,i.SU)(r),rules:c},{default:(0,i.w5)(()=>[(0,i.Wm)(v,{label:"供应商标题",path:"supplierTitle"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).supplierTitle,"onUpdate:value":t[0]||(t[0]=e=>(0,i.SU)(r).supplierTitle=e)},null,8,["value"])]),_:1}),(0,i.Wm)(v,{label:"供应商名称",path:"supplierName"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).supplierName,"onUpdate:value":t[1]||(t[1]=e=>(0,i.SU)(r).supplierName=e)},null,8,["value"])]),_:1}),(0,i.Wm)(v,{label:"接口地址",path:"baseUrl"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).baseUrl,"onUpdate:value":t[2]||(t[2]=e=>(0,i.SU)(r).baseUrl=e)},null,8,["value"])]),_:1}),(0,i.Wm)(v,{label:"密钥",path:"apiKey"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).apiKey,"onUpdate:value":t[3]||(t[3]=e=>(0,i.SU)(r).apiKey=e)},null,8,["value"])]),_:1})]),_:1},8,["model"])]),_:1},8,["show"])}}});var N=a(29283),C=a(8490);let P={class:"flex justify-end gap-5"},B=(0,i.aZ)({__name:"AddModel",setup(e,t){let{expose:a}=t,l=(0,i.f3)("modelStore"),{addModelFormData:r,addModelFormRef:n}=l,u=(0,i.iH)(!1),c=[{label:"llm",value:"llm"},{label:"vision",value:"vision"},{label:"tools",value:"tools"},{label:"text-to-image",value:"text-to-image"}],v={modelId:[{required:!0,message:"modelId is required",trigger:"blur"}],title:[{required:!0,message:"title is required",trigger:"blur"}],capability:[{validator:(e,t)=>!!t.length||Error("capability is required"),trigger:"blur"}],max_tokens:[{validator:(e,t)=>!!t||Error("max_tokens is required"),trigger:"blur"}]};function b(){u.value=!1,r.value.capability=[],r.value.max_tokens=0,r.value.modelId="",r.value.title=""}return a({open:function(){u.value=!0},close:b}),(e,t)=>{let a=p.Z,f=o.ZP,h=C.Z,g=N.Z,w=s.Z,y=d.ZP,U=m.Z;return(0,i.wg)(),(0,i.j4)(U,{show:(0,i.SU)(u),"onUpdate:show":t[5]||(t[5]=e=>(0,i.dq)(u)?u.value=e:null),preset:"card",draggable:"","close-on-esc":!1,"mask-closable":!1,title:"Add Model",class:"w-100"},{footer:(0,i.w5)(()=>[(0,i._)("div",P,[(0,i.Wm)(y,{onClick:b},{default:(0,i.w5)(()=>t[6]||(t[6]=[(0,i.Uk)("Cancel")])),_:1,__:[6]}),(0,i.Wm)(y,{type:"primary",onClick:t[4]||(t[4]=e=>(0,i.SU)(_)((0,i.SU)(l)))},{default:(0,i.w5)(()=>t[7]||(t[7]=[(0,i.Uk)("Confirm")])),_:1,__:[7]})])]),default:(0,i.w5)(()=>[(0,i.Wm)(w,{ref_key:"addModelFormRef",ref:n,model:(0,i.SU)(r),rules:v},{default:(0,i.w5)(()=>[(0,i.Wm)(f,{label:"模型ID",path:"modelId"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).modelId,"onUpdate:value":t[0]||(t[0]=e=>(0,i.SU)(r).modelId=e)},null,8,["value"])]),_:1}),(0,i.Wm)(f,{label:"模型别名",path:"title"},{default:(0,i.w5)(()=>[(0,i.Wm)(a,{value:(0,i.SU)(r).title,"onUpdate:value":t[1]||(t[1]=e=>(0,i.SU)(r).title=e)},null,8,["value"])]),_:1}),(0,i.Wm)(f,{label:"模型功能",path:"capability"},{default:(0,i.w5)(()=>[(0,i.Wm)(h,{value:(0,i.SU)(r).capability,"onUpdate:value":t[2]||(t[2]=e=>(0,i.SU)(r).capability=e),options:c,multiple:""},null,8,["value"])]),_:1}),(0,i.Wm)(f,{label:"最大上下文长度",path:"max_tokens"},{default:(0,i.w5)(()=>[(0,i.Wm)(g,{value:(0,i.SU)(r).max_tokens,"onUpdate:value":t[3]||(t[3]=e=>(0,i.SU)(r).max_tokens=e),class:"w-100%"},null,8,["value"])]),_:1})]),_:1},8,["model"])]),_:1},8,["show"])}}}),j={class:"manager-wrapper"},H={class:"left-provider"},K={class:"provider-list"},T=["onClick"],O={class:"item-icon"},q={key:1,class:"i-ai:big-model w-6.5 h-6.5"},I={class:"tit"},z={class:"add-privider"},M={class:"right-configuration"},$={class:"top-switch"},A={class:"switch-info"},D={class:"center-api"},E={class:"w-100%"},R={class:"w-100%"},F={class:"bottom-model-list"},V={class:"model-item"},Y={class:"model-name"},L={class:"operation"},J=(0,i.aZ)({__name:"ModelManager",setup(e){let t=(0,i.f3)("modelStore"),{providerList:a,modelList:l,currentProvider:m,configurationLoading:v,addProviderRef:b,addModelRef:g}=t;function _(){b.value.open()}return f(t),(e,f)=>{let Z=c.Z,N=d.ZP,C=u.Z,P=p.Z,J=n.Z,G=o.ZP,X=s.Z,Q=r.Z;return(0,i.wg)(),(0,i.iD)(i.HY,null,[(0,i._)("div",j,[(0,i._)("div",H,[(0,i._)("div",K,[((0,i.wg)(!0),(0,i.iD)(i.HY,null,(0,i.Ko)((0,i.SU)(a),e=>{var a;return(0,i.wg)(),(0,i.iD)("div",{key:e.supplierName,class:(0,i.C_)(["provider-item",{active:e.supplierName==(null==(a=(0,i.SU)(m))?void 0:a.supplierName)}]),onClick:a=>{m.value=e,h(e.supplierName,t)}},[(0,i._)("div",O,[e.icon?((0,i.wg)(),(0,i.j4)(Z,{key:0,src:e.icon,height:"24",width:"24"},null,8,["src"])):((0,i.wg)(),(0,i.iD)("i",q))]),(0,i._)("span",I,(0,i.zw)(e.supplierTitle),1)],10,T)}),128))]),(0,i._)("div",z,[(0,i.Wm)(N,{onClick:_},{icon:(0,i.w5)(()=>f[8]||(f[8]=[(0,i._)("i",{class:"i-ic:baseline-add-circle"},null,-1)])),default:(0,i.w5)(()=>[f[9]||(f[9]=(0,i.Uk)(" Add model service provider "))]),_:1,__:[9]})])]),(0,i.Wm)(Q,{show:(0,i.SU)(v)},{default:(0,i.w5)(()=>[(0,i._)("div",M,[(0,i._)("div",$,[(0,i._)("div",A,[f[10]||(f[10]=(0,i._)("i",{class:"i-ic:baseline-error text-5"},null,-1)),f[11]||(f[11]=(0,i._)("span",{class:"tit"},"DeepSeek",-1)),(0,i.Wm)(C,{value:(0,i.SU)(m).status,"onUpdate:value":[f[0]||(f[0]=e=>(0,i.SU)(m).status=e),f[1]||(f[1]=e=>(0,i.SU)(y)(e,(0,i.SU)(t)))]},null,8,["value"])]),(0,i._)("i",{class:"i-ri:delete-bin-3-line text-5.5 hover:text-red-5 cursor-pointer",onClick:f[2]||(f[2]=e=>(0,i.SU)(x)((0,i.SU)(t)))})]),(0,i._)("div",D,[(0,i.Wm)(X,{class:"mt-15px"},{default:(0,i.w5)(()=>[(0,i.Wm)(G,{label:"API密钥"},{default:(0,i.w5)(()=>[(0,i._)("div",E,[(0,i.Wm)(J,null,{default:(0,i.w5)(()=>[(0,i.Wm)(P,{value:(0,i.SU)(m).apiKey,"onUpdate:value":f[3]||(f[3]=e=>(0,i.SU)(m).apiKey=e)},null,8,["value"]),(0,i.Wm)(N,{onClick:f[4]||(f[4]=e=>(0,i.SU)(U)((0,i.SU)(t)))},{default:(0,i.w5)(()=>f[12]||(f[12]=[(0,i.Uk)("检查")])),_:1,__:[12]})]),_:1}),(0,i.Wm)(N,{text:"",type:"info",class:"mt-5px",onClick:f[5]||(f[5]=e=>(0,i.SU)(w)((0,i.SU)(t)))},{default:(0,i.w5)(()=>f[13]||(f[13]=[(0,i.Uk)("点击获取密钥")])),_:1,__:[13]})])]),_:1}),(0,i.Wm)(G,{label:"API地址"},{default:(0,i.w5)(()=>[(0,i._)("div",R,[(0,i.Wm)(P,{value:(0,i.SU)(m).baseUrl,"onUpdate:value":f[6]||(f[6]=e=>(0,i.SU)(m).baseUrl=e)},null,8,["value"]),f[14]||(f[14]=(0,i._)("span",{class:"text-[var(--color-text-3)] mt-5"},"示例: https://api.deepseek.com/v1",-1))])]),_:1})]),_:1}),(0,i.Wm)(N,{type:"primary",onClick:f[7]||(f[7]=e=>(0,i.SU)(k)((0,i.SU)(t)))},{default:(0,i.w5)(()=>f[15]||(f[15]=[(0,i.Uk)("保存API")])),_:1,__:[15]})]),(0,i._)("div",F,[f[16]||(f[16]=(0,i._)("div",{class:"header-tit"},[(0,i._)("div",null,[(0,i._)("span",{class:"tit"},"模型"),(0,i._)("span",{class:"sub-tit"},"默认从/models获取所有模型")])],-1)),((0,i.wg)(!0),(0,i.iD)(i.HY,null,(0,i.Ko)((0,i.SU)(l),e=>((0,i.wg)(),(0,i.iD)("div",{key:e.title,class:"model-list"},[(0,i._)("div",V,[(0,i._)("span",Y,(0,i.zw)(e.title),1),(0,i._)("div",L,[(0,i.Wm)(C,{value:e.status,"onUpdate:value":[t=>e.status=t,a=>(0,i.SU)(S)(e.modelId,a,(0,i.SU)(t))]},null,8,["value","onUpdate:value"])])])]))),128))])])]),_:1},8,["show"])]),(0,i.Wm)(W,{ref_key:"addProviderRef",ref:b},null,512),(0,i.Wm)(B,{ref_key:"addModelRef",ref:g},null,512)],64)}}});var G=a(41748);let X=(0,G.default)(J,[["__scopeId","data-v-1a0f8780"]]),Q={class:"max-w-4xl mx-auto px-4 py-8"},ee=(0,i.aZ)({__name:"index",setup(e){let t=function(){let e=(0,i.iH)([]),t=(0,i.iH)([]),a=(0,i.iH)(),l=(0,i.iH)({supplierTitle:"",supplierName:"",baseUrl:"",baseUrlExample:"",isUseUrlExample:!1,apiKey:"",home:"",help:"",status:!1,icon:"",sort:0}),r=(0,i.iH)(!1),s=(0,i.iH)({supplierTitle:"",supplierName:"",baseUrl:"",apiKey:""}),o=(0,i.iH)(),n=(0,i.iH)({title:"",modelId:"",max_tokens:0,capability:[]});return{providerList:e,modelList:t,addProviderRef:a,currentProvider:l,configurationLoading:r,addProviderFormData:s,addProviderFormDataRef:o,addModelFormData:n,addModelFormRef:(0,i.iH)(),addModelRef:(0,i.iH)()}}();return(0,i.JJ)("modelStore",t),(e,t)=>{let a=l.ZP;return(0,i.wg)(),(0,i.iD)("div",Q,[(0,i.Wm)(a,{class:"mb-5"},{default:(0,i.w5)(()=>t[0]||(t[0]=[(0,i._)("div",{class:"page-tit",style:{margin:"0"}},[(0,i._)("div",{class:"left-tit"},[(0,i._)("div",{class:"back-tool"},[(0,i._)("i",{class:"i-ri:apps-fill text-6"})]),(0,i._)("span",{class:"tit-content"}," AI Model ")])],-1)])),_:1,__:[0]}),(0,i.Wm)(a,null,{default:(0,i.w5)(()=>[(0,i.Wm)(X)]),_:1})])}}}),et=(0,G.default)(ee,[["__scopeId","data-v-c50590a6"]])}}]);