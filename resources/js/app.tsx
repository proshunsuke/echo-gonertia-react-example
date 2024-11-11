import { Layout } from "@/Components/Layout/index";
import { createInertiaApp } from "@inertiajs/react";
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers";
import { createRoot } from "react-dom/client";
import "../css/app.css";

createInertiaApp({
	resolve: (name) =>
		resolvePageComponent(
			`./Pages/${name}.tsx`,
			import.meta.glob("./Pages/**/*.tsx"),
		),
	setup({ el, App, props }) {
		const root = createRoot(el);

		root.render(
			<Layout>
				<App {...props} />
			</Layout>,
		);
	},
});
