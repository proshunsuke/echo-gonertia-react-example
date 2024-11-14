import { Layout } from "@/Components/Layout/index";
import type { Page } from "@inertiajs/core";
import { createInertiaApp } from "@inertiajs/react";
import type { ReactElement } from "react";
import { createRoot } from "react-dom/client";
import "../css/app.css";

type PageType = Page & {
	default: {
		layout?: (page: ReactElement) => ReactElement;
	};
};

createInertiaApp({
	resolve: async (name) => {
		const pages = import.meta.glob<Promise<PageType>>("./Pages/**/*.tsx");
		const pageImport = pages[`./Pages/${name}.tsx`];
		const page = await pageImport();
		page.default.layout =
			page.default.layout || ((page: ReactElement) => <Layout>{page}</Layout>);
		return page;
	},
	setup({ el, App, props }) {
		const root = createRoot(el);

		root.render(<App {...props} />);
	},
});
