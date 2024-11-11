import Home from "@/Pages/index";
import { expect, test } from "vitest";
import { render } from "vitest-browser-react";

test("renders text", async () => {
	const { getByText } = render(<Home text="title" />);

	await expect.element(getByText("title")).toBeInTheDocument();
});
