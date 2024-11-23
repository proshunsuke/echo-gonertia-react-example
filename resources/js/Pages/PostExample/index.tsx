import { Head, useForm } from "@inertiajs/react";
import type { FC, FormEvent } from "react";

type Props = {
	text: string;
};

const PostExample: FC<Props> = ({ text }) => {
	const { data, setData, post, processing, errors } = useForm({
		postData: "",
	});

	const submit = (e: FormEvent) => {
		e.preventDefault();
		post("/post-example");
	};

	return (
		<>
			<Head title="PostExample" />
			<div className="grow">
				<h1 className="mb-1 font-bold text-xl">{text}</h1>
				<div className="inset-y-0 start-0 my-px ms-px w-80 rounded-l-lg text-neutral-500">
					<form onSubmit={submit}>
						<div className="flex gap-2">
							<input
								type="text"
								placeholder={"input postData"}
								value={data.postData}
								className="block w-full rounded-lg border border-neutral-200 py-2 ps-2 pe-3 leading-6 placeholder-neutral-500 focus:border-neutral-500 focus:ring focus:ring-neutral-500/25"
								onChange={(e) => setData("postData", e.target.value)}
							/>

							<button
								type="submit"
								className="inline-flex items-center justify-center gap-1 rounded-lg border border-violet-700 bg-violet-700 px-3 py-2 font-semibold text-sm text-white leading-5 hover:border-violet-600 hover:bg-violet-600 hover:text-white focus:ring focus:ring-violet-400/50 active:border-violet-700 active:bg-violet-700 disabled:cursor-not-allowed disabled:opacity-50"
								disabled={processing}
							>
								Submit
							</button>
						</div>
						<div className="flex gap-2">
							{errors.postData && (
								<div className="text-red-500">{errors.postData}</div>
							)}
						</div>
					</form>
				</div>
			</div>
		</>
	);
};

export default PostExample;
