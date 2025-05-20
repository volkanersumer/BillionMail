import { defineConfig } from 'unocss'
import presetUno from '@unocss/preset-uno'
import presetIcons from '@unocss/preset-icons'
import presetAttributify from '@unocss/preset-attributify'
import transformerDirectives from '@unocss/transformer-directives'

export default defineConfig({
	presets: [
		presetUno({ preflight: false }),
		presetAttributify(),
		presetIcons({
			prefix: 'i-',
			extraProperties: {
				display: 'inline-flex',
			},
			collections: {
				mdi: () => import('@iconify-json/mdi/icons.json').then(i => i.default),
				custom: {
					help: '<svg aria-hidden="true" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg"><path fill="currentColor" fill-rule="evenodd" clip-rule="evenodd" d="M6.886 4.923c-.232.25-.386.63-.386 1.122a.75.75 0 1 1-1.5 0c0-.803.254-1.57.788-2.144C6.33 3.32 7.1 3 8 3c.9 0 1.67.32 2.212.901.534.574.788 1.341.788 2.144 0 1.234-.869 1.922-1.436 2.332-.094.068-.178.127-.255.181-.19.134-.33.233-.449.343a.64.64 0 0 0-.11.121V9.5a.75.75 0 0 1-1.5 0V9c0-.563.326-.956.593-1.202.198-.183.457-.364.672-.516l.17-.121c.559-.404.815-.693.815-1.116 0-.492-.154-.872-.386-1.122C8.888 4.68 8.533 4.5 8 4.5c-.533 0-.888.18-1.114.423Z"></path><path d="M9 12a1.001 1.001 0 0 1-2 0 1.001 1.001 0 0 1 2 0Z"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M8 14.5A6.5 6.5 0 1 0 1.5 8c0 1.022.179 1.608.603 2.399.316.59.407 1.313.178 2.003l-.66 1.976 1.977-.659a2.541 2.541 0 0 1 2.004.178c.79.424 1.376.603 2.398.603ZM8 16a8 8 0 1 0-8-8c0 1.29.25 2.117.78 3.107.136.252.168.549.078.82l-.807 2.42a1 1 0 0 0-.051.315V15a1 1 0 0 0 1 1h.338a1 1 0 0 0 .316-.051l2.419-.807c.271-.09.568-.057.82.078.99.53 1.817.78 3.107.78Z"></path></svg>',
					smtp: '<svg t="1747640620453" viewBox="0 0 1025 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3817" xmlns:xlink="http://www.w3.org/1999/xlink"><path d="M841.142857 73.142857a109.714286 109.714286 0 0 1 109.202286 99.181714L950.857143 182.857143v294.107428a36.571429 36.571429 0 0 1-72.557714 6.582858L877.714286 476.964571V194.340571L569.051429 459.776c-32.182857 19.529143-71.826286 21.211429-108.909715 2.121143l-10.020571-5.705143L146.285714 194.56v573.44a36.571429 36.571429 0 0 0 29.988572 35.986286L182.857143 804.571429h221.842286a36.571429 36.571429 0 0 1 6.582857 72.557714l-6.582857 0.585143H182.857143a109.714286 109.714286 0 0 1-109.202286-99.181715L73.142857 768v-585.142857a109.714286 109.714286 0 0 1 99.181714-109.202286L182.857143 73.142857h658.285714z m-19.675428 73.142857H202.24L492.982857 397.238857a40.740571 40.740571 0 0 0 29.257143 5.046857l3.949714-1.462857L821.467429 146.285714z" fill="currentColor" p-id="3818"></path><path d="M857.526857 626.395429l5.632 3.510857 146.651429 109.421714a36.571429 36.571429 0 0 1 5.558857 53.613714l-5.266286 4.827429-142.482286 108.544a36.571429 36.571429 0 0 1-49.225142-53.76l4.900571-4.388572 56.32-43.008-257.755429 0.073143a36.571429 36.571429 0 0 1-6.582857-72.557714l6.582857-0.585143H877.714286l-58.294857-43.593143a36.571429 36.571429 0 0 1 38.107428-62.098285z" fill="currentColor" p-id="3819"></path></svg>',
				},
			},
		}),
	],
	transformers: [transformerDirectives({ enforce: 'pre' })],
	theme: {
		colors: {
			primary: 'var(--color-primary-1)',
			primaryHover: 'var(--color-primary-hover-1)',
			error: 'var(--color-error-1)',
			warning: 'var(--color-warning-1)',
			desc: 'var(--color-text-3)',
		},
	},
	shortcuts: {
		'flex-center': 'flex items-center justify-center',
	},
})
