"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["645"],{618:function(e,t,a){a.d(t,{Z:()=>o});var r=a(209);let o=(0,r.aZ)({name:"Add",render:()=>(0,r.h)("svg",{width:"512",height:"512",viewBox:"0 0 512 512",fill:"none",xmlns:"http://www.w3.org/2000/svg"},(0,r.h)("path",{d:"M256 112V400M400 256H112",stroke:"currentColor","stroke-width":"32","stroke-linecap":"round","stroke-linejoin":"round"}))})},6720:function(e,t,a){a.d(t,{Z:()=>l,i:()=>i});var r=a(209),o=a(7475),n=a(8965);let i={tab:[String,Number,Object,Function],name:{type:[String,Number],required:!0},disabled:Boolean,displayDirective:{type:String,default:"if"},closable:{type:Boolean,default:void 0},tabProps:Object,label:[String,Number,Object,Function]},l=(0,r.aZ)({__TAB_PANE__:!0,name:"TabPane",alias:["TabPanel"],props:i,slots:Object,setup(e){let t=(0,r.f3)(n.x,null);return t||(0,o._y)("tab-pane","`n-tab-pane` must be placed inside `n-tabs`."),{style:t.paneStyleRef,class:t.paneClassRef,mergedClsPrefix:t.mergedClsPrefixRef}},render(){return(0,r.h)("div",{class:[`${this.mergedClsPrefix}-tab-pane`,this.class],style:this.style},this.$slots)}})},1664:function(e,t,a){a.d(t,{Z:()=>V});var r=a(6129),o=a(651);let n=function(){return o.Z.Date.now()};var i=a(157),l=Math.max,s=Math.min;let d=function(e,t,a){var o,d,c,b,p,f,v=0,u=!1,h=!1,g=!0;if("function"!=typeof e)throw TypeError("Expected a function");function x(t){var a=o,r=d;return o=d=void 0,v=t,b=e.apply(r,a)}function m(e){var a=e-f,r=e-v;return void 0===f||a>=t||a<0||h&&r>=c}function y(){var e,a,r,o=n();if(m(o))return w(o);p=setTimeout(y,(e=o-f,a=o-v,r=t-e,h?s(r,c-a):r))}function w(e){return(p=void 0,g&&o)?x(e):(o=d=void 0,b)}function B(){var e,a=n(),r=m(a);if(o=arguments,d=this,f=a,r){if(void 0===p)return v=e=f,p=setTimeout(y,t),u?x(e):b;if(h)return clearTimeout(p),p=setTimeout(y,t),x(f)}return void 0===p&&(p=setTimeout(y,t)),b}return t=(0,i.Z)(t)||0,(0,r.Z)(a)&&(u=!!a.leading,c=(h="maxWait"in a)?l((0,i.Z)(a.maxWait)||0,t):c,g="trailing"in a?!!a.trailing:g),B.cancel=function(){void 0!==p&&clearTimeout(p),v=0,o=f=d=p=void 0},B.flush=function(){return void 0===p?b:w(n())},B},c=function(e,t,a){var o=!0,n=!0;if("function"!=typeof e)throw TypeError("Expected a function");return(0,r.Z)(a)&&(o="leading"in a?!!a.leading:o,n="trailing"in a?!!a.trailing:n),d(e,t,{leading:o,maxWait:t,trailing:n})};var b=a(5083),p=a(2370),f=a(9226),v=a(2090),u=a(209),h=a(9079),g=a(7035),x=a(2382);let m=(0,x.c)(".v-x-scroll",{overflow:"auto",scrollbarWidth:"none"},[(0,x.c)("&::-webkit-scrollbar",{width:0,height:0})]),y=(0,u.aZ)({name:"XScroll",props:{disabled:Boolean,onScroll:Function},setup(){let e=(0,u.iH)(null),t=(0,g.O)();return m.mount({id:"vueuc/x-scroll",head:!0,anchorMetaName:x.A,ssr:t}),Object.assign({selfRef:e,handleWheel:function(e){e.currentTarget.offsetWidth<e.currentTarget.scrollWidth&&0!==e.deltaY&&(e.currentTarget.scrollLeft+=e.deltaY+e.deltaX,e.preventDefault())}},{scrollTo(...t){var a;null===(a=e.value)||void 0===a||a.scrollTo(...t)}})},render(){return(0,u.h)("div",{ref:"selfRef",onScroll:this.onScroll,onWheel:this.disabled?void 0:this.handleWheel,class:"v-x-scroll"},this.$slots)}});var w=a(1321),B=a(4124),$=a(6169),z=a(6253),T=a(1844),S=a(2249),C=a(8282),M=a(1011),P=a(8965);let R=(0,S.cB)("tabs",`
 box-sizing: border-box;
 width: 100%;
 display: flex;
 flex-direction: column;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
