import { Link } from "@inertiajs/react";
import type React from "react";
import type { FC } from "react";

type Props = {
	children: React.JSX.Element;
};

export const Layout: FC<Props> = ({ children }) => (
	<div
		id="page-container"
		className="mx-auto flex min-h-screen w-full min-w-[320px] flex-col bg-slate-50"
	>
		<header id="page-header" className="z-1 flex flex-none items-center pt-5">
			<div className="container mx-auto px-4 lg:px-8 xl:max-w-7xl">
				<div className="-mx-4 lg:-mx-6 border-slate-200 border-y bg-white px-4 shadow-sm sm:rounded-lg sm:border lg:px-6">
					<div className="flex justify-between py-2.5 lg:py-3.5">
						<div className="flex items-center gap-2 lg:gap-6">
							<nav className="hidden items-center gap-1.5 lg:flex">
								<Link
									href="/"
									className="group flex items-center gap-2 rounded-lg px-2.5 py-1.5 font-medium text-slate-800 text-sm hover:bg-violet-50 hover:text-violet-600 active:border-violet-100"
								>
									<span>HOME</span>
								</Link>
								<Link
									href="/post-example"
									className="group flex items-center gap-2 rounded-lg px-2.5 py-1.5 font-medium text-slate-800 text-sm hover:bg-violet-50 hover:text-violet-600 active:border-violet-100"
								>
									<span>Post Example</span>
								</Link>
								<Link
									href="/404-error-page"
									className="group flex items-center gap-2 rounded-lg px-2.5 py-1.5 font-medium text-slate-800 text-sm hover:bg-violet-50 hover:text-violet-600 active:border-violet-100"
								>
									<span>404 Error Page</span>
								</Link>
							</nav>
						</div>
					</div>
				</div>
			</div>
		</header>
		<main id="page-content" className="flex max-w-full flex-auto flex-col">
			<div className="container mx-auto px-4 pt-6 lg:px-8 lg:pt-8 xl:max-w-7xl">
				{children}
			</div>
		</main>
	</div>
);
