"use strict";(self.webpackChunkfrontend=self.webpackChunkfrontend||[]).push([["81"],{6768:function(e,n,t){function r(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter(function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable})),t.push.apply(t,r)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?r(Object(t),!0).forEach(function(n){var r;r=t[n],n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):r(Object(t)).forEach(function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))})}return e}function i(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,r=Array(n);t<n;t++)r[t]=e[t];return r}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter(function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable})),t.push.apply(t,r)}return t}function c(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach(function(n){var r;r=t[n],n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach(function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))})}return e}function l(e){return function n(){for(var t=this,r=arguments.length,o=Array(r),i=0;i<r;i++)o[i]=arguments[i];return o.length>=e.length?e.apply(this,o):function(){for(var e=arguments.length,r=Array(e),i=0;i<e;i++)r[i]=arguments[i];return n.apply(t,[].concat(o,r))}}}function s(e){return({}).toString.call(e).includes("Object")}function u(e){return"function"==typeof e}t.d(n,{Z:()=>R});var d,f,p=l(function(e,n){throw Error(e[n]||e.default)})({initialIsRequired:"initial state is required",initialType:"initial state should be an object",initialContent:"initial state shouldn't be an empty object",handlerType:"handler should be an object or a function",handlersType:"all handlers should be a functions",selectorType:"selector should be a function",changeType:"provided value of changes should be an object",changeField:'it seams you want to change a field in the state which is not specified in the "initial" state',default:"an unknown error accured in `state-local` package"}),h=function(e,n){return s(n)||p("changeType"),Object.keys(n).some(function(n){return!Object.prototype.hasOwnProperty.call(e,n)})&&p("changeField"),n},v=function(e){u(e)||p("selectorType")},m=function(e){u(e)||s(e)||p("handlerType"),s(e)&&Object.values(e).some(function(e){return!u(e)})&&p("handlersType")},b=function(e){e||p("initialIsRequired"),s(e)||p("initialType"),!Object.keys(e).length&&p("initialContent")};function g(e,n){return u(n)?n(e.current):n}function y(e,n){return e.current=c(c({},e.current),n),n}function j(e,n,t){return u(n)?n(e.current):Object.keys(t).forEach(function(t){var r;return null===(r=n[t])||void 0===r?void 0:r.call(n,e.current[t])}),t}var O={configIsRequired:"the configuration object is required",configType:"the configuration object should be an object",default:"an unknown error accured in `@monaco-editor/loader` package",deprecation:"Deprecation warning!\n    You are using deprecated way of configuration.\n\n    Instead of using\n      monaco.config({ urls: { monacoBase: '...' } })\n    use\n      monaco.config({ paths: { vs: '...' } })\n\n    For more please check the link https://github.com/suren-atoyan/monaco-loader#config\n  "},w=(d=function(e,n){throw Error(e[n]||e.default)},function e(){for(var n=this,t=arguments.length,r=Array(t),o=0;o<t;o++)r[o]=arguments[o];return r.length>=d.length?d.apply(this,r):function(){for(var t=arguments.length,o=Array(t),i=0;i<t;i++)o[i]=arguments[i];return e.apply(n,[].concat(r,o))}})(O);let k=function(e){return e||w("configIsRequired"),!({}).toString.call(e).includes("Object")&&w("configType"),e.urls?(console.warn(O.deprecation),{paths:{vs:e.urls.monacoBase}}):e},x=function(){for(var e=arguments.length,n=Array(e),t=0;t<e;t++)n[t]=arguments[t];return function(e){return n.reduceRight(function(e,n){return n(e)},e)}};var P={type:"cancelation",msg:"operation is manually canceled"};let S=function(e){var n=!1,t=new Promise(function(t,r){e.then(function(e){return n?r(P):t(e)}),e.catch(r)});return t.cancel=function(){return n=!0},t};var z=function(e){if(Array.isArray(e))return e}(f=({create:function(e){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:{};b(e),m(n);var t={current:e},r=l(j)(t,n),o=l(y)(t),i=l(h)(e),a=l(g)(t);return[function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:function(e){return e};return v(e),e(t.current)},function(e){(function(){for(var e=arguments.length,n=Array(e),t=0;t<e;t++)n[t]=arguments[t];return function(e){return n.reduceRight(function(e,n){return n(e)},e)}})(r,o,i,a)(e)}]}}).create({config:{paths:{vs:"https://cdn.jsdelivr.net/npm/monaco-editor@0.52.2/min/vs"}},isInitialized:!1,resolve:null,reject:null,monaco:null}))||function(e,n){if("undefined"!=typeof Symbol&&Symbol.iterator in Object(e)){var t=[],r=!0,o=!1,i=void 0;try{for(var a,c=e[Symbol.iterator]();!(r=(a=c.next()).done)&&(t.push(a.value),t.length!==n);r=!0);}catch(e){o=!0,i=e}finally{try{r||null==c.return||c.return()}finally{if(o)throw i}}return t}}(f,2)||function(e,n){if(e){if("string"==typeof e)return i(e,2);var t=Object.prototype.toString.call(e).slice(8,-1);if("Object"===t&&e.constructor&&(t=e.constructor.name),"Map"===t||"Set"===t)return Array.from(e);if("Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return i(e,n)}}(f,2)||function(){throw TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}(),C=z[0],B=z[1];function E(e){return document.body.appendChild(e)}function I(e){var n,t,r=C(function(e){return{config:e.config,reject:e.reject}}),o=(n="".concat(r.config.paths.vs,"/loader.js"),t=document.createElement("script"),n&&(t.src=n),t);return o.onload=function(){return e()},o.onerror=r.reject,o}function T(){var e=C(function(e){return{config:e.config,resolve:e.resolve,reject:e.reject}}),n=window.require;n.config(e.config),n(["vs/editor/editor.main"],function(n){Z(n),e.resolve(n)},function(n){e.reject(n)})}function Z(e){C().monaco||B({monaco:e})}var $=new Promise(function(e,n){return B({resolve:e,reject:n})});let R={config:function(e){var n=k(e),t=n.monaco,r=function(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],!(n.indexOf(t)>=0)&&Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}(n,["monaco"]);B(function(e){return{config:function e(n,t){return Object.keys(t).forEach(function(r){t[r]instanceof Object&&n[r]&&Object.assign(t[r],e(n[r],t[r]))}),o(o({},n),t)}(e.config,r),monaco:t}})},init:function(){var e=C(function(e){return{monaco:e.monaco,isInitialized:e.isInitialized,resolve:e.resolve}});if(!e.isInitialized){if(B({isInitialized:!0}),e.monaco)return e.resolve(e.monaco),S($);if(window.monaco&&window.monaco.editor)return Z(window.monaco),e.resolve(window.monaco),S($);x(E,I)(T)}return S($)},__getMonacoInstance:function(){return C(function(e){return e.monaco})}}},4944:function(e,n,t){t.d(n,{L:()=>r});function r(e,n){if(!e)return;let t=document.createElement("a");t.href=e,void 0!==n&&(t.download=n),document.body.appendChild(t),t.click(),document.body.removeChild(t)}},8439:function(e,n,t){t.d(n,{W3:()=>d,ZP:()=>p});var r=t(209),o=t(1321),i=t(4124),a=t(6169),c=t(1579),l=t(3723),s=t(2249);let u=(0,s.cB)("breadcrumb",`
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
 `)])])]),d=(0,c.U)("n-breadcrumb"),f=Object.assign(Object.assign({},o.Z.props),{separator:{type:String,default:"/"}}),p=(0,r.aZ)({name:"Breadcrumb",props:f,setup(e){let{mergedClsPrefixRef:n,inlineThemeDisabled:t}=(0,i.ZP)(e),c=(0,o.Z)("Breadcrumb","-breadcrumb",u,l.Z,e,n);(0,r.JJ)(d,{separatorRef:(0,r.Vh)(e,"separator"),mergedClsPrefixRef:n});let s=(0,r.Fl)(()=>{let{common:{cubicBezierEaseInOut:e},self:{separatorColor:n,itemTextColor:t,itemTextColorHover:r,itemTextColorPressed:o,itemTextColorActive:i,fontSize:a,fontWeightActive:l,itemBorderRadius:s,itemColorHover:u,itemColorPressed:d,itemLineHeight:f}}=c.value;return{"--n-font-size":a,"--n-bezier":e,"--n-item-text-color":t,"--n-item-text-color-hover":r,"--n-item-text-color-pressed":o,"--n-item-text-color-active":i,"--n-separator-color":n,"--n-item-color-hover":u,"--n-item-color-pressed":d,"--n-item-border-radius":s,"--n-font-weight-active":l,"--n-item-line-height":f}}),f=t?(0,a.F)("breadcrumb",void 0,s,e):void 0;return{mergedClsPrefix:n,cssVars:t?void 0:s,themeClass:null==f?void 0:f.themeClass,onRender:null==f?void 0:f.onRender}},render(){var e;return null===(e=this.onRender)||void 0===e||e.call(this),(0,r.h)("nav",{class:[`${this.mergedClsPrefix}-breadcrumb`,this.themeClass],style:this.cssVars,"aria-label":"Breadcrumb"},(0,r.h)("ul",null,this.$slots))}})},6666:function(e,n,t){t.d(n,{Z:()=>c});var r=t(209),o=t(8282),i=t(4934),a=t(8439);let c=(0,r.aZ)({name:"BreadcrumbItem",props:{separator:String,href:String,clickable:{type:Boolean,default:!0},onClick:Function},slots:Object,setup(e,{slots:n}){let t=(0,r.f3)(a.W3,null);if(!t)return()=>null;let{separatorRef:c,mergedClsPrefixRef:l}=t,s=function(e=i.j?window:null){let n=()=>{let{hash:n,host:t,hostname:r,href:o,origin:i,pathname:a,port:c,protocol:l,search:s}=(null==e?void 0:e.location)||{};return{hash:n,host:t,hostname:r,href:o,origin:i,pathname:a,port:c,protocol:l,search:s}},t=(0,r.iH)(n()),o=()=>{t.value=n()};return(0,r.bv)(()=>{e&&(e.addEventListener("popstate",o),e.addEventListener("hashchange",o))}),(0,r.SK)(()=>{e&&(e.removeEventListener("popstate",o),e.removeEventListener("hashchange",o))}),t}(),u=(0,r.Fl)(()=>e.href?"a":"span"),d=(0,r.Fl)(()=>s.value.href===e.href?"location":null);return()=>{let{value:t}=l;return(0,r.h)("li",{class:[`${t}-breadcrumb-item`,e.clickable&&`${t}-breadcrumb-item--clickable`]},(0,r.h)(u.value,{class:`${t}-breadcrumb-item__link`,"aria-current":d.value,href:e.href,onClick:e.onClick},n),(0,r.h)("span",{class:`${t}-breadcrumb-item__separator`,"aria-hidden":"true"},(0,o.gI)(n.separator,()=>{var n;return[null!==(n=e.separator)&&void 0!==n?n:c.value]})))}}})},8574:function(e,n,t){t.d(n,{Z:()=>m});var r=t(5083),o=t(2370),i=t(209),a=t(4131),c=t(1321),l=t(4124),s=t(6169),u=t(2249),d=t(2379),f=t(3291);let p=(0,u.c)([(0,u.c)("@keyframes spin-rotate",`
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
 `,[(0,f.h)()])]),(0,u.cB)("spin-body",`
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
 `)])]),h={small:20,medium:18,large:16},v=Object.assign(Object.assign({},c.Z.props),{contentClass:String,contentStyle:[Object,String],description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0},delay:Number}),m=(0,i.aZ)({name:"Spin",props:v,slots:Object,setup(e){let{mergedClsPrefixRef:n,inlineThemeDisabled:t}=(0,l.ZP)(e),a=(0,c.Z)("Spin","-spin",p,d.Z,e,n),f=(0,i.Fl)(()=>{let{size:n}=e,{common:{cubicBezierEaseInOut:t},self:o}=a.value,{opacitySpinning:i,color:c,textColor:l}=o;return{"--n-bezier":t,"--n-opacity-spinning":i,"--n-size":"number"==typeof n?(0,r.BL)(n):o[(0,u.Tl)("size",n)],"--n-color":c,"--n-text-color":l}}),v=t?(0,s.F)("spin",(0,i.Fl)(()=>{let{size:n}=e;return"number"==typeof n?String(n):n[0]}),f,e):void 0,m=(0,o.Z)(e,["spinning","show"]),b=(0,i.iH)(!1);return(0,i.m0)(n=>{let t;if(m.value){let{delay:r}=e;if(r){t=window.setTimeout(()=>{b.value=!0},r),n(()=>{clearTimeout(t)});return}}b.value=m.value}),{mergedClsPrefix:n,active:b,mergedStrokeWidth:(0,i.Fl)(()=>{let{strokeWidth:n}=e;if(void 0!==n)return n;let{size:t}=e;return h["number"==typeof t?"medium":t]}),cssVars:t?void 0:f,themeClass:null==v?void 0:v.themeClass,onRender:null==v?void 0:v.onRender}},render(){var e,n;let{$slots:t,mergedClsPrefix:r,description:o}=this,c=t.icon&&this.rotate,l=(o||t.description)&&(0,i.h)("div",{class:`${r}-spin-description`},o||(null===(e=t.description)||void 0===e?void 0:e.call(t))),s=t.icon?(0,i.h)("div",{class:[`${r}-spin-body`,this.themeClass]},(0,i.h)("div",{class:[`${r}-spin`,c&&`${r}-spin--rotate`],style:t.default?"":this.cssVars},t.icon()),l):(0,i.h)("div",{class:[`${r}-spin-body`,this.themeClass]},(0,i.h)(a.Z,{clsPrefix:r,style:t.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${r}-spin`}),l);return null===(n=this.onRender)||void 0===n||n.call(this),t.default?(0,i.h)("div",{class:[`${r}-spin-container`,this.themeClass],style:this.cssVars},(0,i.h)("div",{class:[`${r}-spin-content`,this.active&&`${r}-spin-content--spinning`,this.contentClass],style:this.contentStyle},t),(0,i.h)(i.uT,{name:"fade-in-transition"},{default:()=>this.active?s:null})):s}})}}]);