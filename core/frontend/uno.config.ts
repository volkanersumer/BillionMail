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
