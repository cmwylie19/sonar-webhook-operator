import React from "react";
import ResultList from "../components/ResultList";
import { shallow } from "enzyme";

test("ResultList renders", () => {
  const wrapper = shallow(<ResultList sonarResults={[]} />);

  expect(wrapper).toBeDefined();
});
