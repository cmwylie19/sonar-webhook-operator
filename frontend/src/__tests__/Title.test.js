import React from "react";
import { render, screen } from "@testing-library/react";
import Title from "../components/Title";

test("renders the title", () => {
  render(<Title>ABCD 0123</Title>);
  const linkElement = screen.getByText(/ABCD 0123/i);
  expect(linkElement).toBeInTheDocument();
});
