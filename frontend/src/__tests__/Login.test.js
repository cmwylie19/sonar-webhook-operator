import React from "react";
import Login from "../components/Login";
import { shallow } from "enzyme";

it("Tab is changed when tab is clicked", () => {
  const wrapper = shallow(<Login />);

  // prints tabs component
  let description = wrapper.find("#description");
  console.log(wrapper.find("#description").debug());
  expect(description.text()).toBe("Whats the deal");
  expect(wrapper).toBeDefined();
});
