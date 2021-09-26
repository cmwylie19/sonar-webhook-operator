import React from "react";
import Signup from "../components/Signup";
import { shallow } from "enzyme";

it("Signup Component", () => {
  const wrapper = shallow(<Signup />);

  let signin = wrapper.find("#signin");
  let submit = wrapper.find("#submit");

  console.log(wrapper.find("#signin").debug());

  expect(signin.text()).toBe("Already have an account? Sign in");
  expect(submit.text()).toBe("Sign Up");
  expect(wrapper).toBeDefined();
});
