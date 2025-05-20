"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["798"],{618:function(e,t,i){i.d(t,{Z:()=>l});var r=i(209);let l=(0,r.aZ)({name:"Add",render:()=>(0,r.h)("svg",{width:"512",height:"512",viewBox:"0 0 512 512",fill:"none",xmlns:"http://www.w3.org/2000/svg"},(0,r.h)("path",{d:"M256 112V400M400 256H112",stroke:"currentColor","stroke-width":"32","stroke-linecap":"round","stroke-linejoin":"round"}))})},5476:function(e,t,i){i.d(t,{HX:()=>d,Ox:()=>c,ZP:()=>h,uv:()=>u});var r=i(209),l=i(1321),n=i(4124),o=i(2559),a=i(2284),s=i(9872);function d(e){return`${e}-ellipsis--line-clamp`}function c(e,t){return`${e}-ellipsis--cursor-${t}`}let u=Object.assign(Object.assign({},l.Z.props),{expandTrigger:String,lineClamp:[Number,String],tooltip:{type:[Boolean,Object],default:!0}}),h=(0,r.aZ)({name:"Ellipsis",inheritAttrs:!1,props:u,slots:Object,setup(e,{slots:t,attrs:i}){let o=(0,n.hJ)(),u=(0,l.Z)("Ellipsis","-ellipsis",s.Z,a.Z,e,o),h=(0,r.iH)(null),p=(0,r.iH)(null),v=(0,r.iH)(null),f=(0,r.iH)(!1),b=(0,r.Fl)(()=>{let{lineClamp:t}=e,{value:i}=f;return void 0!==t?{textOverflow:"","-webkit-line-clamp":i?"":t}:{textOverflow:i?"":"ellipsis","-webkit-line-clamp":""}});function g(){let t=!1,{value:i}=f;if(i)return!0;let{value:r}=h;if(r){let{lineClamp:i}=e;if(function(t){if(!t)return;let i=b.value,r=d(o.value);for(let l in void 0!==e.lineClamp?w(t,r,"add"):w(t,r,"remove"),i)t.style[l]!==i[l]&&(t.style[l]=i[l])}(r),void 0!==i)t=r.scrollHeight<=r.offsetHeight;else{let{value:e}=p;e&&(t=e.getBoundingClientRect().width<=r.getBoundingClientRect().width)}!function(t,i){let r=c(o.value,"pointer");"click"!==e.expandTrigger||i?w(t,r,"remove"):w(t,r,"add")}(r,t)}return t}let m=(0,r.Fl)(()=>"click"===e.expandTrigger?()=>{var e;let{value:t}=f;t&&(null===(e=v.value)||void 0===e||e.setShow(!1)),f.value=!t}:void 0);function w(e,t,i){"add"===i?e.classList.contains(t)||e.classList.add(t):e.classList.contains(t)&&e.classList.remove(t)}return(0,r.se)(()=>{var t;e.tooltip&&(null===(t=v.value)||void 0===t||t.setShow(!1))}),{mergedTheme:u,triggerRef:h,triggerInnerRef:p,tooltipRef:v,handleClick:m,renderTrigger:()=>(0,r.h)("span",Object.assign({},(0,r.dG)(i,{class:[`${o.value}-ellipsis`,void 0!==e.lineClamp?d(o.value):void 0,"click"===e.expandTrigger?c(o.value,"pointer"):void 0],style:b.value}),{ref:"triggerRef",onClick:m.value,onMouseenter:"click"===e.expandTrigger?g:void 0}),e.lineClamp?t:(0,r.h)("span",{ref:"triggerInnerRef"},t)),getTooltipDisabled:g}},render(){var e;let{tooltip:t,renderTrigger:i,$slots:l}=this;if(!t)return i();{let{mergedTheme:n}=this;return(0,r.h)(o.Z,Object.assign({ref:"tooltipRef",placement:"top"},t,{getDisabled:this.getTooltipDisabled,theme:n.peers.Tooltip,themeOverrides:n.peerOverrides.Tooltip}),{trigger:i,default:null!==(e=l.tooltip)&&void 0!==e?e:l.default})}}})},9872:function(e,t,i){i.d(t,{Z:()=>l});var r=i(2249);let l=(0,r.cB)("ellipsis",{overflow:"hidden"},[(0,r.u4)("line-clamp",`
 white-space: nowrap;
 display: inline-block;
 vertical-align: bottom;
 max-width: 100%;
 `),(0,r.cM)("line-clamp",`
 display: -webkit-inline-box;
 -webkit-box-orient: vertical;
 `),(0,r.cM)("cursor-pointer",`
 cursor: pointer;
 `)])},105:function(e,t,i){i.d(t,{ZP:()=>p});var r=i(209),l=i(1140),n=i(3898),o=i(5083),a=i(3615);let s={span:{type:[Number,String],default:1},offset:{type:[Number,String],default:0},suffix:Boolean,privateOffset:Number,privateSpan:Number,privateColStart:Number,privateShow:{type:Boolean,default:!0}},d=(0,l.u)(s),c=(0,r.aZ)({__GRID_ITEM__:!0,name:"GridItem",alias:["Gi"],props:s,setup(){let{isSsrRef:e,xGapRef:t,itemStyleRef:i,overflowRef:l,layoutShiftDisabledRef:n}=(0,r.f3)(a.r),s=(0,r.FN)();return{overflow:l,itemStyle:i,layoutShiftDisabled:n,mergedXGap:(0,r.Fl)(()=>(0,o.BL)(t.value||0)),deriveStyle:()=>{e.value;let{privateSpan:i=1,privateShow:r=!0,privateColStart:l,privateOffset:n=0}=s.vnode.props,{value:a}=t,d=(0,o.BL)(a||0);return{display:r?"":"none",gridColumn:`${null!=l?l:`span ${i}`} / span ${i}`,marginLeft:n?`calc((100% - (${i} - 1) * ${d}) / ${i} * ${n} + ${d} * ${n})`:""}}}},render(){var e,t;if(this.layoutShiftDisabled){let{span:e,offset:t,mergedXGap:i}=this;return(0,r.h)("div",{style:{gridColumn:`span ${e} / span ${e}`,marginLeft:t?`calc((100% - (${e} - 1) * ${i}) / ${e} * ${t} + ${i} * ${t})`:""}},this.$slots)}return(0,r.h)("div",{style:[this.itemStyle,this.deriveStyle()]},null===(t=(e=this.$slots).default)||void 0===t?void 0:t.call(e,{overflow:this.overflow}))}});var u=i(6883);let h=Object.assign(Object.assign({},s),u.Cn);(0,l.u)(h);let p=(0,r.aZ)({__GRID_ITEM__:!0,name:"FormItemGridItem",alias:["FormItemGi"],props:h,setup(){let e=(0,r.iH)(null);return{formItemInstRef:e,validate:(...t)=>{let{value:i}=e;if(i)return i.validate(...t)},restoreValidation:()=>{let{value:t}=e;t&&t.restoreValidation()}}},render(){return(0,r.h)(c,(0,n.C)(this.$.vnode.props||{},d),{default:()=>{let e=(0,n.C)(this.$props,u.eh);return(0,r.h)(u.ZP,Object.assign({ref:"formItemInstRef"},e),this.$slots)}})}})},1948:function(e,t,i){function r(e,t){var i;if(null==e)return;let r=function(e){if("number"==typeof e)return{"":e.toString()};let t={};return e.split(/ +/).forEach(e=>{if(""===e)return;let[i,r]=e.split(":");void 0===r?t[""]=i:t[i]=r}),t}(e);if(void 0===t)return r[""];if("string"==typeof t)return null!==(i=r[t])&&void 0!==i?i:r[""];if(Array.isArray(t)){for(let e=t.length-1;e>=0;--e){let i=t[e];if(i in r)return r[i]}return r[""]}{let e;let i=-1;return Object.keys(r).forEach(l=>{let n=Number(l);!Number.isNaN(n)&&t>=n&&n>=i&&(i=n,e=r[l])}),e}}i.d(t,{Z:()=>y});var l=i(6072),n=i(5083),o=i(209),a=i(7231);let s={xs:0,s:640,m:1024,l:1280,xl:1536,"2xl":1920},d={},c=function(e=s){if(!a.j||"function"!=typeof window.matchMedia)return(0,o.Fl)(()=>[]);let t=(0,o.iH)({}),i=Object.keys(e),r=(e,i)=>{e.matches?t.value[i]=!0:t.value[i]=!1};return i.forEach(t=>{let i,l;let n=e[t];if(void 0===d[n])(i=window.matchMedia(`(min-width: ${n}px)`)).addEventListener?i.addEventListener("change",e=>{l.forEach(i=>{i(e,t)})}):i.addListener&&i.addListener(e=>{l.forEach(i=>{i(e,t)})}),l=new Set,d[n]={mql:i,cbs:l};else i=d[n].mql,l=d[n].cbs;l.add(r),i.matches&&l.forEach(e=>{e(i,t)})}),(0,o.Jd)(()=>{i.forEach(t=>{let{cbs:i}=d[e[t]];i.has(r)&&i.delete(r)})}),(0,o.Fl)(()=>{let{value:e}=t;return i.filter(t=>e[t])})};var u=i(1367),h=i(9079),p=i(4124),v=i(4934),f=i(6253),b=i(3876);let g={xs:0,s:640,m:1024,l:1280,xl:1536,xxl:1920};var m=i(3615);let w="__ssr__",y=(0,o.aZ)({name:"Grid",inheritAttrs:!1,props:{layoutShiftDisabled:Boolean,responsive:{type:[String,Boolean],default:"self"},cols:{type:[Number,String],default:24},itemResponsive:Boolean,collapsed:Boolean,collapsedRows:{type:Number,default:1},itemStyle:[Object,String],xGap:{type:[Number,String],default:0},yGap:{type:[Number,String],default:0}},setup(e){let{mergedClsPrefixRef:t,mergedBreakpointsRef:i}=(0,p.ZP)(e),a=/^\d+$/,s=(0,o.iH)(void 0),d=c((null==i?void 0:i.value)||g),h=(0,u.Z)(()=>!(!e.itemResponsive&&a.test(e.cols.toString())&&a.test(e.xGap.toString())&&a.test(e.yGap.toString()))),f=(0,o.Fl)(()=>{if(h.value)return"self"===e.responsive?s.value:d.value}),b=(0,u.Z)(()=>{var t;return null!==(t=Number(r(e.cols.toString(),f.value)))&&void 0!==t?t:24}),y=(0,u.Z)(()=>r(e.xGap.toString(),f.value)),S=(0,u.Z)(()=>r(e.yGap.toString(),f.value)),x=e=>{s.value=e.contentRect.width},k=e=>{(0,l.J)(x,e)},$=(0,o.iH)(!1),B=(0,o.Fl)(()=>{if("self"===e.responsive)return k}),C=(0,o.iH)(!1),E=(0,o.iH)();return(0,o.bv)(()=>{let{value:e}=E;e&&e.hasAttribute(w)&&(e.removeAttribute(w),C.value=!0)}),(0,o.JJ)(m.r,{layoutShiftDisabledRef:(0,o.Vh)(e,"layoutShiftDisabled"),isSsrRef:C,itemStyleRef:(0,o.Vh)(e,"itemStyle"),xGapRef:y,overflowRef:$}),{isSsr:!v.j,contentEl:E,mergedClsPrefix:t,style:(0,o.Fl)(()=>e.layoutShiftDisabled?{width:"100%",display:"grid",gridTemplateColumns:`repeat(${e.cols}, minmax(0, 1fr))`,columnGap:(0,n.BL)(e.xGap),rowGap:(0,n.BL)(e.yGap)}:{width:"100%",display:"grid",gridTemplateColumns:`repeat(${b.value}, minmax(0, 1fr))`,columnGap:(0,n.BL)(y.value),rowGap:(0,n.BL)(S.value)}),isResponsive:h,responsiveQuery:f,responsiveCols:b,handleResize:B,overflow:$}},render(){if(this.layoutShiftDisabled)return(0,o.h)("div",(0,o.dG)({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style},this.$attrs),this.$slots);let e=()=>{var e,t,i,l,n,a,s;this.overflow=!1;let d=(0,f.x)((0,b.z)(this)),c=[],{collapsed:u,collapsedRows:h,responsiveCols:p,responsiveQuery:v}=this;d.forEach(e=>{var t,i,l,n,a;if((null===(t=null==e?void 0:e.type)||void 0===t?void 0:t.__GRID_ITEM__)!==!0)return;if(function(e){var t;let i=null===(t=e.dirs)||void 0===t?void 0:t.find(({dir:e})=>e===o.F8);return!!(i&&!1===i.value)}(e)){let t=(0,o.Ho)(e);t.props?t.props.privateShow=!1:t.props={privateShow:!1},c.push({child:t,rawChildSpan:0});return}e.dirs=(null===(i=e.dirs)||void 0===i?void 0:i.filter(({dir:e})=>e!==o.F8))||null,(null===(l=e.dirs)||void 0===l?void 0:l.length)===0&&(e.dirs=null);let s=(0,o.Ho)(e),d=Number(null!==(a=r(null===(n=s.props)||void 0===n?void 0:n.span,v))&&void 0!==a?a:m.L);0!==d&&c.push({child:s,rawChildSpan:d})});let g=0,y=null===(e=c[c.length-1])||void 0===e?void 0:e.child;if(null==y?void 0:y.props){let e=null===(t=y.props)||void 0===t?void 0:t.suffix;void 0!==e&&!1!==e&&(g=Number(null!==(l=r(null===(i=y.props)||void 0===i?void 0:i.span,v))&&void 0!==l?l:m.L),y.props.privateSpan=g,y.props.privateColStart=p+1-g,y.props.privateShow=null===(n=y.props.privateShow)||void 0===n||n)}let S=0,x=!1;for(let{child:e,rawChildSpan:t}of c){if(x&&(this.overflow=!0),!x){let i=Number(null!==(s=r(null===(a=e.props)||void 0===a?void 0:a.offset,v))&&void 0!==s?s:0),l=Math.min(t+i,p);if(e.props?(e.props.privateSpan=l,e.props.privateOffset=i):e.props={privateSpan:l,privateOffset:i},u){let e=S%p;l+e>p&&(S+=p-e),l+S+g>h*p?x=!0:S+=l}}x&&(e.props?!0!==e.props.privateShow&&(e.props.privateShow=!1):e.props={privateShow:!1})}return(0,o.h)("div",(0,o.dG)({ref:"contentEl",class:`${this.mergedClsPrefix}-grid`,style:this.style,[w]:this.isSsr||void 0},this.$attrs),c.map(({child:e})=>e))};return this.isResponsive&&"self"===this.responsive?(0,o.h)(h.Z,{onResize:this.handleResize},{default:e}):e()}})},3615:function(e,t,i){i.d(t,{L:()=>l,r:()=>n});var r=i(1579);let l=1,n=(0,r.U)("n-grid")},304:function(e,t,i){i.d(t,{Z:()=>v});var r=i(5083),l=i(7102),n=i(209),o=i(1321),a=i(4124),s=i(4934);let d=!1;var c=i(2249);let u={name:"Skeleton",common:i(8755).Z,self:function(e){let{heightSmall:t,heightMedium:i,heightLarge:r,borderRadius:l}=e;return{color:"#eee",colorEnd:"#ddd",borderRadius:l,heightSmall:t,heightMedium:i,heightLarge:r}}},h=(0,c.c)([(0,c.cB)("skeleton",`
 height: 1em;
 width: 100%;
 transition:
 --n-color-start .3s var(--n-bezier),
 --n-color-end .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 animation: 2s skeleton-loading infinite cubic-bezier(0.36, 0, 0.64, 1);
 background-color: var(--n-color-start);
 `),(0,c.c)("@keyframes skeleton-loading",`
 0% {
 background: var(--n-color-start);
 }
 40% {
 background: var(--n-color-end);
 }
 80% {
 background: var(--n-color-start);
 }
 100% {
 background: var(--n-color-start);
 }
 `)]),p=Object.assign(Object.assign({},o.Z.props),{text:Boolean,round:Boolean,circle:Boolean,height:[String,Number],width:[String,Number],size:String,repeat:{type:Number,default:1},animated:{type:Boolean,default:!0},sharp:{type:Boolean,default:!0}}),v=(0,n.aZ)({name:"Skeleton",inheritAttrs:!1,props:p,setup(e){!function(){if(s.j&&window.CSS&&!d&&(d=!0,"registerProperty"in(null==window?void 0:window.CSS)))try{CSS.registerProperty({name:"--n-color-start",syntax:"<color>",inherits:!1,initialValue:"#0000"}),CSS.registerProperty({name:"--n-color-end",syntax:"<color>",inherits:!1,initialValue:"#0000"})}catch(e){}}();let{mergedClsPrefixRef:t}=(0,a.ZP)(e),i=(0,o.Z)("Skeleton","-skeleton",h,u,e,t);return{mergedClsPrefix:t,style:(0,n.Fl)(()=>{var t,l;let n;let o=i.value,{common:{cubicBezierEaseInOut:a}}=o,s=o.self,{color:d,colorEnd:u,borderRadius:h}=s,{circle:p,sharp:v,round:f,width:b,height:g,size:m,text:w,animated:y}=e;void 0!==m&&(n=s[(0,c.Tl)("height",m)]);let S=p?null!==(t=null!=b?b:g)&&void 0!==t?t:n:b,x=null!==(l=p&&null!=b?b:g)&&void 0!==l?l:n;return{display:w?"inline-block":"",verticalAlign:w?"-0.125em":"",borderRadius:p?"50%":f?"4096px":v?"":h,width:"number"==typeof S?(0,r.BL)(S):S,height:"number"==typeof x?(0,r.BL)(x):x,animation:y?"":"none","--n-bezier":a,"--n-color-start":d,"--n-color-end":u}})}},render(){let{repeat:e,style:t,mergedClsPrefix:i,$attrs:r}=this,o=(0,n.h)("div",(0,n.dG)({class:`${i}-skeleton`,style:t},r));return e>1?(0,n.h)(n.HY,null,(0,l.rx)(e,null).map(e=>[o,"\n"])):o}})},9946:function(e,t,i){let r;i.d(t,{Z:()=>k});var l=i(5083),n=i(9226),o=i(209),a=i(6154),s=i(4131),d=i(1321),c=i(4124),u=i(9241),h=i(6169),p=i(1844),v=i(2249),f=i(8282),b=i(363),g=i(8755),m=i(2909);let w={name:"Switch",common:g.Z,self:function(e){let{primaryColor:t,opacityDisabled:i,borderRadius:r,textColor3:l}=e;return Object.assign(Object.assign({},m.Z),{iconColor:l,textColor:"white",loadingColor:t,opacityDisabled:i,railColor:"rgba(0, 0, 0, .14)",railColorActive:t,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:r,railBorderRadiusMedium:r,railBorderRadiusLarge:r,buttonBorderRadiusSmall:r,buttonBorderRadiusMedium:r,buttonBorderRadiusLarge:r,boxShadowFocus:`0 0 0 2px ${(0,b.zX)(t,{alpha:.2})}`})}};var y=i(8758);let S=(0,v.cB)("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[(0,v.cE)("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),(0,v.cE)("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),(0,v.cE)("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),(0,v.cB)("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[(0,y.c)({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),(0,v.cE)("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),(0,v.cE)("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),(0,v.cE)("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),(0,v.c)("&:focus",[(0,v.cE)("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),(0,v.cM)("round",[(0,v.cE)("rail","border-radius: calc(var(--n-rail-height) / 2);",[(0,v.cE)("button","border-radius: calc(var(--n-button-height) / 2);")])]),(0,v.u4)("disabled",[(0,v.u4)("icon",[(0,v.cM)("rubber-band",[(0,v.cM)("pressed",[(0,v.cE)("rail",[(0,v.cE)("button","max-width: var(--n-button-width-pressed);")])]),(0,v.cE)("rail",[(0,v.c)("&:active",[(0,v.cE)("button","max-width: var(--n-button-width-pressed);")])]),(0,v.cM)("active",[(0,v.cM)("pressed",[(0,v.cE)("rail",[(0,v.cE)("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),(0,v.cE)("rail",[(0,v.c)("&:active",[(0,v.cE)("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),(0,v.cM)("active",[(0,v.cE)("rail",[(0,v.cE)("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),(0,v.cE)("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[(0,v.cE)("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[(0,y.c)()]),(0,v.cE)("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),(0,v.cM)("active",[(0,v.cE)("rail","background-color: var(--n-rail-color-active);")]),(0,v.cM)("loading",[(0,v.cE)("rail",`
 cursor: wait;
 `)]),(0,v.cM)("disabled",[(0,v.cE)("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),x=Object.assign(Object.assign({},d.Z.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]}),k=(0,o.aZ)({name:"Switch",props:x,slots:Object,setup(e){void 0===r&&(r="undefined"==typeof CSS||void 0!==CSS.supports&&CSS.supports("width","max(1px)"));let{mergedClsPrefixRef:t,inlineThemeDisabled:i}=(0,c.ZP)(e),a=(0,d.Z)("Switch","-switch",S,w,e,t),s=(0,u.Z)(e),{mergedSizeRef:f,mergedDisabledRef:b}=s,g=(0,o.iH)(e.defaultValue),m=(0,o.Vh)(e,"value"),y=(0,n.Z)(m,g),x=(0,o.Fl)(()=>y.value===e.checkedValue),k=(0,o.iH)(!1),$=(0,o.iH)(!1),B=(0,o.Fl)(()=>{let{railStyle:t}=e;if(t)return t({focused:$.value,checked:x.value})});function C(t){let{"onUpdate:value":i,onChange:r,onUpdateValue:l}=e,{nTriggerFormInput:n,nTriggerFormChange:o}=s;i&&(0,p.R)(i,t),l&&(0,p.R)(l,t),r&&(0,p.R)(r,t),g.value=t,n(),o()}let E=(0,o.Fl)(()=>{let e,t,i;let{value:n}=f,{self:{opacityDisabled:o,railColor:s,railColorActive:d,buttonBoxShadow:c,buttonColor:u,boxShadowFocus:h,loadingColor:p,textColor:b,iconColor:g,[(0,v.Tl)("buttonHeight",n)]:m,[(0,v.Tl)("buttonWidth",n)]:w,[(0,v.Tl)("buttonWidthPressed",n)]:y,[(0,v.Tl)("railHeight",n)]:S,[(0,v.Tl)("railWidth",n)]:x,[(0,v.Tl)("railBorderRadius",n)]:k,[(0,v.Tl)("buttonBorderRadius",n)]:$},common:{cubicBezierEaseInOut:B}}=a.value;return r?(e=`calc((${S} - ${m}) / 2)`,t=`max(${S}, ${m})`,i=`max(${x}, calc(${x} + ${m} - ${S}))`):(e=(0,l.BL)(((0,l.fQ)(S)-(0,l.fQ)(m))/2),t=(0,l.BL)(Math.max((0,l.fQ)(S),(0,l.fQ)(m))),i=(0,l.fQ)(S)>(0,l.fQ)(m)?x:(0,l.BL)((0,l.fQ)(x)+(0,l.fQ)(m)-(0,l.fQ)(S))),{"--n-bezier":B,"--n-button-border-radius":$,"--n-button-box-shadow":c,"--n-button-color":u,"--n-button-width":w,"--n-button-width-pressed":y,"--n-button-height":m,"--n-height":t,"--n-offset":e,"--n-opacity-disabled":o,"--n-rail-border-radius":k,"--n-rail-color":s,"--n-rail-color-active":d,"--n-rail-height":S,"--n-rail-width":x,"--n-width":i,"--n-box-shadow-focus":h,"--n-loading-color":p,"--n-text-color":b,"--n-icon-color":g}}),_=i?(0,h.F)("switch",(0,o.Fl)(()=>f.value[0]),E,e):void 0;return{handleClick:function(){e.loading||b.value||(y.value!==e.checkedValue?C(e.checkedValue):C(e.uncheckedValue))},handleBlur:function(){$.value=!1,function(){let{nTriggerFormBlur:e}=s;e()}(),k.value=!1},handleFocus:function(){$.value=!0,function(){let{nTriggerFormFocus:e}=s;e()}()},handleKeyup:function(t){e.loading||b.value||" "!==t.key||(y.value!==e.checkedValue?C(e.checkedValue):C(e.uncheckedValue),k.value=!1)},handleKeydown:function(t){e.loading||b.value||" "!==t.key||(t.preventDefault(),k.value=!0)},mergedRailStyle:B,pressed:k,mergedClsPrefix:t,mergedValue:y,checked:x,mergedDisabled:b,cssVars:i?void 0:E,themeClass:null==_?void 0:_.themeClass,onRender:null==_?void 0:_.onRender}},render(){let{mergedClsPrefix:e,mergedDisabled:t,checked:i,mergedRailStyle:r,onRender:l,$slots:n}=this;null==l||l();let{checked:d,unchecked:c,icon:u,"checked-icon":h,"unchecked-icon":p}=n,v=!((0,f.aD)(u)&&(0,f.aD)(h)&&(0,f.aD)(p));return(0,o.h)("div",{role:"switch","aria-checked":i,class:[`${e}-switch`,this.themeClass,v&&`${e}-switch--icon`,i&&`${e}-switch--active`,t&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},(0,o.h)("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:r},(0,f.K9)(d,t=>(0,f.K9)(c,i=>t||i?(0,o.h)("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},(0,o.h)("div",{class:`${e}-switch__rail-placeholder`},(0,o.h)("div",{class:`${e}-switch__button-placeholder`}),t),(0,o.h)("div",{class:`${e}-switch__rail-placeholder`},(0,o.h)("div",{class:`${e}-switch__button-placeholder`}),i)):null)),(0,o.h)("div",{class:`${e}-switch__button`},(0,f.K9)(u,t=>(0,f.K9)(h,i=>(0,f.K9)(p,r=>(0,o.h)(a.Z,null,{default:()=>this.loading?(0,o.h)(s.Z,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(i||t)?(0,o.h)("div",{class:`${e}-switch__button-icon`,key:i?"checked-icon":"icon"},i||t):!this.checked&&(r||t)?(0,o.h)("div",{class:`${e}-switch__button-icon`,key:r?"unchecked-icon":"icon"},r||t):null})))),(0,f.K9)(d,t=>t&&(0,o.h)("div",{key:"checked",class:`${e}-switch__checked`},t)),(0,f.K9)(c,t=>t&&(0,o.h)("div",{key:"unchecked",class:`${e}-switch__unchecked`},t)))))}})}}]);