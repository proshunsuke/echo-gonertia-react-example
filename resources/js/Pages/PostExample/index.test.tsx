import PostExample from "@/Pages/PostExample/index";
import { expect, test } from "vitest";
import { render } from "vitest-browser-react";

test("renders text", async () => {
	const { getByText } = render(<PostExample text="title" />);

	await expect.element(getByText("title")).toBeInTheDocument();
});
