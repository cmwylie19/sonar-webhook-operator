import React from "react";
import { shallow } from "enzyme";
import Theme, { pinkTheme } from "../components/Theme";

describe("ThemeContainer", () => {
  let wrapper = shallow(
    <Theme theme={pinkTheme}>
      <div id="child">child</div>
    </Theme>
  );
  it("mounts", () => {
    expect(wrapper).toBeDefined();
  });

  it("matches snapshot", () => {
    expect(wrapper).toMatchSnapshot();
  });
});
