import React from "react";
import ReactDOM from "react-dom";
import Root from "../Root";
import { shallow } from "enzyme";
import { RenderApp } from "../index";
jest.mock("react-dom", () => ({ render: jest.fn() }));

describe("Application root", () => {
  const wrapper = shallow(<Root />);
  it("should render without crashing", () => {
    const div = document.createElement("div");
    div.id = "root";
    document.body.appendChild(div);
    let spy = jest.spyOn(ReactDOM, "render");
    expect(spy).toHaveBeenCalledTimes(0);
  });
  expect(wrapper).toBeDefined();
  expect(wrapper).toMatchSnapshot();
  expect(wrapper.find("#close")).toBeDefined();

  it("should render", () => {
    let renderSpy = jest.spyOn(ReactDOM, "render");
    RenderApp();
    expect(renderSpy).toHaveBeenCalled();
  });
});
