import { Head, useForm } from "@inertiajs/react";
import React, { type FC, type FormEvent } from "react";

type Props = {
	text: string;
};

const PostExample: FC<Props> = ({ text }) => {
	const { data, setData, post, processing, errors } = useForm({
		post_data: "",
	});

	const submit = (e: FormEvent) => {
		e.preventDefault();
		post("/post-example");
	};

	return (
		<>
			<Head title="PostExample" />
			<div className="grow">
				<h1 className="mb-1 text-xl font-bold">{text}</h1>
				<div className="inset-y-0 start-0 my-px ms-px w-80 rounded-l-lg text-neutral-500">
					<form onSubmit={submit}>
						<div className="flex gap-2">
							<input
								type="text"
								placeholder={"input post_data"}
								value={data.post_data}
								className="block w-full rounded-lg border border-neutral-200 py-2 pe-3 ps-2 leading-6 placeholder-neutral-500 focus:border-neutral-500 focus:ring focus:ring-neutral-500/25"
								onChange={(e) => setData("post_data", e.target.value)}
							/>

							<button
								type="submit"
								className="inline-flex items-center justify-center gap-1 rounded-lg border border-violet-700 bg-violet-700 px-3 py-2 text-sm font-semibold leading-5 text-white hover:border-violet-600 hover:bg-violet-600 hover:text-white focus:ring focus:ring-violet-400/50 active:border-violet-700 active:bg-violet-700 disabled:opacity-50 disabled:cursor-not-allowed"
								disabled={processing}
							>
								Submit
							</button>
						</div>
						<div className="flex gap-2">
							{errors.post_data && (
								<div className="text-red-500">{errors.post_data}</div>
							)}
						</div>
					</form>
				</div>
			</div>
		</>
	);
};

export default PostExample;