`,[(0,S.cM)("segment-type",[(0,S.cB)("tabs-rail",[(0,S.c)("&.transition-disabled",[(0,S.cB)("tabs-capsule",`
 transition: none;
 `)])])]),(0,S.cM)("top",[(0,S.cB)("tab-pane",`
 padding: var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left);
 `)]),(0,S.cM)("left",[(0,S.cB)("tab-pane",`
 padding: var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left) var(--n-pane-padding-top);
 `)]),(0,S.cM)("left, right",`
 flex-direction: row;
 `,[(0,S.cB)("tabs-bar",`
 width: 2px;
 right: 0;
 transition:
 top .2s var(--n-bezier),
 max-height .2s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-tab",`
 padding: var(--n-tab-padding-vertical); 
 `)]),(0,S.cM)("right",`
 flex-direction: row-reverse;
 `,[(0,S.cB)("tab-pane",`
 padding: var(--n-pane-padding-left) var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom);
 `),(0,S.cB)("tabs-bar",`
 left: 0;
 `)]),(0,S.cM)("bottom",`
 flex-direction: column-reverse;
 justify-content: flex-end;
 `,[(0,S.cB)("tab-pane",`
 padding: var(--n-pane-padding-bottom) var(--n-pane-padding-right) var(--n-pane-padding-top) var(--n-pane-padding-left);
 `),(0,S.cB)("tabs-bar",`
 top: 0;
 `)]),(0,S.cB)("tabs-rail",`
 position: relative;
 padding: 3px;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 background-color: var(--n-color-segment);
 transition: background-color .3s var(--n-bezier);
 display: flex;
 align-items: center;
 `,[(0,S.cB)("tabs-capsule",`
 border-radius: var(--n-tab-border-radius);
 position: absolute;
 pointer-events: none;
 background-color: var(--n-tab-color-segment);
 box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .08);
 transition: transform 0.3s var(--n-bezier);
 `),(0,S.cB)("tabs-tab-wrapper",`
 flex-basis: 0;
 flex-grow: 1;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[(0,S.cB)("tabs-tab",`
 overflow: hidden;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[(0,S.cM)("active",`
 font-weight: var(--n-font-weight-strong);
 color: var(--n-tab-text-color-active);
 `),(0,S.c)("&:hover",`
 color: var(--n-tab-text-color-hover);
 `)])])]),(0,S.cM)("flex",[(0,S.cB)("tabs-nav",`
 width: 100%;
 position: relative;
 `,[(0,S.cB)("tabs-wrapper",`
 width: 100%;
 `,[(0,S.cB)("tabs-tab",`
 margin-right: 0;
 `)])])]),(0,S.cB)("tabs-nav",`
 box-sizing: border-box;
 line-height: 1.5;
 display: flex;
 transition: border-color .3s var(--n-bezier);
 `,[(0,S.cE)("prefix, suffix",`
 display: flex;
 align-items: center;
 `),(0,S.cE)("prefix","padding-right: 16px;"),(0,S.cE)("suffix","padding-left: 16px;")]),(0,S.cM)("top, bottom",[(0,S.cB)("tabs-nav-scroll-wrapper",[(0,S.c)("&::before",`
 top: 0;
 bottom: 0;
 left: 0;
 width: 20px;
 `),(0,S.c)("&::after",`
 top: 0;
 bottom: 0;
 right: 0;
 width: 20px;
 `),(0,S.cM)("shadow-start",[(0,S.c)("&::before",`
 box-shadow: inset 10px 0 8px -8px rgba(0, 0, 0, .12);
 `)]),(0,S.cM)("shadow-end",[(0,S.c)("&::after",`
 box-shadow: inset -10px 0 8px -8px rgba(0, 0, 0, .12);
 `)])])]),(0,S.cM)("left, right",[(0,S.cB)("tabs-nav-scroll-content",`
 flex-direction: column;
 `),(0,S.cB)("tabs-nav-scroll-wrapper",[(0,S.c)("&::before",`
 top: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),(0,S.c)("&::after",`
 bottom: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),(0,S.cM)("shadow-start",[(0,S.c)("&::before",`
 box-shadow: inset 0 10px 8px -8px rgba(0, 0, 0, .12);
 `)]),(0,S.cM)("shadow-end",[(0,S.c)("&::after",`
 box-shadow: inset 0 -10px 8px -8px rgba(0, 0, 0, .12);
 `)])])]),(0,S.cB)("tabs-nav-scroll-wrapper",`
 flex: 1;
 position: relative;
 overflow: hidden;
 `,[(0,S.cB)("tabs-nav-y-scroll",`
 height: 100%;
 width: 100%;
 overflow-y: auto; 
 scrollbar-width: none;
 `,[(0,S.c)("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",`
 width: 0;
 height: 0;
 display: none;
 `)]),(0,S.c)("&::before, &::after",`
 transition: box-shadow .3s var(--n-bezier);
 pointer-events: none;
 content: "";
 position: absolute;
 z-index: 1;
 `)]),(0,S.cB)("tabs-nav-scroll-content",`
 display: flex;
 position: relative;
 min-width: 100%;
 min-height: 100%;
 width: fit-content;
 box-sizing: border-box;
 `),(0,S.cB)("tabs-wrapper",`
 display: inline-flex;
 flex-wrap: nowrap;
 position: relative;
 `),(0,S.cB)("tabs-tab-wrapper",`
 display: flex;
 flex-wrap: nowrap;
 flex-shrink: 0;
 flex-grow: 0;
 `),(0,S.cB)("tabs-tab",`
 cursor: pointer;
 white-space: nowrap;
 flex-wrap: nowrap;
 display: inline-flex;
 align-items: center;
 color: var(--n-tab-text-color);
 font-size: var(--n-tab-font-size);
 background-clip: padding-box;
 padding: var(--n-tab-padding);
 transition:
 box-shadow .3s var(--n-bezier),
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[(0,S.cM)("disabled",{cursor:"not-allowed"}),(0,S.cE)("close",`
 margin-left: 6px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),(0,S.cE)("label",`
 display: flex;
 align-items: center;
 z-index: 1;
 `)]),(0,S.cB)("tabs-bar",`
 position: absolute;
 bottom: 0;
 height: 2px;
 border-radius: 1px;
 background-color: var(--n-bar-color);
 transition:
 left .2s var(--n-bezier),
 max-width .2s var(--n-bezier),
 opacity .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `,[(0,S.c)("&.transition-disabled",`
 transition: none;
 `),(0,S.cM)("disabled",`
 background-color: var(--n-tab-text-color-disabled)
 `)]),(0,S.cB)("tabs-pane-wrapper",`
 position: relative;
 overflow: hidden;
 transition: max-height .2s var(--n-bezier);
 `),(0,S.cB)("tab-pane",`
 color: var(--n-pane-text-color);
 width: 100%;
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 opacity .2s var(--n-bezier);
 left: 0;
 right: 0;
 top: 0;
 `,[(0,S.c)("&.next-transition-leave-active, &.prev-transition-leave-active, &.next-transition-enter-active, &.prev-transition-enter-active",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 transform .2s var(--n-bezier),
 opacity .2s var(--n-bezier);
 `),(0,S.c)("&.next-transition-leave-active, &.prev-transition-leave-active",`
 position: absolute;
 `),(0,S.c)("&.next-transition-enter-from, &.prev-transition-leave-to",`
 transform: translateX(32px);
 opacity: 0;
 `),(0,S.c)("&.next-transition-leave-to, &.prev-transition-enter-from",`
 transform: translateX(-32px);
 opacity: 0;
 `),(0,S.c)("&.next-transition-leave-from, &.next-transition-enter-to, &.prev-transition-leave-from, &.prev-transition-enter-to",`
 transform: translateX(0);
 opacity: 1;
 `)]),(0,S.cB)("tabs-tab-pad",`
 box-sizing: border-box;
 width: var(--n-tab-gap);
 flex-grow: 0;
 flex-shrink: 0;
 `),(0,S.cM)("line-type, bar-type",[(0,S.cB)("tabs-tab",`
 font-weight: var(--n-tab-font-weight);
 box-sizing: border-box;
 vertical-align: bottom;
 `,[(0,S.c)("&:hover",{color:"var(--n-tab-text-color-hover)"}),(0,S.cM)("active",`
 color: var(--n-tab-text-color-active);
 font-weight: var(--n-tab-font-weight-active);
 `),(0,S.cM)("disabled",{color:"var(--n-tab-text-color-disabled)"})])]),(0,S.cB)("tabs-nav",[(0,S.cM)("line-type",[(0,S.cM)("top",[(0,S.cE)("prefix, suffix",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-nav-scroll-content",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-bar",`
 bottom: -1px;
 `)]),(0,S.cM)("left",[(0,S.cE)("prefix, suffix",`
 border-right: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-nav-scroll-content",`
 border-right: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-bar",`
 right: -1px;
 `)]),(0,S.cM)("right",[(0,S.cE)("prefix, suffix",`
 border-left: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-nav-scroll-content",`
 border-left: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-bar",`
 left: -1px;
 `)]),(0,S.cM)("bottom",[(0,S.cE)("prefix, suffix",`
 border-top: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-nav-scroll-content",`
 border-top: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-bar",`
 top: -1px;
 `)]),(0,S.cE)("prefix, suffix",`
 transition: border-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-nav-scroll-content",`
 transition: border-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-bar",`
 border-radius: 0;
 `)]),(0,S.cM)("card-type",[(0,S.cE)("prefix, suffix",`
 transition: border-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-pad",`
 flex-grow: 1;
 transition: border-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-tab-pad",`
 transition: border-color .3s var(--n-bezier);
 `),(0,S.cB)("tabs-tab",`
 font-weight: var(--n-tab-font-weight);
 border: 1px solid var(--n-tab-border-color);
 background-color: var(--n-tab-color);
 box-sizing: border-box;
 position: relative;
 vertical-align: bottom;
 display: flex;
 justify-content: space-between;
 font-size: var(--n-tab-font-size);
 color: var(--n-tab-text-color);
 `,[(0,S.cM)("addable",`
 padding-left: 8px;
 padding-right: 8px;
 font-size: 16px;
 justify-content: center;
 `,[(0,S.cE)("height-placeholder",`
 width: 0;
 font-size: var(--n-tab-font-size);
 `),(0,S.u4)("disabled",[(0,S.c)("&:hover",`
 color: var(--n-tab-text-color-hover);
 `)])]),(0,S.cM)("closable","padding-right: 8px;"),(0,S.cM)("active",`
 background-color: #0000;
 font-weight: var(--n-tab-font-weight-active);
 color: var(--n-tab-text-color-active);
 `),(0,S.cM)("disabled","color: var(--n-tab-text-color-disabled);")])]),(0,S.cM)("left, right",`
 flex-direction: column; 
 `,[(0,S.cE)("prefix, suffix",`
 padding: var(--n-tab-padding-vertical);
 `),(0,S.cB)("tabs-wrapper",`
 flex-direction: column;
 `),(0,S.cB)("tabs-tab-wrapper",`
 flex-direction: column;
 `,[(0,S.cB)("tabs-tab-pad",`
 height: var(--n-tab-gap-vertical);
 width: 100%;
 `)])]),(0,S.cM)("top",[(0,S.cM)("card-type",[(0,S.cB)("tabs-scroll-padding","border-bottom: 1px solid var(--n-tab-border-color);"),(0,S.cE)("prefix, suffix",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-tab",`
 border-top-left-radius: var(--n-tab-border-radius);
 border-top-right-radius: var(--n-tab-border-radius);
 `,[(0,S.cM)("active",`
 border-bottom: 1px solid #0000;
 `)]),(0,S.cB)("tabs-tab-pad",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-pad",`
 border-bottom: 1px solid var(--n-tab-border-color);
 `)])]),(0,S.cM)("left",[(0,S.cM)("card-type",[(0,S.cB)("tabs-scroll-padding","border-right: 1px solid var(--n-tab-border-color);"),(0,S.cE)("prefix, suffix",`
 border-right: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-tab",`
 border-top-left-radius: var(--n-tab-border-radius);
 border-bottom-left-radius: var(--n-tab-border-radius);
 `,[(0,S.cM)("active",`
 border-right: 1px solid #0000;
 `)]),(0,S.cB)("tabs-tab-pad",`
 border-right: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-pad",`
 border-right: 1px solid var(--n-tab-border-color);
 `)])]),(0,S.cM)("right",[(0,S.cM)("card-type",[(0,S.cB)("tabs-scroll-padding","border-left: 1px solid var(--n-tab-border-color);"),(0,S.cE)("prefix, suffix",`
 border-left: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-tab",`
 border-top-right-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[(0,S.cM)("active",`
 border-left: 1px solid #0000;
 `)]),(0,S.cB)("tabs-tab-pad",`
 border-left: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-pad",`
 border-left: 1px solid var(--n-tab-border-color);
 `)])]),(0,S.cM)("bottom",[(0,S.cM)("card-type",[(0,S.cB)("tabs-scroll-padding","border-top: 1px solid var(--n-tab-border-color);"),(0,S.cE)("prefix, suffix",`
 border-top: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-tab",`
 border-bottom-left-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[(0,S.cM)("active",`
 border-top: 1px solid #0000;
 `)]),(0,S.cB)("tabs-tab-pad",`
 border-top: 1px solid var(--n-tab-border-color);
 `),(0,S.cB)("tabs-pad",`
 border-top: 1px solid var(--n-tab-border-color);
 `)])])])]);var k=a(8822),E=a(9653),L=a(618),j=a(2487),H=a(3772),_=a(6720);let A=Object.assign({internalLeftPadded:Boolean,internalAddable:Boolean,internalCreatedByPane:Boolean},(0,j.C)(_.i,["displayDirective"])),W=(0,u.aZ)({__TAB__:!0,inheritAttrs:!1,name:"Tab",props:A,setup(e){let{mergedClsPrefixRef:t,valueRef:a,typeRef:r,closableRef:o,tabStyleRef:n,addTabStyleRef:i,tabClassRef:l,addTabClassRef:s,tabChangeIdRef:d,onBeforeLeaveRef:c,triggerRef:b,handleAdd:p,activateTab:f,handleClose:v}=(0,u.f3)(P.x);return{trigger:b,mergedClosable:(0,u.Fl)(()=>{if(e.internalAddable)return!1;let{closable:t}=e;return void 0===t?o.value:t}),style:n,addStyle:i,tabClass:l,addTabClass:s,clsPrefix:t,value:a,type:r,handleClose(t){t.stopPropagation(),e.disabled||v(e.name)},activateTab(){if(e.disabled)return;if(e.internalAddable){p();return}let{name:t}=e,r=++d.id;if(t!==a.value){let{value:o}=c;o?Promise.resolve(o(e.name,a.value)).then(e=>{e&&d.id===r&&f(t)}):f(t)}}}},render(){let{internalAddable:e,clsPrefix:t,name:a,disabled:r,label:o,tab:n,value:i,mergedClosable:l,trigger:s,$slots:{default:d}}=this,c=null!=o?o:n;return(0,u.h)("div",{class:`${t}-tabs-tab-wrapper`},this.internalLeftPadded?(0,u.h)("div",{class:`${t}-tabs-tab-pad`}):null,(0,u.h)("div",Object.assign({key:a,"data-name":a,"data-disabled":!!r||void 0},(0,u.dG)({class:[`${t}-tabs-tab`,i===a&&`${t}-tabs-tab--active`,r&&`${t}-tabs-tab--disabled`,l&&`${t}-tabs-tab--closable`,e&&`${t}-tabs-tab--addable`,e?this.addTabClass:this.tabClass],onClick:"click"===s?this.activateTab:void 0,onMouseenter:"hover"===s?this.activateTab:void 0,style:e?this.addStyle:this.style},this.internalCreatedByPane?this.tabProps||{}:this.$attrs)),(0,u.h)("span",{class:`${t}-tabs-tab__label`},e?(0,u.h)(u.HY,null,(0,u.h)("div",{class:`${t}-tabs-tab__height-placeholder`},"\xa0"),(0,u.h)(k.Z,{clsPrefix:t},{default:()=>(0,u.h)(L.Z,null)})):d?d():"object"==typeof c?c:(0,H.s)(null!=c?c:a)),l&&"card"===this.type?(0,u.h)(E.Z,{clsPrefix:t,class:`${t}-tabs-tab__close`,onClick:this.handleClose,disabled:r}):null))}}),Z=Object.assign(Object.assign({},w.Z.props),{value:[String,Number],defaultValue:[String,Number],trigger:{type:String,default:"click"},type:{type:String,default:"bar"},closable:Boolean,justifyContent:String,size:{type:String,default:"medium"},placement:{type:String,default:"top"},tabStyle:[String,Object],tabClass:String,addTabStyle:[String,Object],addTabClass:String,barWidth:Number,paneClass:String,paneStyle:[String,Object],paneWrapperClass:String,paneWrapperStyle:[String,Object],addable:[Boolean,Object],tabsPadding:{type:Number,default:0},animated:Boolean,onBeforeLeave:Function,onAdd:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onClose:[Function,Array],labelSize:String,activeName:[String,Number],onActiveNameChange:[Function,Array]}),V=(0,u.aZ)({name:"Tabs",props:Z,slots:Object,setup(e,{slots:t}){var a,r,o,n;let{mergedClsPrefixRef:i,inlineThemeDisabled:l}=(0,B.ZP)(e),s=(0,w.Z)("Tabs","-tabs",R,M.Z,e,i),d=(0,u.iH)(null),h=(0,u.iH)(null),g=(0,u.iH)(null),x=(0,u.iH)(null),m=(0,u.iH)(null),y=(0,u.iH)(null),C=(0,u.iH)(!0),k=(0,u.iH)(!0),E=(0,p.Z)(e,["labelSize","size"]),L=(0,p.Z)(e,["activeName","value"]),j=(0,u.iH)(null!==(r=null!==(a=L.value)&&void 0!==a?a:e.defaultValue)&&void 0!==r?r:t.default?null===(n=null===(o=(0,z.x)(t.default())[0])||void 0===o?void 0:o.props)||void 0===n?void 0:n.name:null),H=(0,f.Z)(L,j),_={id:0},A=(0,u.Fl)(()=>{if(e.justifyContent&&"card"!==e.type)return{display:"flex",justifyContent:e.justifyContent}});function W(){var e;let{value:t}=H;return null===t?null:null===(e=d.value)||void 0===e?void 0:e.querySelector(`[data-name="${t}"]`)}function Z(e){let{value:t}=h;if(t)for(let a of e)t.style[a]=""}function V(){if("card"===e.type)return;let t=W();t?function(t){if("card"===e.type)return;let{value:a}=h;if(!a)return;let r="0"===a.style.opacity;if(t){let o=`${i.value}-tabs-bar--disabled`,{barWidth:n,placement:l}=e;if("true"===t.dataset.disabled?a.classList.add(o):a.classList.remove(o),["top","bottom"].includes(l)){if(Z(["top","maxHeight","height"]),"number"==typeof n&&t.offsetWidth>=n){let e=Math.floor((t.offsetWidth-n)/2)+t.offsetLeft;a.style.left=`${e}px`,a.style.maxWidth=`${n}px`}else a.style.left=`${t.offsetLeft}px`,a.style.maxWidth=`${t.offsetWidth}px`;a.style.width="8192px",r&&(a.style.transition="none"),a.offsetWidth,r&&(a.style.transition="",a.style.opacity="1")}else{if(Z(["left","maxWidth","width"]),"number"==typeof n&&t.offsetHeight>=n){let e=Math.floor((t.offsetHeight-n)/2)+t.offsetTop;a.style.top=`${e}px`,a.style.maxHeight=`${n}px`}else a.style.top=`${t.offsetTop}px`,a.style.maxHeight=`${t.offsetHeight}px`;a.style.height="8192px",r&&(a.style.transition="none"),a.offsetHeight,r&&(a.style.transition="",a.style.opacity="1")}}}(t):function(){if("card"===e.type)return;let{value:t}=h;t&&(t.style.opacity="0")}()}function O(){var e;let t=null===(e=m.value)||void 0===e?void 0:e.$el;if(!t)return;let a=W();if(!a)return;let{scrollLeft:r,offsetWidth:o}=t,{offsetLeft:n,offsetWidth:i}=a;r>n?t.scrollTo({top:0,left:n,behavior:"smooth"}):n+i>r+o&&t.scrollTo({top:0,left:n+i-o,behavior:"smooth"})}(0,u.YP)(H,()=>{_.id=0,V(),O()});let N=(0,u.iH)(null),F=0,D=null,Y={value:[]},X=(0,u.iH)("next"),Q=!0;function U(){let{value:e}=h;if(!e)return;Q||(Q=!1);let t="transition-disabled";e.classList.add(t),V(),e.classList.remove(t)}let G=(0,u.iH)(null);function q({transitionDisabled:e}){let t=d.value;if(!t)return;e&&t.classList.add("transition-disabled");let a=W();a&&G.value&&(G.value.style.width=`${a.offsetWidth}px`,G.value.style.height=`${a.offsetHeight}px`,G.value.style.transform=`translateX(${a.offsetLeft-(0,b.fQ)(getComputedStyle(t).paddingLeft)}px)`,e&&G.value.offsetWidth),e&&t.classList.remove("transition-disabled")}(0,u.YP)([H],()=>{"segment"===e.type&&(0,u.Y3)(()=>{q({transitionDisabled:!1})})}),(0,u.bv)(()=>{"segment"===e.type&&q({transitionDisabled:!0})});let I=0,J=c(function(t){var a,r;if(0===t.contentRect.width&&0===t.contentRect.height||I===t.contentRect.width)return;I=t.contentRect.width;let{type:o}=e;if(("line"===o||"bar"===o)&&(Q||(null===(a=e.justifyContent)||void 0===a?void 0:a.startsWith("space")))&&U(),"segment"!==o){let{placement:t}=e;et(("top"===t||"bottom"===t?null===(r=m.value)||void 0===r?void 0:r.$el:y.value)||null)}},64);(0,u.YP)([()=>e.justifyContent,()=>e.size],()=>{(0,u.Y3)(()=>{let{type:t}=e;("line"===t||"bar"===t)&&U()})});let K=(0,u.iH)(!1),ee=c(function(t){var a;let{target:r,contentRect:{width:o,height:n}}=t,i=r.parentElement.parentElement.offsetWidth,l=r.parentElement.parentElement.offsetHeight,{placement:s}=e;if(K.value){let{value:e}=x;if(!e)return;"top"===s||"bottom"===s?i-o>e.$el.offsetWidth&&(K.value=!1):l-n>e.$el.offsetHeight&&(K.value=!1)}else"top"===s||"bottom"===s?i<o&&(K.value=!0):l<n&&(K.value=!0);et((null===(a=m.value)||void 0===a?void 0:a.$el)||null)},64);function et(t){if(!t)return;let{placement:a}=e;if("top"===a||"bottom"===a){let{scrollLeft:e,scrollWidth:a,offsetWidth:r}=t;C.value=e<=0,k.value=e+r>=a}else{let{scrollTop:e,scrollHeight:a,offsetHeight:r}=t;C.value=e<=0,k.value=e+r>=a}}let ea=c(e=>{et(e.target)},64);(0,u.JJ)(P.x,{triggerRef:(0,u.Vh)(e,"trigger"),tabStyleRef:(0,u.Vh)(e,"tabStyle"),tabClassRef:(0,u.Vh)(e,"tabClass"),addTabStyleRef:(0,u.Vh)(e,"addTabStyle"),addTabClassRef:(0,u.Vh)(e,"addTabClass"),paneClassRef:(0,u.Vh)(e,"paneClass"),paneStyleRef:(0,u.Vh)(e,"paneStyle"),mergedClsPrefixRef:i,typeRef:(0,u.Vh)(e,"type"),closableRef:(0,u.Vh)(e,"closable"),valueRef:H,tabChangeIdRef:_,onBeforeLeaveRef:(0,u.Vh)(e,"onBeforeLeave"),activateTab:function(t){let a=H.value,r="next";for(let e of Y.value){if(e===a)break;if(e===t){r="prev";break}}X.value=r,function(t){let{onActiveNameChange:a,onUpdateValue:r,"onUpdate:value":o}=e;a&&(0,T.R)(a,t),r&&(0,T.R)(r,t),o&&(0,T.R)(o,t),j.value=t}(t)},handleClose:function(t){let{onClose:a}=e;a&&(0,T.R)(a,t)},handleAdd:function(){let{onAdd:t}=e;t&&t(),(0,u.Y3)(()=>{let e=W(),{value:t}=m;e&&t&&t.scrollTo({left:e.offsetLeft,top:0,behavior:"smooth"})})}}),(0,v.Z)(()=>{V(),O()}),(0,u.m0)(()=>{let{value:e}=g;if(!e)return;let{value:t}=i,a=`${t}-tabs-nav-scroll-wrapper--shadow-start`,r=`${t}-tabs-nav-scroll-wrapper--shadow-end`;C.value?e.classList.remove(a):e.classList.add(a),k.value?e.classList.remove(r):e.classList.add(r)});let er=(0,u.Fl)(()=>{let{value:t}=E,{type:a}=e,r={card:"Card",bar:"Bar",line:"Line",segment:"Segment"}[a],o=`${t}${r}`,{self:{barColor:n,closeIconColor:i,closeIconColorHover:l,closeIconColorPressed:d,tabColor:c,tabBorderColor:p,paneTextColor:f,tabFontWeight:v,tabBorderRadius:u,tabFontWeightActive:h,colorSegment:g,fontWeightStrong:x,tabColorSegment:m,closeSize:y,closeIconSize:w,closeColorHover:B,closeColorPressed:$,closeBorderRadius:z,[(0,S.Tl)("panePadding",t)]:T,[(0,S.Tl)("tabPadding",o)]:C,[(0,S.Tl)("tabPaddingVertical",o)]:M,[(0,S.Tl)("tabGap",o)]:P,[(0,S.Tl)("tabGap",`${o}Vertical`)]:R,[(0,S.Tl)("tabTextColor",a)]:k,[(0,S.Tl)("tabTextColorActive",a)]:L,[(0,S.Tl)("tabTextColorHover",a)]:j,[(0,S.Tl)("tabTextColorDisabled",a)]:H,[(0,S.Tl)("tabFontSize",t)]:_},common:{cubicBezierEaseInOut:A}}=s.value;return{"--n-bezier":A,"--n-color-segment":g,"--n-bar-color":n,"--n-tab-font-size":_,"--n-tab-text-color":k,"--n-tab-text-color-active":L,"--n-tab-text-color-disabled":H,"--n-tab-text-color-hover":j,"--n-pane-text-color":f,"--n-tab-border-color":p,"--n-tab-border-radius":u,"--n-close-size":y,"--n-close-icon-size":w,"--n-close-color-hover":B,"--n-close-color-pressed":$,"--n-close-border-radius":z,"--n-close-icon-color":i,"--n-close-icon-color-hover":l,"--n-close-icon-color-pressed":d,"--n-tab-color":c,"--n-tab-font-weight":v,"--n-tab-font-weight-active":h,"--n-tab-padding":C,"--n-tab-padding-vertical":M,"--n-tab-gap":P,"--n-tab-gap-vertical":R,"--n-pane-padding-left":(0,b.tQ)(T,"left"),"--n-pane-padding-right":(0,b.tQ)(T,"right"),"--n-pane-padding-top":(0,b.tQ)(T,"top"),"--n-pane-padding-bottom":(0,b.tQ)(T,"bottom"),"--n-font-weight-strong":x,"--n-tab-color-segment":m}}),eo=l?(0,$.F)("tabs",(0,u.Fl)(()=>`${E.value[0]}${e.type[0]}`),er,e):void 0;return Object.assign({mergedClsPrefix:i,mergedValue:H,renderedNames:new Set,segmentCapsuleElRef:G,tabsPaneWrapperRef:N,tabsElRef:d,barElRef:h,addTabInstRef:x,xScrollInstRef:m,scrollWrapperElRef:g,addTabFixed:K,tabWrapperStyle:A,handleNavResize:J,mergedSize:E,handleScroll:ea,handleTabsResize:ee,cssVars:l?void 0:er,themeClass:null==eo?void 0:eo.themeClass,animationDirection:X,renderNameListRef:Y,yScrollElRef:y,handleSegmentResize:()=>{q({transitionDisabled:!0})},onAnimationBeforeLeave:function(e){let t=N.value;if(t){F=e.getBoundingClientRect().height;let a=`${F}px`,r=()=>{t.style.height=a,t.style.maxHeight=a};D?(r(),D(),D=null):D=r}},onAnimationEnter:function(e){let t=N.value;if(t){let a=e.getBoundingClientRect().height,r=()=>{document.body.offsetHeight,t.style.maxHeight=`${a}px`,t.style.height=`${Math.max(F,a)}px`};D?(D(),D=null,r()):D=r}},onAnimationAfterEnter:function(){let t=N.value;if(t){t.style.maxHeight="",t.style.height="";let{paneWrapperStyle:a}=e;if("string"==typeof a)t.style.cssText=a;else if(a){let{maxHeight:e,height:r}=a;void 0!==e&&(t.style.maxHeight=e),void 0!==r&&(t.style.height=r)}}},onRender:null==eo?void 0:eo.onRender},{syncBarPosition:()=>{V()}})},render(){let{mergedClsPrefix:e,type:t,placement:a,addTabFixed:r,addable:o,mergedSize:n,renderNameListRef:i,onRender:l,paneWrapperClass:s,paneWrapperStyle:d,$slots:{default:c,prefix:b,suffix:p}}=this;null==l||l();let f=c?(0,z.x)(c()).filter(e=>!0===e.type.__TAB_PANE__):[],v=c?(0,z.x)(c()).filter(e=>!0===e.type.__TAB__):[],g=!v.length,x="card"===t,m="segment"===t,w=!x&&!m&&this.justifyContent;i.value=[];let B=()=>{let t=(0,u.h)("div",{style:this.tabWrapperStyle,class:`${e}-tabs-wrapper`},w?null:(0,u.h)("div",{class:`${e}-tabs-scroll-padding`,style:"top"===a||"bottom"===a?{width:`${this.tabsPadding}px`}:{height:`${this.tabsPadding}px`}}),g?f.map((e,t)=>(i.value.push(e.props.name),D((0,u.h)(W,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:0!==t&&(!w||"center"===w||"start"===w||"end"===w)}),e.children?{default:e.children.tab}:void 0)))):v.map((e,t)=>(i.value.push(e.props.name),0===t||w)?D(e):D(F(e))),!r&&o&&x?N(o,(g?f.length:v.length)!==0):null,w?null:(0,u.h)("div",{class:`${e}-tabs-scroll-padding`,style:{width:`${this.tabsPadding}px`}}));return(0,u.h)("div",{ref:"tabsElRef",class:`${e}-tabs-nav-scroll-content`},x&&o?(0,u.h)(h.Z,{onResize:this.handleTabsResize},{default:()=>t}):t,x?(0,u.h)("div",{class:`${e}-tabs-pad`}):null,x?null:(0,u.h)("div",{ref:"barElRef",class:`${e}-tabs-bar`}))},$=m?"top":a;return(0,u.h)("div",{class:[`${e}-tabs`,this.themeClass,`${e}-tabs--${t}-type`,`${e}-tabs--${n}-size`,w&&`${e}-tabs--flex`,`${e}-tabs--${$}`],style:this.cssVars},(0,u.h)("div",{class:[`${e}-tabs-nav--${t}-type`,`${e}-tabs-nav--${$}`,`${e}-tabs-nav`]},(0,C.K9)(b,t=>t&&(0,u.h)("div",{class:`${e}-tabs-nav__prefix`},t)),m?(0,u.h)(h.Z,{onResize:this.handleSegmentResize},{default:()=>(0,u.h)("div",{class:`${e}-tabs-rail`,ref:"tabsElRef"},(0,u.h)("div",{class:`${e}-tabs-capsule`,ref:"segmentCapsuleElRef"},(0,u.h)("div",{class:`${e}-tabs-wrapper`},(0,u.h)("div",{class:`${e}-tabs-tab`}))),g?f.map((e,t)=>(i.value.push(e.props.name),(0,u.h)(W,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:0!==t}),e.children?{default:e.children.tab}:void 0))):v.map((e,t)=>(i.value.push(e.props.name),0===t)?e:F(e)))}):(0,u.h)(h.Z,{onResize:this.handleNavResize},{default:()=>(0,u.h)("div",{class:`${e}-tabs-nav-scroll-wrapper`,ref:"scrollWrapperElRef"},["top","bottom"].includes($)?(0,u.h)(y,{ref:"xScrollInstRef",onScroll:this.handleScroll},{default:B}):(0,u.h)("div",{class:`${e}-tabs-nav-y-scroll`,onScroll:this.handleScroll,ref:"yScrollElRef"},B()))}),r&&o&&x?N(o,!0):null,(0,C.K9)(p,t=>t&&(0,u.h)("div",{class:`${e}-tabs-nav__suffix`},t))),g&&(this.animated&&("top"===$||"bottom"===$)?(0,u.h)("div",{ref:"tabsPaneWrapperRef",style:d,class:[`${e}-tabs-pane-wrapper`,s]},O(f,this.mergedValue,this.renderedNames,this.onAnimationBeforeLeave,this.onAnimationEnter,this.onAnimationAfterEnter,this.animationDirection)):O(f,this.mergedValue,this.renderedNames)))}});function O(e,t,a,r,o,n,i){let l=[];return(e.forEach(e=>{let r,o;let{name:n,displayDirective:i,"display-directive":s}=e.props,d=t===n;if(void 0!==e.key&&(e.key=n),d||i===(r="show")||s===r||(i===(o="show:lazy")||s===o)&&a.has(n)){a.has(n)||a.add(n);l.push("if"!==i&&"if"!==s?(0,u.wy)(e,[[u.F8,d]]):e)}}),i)?(0,u.h)(u.W3,{name:`${i}-transition`,onBeforeLeave:r,onEnter:o,onAfterEnter:n},{default:()=>l}):l}function N(e,t){return(0,u.h)(W,{ref:"addTabInstRef",key:"__addable",name:"__addable",internalCreatedByPane:!0,internalAddable:!0,internalLeftPadded:t,disabled:"object"==typeof e&&e.disabled})}function F(e){let t=(0,u.Ho)(e);return t.props?t.props.internalLeftPadded=!0:t.props={internalLeftPadded:!0},t}function D(e){return Array.isArray(e.dynamicProps)?e.dynamicProps.includes("internalLeftPadded")||e.dynamicProps.push("internalLeftPadded"):e.dynamicProps=["internalLeftPadded"],e}},8965:function(e,t,a){a.d(t,{x:()=>r});let r=(0,a(1579).U)("n-tabs")}}]);