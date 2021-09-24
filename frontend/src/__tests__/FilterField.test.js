import React from "react";
import { render, screen } from "@testing-library/react";
import FilterField from "../components/FilterField";

test("renders the title", () => {
  render(<FilterField/>);
  const linkElement = screen.getByText(/Filter results/i);
  expect(linkElement).toBeInTheDocument();
});
