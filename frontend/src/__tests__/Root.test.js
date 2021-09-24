import React from "react";
import { render, screen } from "@testing-library/react";
import Root from "../Root";

test("renders the title", () => {
  render(<Root />);
  const linkElement = screen.getByText(/Sonar Webhook/i);
  expect(linkElement).toBeInTheDocument();
});
