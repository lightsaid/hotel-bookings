/// <reference types="vite/client" />


/// <reference types="vite/client" />
 
declare module '*.vue' {
	import { DefineComponent } from 'vue';
	// eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
	const component: DefineComponent<{}, {}, any>;
	export default component;
}
 
// 环境变量 TypeScript的智能提示
interface ImportMetaEnv {
	// VITE_APP_XXX: string;
}
 
interface ImportMeta {
	readonly env: ImportMetaEnv;
}